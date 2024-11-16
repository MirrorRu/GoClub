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
	httpserver.HTTPAPI
}

type AppServer interface {
	GRPCAPIs() []grpcserver.GRPCAPI
	HTTPAPIs() []httpserver.HTTPAPI
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

func (srv *appServer) APIs() []APIServer {
	return []APIServer{
		srv.clubServer,
		srv.memberServer,
		srv.roomServer,
	}
}

func (srv *appServer) GRPCAPIs() (result []grpcserver.GRPCAPI) {
	apis := srv.APIs()
	result = make([]grpcserver.GRPCAPI, len(apis))
	for idx := range apis {
		result[idx] = apis[idx]
	}
	return result
}

func (srv *appServer) HTTPAPIs() (result []httpserver.HTTPAPI) {
	apis := srv.APIs()
	result = make([]httpserver.HTTPAPI, len(apis))
	for idx := range apis {
		result[idx] = apis[idx]
	}
	return result
}
