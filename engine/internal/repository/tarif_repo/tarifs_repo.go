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
	TableName         = "dict.tarifs"
	IDFieldKey        = "id"
	NameFieldKey      = "name"
	StartDateFieldKey = "start_date"
	EndDateFieldKey   = "end_date"
)

var (
	tableFields = map[string]common.TableField{
		IDFieldKey:        {Name: IDFieldKey, Title: "Код тарифа", IsKey: true},
		NameFieldKey:      {Name: NameFieldKey, Title: "Код тарифа"},
		StartDateFieldKey: {Name: StartDateFieldKey, Title: "Начало действия"},
		EndDateFieldKey:   {Name: EndDateFieldKey, Title: "Окончание действия"},
	}

	allFieldIds    = common.CollectIds(tableFields, func(fld common.TableField) (canAdd bool) { return true })
	keyFieldIds    = common.CollectIds(tableFields, func(fld common.TableField) (canAdd bool) { return fld.IsKey })
	nonKeyFieldIds = common.CollectIds(tableFields, func(fld common.TableField) (canAdd bool) { return !fld.IsKey })
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

func (repo *tarifsRepo) GetFieldAddress(fieldName string, data *tarifs.Tarif) (pointer any) {
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

func (repo *tarifsRepo) QueryArgsForInsert(data *tarifs.Tarif) (pointers []any) {
	pointers = make([]any, len(nonKeyFieldIds))
	for idx, id := range nonKeyFieldIds {
		pointers[idx] = repo.GetFieldAddress(id, data)
	}
	return pointers
}

func (repo *tarifsRepo) QueryArgsForUpdate(data *tarifs.Tarif) (fields []any) {
	updates := make([]any, len(nonKeyFieldIds))
	for idx, id := range nonKeyFieldIds {
		updates[idx] = repo.GetFieldAddress(id, data)
	}
	keys := make([]any, len(keyFieldIds))
	for idx, id := range keyFieldIds {
		keys[idx] = repo.GetFieldAddress(id, data)
	}

	return slices.Concat(updates, keys)
}

func (repo *tarifsRepo) QueryArgsForKey(data *tarifs.Tarif) (fields []any) {
	keys := make([]any, len(keyFieldIds))
	for idx, id := range keyFieldIds {
		keys[idx] = repo.GetFieldAddress(id, data)
	}
	return keys
}

func (repo *tarifsRepo) QueryArgsForList(filter any) (fields []any) {
	return nil
}

func (repo *tarifsRepo) ScanArgsForAllFields(data *tarifs.Tarif) (pointers []any) {
	pointers = make([]any, len(allFieldIds))
	for idx, id := range allFieldIds {
		pointers[idx] = repo.GetFieldAddress(id, data)
	}
	return pointers
}

func (repo *tarifsRepo) Create(ctx context.Context, data *tarifs.Tarif) (result common.CRUDResult[*tarifs.Tarif]) {
	const fnName = "Create"
	query := dbext.QueryTextForInsert(TableName, nonKeyFieldIds, allFieldIds)
	logger.Debug(ctx, fnName, "query", query)

	queryArgs := repo.QueryArgsForInsert(data)
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
			scanArgs = repo.ScanArgsForAllFields(result.Data)
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
	query := dbext.QueryTextForDelete(TableName, keyFieldIds)
	logger.Debug(ctx, fnName, "query", query)

	queryArgs := repo.QueryArgsForKey(keys)

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
	query := dbext.QueryTextForUpdate(TableName, nonKeyFieldIds, keyFieldIds, allFieldIds)
	logger.Debug(ctx, fnName, "query", query)

	queryArgs := repo.QueryArgsForUpdate(data)
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
			scanArgs = repo.ScanArgsForAllFields(result.Data)
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
	query := dbext.QueryTextForRead(TableName, keyFieldIds, allFieldIds)
	logger.Debug(ctx, fnName, "query", query)

	queryArgs := repo.QueryArgsForKey(keys)
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
			scanArgs = repo.ScanArgsForAllFields(result.Data)
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
	query := dbext.QueryTextForList(TableName, keyFieldIds, allFieldIds, filter)
	logger.Debug(ctx, fnName, "query", query)

	queryArgs := repo.QueryArgsForList(filter)
	logger.Debug(ctx, fnName, "queryArgs", queryArgs)

	var rows *sql.Rows
	rows, result.Error = repo.readHandler.Query(ctx, query, queryArgs...)
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
			scanArgs = repo.ScanArgsForAllFields(data)
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
