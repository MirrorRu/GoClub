package server

import (
	grpcserver "goclub/engine/internal/grpc_server"
	httpserver "goclub/engine/internal/http_server"
	clubsrv "goclub/engine/internal/server/club_srv"
	membersrv "goclub/engine/internal/server/member_srv"
	roomsrv "goclub/engine/internal/server/room_srv"
	tarifsrv "goclub/engine/internal/server/tarif_srv"
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
	tarifServer  APIServer
}

func NewAppServer(controller service.AppService) *appServer {
	return &appServer{
		clubServer:   clubsrv.NewClubServer(controller),
		memberServer: membersrv.NewMemberServer(controller),
		roomServer:   roomsrv.NewRoomServer(controller),
		tarifServer:  tarifsrv.NewTarifServer(controller),
	}
}

func (srv *appServer) APIs() []APIServer {
	return []APIServer{
		srv.clubServer,
		srv.memberServer,
		srv.roomServer,
		srv.tarifServer,
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
