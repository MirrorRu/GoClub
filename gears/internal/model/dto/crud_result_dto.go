package dto

import (
	"goclub/gears/internal/model"
	gears "goclub/gears/pkg/api/v1"
)

type crudResultDTO struct{}

var CRUDResult crudResultDTO

func (crudResultDTO) Resp(src *model.CRUDResult) *gears.CRUDResult {
	return &gears.CRUDResult{
		RowsAffected: src.RowsAffected,
	}
}
