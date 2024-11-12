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
		ID       ID       `json:"id" db:"groups=key,base"`
		Name     Name     `json:"name" db:"groups=base"`
		Birthday Birthday `json:"birthday" db:"groups=base"`
		Notes    Notes    `json:"notes" db:"groups=aux"`
	}
)
