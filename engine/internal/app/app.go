package app

import (
	"context"
	"fmt"
	"goclub/common/logger"
	"goclub/engine/internal/config"
	grpcserver "goclub/engine/internal/grpc_server"
	httpserver "goclub/engine/internal/http_server"
	"goclub/engine/internal/repository"
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
		appSvc     service.AppService
		appSrv     server.AppServer
		errGrp     *errgroup.Group
		errGrpCtx  context.Context
		grpcSrv    grpcserver.GRPCServer
		httpSrv    httpserver.HTTPServer
		cancelFunc context.CancelFunc
		repo       repository.Repository
	}
)

func NewApp() app {
	return app{}
}

// AppService внутренний контролер логики
func (a *app) AppService(ctc context.Context) service.AppService {
	if a.appSvc == nil {
		a.appSvc = service.NewAppService(a.Repo())
	}
	return a.appSvc
}

// AppServer доступ к сервисам приложения
func (a *app) AppServer(ctx context.Context) server.AppServer {
	if a.appSrv == nil {
		a.appSrv = server.NewAppServer(a.AppService(ctx))
	}
	return a.appSrv
}

func (a *app) Repo() repository.Repository {
	if a.repo == nil {
		a.repo = repository.NewPgDBRepo(config.Cfg.DSNs)
	}
	return a.repo
}

// runComponents запускает с работу компоненты системы с возможностью остановки по отмене контекста
func (a *app) runComponent(ctx context.Context, ssi AppComponentInfo) {
	if ssi.cargo.Start != nil {
		a.errGrp.Go(func() error {
			logger.Info(ctx, ssi.title+" - Launch component")
			err := ssi.cargo.Start()
			logger.Info(ctx, ssi.title+" - Component stopped", "error", err)
			a.cancelFunc()
			return err
		})
	}
	if ssi.cargo.Stop != nil {
		a.errGrp.Go(func() error {
			<-a.errGrpCtx.Done()
			logger.Info(ctx, ssi.title+" - Shutdowning")
			if err := ssi.cargo.Stop(); err != nil {
				return fmt.Errorf(ssi.title+" - shutdown error: %w", err)
			}
			return nil
		})
	}
}

func (a *app) Run(ctx context.Context, cfg *config.AppConfig) (err error) {
	const fnName = "Run"
	logger.Debug(ctx, fmt.Sprintf(pkgName+fnName+": cfg=%v", *cfg))

	//Создаем контекст с отменой по прерыванию со стороны ОС
	ctx, a.cancelFunc = signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer a.cancelFunc()

	//a.AppServer(ctx) // Инициализаруем объекты по цепочке Server-Service-Repo-...

	if err = a.Repo().Open(); err != nil {
		return err
	}
	//Создаем "группу контроля работы" errgroup.Group и констекст для остановки смежных процессов
	a.errGrp, a.errGrpCtx = errgroup.WithContext(ctx)

	// Запускаем репозиторий
	repoStartStoper := NewStartStopAdpt(ctx, nil, a.Repo().Close)
	a.runComponent(ctx, AppComponentInfo{cargo: repoStartStoper, title: "Repository"})
	// <<<

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
