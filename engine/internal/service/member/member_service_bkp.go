package memberservice

// import (
// 	"context"
// 	"goclub/engine/internal/repository"
// 	"goclub/model/common"
// 	"goclub/model/members"
// )

// type MembersService interface {
// 	common.CRUDService[*members.Member]
// }

// type memberService struct {
// 	repo repository.Repository
// }

// func NewMemberService(repo repository.Repository) MembersService {
// 	return &memberService{
// 		repo: repo,
// 	}
// }

// func (svc *memberService) Create(ctx context.Context, member *members.Member) (result common.CRUDResult[*members.Member]) {
// 	return svc.repo.Members().Create(ctx, member)
// }

// func (svc *memberService) Update(ctx context.Context, member *members.Member) (result common.CRUDResult[*members.Member]) {
// 	return svc.repo.Members().Update(ctx, member)
// }

// func (svc *memberService) Delete(ctx context.Context, keys ...any) (result common.CRUDResult[struct{}]) {
// 	return svc.repo.Members().Delete(ctx, keys...)
// }

// func (svc *memberService) Read(ctx context.Context, keys ...any) (result common.CRUDResult[*members.Member]) {
// 	return svc.repo.Members().Read(ctx, keys...)
// }

// func (svc *memberService) Listing(ctx context.Context, filter any) (result common.CRUDResult[[]*members.Member]) {
// 	return svc.repo.Members().List(ctx, filter)
// }
