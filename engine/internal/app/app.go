package app

import (
	"context"
	"fmt"
	"goclub/common/logger"
	"goclub/engine/internal/config"
	grpcserver "goclub/engine/internal/grpc_server"
	httpserver "goclub/engine/internal/http_server"
	"goclub/engine/internal/server"
	"goclub/engine/internal/service"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

const (
	pkgName                = "app."
	shutdownTimeoutSeconds = 3
)

type StarterStoper interface {
	Start() error
	Stop() error
}

type AppComponentInfo struct {
	title string
	cargo StarterStoper
}

type AppService interface {
	APIServers() []server.APIServer
}

type (
	app struct {
		//Providers, Adapters and Services
		controller service.AppService
		appSvc     server.AppService
		errGrp     *errgroup.Group
		errGrpCtx  context.Context
		grpcSrv    grpcserver.GRPCServer
		httpSrv    httpserver.HTTPServer
		cancelFunc context.CancelFunc
	}
)

func NewApp() app {
	return app{}
}

// Controller внутренний контролер логики
func (a *app) Controller(ctc context.Context) service.AppService {
	if a.controller == nil {
		a.controller = service.NewService()
	}
	return a.controller
}

// AppServer доступ к сервисам приложения
func (a *app) AppServer(ctx context.Context) server.AppService {
	if a.appSvc == nil {
		a.appSvc = server.NewAppServer(a.Controller(ctx))
	}
	return a.appSvc
}

// runComponents запускает с работу компоненты системы с возможностью остановки по отмене контекста
func (a *app) runComponent(ctx context.Context, ssi AppComponentInfo) {
	a.errGrp.Go(func() error {
		logger.Info(ctx, ssi.title+" - Launch component")
		err := ssi.cargo.Start()
		logger.Info(ctx, ssi.title+" - Component stopped", "error", err)
		a.cancelFunc()
		return err
	})
	a.errGrp.Go(func() error {
		<-a.errGrpCtx.Done()
		logger.Info(ctx, ssi.title+" - Shutdowning")
		if err := ssi.cargo.Stop(); err != nil {
			return fmt.Errorf(ssi.title+" - shutdown error: %w", err)
		}
		return nil
	})
}

func (a *app) Run(ctx context.Context, cfg *config.AppConfig) (err error) {
	const fnName = "Run"
	logger.Debug(ctx, fmt.Sprintf(pkgName+fnName+": cfg=%v", *cfg))

	//Создаем контекст с отменой по прерыванию со стороны ОС
	ctx, a.cancelFunc = signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer a.cancelFunc()

	//Создаем "группу контроля работы" errgroup.Group и констекст для остановки смежных процессов
	a.errGrp, a.errGrpCtx = errgroup.WithContext(ctx)

	// Создаем и запускаем gRPC-сервер
	a.grpcSrv = grpcserver.NewGRPCServer(ctx, cfg.GRPCAddress)
	a.grpcSrv.RegisterAPI(a.AppServer(ctx).GRPCAPIs())
	a.runComponent(ctx, AppComponentInfo{cargo: a.grpcSrv, title: "gRPC server"})
	// <<<

	// Создаем и запускаем HTTP-сервер
	if a.httpSrv, err = httpserver.NewServer(ctx, cfg.HTTPAddress, cfg.GRPCAddress); err != nil {
		return fmt.Errorf(pkgName+fnName+" - httpserver.NewServer fail: %w", err)
	}
	if err = a.httpSrv.RegisterAPI(a.AppServer(ctx).HTTPAPIs()); err != nil {
		return fmt.Errorf(pkgName+fnName+" - httpserver.RegisterAPI fail: %w", err)
	}
	a.runComponent(ctx, AppComponentInfo{cargo: a.httpSrv, title: "HTTP server"})
	// <<<

	logger.Info(ctx, "Application started")
	err = a.errGrp.Wait()
	return err
}
