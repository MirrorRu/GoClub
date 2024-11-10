package membersrepo

import (
	"database/sql"
	"goclub/model/members"
)

type (
	membersDbRepo struct {
		writeHandle *sql.DB
		readHandle  *sql.DB
	}
)

func NewMembersDbRepo(writeHandle *sql.DB, readHandle *sql.DB) *membersDbRepo {
	return &membersDbRepo{
		writeHandle: writeHandle,
		readHandle:  readHandle,
	}
}

func (repo *membersDbRepo) MemberCreate(*members.Member) (result *CreateResp, err error) {
	return &CreateResp{}, nil
}

func (repo *membersDbRepo) MemberUpdate(*members.Member) (result *UpdateResp, err error) {
	return &UpdateResp{}, nil
}

func (repo *membersDbRepo) MemberDelete(*members.ID) (result *DeleteResp, err error) {
	return &DeleteResp{}, nil
}

func (repo *membersDbRepo) MemberRead(*members.ID) (result *members.Member, err error) {
	return &members.Member{}, nil
}

func (repo *membersDbRepo) MemberList(*ListReq) (result *ListResp, err error) {
	return &ListResp{}, nil
}
