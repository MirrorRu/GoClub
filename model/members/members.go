package members

import (
	"goclub/model/common"
)

type (
	ID       common.ID
	FullName common.Name
	Notes    common.Notes
	Birthday common.Date //common.Option[common.Date]

	Member struct {
		ID       ID       `json:"id" db:"groups=key;auto;base|title=Код"`
		FullName FullName `json:"name" db:"groups=base|title=Имя|name=name"`
		Birthday Birthday `json:"birthday" db:"groups=base|title=День рождения"`
		Notes    Notes    `json:"notes" db:"groups=aux|title=Примечание"`
	}
)
