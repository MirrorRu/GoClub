package membersrv

import (
	"context"
	"fmt"
	"goclub/engine/internal/api"
	"goclub/engine/internal/repack"
	"goclub/engine/internal/service"
	membersvc "goclub/engine/internal/service/member_svc"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"google.golang.org/grpc"
)

const pkgName = "MemberServer"

type memberServer struct {
	api.UnimplementedMembersServer
	actor membersvc.MembersService
}

func NewMemberServer(
	controller service.AppService,
) *memberServer {
	return &memberServer{
		actor: controller.Members(),
	}
}

func (srv *memberServer) RegisterGrpcServer(server *grpc.Server) {
	api.RegisterMembersServer(server, srv)
}

func (srv *memberServer) RegisterHttpServer(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	const fnName = ".RegisterHttpServer"
	if err := api.RegisterMembersHandler(ctx, mux, conn); err != nil {
		return fmt.Errorf(pkgName+fnName+" - failed to call RegisterMemberHandler: %w", err)
	}
	return nil
}

func (srv *memberServer) MemberCreate(ctx context.Context, req *api.MemberCreateRequest) (resp *api.MemberCreateResponse, err error) {
	crudRes := srv.actor.Create(ctx, repack.Member.FromInfo(req.GetObject()))
	if crudRes.Error != nil {
		return nil, crudRes.Error
	}

	resp = &api.MemberCreateResponse{
		CrudResult: repack.CRUD.ToResp(crudRes.CRUDInfo),
		Object:     repack.Member.ToInfo(crudRes.Data),
	}
	return resp, err
}

func (srv *memberServer) MemberUpdate(ctx context.Context, req *api.MemberUpdateRequest) (resp *api.MemberUpdateResponse, err error) {
	crudRes := srv.actor.Update(ctx, repack.Member.FromInfo(req.GetObject()))
	if crudRes.Error != nil {
		return nil, crudRes.Error
	}

	resp = &api.MemberUpdateResponse{
		CrudResult: repack.CRUD.ToResp(crudRes.CRUDInfo),
		Object:     repack.Member.ToInfo(crudRes.Data),
	}
	return resp, err
}

func (srv *memberServer) MemberDelete(ctx context.Context, req *api.MemberDeleteRequest) (resp *api.MemberDeleteResponse, err error) {
	crudRes := srv.actor.Delete(ctx, req.GetId())
	if crudRes.Error != nil {
		return nil, crudRes.Error
	}

	resp = &api.MemberDeleteResponse{
		CrudResult: repack.CRUD.ToResp(crudRes.CRUDInfo),
	}
	return resp, err
}

func (srv *memberServer) MemberListing(ctx context.Context, req *api.MemberListingRequest) (resp *api.MemberListingResponse, err error) {
	crudRes := srv.actor.Listing(ctx, nil)
	if crudRes.Error != nil {
		return nil, crudRes.Error
	}

	resp = &api.MemberListingResponse{
		CrudResult: repack.CRUD.ToResp(crudRes.CRUDInfo),
		Objects:    repack.Member.ToInfos(crudRes.Data),
	}
	return resp, err
}

func (srv *memberServer) MemberRead(ctx context.Context, req *api.MemberReadRequest) (resp *api.MemberReadResponse, err error) {
	crudRes := srv.actor.Read(ctx, req.GetId())
	if crudRes.Error != nil {
		return nil, crudRes.Error
	}

	resp = &api.MemberReadResponse{
		CrudResult: repack.CRUD.ToResp(crudRes.CRUDInfo),
		Object:     repack.Member.ToInfo(crudRes.Data),
	}
	return resp, err
}

/*
func (srv *memberServer) MemberCreate(ctx context.Context, res *api.MemberCreateRequest) (*api.MemberCreateResponse, error) {
	obj := res.Object
	member := &members.Member{
		Name: members.Name(obj.GetFullName()),
	}

	createdMember, err := srv.actor.Create(ctx, member)
	resp := &api.MemberCreateResponse{
		MemberId: int64(id),
	}
	return resp, err
}

func (srv *memberServer) MemberDelete(context.Context, *api.MemberDeleteRequest) (*api.MemberDeleteResponse, error) {
	return nil, nil
}

func (srv *memberServer) MemberListing(context.Context, *api.MemberListingRequest) (*api.MemberListingResponse, error) {
	//srv.actor.
	return nil, nil
}

func (srv *memberServer) MemberRead(ctx context.Context, req *api.MemberReadRequest) (*api.MemberReadResponse, error) {
	member, err := srv.actor.Read(ctx, members.ID(req.GetMemberId()))
	if err != nil {
		return nil, err
	}

	return &api.MemberReadResponse{
		Object: &api.MemberInfo{MemberId: int64(member.ID), FullName: string(member.Name)},
	}, nil
}

func (srv *memberServer) MemberUpdate(context.Context, *api.MemberUpdateRequest) (*api.MemberUpdateResponse, error) {
	member, err := srv.actor.Update(ctx, members.ID(req.GetMemberId()))
	if err != nil {
		return nil, err
	}

	return &api.MemberReadResponse{
		Object: &api.MemberInfo{MemberId: int64(member.ID), FullName: string(member.Name)},
	}, nil
}
*/
