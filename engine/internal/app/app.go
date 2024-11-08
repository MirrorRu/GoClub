package app

import (
	"context"
	"fmt"
	"goclub/common/logger"
	"goclub/engine/internal/config"
	"goclub/engine/internal/controller"
	grpcserver "goclub/engine/internal/grpc_server"
	grpcservice "goclub/engine/internal/grpc_service"
	httpserver "goclub/engine/internal/http_server"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

const (
	pkgName                = "app."
	shutdownTimeoutSeconds = 3
)

type (
	app struct {
		//Providers, Adapters and Services
		controller  controller.Controller
		grpcService grpcserver.GRPCService
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
func (a *app) GRPCService(ctx context.Context) grpcserver.GRPCService {
	if a.grpcService == nil {
		a.grpcService = grpcservice.NewAPI(a.Controller(ctx))
	}
	return a.grpcService
}

func (a *app) Run(ctx context.Context, cfg *config.AppConfig) error {
	const fnName = "Run"
	logger.Debug(ctx, fmt.Sprintf(pkgName+fnName+": cfg=%v", *cfg))
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	g, gCtx := errgroup.WithContext(ctx)

	// Grpc Server >>>
	grpcServer := grpcserver.NewGRPCServer(ctx, cfg.GRPCAddress)
	grpcServer.RegisterAPI([]grpcserver.GRPCService{a.GRPCService(ctx)})

	g.Go(func() error {
		err := grpcServer.Start()
		logger.Info(ctx, pkgName+fnName+" - grpcService ends", "error", err)
		stop()
		return err
	})
	g.Go(func() error {
		<-gCtx.Done()
		logger.Info(ctx, pkgName+fnName+" - grpc service shutdown!")
		return grpcServer.Stop()
	})
	// <<< gRPC Server

	// Http Server >>>
	httpServer, err := httpserver.NewServer(ctx, cfg.HTTPAddress, cfg.GRPCAddress)
	if err != nil {
		logger.Error(ctx, pkgName+fnName+" - httpserver.NewServer fail", err)
		stop()
	}
	err = httpServer.RegisterAPI([]httpserver.HTTPService{a.GRPCService(ctx)})
	if err != nil {
		logger.Error(ctx, pkgName+fnName+" - httpserver.RegisterAPI fail", err)
		stop()
	}

	g.Go(func() error {
		err := httpServer.Start()
		stop()
		return fmt.Errorf(pkgName+fnName+" - httpServer ends: %w", err)
	})

	g.Go(func() error {
		<-gCtx.Done()
		return fmt.Errorf(pkgName+fnName+" - httpServer shutdown: %w", httpServer.Stop())
	})
	// <<< Http Server

	// mux := http.NewServeMux()
	// mux.HandleFunc("GET /", func(response http.ResponseWriter, request *http.Request) {
	// 	// Приветствие первой страницы
	// 	logger.Debug(ctx, "Hello()", "PATH", request.URL.Path)
	// 	response.Write([]byte("Привет! " + request.URL.Path))
	// })

	// httpServer := &http.Server{Addr: cfg.HTTPAddress, Handler: mux}

	// g.Go(func() error {
	// 	return httpServer.ListenAndServe()
	// })

	// g.Go(func() error {
	// 	<-gCtx.Done()
	// 	timeoutCtx, timeoutCancel := context.WithTimeout(context.Background(), shutdownTimeoutSeconds*time.Second)
	// 	defer timeoutCancel()
	// 	return httpServer.Shutdown(timeoutCtx)
	// })

	logger.Info(ctx, "Application started")
	return g.Wait()
}
