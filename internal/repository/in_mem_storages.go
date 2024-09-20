package repository

import (
	"goclub/internal/model"
)

type ActivityInmemStore struct {
	TableMemoryStore[model.ActivityID, *model.Activity]
}

type ClubmanInmemStore struct {
	TableMemoryStore[model.ClubmanID, *model.Clubman]
}
