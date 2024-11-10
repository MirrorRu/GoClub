package clubserver

import (
	"context"
	"fmt"

	"goclub/engine/internal/api"
	"goclub/engine/internal/service"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

const pkgName = "ClubServer"

type clubServer struct {
	api.UnimplementedClubServer
	controller service.AppService
}

func NewClubServer(
	controller service.AppService,
) *clubServer {
	return &clubServer{
		controller: controller,
	}
}

func (a *clubServer) RegisterGrpcServer(server *grpc.Server) {
	api.RegisterClubServer(server, a)
}

func (a *clubServer) RegisterHttpServer(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	const fnName = ".RegisterHttpServer"
	if err := api.RegisterClubHandler(ctx, mux, conn); err != nil {
		return fmt.Errorf(pkgName+fnName+" - failed to call RegisterClubHandler: %w", err)
	}
	return nil
}

func (a *clubServer) Ping(ctx context.Context, req *emptypb.Empty) (*api.PingResponse, error) {
	response := api.PingResponse{
		Message: a.controller.Ping(ctx),
	}

	return &response, nil
}
