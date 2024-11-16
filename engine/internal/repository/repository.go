package repository

import (
	membersrepo "goclub/engine/internal/repository/members_repo"
	roomsrepo "goclub/engine/internal/repository/rooms_repo"
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
