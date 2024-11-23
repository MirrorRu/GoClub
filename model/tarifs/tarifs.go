package tarifs

import "goclub/model/common"

type (
	ID   common.ID
	Name common.Name
	Date common.Date

	Tarif struct {
		ID        ID
		Name      Name
		StartDate Date
		EndDate   Date
	}
)
