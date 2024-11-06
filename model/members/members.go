package members

import "goclub/model/common"

type (
	ID       common.ID
	Name     common.Name
	Birthday common.Option[common.Date]

	Member struct {
		ID       ID       `json:"id"`
		Name     Name     `json:"name"`
		Birthday Birthday `json:"birthday"`
	}
)
