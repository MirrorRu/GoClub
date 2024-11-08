package grpcservice

import (
	"context"
	"fmt"

	"goclub/engine/grpc_api"
	"goclub/engine/internal/controller"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

const pkgName = "gearsapi"

type API struct {
	grpc_api.UnimplementedGoClubAPIServer
	controller controller.Controller
}

func NewAPI(
	controller controller.Controller,
) *API {
	return &API{
		controller: controller,
	}
}

func (a *API) RegisterGrpcServer(server *grpc.Server) {
	grpc_api.RegisterGoClubAPIServer(server, a)
}

func (a *API) RegisterHttpHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	const fnName = ".RegisterHttpHandler"
	err := grpc_api.RegisterGoClubAPIHandler(ctx, mux, conn)
	if err != nil {
		return fmt.Errorf(pkgName+fnName+" - failed to register gateway: %w", err)
	}
	return nil
}

func (a *API) Ping(ctx context.Context, req *emptypb.Empty) (*grpc_api.PingResponse, error) {
	response := grpc_api.PingResponse{
		Message: a.controller.Ping(ctx),
	}

	return &response, nil
}
