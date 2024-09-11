package clubapi

import (
	"context"
	"fmt"

	clubservice "goclub/gears/internal/club_service"
	"goclub/gears/internal/logger"
	"goclub/gears/internal/model/dto"

	gears "goclub/gears/pkg/api/v1"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

const pkgLogName = "gearsapi"

type API struct {
	gears.UnimplementedGoClubAPIServer
	clubService clubservice.ClubService
}

func NewAPI(
	clubService clubservice.ClubService,
) *API {
	return &API{
		clubService: clubService,
	}
}

func (a *API) RegisterGrpcServer(server *grpc.Server) {
	gears.RegisterGoClubAPIServer(server, a)
}

func (a *API) RegisterHttpHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	const fnLogName = ".RegisterHttpHandler"
	err := gears.RegisterGoClubAPIHandler(ctx, mux, conn)
	if err != nil {
		return fmt.Errorf(pkgLogName+fnLogName+" - failed to register gateway: %w", err)
	}
	return nil
}

func (a *API) Ping(ctx context.Context, req *emptypb.Empty) (*gears.PingResponse, error) {
	response := gears.PingResponse{
		Message: a.clubService.Ping(ctx),
	}

	return &response, nil
}

func (a *API) MemberCreate(ctx context.Context, req *gears.MemberCreateRequest) (resp *gears.MemberCreateResponse, err error) {
	const fnLogName = ".MemberCreate"
	opRes, err := a.clubService.MemberCreate(ctx, dto.Members.CreateReq(req))
	if err != nil {
		logger.Error(ctx, pkgLogName+fnLogName, err)
		return resp, err
	}
	return dto.Members.CreateResp(opRes), nil
}

func (a *API) MemberUpdate(ctx context.Context, req *gears.MemberUpdateRequest) (resp *gears.MemberUpdateResponse, err error) {
	const fnLogName = ".MemberUpdate"
	opRes, err := a.clubService.MemberUpdate(ctx, dto.Members.UpdateReq(req))
	if err != nil {
		logger.Error(ctx, pkgLogName+fnLogName, err)
		return resp, err
	}
	return dto.Members.UpdateResp(opRes), nil
}

func (a *API) MemberDelete(ctx context.Context, req *gears.MemberDeleteRequest) (resp *gears.MemberDeleteResponse, err error) {
	const fnLogName = ".MemberDelete"
	opRes, err := a.clubService.MemberDelete(ctx, dto.Members.DeleteReq(req))
	if err != nil {
		logger.Error(ctx, pkgLogName+fnLogName, err)
		return resp, err
	}
	return dto.Members.DeleteResp(opRes), nil
}

func (a *API) MemberRead(ctx context.Context, req *gears.MemberReadRequest) (resp *gears.MemberReadResponse, err error) {
	const fnLogName = ".MemberRead"
	opRes, err := a.clubService.MemberRead(ctx, dto.Members.ReadReq(req))
	if err != nil {
		logger.Error(ctx, pkgLogName+fnLogName, err)
		return resp, err
	}
	return dto.Members.ReadResp(opRes), nil
}

func (a *API) MemberListing(ctx context.Context, req *gears.MemberListingRequest) (resp *gears.MemberListingResponse, err error) {
	const fnLogName = ".MemberListing"
	opRes, err := a.clubService.MemberListing(ctx, dto.Members.ListingReq(req))
	if err != nil {
		logger.Error(ctx, pkgLogName+fnLogName, err)
		return resp, err
	}
	return dto.Members.ListingResp(opRes), nil
}
