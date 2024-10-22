package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("GET /ping", http.HandlerFunc(ping))

	// Инициализируем FileServer, он будет обрабатывать
	// HTTP-запросы к статическим файлам из папки "./ui/static".
	// Обратите внимание, что переданный в функцию http.Dir путь
	// является относительным корневой папке проекта
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Используем функцию mux.Handle() для регистрации обработчика для
	// всех запросов, которые начинаются с "/static/". Мы убираем
	// префикс "/static" перед тем как запрос достигнет http.FileServer
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	err := http.ListenAndServe(":8000", mux)
	log.Fatal(err)
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}
