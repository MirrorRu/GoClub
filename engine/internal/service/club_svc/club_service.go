package clubsvc

import "goclub/engine/internal/repository"

type (
	clubService struct{}
)

func NewClubService(repo repository.Repository) *clubService {
	return new(clubService)
}

func (svc *clubService) Do() {

}

func (svc *clubService) Ping() string {
	return "Pong"
}
