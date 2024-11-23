package repack

import (
	"goclub/engine/internal/api"
	"goclub/model/rooms"
)

type room int

const Room room = 0

func (room) FromInfo(info *api.RoomInfo) *rooms.Room {
	return &rooms.Room{
		ID:   rooms.ID(info.GetId()),
		Name: rooms.Name(info.GetName()),
	}
}

func (room) ToInfo(room *rooms.Room) *api.RoomInfo {
	if room == nil {
		return nil
	}
	return &api.RoomInfo{
		Id:   int64(room.ID),
		Name: string(room.Name),
	}
}

func (room) ToInfos(rooms []*rooms.Room) []*api.RoomInfo {
	if rooms == nil {
		return nil
	}
	res := make([]*api.RoomInfo, len(rooms))
	for idx := range rooms {
		res[idx] = Room.ToInfo(rooms[idx])
	}
	return res
}
