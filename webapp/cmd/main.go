package main

import (
	"goclub/webapp/internal/app"
	"goclub/webapp/internal/config"
	"goclub/webapp/internal/logs"
	"os"
)

const runErrorExitCode = -1

func main() {
	app := app.NewApp()
	appErr := app.Run(config.Cfg)
	if appErr != nil {
		logs.AppError(appErr)
		os.Exit(int(appErr.Code()))
	}
}
