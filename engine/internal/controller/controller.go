package controller

import (
	"context"
)

type Controller interface {
	Ping(ctx context.Context) string
	Do()
}

type (
	controller struct{}
)

func NewController() Controller {
	return new(controller)
}

func (c *controller) Do() {

}

func (c *controller) Ping(context.Context) string {
	return "Pong"
}
