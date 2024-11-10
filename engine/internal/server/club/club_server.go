package clubserver

import (
	"context"
	"fmt"

	"goclub/engine/internal/api"
	"goclub/engine/internal/service"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

const pkgName = "ClubServer"

type clubServer struct {
	api.UnimplementedClubServer
	application service.AppService
	club        service.ClubService
}

func NewClubServer(
	appSvc service.AppService,
) *clubServer {
	return &clubServer{
		application: appSvc,
		club:        appSvc.Club(),
	}
}

func (srv *clubServer) RegisterGrpcServer(server *grpc.Server) {
	api.RegisterClubServer(server, srv)
}

func (srv *clubServer) RegisterHttpServer(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	const fnName = ".RegisterHttpServer"
	if err := api.RegisterClubHandler(ctx, mux, conn); err != nil {
		return fmt.Errorf(pkgName+fnName+" - failed to call RegisterClubHandler: %w", err)
	}

	return nil
}

func (srv *clubServer) Ping(ctx context.Context, req *emptypb.Empty) (*api.PingResponse, error) {
	response := api.PingResponse{
		Message: srv.club.Ping(),
	}

	return &response, nil
}
