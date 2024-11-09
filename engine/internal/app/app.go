package app

import (
	"context"
	"fmt"
	"goclub/common/logger"
	"goclub/engine/internal/config"
	"goclub/engine/internal/controller"
	grpcserver "goclub/engine/internal/grpc_server"
	httpserver "goclub/engine/internal/http_server"
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

type (
	app struct {
		//Providers, Adapters and Services
		controller controller.Controller
		appSvc     service.AppService
		grpcSvc    grpcserver.GRPCService
	}
)

func NewApp() app {
	return app{}
}

func (a *app) Controller(ctc context.Context) controller.Controller {
	if a.controller == nil {
		a.controller = controller.NewController()
	}
	return a.controller
}

func (a *app) AppService(ctx context.Context) service.AppService {
	if a.appSvc == nil {
		a.appSvc = service.NewAPI(a.Controller(ctx))
	}
	return a.appSvc
}

func (a *app) Run(ctx context.Context, cfg *config.AppConfig) error {
	const fnName = "Run"
	logger.Debug(ctx, fmt.Sprintf(pkgName+fnName+": cfg=%v", *cfg))
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	g, gCtx := errgroup.WithContext(ctx)

	// Grpc Server >>>
	grpcSrv := grpcserver.NewGRPCServer(ctx, cfg.GRPCAddress)
	grpcSrv.RegisterAPI([]grpcserver.GRPCService{a.AppService(ctx)})

	g.Go(func() error {
		err := grpcSrv.Start()
		logger.Info(ctx, pkgName+fnName+" - service ends", "error", err)
		cancel()
		return err
	})
	g.Go(func() error {
		<-gCtx.Done()
		logger.Info(ctx, pkgName+fnName+" - gPRC server shutdown!")
		if err := grpcSrv.Stop(); err != nil {
			return fmt.Errorf(pkgName+fnName+" - gPRC server shutdown error: %w", err)
		}
		return grpcSrv.Stop()
	})
	// <<< gRPC Server

	// Http Server >>>
	httpSrv, err := httpserver.NewServer(ctx, cfg.HTTPAddress, cfg.GRPCAddress)
	if err != nil {
		logger.Error(ctx, pkgName+fnName+" - httpserver.NewServer fail", err)
		cancel()
	}
	err = httpSrv.RegisterAPI([]httpserver.HTTPService{a.AppService(ctx)})
	if err != nil {
		logger.Error(ctx, pkgName+fnName+" - httpserver.RegisterAPI fail", err)
		cancel()
	}

	g.Go(func() error {
		defer cancel()
		if err := httpSrv.Start(); err != nil {
			return fmt.Errorf(pkgName+fnName+" - httpServer ends: %w", err)
		}
		return nil
	})

	g.Go(func() error {
		<-gCtx.Done()
		logger.Info(ctx, pkgName+fnName+" - HTTP server shutdown!")
		if err := httpSrv.Stop(); err != nil {
			return fmt.Errorf(pkgName+fnName+" - HTTP server shutdown error: %w", err)
		}
		return nil
	})
	// <<< Http Server

	logger.Info(ctx, "Application started")
	return g.Wait()
}
