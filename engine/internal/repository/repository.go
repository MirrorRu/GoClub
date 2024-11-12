package repository

import (
	membersrepo "goclub/engine/internal/repository/members_repo"
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
	membersrepo.MembersRepo
}
