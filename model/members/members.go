package members

import (
	"goclub/model/common"
)

type (
	ID       common.ID
	Name     common.Name
	Notes    common.Notes
	Birthday common.Option[common.Date]

	Member struct {
		ID       ID       `json:"id" db:"group=key"`
		Name     Name     `json:"name" db:"group=base"`
		Birthday Birthday `json:"birthday" db:"group=base"`
		Notes    Notes    `json:"notes" db:"group=aux"`
	}
)
