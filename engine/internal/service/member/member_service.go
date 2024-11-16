package memberservice

import (
	"goclub/model/common"
	"goclub/model/members"
)

type MembersService interface {
	common.CRUDService[*members.Member]
}
