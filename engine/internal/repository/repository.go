package repository

import (
	"goclub/model/members"
)

type (
	CreateResp struct{}
	UpdateResp struct{}
	DeleteResp struct{}
	ListReq    struct{}
	ListResp   struct{}

	MembersRepo interface {
		Create(*members.Member) (result *CreateResp, err error)
		Update(*members.Member) (result *UpdateResp, err error)
		Delete(*members.ID) (result *DeleteResp, err error)
		Read(*members.ID) (result *members.Member, err error)
		List(*ListReq) (result *ListResp, err error)
	}

	membersRepo struct{}
)

func NewMembersRepo() MembersRepo {
	return &membersRepo{}
}

func (repo *membersRepo) Create(*members.Member) (result *CreateResp, err error) {
	return &CreateResp{}, nil
}

func (repo *membersRepo) Update(*members.Member) (result *UpdateResp, err error) {
	return &UpdateResp{}, nil
}

func (repo *membersRepo) Delete(*members.ID) (result *DeleteResp, err error) {
	return &DeleteResp{}, nil
}

func (repo *membersRepo) Read(*members.ID) (result *members.Member, err error) {
	return &members.Member{}, nil
}

func (repo *membersRepo) List(*ListReq) (result *ListResp, err error) {
	return &ListResp{}, nil
}
