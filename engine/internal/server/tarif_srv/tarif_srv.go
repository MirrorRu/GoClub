package tarifsrv

import (
	"context"
	"fmt"
	"goclub/engine/internal/api"
	"goclub/engine/internal/repack"
	"goclub/engine/internal/service"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"google.golang.org/grpc"
)

const pkgName = "TarifServer"

type tarifsServer struct {
	api.UnimplementedTarifsServer
	actor service.AppService
}

func NewTarifServer(
	controller service.AppService,
) *tarifsServer {
	return &tarifsServer{
		actor: controller,
	}
}

func (srv *tarifsServer) RegisterGrpcServer(server *grpc.Server) {
	api.RegisterTarifsServer(server, srv)
}

func (srv *tarifsServer) RegisterHttpServer(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	const fnName = ".RegisterHttpServer"
	if err := api.RegisterTarifsHandler(ctx, mux, conn); err != nil {
		return fmt.Errorf(pkgName+fnName+" - failed to call RegisterTarifHandler: %w", err)
	}
	return nil
}

func (srv *tarifsServer) TarifCreate(ctx context.Context, req *api.TarifCreateRequest) (resp *api.TarifCreateResponse, err error) {
	crudRes := srv.actor.Tarifs().Create(ctx, repack.Tarif.FromInfo(req.GetObject()))
	if crudRes.Error != nil {
		return nil, crudRes.Error
	}

	resp = &api.TarifCreateResponse{
		CrudResult: repack.CRUD.ToResp(crudRes.CRUDInfo),
		Object:     repack.Tarif.ToInfo(crudRes.Data),
	}
	return resp, err
}

func (srv *tarifsServer) TarifUpdate(ctx context.Context, req *api.TarifUpdateRequest) (resp *api.TarifUpdateResponse, err error) {
	crudRes := srv.actor.Tarifs().Update(ctx, repack.Tarif.FromInfo(req.GetObject()))
	if crudRes.Error != nil {
		return nil, crudRes.Error
	}

	resp = &api.TarifUpdateResponse{
		CrudResult: repack.CRUD.ToResp(crudRes.CRUDInfo),
		Object:     repack.Tarif.ToInfo(crudRes.Data),
	}
	return resp, err
}

func (srv *tarifsServer) TarifDelete(ctx context.Context, req *api.TarifDeleteRequest) (resp *api.TarifDeleteResponse, err error) {
	crudRes := srv.actor.Tarifs().Delete(ctx, req.GetId())
	if crudRes.Error != nil {
		return nil, crudRes.Error
	}

	resp = &api.TarifDeleteResponse{
		CrudResult: repack.CRUD.ToResp(crudRes.CRUDInfo),
	}
	return resp, err
}

func (srv *tarifsServer) TarifListing(ctx context.Context, req *api.TarifListingRequest) (resp *api.TarifListingResponse, err error) {
	crudRes := srv.actor.Tarifs().Listing(ctx, nil)
	if crudRes.Error != nil {
		return nil, crudRes.Error
	}

	resp = &api.TarifListingResponse{
		CrudResult: repack.CRUD.ToResp(crudRes.CRUDInfo),
		Objects:    repack.Tarif.ToInfos(crudRes.Data),
	}
	return resp, err
}

func (srv *tarifsServer) TarifRead(ctx context.Context, req *api.TarifReadRequest) (resp *api.TarifReadResponse, err error) {
	crudRes := srv.actor.Tarifs().Read(ctx, req.GetId())
	if crudRes.Error != nil {
		return nil, crudRes.Error
	}

	resp = &api.TarifReadResponse{
		CrudResult: repack.CRUD.ToResp(crudRes.CRUDInfo),
		Object:     repack.Tarif.ToInfo(crudRes.Data),
	}
	return resp, err
}
