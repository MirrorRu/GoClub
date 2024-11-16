package repository

import (
	membersrepo "goclub/engine/internal/repository/memberrepo"
	roomsrepo "goclub/engine/internal/repository/roomrepo"
)

const (
	pkgName = "repository."
)

type RepoOpenCloser interface {
	Open() error
	Close() error
}

type Repository interface {
	RepoOpenCloser
	Members() membersrepo.MembersRepo
	Rooms() roomsrepo.RoomsRepo
}
