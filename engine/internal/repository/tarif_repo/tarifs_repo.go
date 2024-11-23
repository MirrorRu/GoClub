package tarifrepo

import (
	"context"
	"goclub/engine/internal/repository/db"
	"goclub/model/common"
	"goclub/model/tarifs"
	"strconv"
	"strings"
)

const (
	TableName         = "tarifs"
	IDFieldKey        = "id"
	NameFieldKey      = "name"
	StartDateFieldKey = "start_date"
	EndDateFieldKey   = "end_date"
)

func collectIds(src map[string]common.TableField, filter func(fld common.TableField) (canAdd bool)) []string {
	result := make([]string, 0, len(src))
	for k, v := range src {
		if filter(v) {
			result = append(result, k)
		}
	}
	return result
}

var (
	tableFields = map[string]common.TableField{
		IDFieldKey:        {Name: IDFieldKey, Title: "Код тарифа", IsKey: true},
		NameFieldKey:      {Name: NameFieldKey, Title: "Код тарифа"},
		StartDateFieldKey: {Name: StartDateFieldKey, Title: "Начало действия"},
		EndDateFieldKey:   {Name: EndDateFieldKey, Title: "Окончание действия"},
	}

	keyFieldIds    = collectIds(tableFields, func(fld common.TableField) (canAdd bool) { return fld.IsKey })
	nonKeyFieldIds = collectIds(tableFields, func(fld common.TableField) (canAdd bool) { return !fld.IsKey })
)

type TarifsRepo interface {
	common.CRUDRepo[*tarifs.Tarif]
}

type (
	tarifsRepo struct {
		writeHandler db.DbHandler
		readHandler  db.DbHandler
	}
)

func NewTarifsRepo(writeHandler db.DbHandler, readHandler db.DbHandler) TarifsRepo {
	return &tarifsRepo{
		writeHandler: writeHandler,
		readHandler:  readHandler,
	}
}

func (repo *tarifsRepo) QueryTextForCreate(_ *tarifs.Tarif) string {
	parmNums := make([]string, len(nonKeyFieldIds))
	for idx := range nonKeyFieldIds {
		parmNums[idx] = "$" + strconv.Itoa(idx+1)
	}
	q := "insert into " + TableName + " (" + strings.Join(nonKeyFieldIds, ", ") + ") values (" + strings.Join(parmNums, ", ") + ")"
	return q
}

func (repo *tarifsRepo) Create(ctx context.Context, data *tarifs.Tarif) (result common.CRUDResult[*tarifs.Tarif]) {

}
