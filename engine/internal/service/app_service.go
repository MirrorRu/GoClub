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

type MembersService interface {
	Do()
}

type AppService interface {
	Club() ClubService
	Members() MembersService
}

type (
	appService struct {
		club    ClubService
		members MembersService
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

func (svc *appService) Members() MembersService {
	return svc.members
}
