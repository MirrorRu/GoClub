package service

import (
	"context"
	"fmt"

	"goclub/engine/internal/api"
	"goclub/engine/internal/controller"
	grpcserver "goclub/engine/internal/grpc_server"
	httpserver "goclub/engine/internal/http_server"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

const pkgName = "gearsapi"

type AppService interface {
	grpcserver.GRPCService
	httpserver.HTTPService
}

type API struct {
	api.UnimplementedGoClubAPIServer
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
	api.RegisterGoClubAPIServer(server, a)
}

func (a *API) RegisterHttpServer(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	const fnName = ".RegisterHttpServer"
	err := api.RegisterGoClubAPIHandler(ctx, mux, conn)
	if err != nil {
		return fmt.Errorf(pkgName+fnName+" - failed to register gateway: %w", err)
	}
	return nil
}

func (a *API) Ping(ctx context.Context, req *emptypb.Empty) (*api.PingResponse, error) {
	response := api.PingResponse{
		Message: a.controller.Ping(ctx),
	}

	return &response, nil
}
