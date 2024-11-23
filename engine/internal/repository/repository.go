package repository

import (
	"goclub/engine/internal/repository/member_repo"
	"goclub/engine/internal/repository/room_repo"
	"goclub/engine/internal/repository/tarif_repo"
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
	Members() memberrepo.MembersRepo
	Rooms() roomrepo.RoomsRepo
	Tarifs() tarifrepo.TarifsRepo
}
