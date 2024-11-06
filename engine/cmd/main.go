package main

import (
	"context"
	"fmt"
	"goclub/common/logger"
	"goclub/engine/internal/app"
	"goclub/engine/internal/config"
	"os"
)

func main() {
	cfg := config.Cfg
	ctx := context.Background()

	logger.Info(ctx, fmt.Sprintf("%s %s is starting", cfg.AppName, cfg.AppVersion))

	app := app.NewApp()
	if err := app.Run(ctx, cfg); err != nil {
		logger.Error(ctx, "Application finished with error", err)
		os.Exit(-1)
	}
	logger.Info(ctx, "Application is finished")
}
