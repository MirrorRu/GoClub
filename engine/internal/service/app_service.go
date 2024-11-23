package service

import (
	"goclub/engine/internal/repository"
	clubsvc "goclub/engine/internal/service/club_svc"
	membersvc "goclub/engine/internal/service/member_svc"
	roomsvc "goclub/engine/internal/service/room_svc"
	tarifsvc "goclub/engine/internal/service/tarif_svc"
	"goclub/model/members"
	"goclub/model/rooms"
	"goclub/model/tarifs"
)

type ClubService interface {
	Do()
	Ping() string
}

type AppService interface {
	Club() ClubService
	Members() membersvc.MembersService
	Rooms() roomsvc.RoomsService
	Tarifs() tarifsvc.TarifsService
}

type (
	appService struct {
		club    ClubService
		members membersvc.MembersService
		rooms   roomsvc.RoomsService
		tarifs  tarifsvc.TarifsService
	}
)

func NewAppService(repo repository.Repository) (appSvc *appService) {
	appSvc = &appService{
		club:    clubsvc.NewClubService(repo),
		members: NewDictBaseService[members.Member](repo.Members()),
		rooms:   NewDictBaseService[rooms.Room](repo.Rooms()),
		tarifs:  NewDictBaseService[tarifs.Tarif](repo.Tarifs()),
	}
	return appSvc
}

func (svc *appService) Club() ClubService {
	return svc.club
}

func (svc *appService) Members() membersvc.MembersService {
	return svc.members
}

func (svc *appService) Rooms() roomsvc.RoomsService {
	return svc.rooms
}

func (svc *appService) Tarifs() tarifsvc.TarifsService {
	return svc.tarifs
}
