package face_app

import (
	"context"
	"goclub/v1/common/logger"
	"goclub/v1/face/config"
	"net/http"
)

type (
	FaceApp interface {
		Run(ctx context.Context)
	}

	faceApp struct {
	}
)

var cfg = &config.FaceConfig{
	RunAddress:      ":8080",
	StaticFilesPath: "./jquery-easyui/",
	StaticURLPrefix: "/jquery-easyui/",
}

func NewFaceApp() FaceApp {
	return &faceApp{}
}

func (app *faceApp) Run(ctx context.Context) {
	logger.Info(ctx, "FaceApp Run")

	mux := http.NewServeMux()
	mux.Handle(http.MethodGet+" /ping", http.HandlerFunc(app.ping))

	// Инициализируем FileServer, он будет обрабатывать
	// HTTP-запросы к статическим файлам из папки "./ui/static".
	// Обратите внимание, что переданный в функцию http.Dir путь
	// является относительным корневой папке проекта
	fileServer := http.FileServer(http.Dir(cfg.StaticFilesPath))

	// Используем функцию mux.Handle() для регистрации обработчика для
	// всех запросов, которые начинаются с "/static/". Мы убираем
	// префикс "/static" перед тем как запрос достигнет http.FileServer
	mux.Handle(http.MethodGet+" "+cfg.StaticURLPrefix, http.StripPrefix(cfg.StaticURLPrefix, fileServer))

	err := http.ListenAndServe(cfg.RunAddress, mux)
	logger.Error(ctx, "FaceApp finished", err)
}

func (app *faceApp) ping(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte("pong"))
}
