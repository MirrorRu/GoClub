package repack

import (
	"goclub/engine/internal/api"
	"goclub/model/common"
)

type crud int

const CRUD crud = 0

func (crud) ToResp(src common.CRUDInfo) *api.CRUDResult {
	return &api.CRUDResult{
		RowsAffected: int64(src.RecordsAffected),
	}
}
