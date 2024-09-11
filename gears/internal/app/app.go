package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"goclub/gears/internal/config"
	"goclub/gears/internal/logger"
	grpcserver "goclub/gears/internal/pkg/grpc_server"
	httpserver "goclub/gears/internal/pkg/http_server"
	serviceprovider "goclub/gears/internal/pkg/service_provider"

	"golang.org/x/sync/errgroup"
)

const pkgLogName = "GoClubGears"

type App interface {
	Run() error
}

type app struct {
	ctx context.Context
	sp  *serviceprovider.ServiceProvider
	cfg *config.AppConfig
}

func NewApp(ctx context.Context, cfg *config.AppConfig) App {
	return &app{
		ctx: ctx,
		sp:  serviceprovider.NewServiceProvider(cfg),
		cfg: cfg,
	}
}

func (a *app) Run() (err error) {
	const fnLogName = ".Run"
	logger.Info(a.ctx, "GoClubGears app is starting...")
	defer logger.Info(a.ctx, "app finished")

	if err = a.sp.OpenStorage(a.ctx); err != nil {
		return err
	}

	defer a.sp.CloseStorage(a.ctx)

	ctx, cancel := context.WithCancel(a.ctx)
	defer cancel()

	clubAPI, caErr := a.sp.GetClubAPI(a.ctx)
	if caErr != nil {
		return fmt.Errorf(pkgLogName+fnLogName+" - OrderAPI error: %w", caErr)
	}

	go func() {
		c := make(chan os.Signal, 1) // we need to reserve to buffer size 1, so the notifier are not blocked
		signal.Notify(c, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

		<-c
		logger.Info(ctx, pkgLogName+fnLogName+" - syscall notify")
		cancel()
	}()

	// Grpc Server >>>
	grpcServer := grpcserver.NewServer(a.ctx, a.cfg.GRPCAddress)
	grpcServer.RegisterAPI([]grpcserver.API{clubAPI})
	g, gCtx := errgroup.WithContext(ctx)

	g.Go(func() error {
		err := grpcServer.Start()
		logger.Info(ctx, pkgLogName+fnLogName+" - grpcServer ends", "error", err)
		cancel()
		return err
	})
	g.Go(func() error {
		<-gCtx.Done()
		logger.Info(ctx, pkgLogName+fnLogName+" - grpc server shutdown!")
		return grpcServer.Stop()
	})
	// <<< gRPC Server

	// Http Server >>>
	httpServer, err := httpserver.NewServer(a.ctx, a.cfg.HTTPAddress, a.cfg.GRPCAddress)
	if err != nil {
		logger.Error(ctx, pkgLogName+fnLogName+" - httpserver.NewServer fail", err)
		cancel()
	}
	err = httpServer.RegisterAPI([]httpserver.API{clubAPI})
	if err != nil {
		logger.Error(ctx, pkgLogName+fnLogName+" - httpserver.RegisterAPI fail", err)
		cancel()
	}

	g.Go(func() error {
		err := httpServer.Start()
		cancel()
		return fmt.Errorf(pkgLogName+fnLogName+" - httpServer ends: %w", err)
	})

	g.Go(func() error {
		<-gCtx.Done()
		return fmt.Errorf(pkgLogName+fnLogName+" - httpServer shutdown: %w", httpServer.Stop())
	})
	// <<< Http Server

	logger.Info(a.ctx, "GoClubGears is running")

	//---------------------
	if err := g.Wait(); err != nil {
		logger.Info(ctx, "Exit", "reason", err)
	}

	return err

}
