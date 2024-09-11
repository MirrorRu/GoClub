package app

import (
	"context"
	"goclub/webapp/internal/config"
	errs "goclub/webapp/internal/error"
	"goclub/webapp/internal/handlers"
	"goclub/webapp/internal/storage"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

type App struct {
	storage storage.Storage
}

func NewApp() *App {
	return &App{}
}

func (app *App) Run(_cfg *config.AppConfig) errs.AppError {
	const opName = "Run"
	var (
		appErr errs.AppError
	)
	ctx, cancel := context.WithCancel(_cfg.Ctx)
	defer cancel()

	storage, appErr := storage.NewStorage(ctx, _cfg)
	if appErr != nil {
		return appErr
	}

	appErr = storage.Open()
	if appErr != nil {
		return appErr
	}
	defer storage.Close()

	go func() {
		c := make(chan os.Signal, 1) // we need to reserve to buffer size 1, so the notifier are not blocked
		signal.Notify(c, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

		<-c
		cancel()
	}()

	router := handlers.NewHTTPRouter()
	httpServer := &http.Server{Addr: _cfg.ServerAddress, Handler: router}

	g, gCtx := errgroup.WithContext(ctx)

	g.Go(func() error {
		const goOpName = "httpServer.ListenAndServe"
		err := httpServer.ListenAndServe()
		if err != nil {
			cancel()
			return errs.FromError(goOpName, errs.ECRoutineRun, err)
		}
		return nil
	})

	g.Go(func() error {
		const goOpName = "httpServer.Shutdown"
		<-gCtx.Done()
		if err := httpServer.Shutdown(context.Background()); err != nil {
			return errs.FromError(goOpName, errs.ECRoutineRun, err)
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		switch typedErr := err.(type) {
		case errs.AppError:
			return typedErr
		default:
			return errs.FromError(opName, errs.ECUnknown, err)
		}
	}

	return nil
}
