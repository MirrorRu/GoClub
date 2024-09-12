package activity

import (
	"goclub/internal/model"
	"sync"
)

type ActivityInMemStorage struct {
	locker  sync.Mutex
	lastID  model.ActivityID
	members map[model.ActivityID]model.Activity
}
