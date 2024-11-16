package service

import (
	"goclub/engine/internal/repository"
	clubservice "goclub/engine/internal/service/club"
	memberservice "goclub/engine/internal/service/member"
)

type ClubService interface {
	Do()
	Ping() string
}

type AppService interface {
	Club() ClubService
	Members() memberservice.MembersService
}

type (
	appService struct {
		club    ClubService
		members memberservice.MembersService
	}
)

func NewAppService(repo repository.Repository) *appService {
	return &appService{
		club:    clubservice.NewClubService(repo),
		members: memberservice.NewMemberService(repo),
	}
}

func (svc *appService) Club() ClubService {
	return svc.club
}

func (svc *appService) Members() memberservice.MembersService {
	return svc.members
}
