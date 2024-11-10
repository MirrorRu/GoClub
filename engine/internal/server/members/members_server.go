package memberserver

import (
	"context"
	"fmt"
	"goclub/engine/internal/api"
	"goclub/engine/internal/service"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"google.golang.org/grpc"
)

const pkgName = "MemberServer"

type memberServer struct {
	api.UnimplementedMembersServer
	actor service.AppService
}

func NewMemberServer(
	controller service.AppService,
) *memberServer {
	return &memberServer{
		actor: controller,
	}
}

func (srv *memberServer) RegisterGrpcServer(server *grpc.Server) {
	api.RegisterMembersServer(server, srv)
}

func (srv *memberServer) RegisterHttpServer(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	const fnName = ".RegisterHttpServer"
	if err := api.RegisterMembersHandler(ctx, mux, conn); err != nil {
		return fmt.Errorf(pkgName+fnName+" - failed to call RegisterMemberHandler: %w", err)
	}
	return nil
}

func (srv *memberServer) MemberCreate(context.Context, *api.MemberCreateRequest) (*api.MemberCreateResponse, error) {
	return nil, nil
}

func (srv *memberServer) MemberDelete(context.Context, *api.MemberDeleteRequest) (*api.MemberDeleteResponse, error) {
	return nil, nil
}

func (srv *memberServer) MemberListing(context.Context, *api.MemberListingRequest) (*api.MemberListingResponse, error) {
	return nil, nil
}

func (srv *memberServer) MemberRead(context.Context, *api.MemberReadRequest) (*api.MemberReadResponse, error) {
	return nil, nil
}

func (srv *memberServer) MemberUpdate(context.Context, *api.MemberUpdateRequest) (*api.MemberUpdateResponse, error) {
	return nil, nil
}
