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
	q := "insert into " +
		TableName +
		" (" +
		strings.Join(nonKeyFieldIds, ", ") +
		") values (" +
		strings.Join(parmNums, ", ") +
		")"
	return q
}

func (repo *tarifsRepo) fieldAddress(fieldName string, data *tarifs.Tarif) (pointer any) {
	switch fieldName {
	case IDFieldKey:
		return &data.ID
	case NameFieldKey:
		return &data.Name
	case StartDateFieldKey:
		return &data.StartDate
	case EndDateFieldKey:
		return &data.EndDate
	}
	return nil
}

func (repo *tarifsRepo) QueryArgsForCreate(data *tarifs.Tarif) (pointers []any) {
	pointers = make(any, len(nonKeyFieldIds))
	for idx, id := range nonKeyFieldIds {
		pointers[idx] = repo.fieldAddress(id, data)
	}
	return pointers
}

func (repo *tarifsRepo) Create(ctx context.Context, member *tarifs.Tarif) (result common.CRUDResult[*tarifs.Tarif]) {
	const fnName = "Create"
	query := repo.QueryTextForInsert(member)
	logger.Debug(ctx, fnName, "query", query)

	queryArgs := QueryArgsForInsert(member)
	logger.Debug(ctx, fnName, "queryArgs", queryArgs)

	var rows *sql.Rows
	rows, result.Error = repo.writeHandleLoc.Query(ctx, query, queryArgs...)
	if result.Error != nil {
		return result
	}
	logger.Debug(ctx, fnName, "rows", rows)

	defer rows.Close()
	var scanArgs []any
	for rows.Next() {
		if result.Data == nil {
			result.Data = new(T)
			scanArgs = ScanArgsForAllFields(result.Data)
			logger.Debug(ctx, fnName, "scanArgs", scanArgs)
		}
		result.RecordsAffected++
		rows.Scan(scanArgs...)
	}
	result.Error = rows.Err()

	logger.Debug(ctx, fnName, "member", member)
	return result
}
