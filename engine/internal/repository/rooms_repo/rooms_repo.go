package roomsrepo

import (
	"goclub/model/common"
	"goclub/model/rooms"
)

type (
	RoomsRepo interface {
		common.CRUDRepo[*rooms.Room]
	}
)
