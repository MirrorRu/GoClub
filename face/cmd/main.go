package main

import "goclub/v1/face/internal/face_app"

func main() {
	app := face_app.NewFaceApp()
	app.Run()
}
