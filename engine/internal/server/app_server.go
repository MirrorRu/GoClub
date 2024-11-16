package server

import (
	grpcserver "goclub/engine/internal/grpc_server"
	httpserver "goclub/engine/internal/http_server"
	clubserver "goclub/engine/internal/server/club"
	memberserver "goclub/engine/internal/server/members"
	roomsserver "goclub/engine/internal/server/rooms"
	"goclub/engine/internal/service"
)

const pkgName = "api"

type APIServer interface {
	grpcserver.GRPCAPI
	httpserver.HTTPAPIs
}

type AppServer interface {
	GRPCAPIs() []grpcserver.GRPCAPI
	HTTPAPIs() []httpserver.HTTPAPIs
}

type appServer struct {
	clubServer   APIServer
	memberServer APIServer
	roomServer   APIServer
}

func NewAppServer(controller service.AppService) *appServer {
	return &appServer{
		clubServer:   clubserver.NewClubServer(controller),
		memberServer: memberserver.NewMemberServer(controller),
		roomServer:   roomsserver.NewRoomServer(controller),
	}
}

func (srv *appServer) GRPCAPIs() []grpcserver.GRPCAPI {
	return []grpcserver.GRPCAPI{
		srv.clubServer,
		srv.memberServer,
		srv.roomServer,
	}
}

func (srv *appServer) HTTPAPIs() []httpserver.HTTPAPIs {
	return []httpserver.HTTPAPIs{
		srv.clubServer,
		srv.memberServer,
		srv.roomServer,
	}
}
