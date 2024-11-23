package tarifrepo

import (
	"context"
	"database/sql"
	"goclub/common/logger"
	"goclub/engine/internal/repository/db"
	"goclub/engine/internal/repository/dbext"
	"goclub/model/common"
	"goclub/model/tarifs"
	"slices"
)

// CREATE TABLE dict.tarifs (
// 	id bigserial NOT NULL,
// 	"name" varchar(100) COLLATE "unicode" NOT NULL,
// 	start_date date NULL,
// 	end_date date NULL,
// 	CONSTRAINT tarifs_name_unique UNIQUE (name)
// );

const (
	TableName = "dict.tarifs"
)
const (
	IDFieldKey int = iota
	NameFieldKey
	StartDateFieldKey
	EndDateFieldKey
)

var (
	tableFields = map[int]common.TableField{
		IDFieldKey:        {Name: "id", Title: "Код тарифа", IsKey: true},
		NameFieldKey:      {Name: "name", Title: "Код тарифа"},
		StartDateFieldKey: {Name: "start_date", Title: "Начало действия"},
		EndDateFieldKey:   {Name: "end_date", Title: "Окончание действия"},
	}

	allFieldIds      = common.CollectIds(tableFields, func(fld common.TableField) (canAdd bool) { return true })
	allFieldNames    = common.CollectNames(tableFields, allFieldIds)
	keyFieldIds      = common.CollectIds(tableFields, func(fld common.TableField) (canAdd bool) { return fld.IsKey })
	keyFieldNames    = common.CollectNames(tableFields, keyFieldIds)
	nonKeyFieldIds   = common.CollectIds(tableFields, func(fld common.TableField) (canAdd bool) { return !fld.IsKey })
	nonKeyFieldNames = common.CollectNames(tableFields, nonKeyFieldIds)
)

type TarifsRepo interface {
	common.CRUDRepoExt[*tarifs.Tarif]
}

type (
	tarifsRepo struct {
		writeHandler db.DbHandler
		readHandler  db.DbHandler
	}
)

func NewTarifsRepo(writeHandler db.DbHandler, readHandler db.DbHandler) *tarifsRepo {
	return &tarifsRepo{
		writeHandler: writeHandler,
		readHandler:  readHandler,
	}
}

func (repo *tarifsRepo) GetFieldsAddres(data *tarifs.Tarif, fieldIds []int) (pointers []any) {
	pointers = make([]any, len(fieldIds))
	var a any
	for idx, fieldId := range fieldIds {
		switch fieldId {
		case IDFieldKey:
			a = &data.ID
		case NameFieldKey:
			a = &data.Name
		case StartDateFieldKey:
			a = &data.StartDate
		case EndDateFieldKey:
			a = &data.EndDate
		default:
			a = nil
		}
		pointers[idx] = a
	}
	return pointers
}

func (repo *tarifsRepo) Create(ctx context.Context, data *tarifs.Tarif) (result common.CRUDResult[*tarifs.Tarif]) {
	const fnName = "Create"
	query := dbext.QueryTextForInsert(TableName, nonKeyFieldNames, allFieldNames)
	logger.Debug(ctx, fnName, "query", query)

	queryArgs := repo.GetFieldsAddres(data, nonKeyFieldIds)
	logger.Debug(ctx, fnName, "queryArgs", queryArgs)

	var rows *sql.Rows
	rows, result.Error = repo.writeHandler.Query(ctx, query, queryArgs...)
	if result.Error != nil {
		return result
	}
	logger.Debug(ctx, fnName, "rows", rows)

	defer rows.Close()
	var scanArgs []any
	for rows.Next() {
		if result.Data == nil {
			result.Data = new(tarifs.Tarif)
			scanArgs = repo.GetFieldsAddres(result.Data, allFieldIds)
			logger.Debug(ctx, fnName, "scanArgs", scanArgs)
		}
		result.RecordsAffected++
		rows.Scan(scanArgs...)
	}
	result.Error = rows.Err()

	logger.Debug(ctx, fnName, "data", data)
	return result
}

func (repo *tarifsRepo) Delete(ctx context.Context, keys *tarifs.Tarif) (result common.CRUDResult[struct{}]) {
	const fnName = "Delete"
	query := dbext.QueryTextForDelete(TableName, keyFieldNames)
	logger.Debug(ctx, fnName, "query", query)

	queryArgs := repo.GetFieldsAddres(keys, keyFieldIds)

	var execRes sql.Result
	execRes, result.Error = repo.writeHandler.Exec(ctx, query, queryArgs...)
	if result.Error != nil {
		return result
	}
	result.RecordsAffected, result.Error = execRes.RowsAffected()

	return result
}

// Update(ctx context.Context, data DataType) (result CRUDResult[DataType])
func (repo *tarifsRepo) Update(ctx context.Context, data *tarifs.Tarif) (result common.CRUDResult[*tarifs.Tarif]) {
	const fnName = "Update"
	query := dbext.QueryTextForUpdate(TableName, nonKeyFieldNames, keyFieldNames, allFieldNames)
	logger.Debug(ctx, fnName, "query", query)

	queryArgs := repo.GetFieldsAddres(data, slices.Concat(nonKeyFieldIds, keyFieldIds))
	logger.Debug(ctx, fnName, "queryArgs", queryArgs)

	var rows *sql.Rows
	rows, result.Error = repo.writeHandler.Query(ctx, query, queryArgs...)
	if result.Error != nil {
		return result
	}
	logger.Debug(ctx, fnName, "rows", rows)

	defer rows.Close()
	var scanArgs []any
	for rows.Next() {
		if result.Data == nil {
			result.Data = new(tarifs.Tarif)
			scanArgs = repo.GetFieldsAddres(result.Data, allFieldIds)
			logger.Debug(ctx, fnName, "scanArgs", scanArgs)
		}
		result.RecordsAffected++
		rows.Scan(scanArgs...)
	}
	result.Error = rows.Err()

	logger.Debug(ctx, fnName, "data", data)
	return result
}

// Read(ctx context.Context, keys ...any) (result CRUDResult[DataType])
func (repo *tarifsRepo) Read(ctx context.Context, keys *tarifs.Tarif) (result common.CRUDResult[*tarifs.Tarif]) {
	const fnName = "Read"
	data := new(tarifs.Tarif)
	query := dbext.QueryTextForRead(TableName, keyFieldNames, allFieldNames)
	logger.Debug(ctx, fnName, "query", query)

	queryArgs := repo.GetFieldsAddres(keys, keyFieldIds)
	var rows *sql.Rows
	rows, result.Error = repo.readHandler.Query(ctx, query, queryArgs...)
	if result.Error != nil {
		return result
	}
	logger.Debug(ctx, fnName, "rows", rows)

	defer rows.Close()
	var scanArgs []any
	for rows.Next() {
		if result.Data == nil {
			result.Data = data
			scanArgs = repo.GetFieldsAddres(result.Data, allFieldIds)
			logger.Debug(ctx, fnName, "scanArgs", scanArgs)
		}
		result.RecordsAffected++
		rows.Scan(scanArgs...)
	}
	result.Error = rows.Err()

	logger.Debug(ctx, fnName, "data", data)
	return result
}

// List(ctx context.Context, filter any) (result CRUDResult[[]DataType])
func (repo *tarifsRepo) List(ctx context.Context, filter common.ListFilter) (result common.CRUDResult[[]*tarifs.Tarif]) {
	const fnName = "List"
	query := dbext.QueryTextForList(TableName, keyFieldNames, allFieldNames, filter)
	logger.Debug(ctx, fnName, "query", query)

	//queryArgs := repo.QueryArgsForList(filter)
	//logger.Debug(ctx, fnName, "queryArgs", queryArgs)

	var rows *sql.Rows
	rows, result.Error = repo.readHandler.Query(ctx, query)
	if result.Error != nil {
		return result
	}
	logger.Debug(ctx, fnName, "rows", rows)

	defer rows.Close()
	var scanArgs []any
	var data *tarifs.Tarif
	for rows.Next() {
		if result.Data == nil {
			result.Data = make([]*tarifs.Tarif, 0)
			data = new(tarifs.Tarif)
			scanArgs = repo.GetFieldsAddres(data, allFieldIds)
			logger.Debug(ctx, fnName, "scanArgs", scanArgs)
		}
		result.RecordsAffected++
		rows.Scan(scanArgs...)
		dataCopy := *data
		result.Data = append(result.Data, &dataCopy) // cloning
	}
	result.Error = rows.Err()

	logger.Debug(ctx, fnName, "data", data)
	return result
}
