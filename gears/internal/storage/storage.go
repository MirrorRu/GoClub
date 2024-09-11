package storage

import (
	"context"
	"goclub/gears/internal/config"
	"goclub/gears/internal/model"
)

type (
	MemberStorage interface {
		MemberCreate(ctx context.Context, req *model.MemberCreateReq) (resp *model.MemberCreateResp, err error)
		MemberUpdate(ctx context.Context, req *model.MemberUpdateReq) (resp *model.MemberUpdateResp, err error)
		MemberDelete(ctx context.Context, req *model.MemberDeleteReq) (resp *model.MemberDeleteResp, err error)
		MemberRead(ctx context.Context, req *model.MemberReadReq) (resp *model.MemberReadResp, err error)
		MemberListing(ctx context.Context, req *model.MemberListingReq) (resp *model.MemberListingResp, err error)
	}

	ClubStorage interface {
		Open(dsns config.DSNPair) error // Open storage
		Close() error                   //Close storage
		Migrate() error                 //Migration/updatge for storage structure
		MemberStorage
	}
)
