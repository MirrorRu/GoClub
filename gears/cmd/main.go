package main

import (
	"context"
	"goclub/gears/internal/app"
	"goclub/gears/internal/config"
	"goclub/gears/internal/logger"
	"goclub/gears/internal/tracker"
	"os"
)

func main() {
	cfg := config.Cfg
	ctx := context.Background()
	tracker.InitTracker(ctx, config.Cfg.AppTraceID, config.Cfg.JaegerURL)

	logger.Info(ctx, "GoClub start", "config", *cfg)

	app := app.NewApp(ctx, cfg)
	if err := app.Run(); err != nil {
		logger.Error(ctx, "Application run error", err)
		os.Exit(-1)
	}
}
