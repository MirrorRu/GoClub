package dto

import (
	"goclub/gears/internal/model"
	gears "goclub/gears/pkg/api/v1"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type memberDTO struct{}

var Members memberDTO

func (memberDTO) Resp(src *model.MemberInfo) (result *gears.MemberInfo) {
	result = &gears.MemberInfo{
		MemberId: int64(src.ID),
		FullName: string(src.FullName),
		Birthday: &timestamppb.Timestamp{Seconds: time.Time(src.Birthday).Unix()},
	}
	return result
}

func (_ memberDTO) CreateReq(src *gears.MemberCreateRequest) (result *model.MemberCreateReq) {
	result = &model.MemberCreateReq{
		MemberInfo: model.MemberInfo{
			FullName: model.MemberFullName(src.GetObject().GetFullName()),
			Birthday: model.MemberBirthday(src.GetObject().GetBirthday().AsTime()),
		},
	}
	return result
}

func (_ memberDTO) CreateResp(src *model.MemberCreateResp) (result *gears.MemberCreateResponse) {
	result = &gears.MemberCreateResponse{
		CrudResult: CRUDResult.Resp(&src.CRUDResult),
		MemberId:   int64(src.MemberID),
	}
	return result
}

func (_ memberDTO) UpdateReq(src *gears.MemberUpdateRequest) (result *model.MemberUpdateReq) {
	object := src.GetObject()
	result = &model.MemberUpdateReq{
		MemberInfo: model.MemberInfo{
			ID:       model.MemberID(object.GetMemberId()),
			FullName: model.MemberFullName(object.GetFullName()),
			Birthday: model.MemberBirthday(object.GetBirthday().AsTime()),
		},
	}
	return result
}

func (_ memberDTO) UpdateResp(src *model.MemberUpdateResp) (result *gears.MemberUpdateResponse) {
	result = &gears.MemberUpdateResponse{
		CrudResult: CRUDResult.Resp(&src.CRUDResult),
	}
	return result
}

func (_ memberDTO) DeleteReq(src *gears.MemberDeleteRequest) (result *model.MemberDeleteReq) {
	result = &model.MemberDeleteReq{
		MemberID: model.MemberID(src.GetMemberId()),
	}
	return result
}

func (_ memberDTO) DeleteResp(src *model.MemberDeleteResp) (result *gears.MemberDeleteResponse) {
	result = &gears.MemberDeleteResponse{
		CrudResult: CRUDResult.Resp(&src.CRUDResult),
	}
	return result
}

func (_ memberDTO) ReadReq(src *gears.MemberReadRequest) (result *model.MemberReadReq) {
	result = &model.MemberReadReq{
		MemberID: model.MemberID(src.GetMemberId()),
	}
	return result
}

func (dto memberDTO) ReadResp(src *model.MemberReadResp) (result *gears.MemberReadResponse) {
	result = &gears.MemberReadResponse{
		CrudResult: CRUDResult.Resp(&src.CRUDResult),
		Object:     dto.Resp(&src.MemberInfo),
	}
	return result
}

func (_ memberDTO) ListingReq(src *gears.MemberListingRequest) (result *model.MemberListingReq) {
	result = &model.MemberListingReq{}
	return result
}

func (dto memberDTO) ListingResp(src *model.MemberListingResp) (result *gears.MemberListingResponse) {
	result = &gears.MemberListingResponse{
		CrudResult: CRUDResult.Resp(&src.CRUDResult),
		Objects:    make([]*gears.MemberInfo, len(src.Items)),
	}
	for i := 0; i < len(src.Items); i++ {
		result.Objects[i] = dto.Resp(&src.Items[i])
	}
	return result
}
