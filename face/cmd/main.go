package main

import (
	"context"
	"goclub/v1/face/internal/face_app"
	"os"
	"os/signal"
)

func main() {
	app := face_app.NewFaceApp()
	runCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer stop()
	app.Run(runCtx)
}
