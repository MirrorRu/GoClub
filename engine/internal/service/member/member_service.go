package memberservice

import (
	"context"
	"goclub/engine/internal/repository"
	"goclub/model/common"
	"goclub/model/members"
)

type MembersService interface {
	common.CRUDService[*members.Member, members.ID]
}

type memberService struct {
	repo repository.Repository
}

func NewMemberService(repo repository.Repository) MembersService {
	return &memberService{
		repo: repo,
	}
}

func (svc *memberService) Read(ctx context.Context, id members.ID) (result common.CRUDResult[*members.Member]) {
	return svc.repo.MemberRead(ctx, id)
}

func (svc *memberService) Create(ctx context.Context, member *members.Member) (result common.CRUDResult[*members.Member]) {
	return svc.repo.MemberCreate(ctx, member)
}

func (svc *memberService) Update(ctx context.Context, member *members.Member) (result common.CRUDResult[*members.Member]) {
	return svc.repo.MemberUpdate(ctx, member)
}

func (svc *memberService) Delete(ctx context.Context, id members.ID) (result common.CRUDResult[struct{}]) {
	return svc.repo.MemberDelete(ctx, id)
}

func (svc *memberService) Listing(ctx context.Context, filter any) (result common.CRUDResult[[]*members.Member]) {
	return svc.repo.MemberList(ctx, filter)
}
