package roomservice

import (
	"goclub/model/common"
	"goclub/model/rooms"
)

type RoomsService interface {
	common.CRUDService[*rooms.Room]
}
