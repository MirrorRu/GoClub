package roomsrv

import (
	"context"
	"fmt"
	"goclub/engine/internal/api"
	"goclub/engine/internal/service"
	"goclub/engine/repack"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"google.golang.org/grpc"
)

const pkgName = "RoomServer"

type roomsServer struct {
	api.UnimplementedRoomsServer
	actor service.AppService
}

func NewRoomServer(
	controller service.AppService,
) *roomsServer {
	return &roomsServer{
		actor: controller,
	}
}

func (srv *roomsServer) RegisterGrpcServer(server *grpc.Server) {
	api.RegisterRoomsServer(server, srv)
}

func (srv *roomsServer) RegisterHttpServer(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	const fnName = ".RegisterHttpServer"
	if err := api.RegisterRoomsHandler(ctx, mux, conn); err != nil {
		return fmt.Errorf(pkgName+fnName+" - failed to call RegisterRoomHandler: %w", err)
	}
	return nil
}

func (srv *roomsServer) RoomCreate(ctx context.Context, req *api.RoomCreateRequest) (resp *api.RoomCreateResponse, err error) {
	crudRes := srv.actor.Rooms().Create(ctx, repack.Room.FromInfo(req.GetObject()))
	if crudRes.Error != nil {
		return nil, crudRes.Error
	}

	resp = &api.RoomCreateResponse{
		CrudResult: repack.CRUD.ToResp(crudRes.CRUDInfo),
		Object:     repack.Room.ToInfo(crudRes.Data),
	}
	return resp, err
}

func (srv *roomsServer) RoomUpdate(ctx context.Context, req *api.RoomUpdateRequest) (resp *api.RoomUpdateResponse, err error) {
	crudRes := srv.actor.Rooms().Update(ctx, repack.Room.FromInfo(req.GetObject()))
	if crudRes.Error != nil {
		return nil, crudRes.Error
	}

	resp = &api.RoomUpdateResponse{
		CrudResult: repack.CRUD.ToResp(crudRes.CRUDInfo),
		Object:     repack.Room.ToInfo(crudRes.Data),
	}
	return resp, err
}

func (srv *roomsServer) RoomDelete(ctx context.Context, req *api.RoomDeleteRequest) (resp *api.RoomDeleteResponse, err error) {
	crudRes := srv.actor.Rooms().Delete(ctx, req.GetId())
	if crudRes.Error != nil {
		return nil, crudRes.Error
	}

	resp = &api.RoomDeleteResponse{
		CrudResult: repack.CRUD.ToResp(crudRes.CRUDInfo),
	}
	return resp, err
}

func (srv *roomsServer) RoomListing(ctx context.Context, req *api.RoomListingRequest) (resp *api.RoomListingResponse, err error) {
	crudRes := srv.actor.Rooms().Listing(ctx, nil)
	if crudRes.Error != nil {
		return nil, crudRes.Error
	}

	resp = &api.RoomListingResponse{
		CrudResult: repack.CRUD.ToResp(crudRes.CRUDInfo),
		Objects:    repack.Room.ToInfos(crudRes.Data),
	}
	return resp, err
}

func (srv *roomsServer) RoomRead(ctx context.Context, req *api.RoomReadRequest) (resp *api.RoomReadResponse, err error) {
	crudRes := srv.actor.Rooms().Read(ctx, req.GetId())
	if crudRes.Error != nil {
		return nil, crudRes.Error
	}

	resp = &api.RoomReadResponse{
		CrudResult: repack.CRUD.ToResp(crudRes.CRUDInfo),
		Object:     repack.Room.ToInfo(crudRes.Data),
	}
	return resp, err
}
