package app

import (
	"context"
	"fmt"
	"goclub/common/logger"
	"goclub/engine/internal/config"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

const (
	pkgName                = "app."
	shutdownTimeoutSeconds = 3
)

type (
	app struct {
		//Providers, Adapters and Services
	}
)

func NewApp() app {
	return app{}
}

func (a *app) Run(ctx context.Context, cfg *config.AppConfig) error {
	const fnName = "Run"
	logger.Debug(ctx, fmt.Sprintf(pkgName+fnName+": cfg=%v", *cfg))
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", func(response http.ResponseWriter, request *http.Request) {
		// Приветствие первой страницы
		logger.Debug(ctx, "Hello()", "PATH", request.URL.Path)
		response.Write([]byte("Привет! " + request.URL.Path))
	})

	httpServer := &http.Server{Addr: cfg.HTTPAddress, Handler: mux}

	g, gCtx := errgroup.WithContext(ctx)
	g.Go(func() error {
		return httpServer.ListenAndServe()
	})

	g.Go(func() error {
		<-gCtx.Done()
		timeoutCtx, timeoutCancel := context.WithTimeout(context.Background(), shutdownTimeoutSeconds*time.Second)
		defer timeoutCancel()
		return httpServer.Shutdown(timeoutCtx)
	})
	logger.Info(ctx, "Application started")
	return g.Wait()
}
