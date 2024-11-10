package memberservice

import "goclub/engine/internal/repository"

type memberService struct{}

func NewMemberService(repo repository.Repository) *memberService {
	return new(memberService)
}

func (svc *memberService) Do() {

}
