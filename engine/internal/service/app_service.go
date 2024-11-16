package service

import (
	"goclub/engine/internal/repository"
	clubservice "goclub/engine/internal/service/club"
	memberservice "goclub/engine/internal/service/member"
	roomservice "goclub/engine/internal/service/room"
	"goclub/model/members"
	"goclub/model/rooms"
)

type ClubService interface {
	Do()
	Ping() string
}

type AppService interface {
	Club() ClubService
	Members() memberservice.MembersService
	Rooms() roomservice.RoomsService
}

type (
	appService struct {
		club    ClubService
		members memberservice.MembersService
		rooms   roomservice.RoomsService
	}
)

func NewAppService(repo repository.Repository) (appSvc *appService) {
	appSvc = &appService{
		club:    clubservice.NewClubService(repo),
		members: NewDictBaseService[members.Member](repo.Members()),
		rooms:   NewDictBaseService[rooms.Room](repo.Rooms()),
	}
	return appSvc
}

func (svc *appService) Club() ClubService {
	return svc.club
}

func (svc *appService) Members() memberservice.MembersService {
	return svc.members
}

func (svc *appService) Rooms() roomservice.RoomsService {
	return svc.rooms
}
