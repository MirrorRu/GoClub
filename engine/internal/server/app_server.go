package server

import (
	grpcserver "goclub/engine/internal/grpc_server"
	httpserver "goclub/engine/internal/http_server"
	"goclub/engine/internal/service"
)

const pkgName = "api"

type APIServer interface {
	grpcserver.GRPCRegistrar
	httpserver.HTTPRegistrar
}

type AppServer interface {
	GRPCServers() []grpcserver.GRPCRegistrar
}

type appServer struct {
	clubServer   APIServer
	memberServer APIServer
}

func NewAppServer(controller service.AppService) *appServer {
	return &appServer{}
}

func (srv *appServer) GRPCServers() []grpcserver.GRPCRegistrar {
	return []grpcserver.GRPCRegistrar{
		srv.clubServer,
		srv.memberServer,
	}
}
