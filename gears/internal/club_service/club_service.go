package clubservice

import (
	"context"
	"goclub/gears/internal/logger"
	"goclub/gears/internal/model"
	"goclub/gears/internal/storage"
)

const pkgLogName = "clubservice"

type ClubService interface {
	Ping(ctx context.Context) (message string)
	MemberCreate(ctx context.Context, req *model.MemberCreateReq) (resp *model.MemberCreateResp, err error)
	MemberUpdate(ctx context.Context, req *model.MemberUpdateReq) (resp *model.MemberUpdateResp, err error)
	MemberDelete(ctx context.Context, req *model.MemberDeleteReq) (resp *model.MemberDeleteResp, err error)
	MemberRead(ctx context.Context, req *model.MemberReadReq) (resp *model.MemberReadResp, err error)
	MemberListing(ctx context.Context, req *model.MemberListingReq) (resp *model.MemberListingResp, err error)
}

type clubService struct {
	ctx     context.Context
	storage storage.ClubStorage
}

func NewClubService(ctx context.Context, storage storage.ClubStorage) ClubService {
	return &clubService{
		ctx:     ctx,
		storage: storage,
	}
}

func (*clubService) Ping(ctx context.Context) (message string) {
	const fnLogName = ".Ping"
	logger.Debug(ctx, pkgLogName+fnLogName)
	return "Pong"
}

func (svc *clubService) MemberCreate(ctx context.Context, req *model.MemberCreateReq) (resp *model.MemberCreateResp, err error) {
	return svc.storage.MemberCreate(ctx, req)
}

func (svc *clubService) MemberUpdate(ctx context.Context, req *model.MemberUpdateReq) (resp *model.MemberUpdateResp, err error) {
	return svc.storage.MemberUpdate(ctx, req)
}

func (svc *clubService) MemberDelete(ctx context.Context, req *model.MemberDeleteReq) (resp *model.MemberDeleteResp, err error) {
	return svc.storage.MemberDelete(ctx, req)
}

func (svc *clubService) MemberRead(ctx context.Context, req *model.MemberReadReq) (resp *model.MemberReadResp, err error) {
	return svc.storage.MemberRead(ctx, req)
}

func (svc *clubService) MemberListing(ctx context.Context, req *model.MemberListingReq) (resp *model.MemberListingResp, err error) {
	return svc.storage.MemberListing(ctx, req)
}
