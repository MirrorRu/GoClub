package service

import (
	"context"
)

type ClubService interface {
	Ping(ctx context.Context) string
}

type MembersService interface {
}

type AppService interface {
	ClubService
	MembersService
}

type (
	appService struct{}
)

func NewService() AppService {
	return new(appService)
}

func (c *appService) Ping(context.Context) string {
	return "Pong"
}
