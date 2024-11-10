package membersrepo

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
		MemberCreate(*members.Member) (result *CreateResp, err error)
		MemberUpdate(*members.Member) (result *UpdateResp, err error)
		MemberDelete(*members.ID) (result *DeleteResp, err error)
		MemberRead(*members.ID) (result *members.Member, err error)
		MemberList(*ListReq) (result *ListResp, err error)
	}
)
