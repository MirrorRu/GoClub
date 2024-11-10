package repository

import (
	membersrepo "goclub/engine/internal/repository/members_repo"
)

type RepoOpenCloser interface {
	Open() error
	Close() error
}
type Repository interface {
	RepoOpenCloser
	membersrepo.MembersRepo
}
