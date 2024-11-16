package roomrepo

import (
	"goclub/model/common"
	"goclub/model/members"
)

type (
	MembersRepo interface {
		common.CRUDRepo[*members.Member]
	}
)
