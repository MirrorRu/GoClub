package repack

import (
	"goclub/engine/internal/api"
	"goclub/model/members"
)

type member int

const Member member = 0

func (member) FromInfo(info *api.MemberInfo) *members.Member {
	return &members.Member{
		ID:       members.ID(info.GetId()),
		FullName: members.FullName(info.GetFullName()),
		Birthday: members.Birthday(info.Birthday.GetYyyyMmDd()),
		Notes:    members.Notes(info.GetNotes()),
	}
}

func (member) ToInfo(member *members.Member) *api.MemberInfo {
	if member == nil {
		return nil
	}
	return &api.MemberInfo{
		Id:       int64(member.ID),
		FullName: string(member.FullName),
		Birthday: &api.Date{YyyyMmDd: int32(member.Birthday)},
		Notes:    string(member.Notes),
	}
}

func (member) ToInfos(members []*members.Member) []*api.MemberInfo {
	if members == nil {
		return nil
	}
	res := make([]*api.MemberInfo, len(members))
	for idx := range members {
		res[idx] = Member.ToInfo(members[idx])
	}
	return res
}
