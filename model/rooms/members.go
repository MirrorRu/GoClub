package rooms

import (
	"goclub/model/common"
)

type (
	ID   common.ID
	Name common.Name

	Room struct {
		ID   ID   `json:"id" db:"groups=key;auto;base|title=Код"`
		Name Name `json:"name" db:"groups=base|title=Имя"`
	}
)
