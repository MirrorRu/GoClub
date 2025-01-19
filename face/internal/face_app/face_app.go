package face_app

import "fmt"

type (
	FaceApp interface {
		Run()
	}

	faceApp struct {
	}
)

func NewFaceApp() FaceApp {
	return &faceApp{}
}

func (app *faceApp) Run() {
	fmt.Println("Hello World")
}
