package tarifrepo

import (
	"context"
	"database/sql"
	"goclub/common/logger"
	"goclub/engine/internal/repository/db"
	"goclub/model/common"
	"goclub/model/tarifs"
	"slices"
	"strconv"
	"strings"
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

	allFieldIds    = collectIds(tableFields, func(fld common.TableField) (canAdd bool) { return true })
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

func NewTarifsRepo(writeHandler db.DbHandler, readHandler db.DbHandler) *tarifsRepo {
	return &tarifsRepo{
		writeHandler: writeHandler,
		readHandler:  readHandler,
	}
}

func (repo *tarifsRepo) QueryTextForInsert() string {
	fields := make([]string, len(nonKeyFieldIds))
	for idx := range nonKeyFieldIds {
		fields[idx] = "$" + strconv.Itoa(idx+1)
	}
	q := "insert into " +
		TableName +
		" (" +
		strings.Join(nonKeyFieldIds, ", ") +
		") values (" +
		strings.Join(fields, ", ") +
		") returning " +
		strings.Join(allFieldIds, ", ")
	return q
}

func (repo *tarifsRepo) QueryTextForDelete() string {
	keys := make([]string, len(keyFieldIds))
	for idx, id := range keyFieldIds {
		keys[idx] = id + "=$" + strconv.Itoa(idx+1)
	}
	q := "delete from " +
		TableName +
		" where " +
		strings.Join(keys, " and ")
	return q
}

func (repo *tarifsRepo) QueryTextForUpdate() string {
	fields := make([]string, len(nonKeyFieldIds))
	for idx, id := range nonKeyFieldIds {
		fields[idx] = id + "=$" + strconv.Itoa(idx+1)
	}
	keys := make([]string, len(keyFieldIds))
	for idx, id := range keyFieldIds {
		keys[idx] = id + "=$" + strconv.Itoa(idx+1)
	}
	q := "update " +
		TableName +
		" set " +
		strings.Join(fields, ", ") +
		" where " +
		strings.Join(keys, " and ") +
		" returning " +
		strings.Join(allFieldIds, ", ")
	return q
}

func (repo *tarifsRepo) QueryTextForRead() string {
	keys := make([]string, len(keyFieldIds))
	for idx, id := range keyFieldIds {
		keys[idx] = id + "=$" + strconv.Itoa(idx+1)
	}
	q := "select " +
		strings.Join(allFieldIds, ", ") +
		" from " +
		TableName +
		" where " +
		strings.Join(keys, " and ")
	return q
}

func (repo *tarifsRepo) QueryTextForList(filter any) string {
	q := "select " +
		strings.Join(allFieldIds, ", ") +
		" from " +
		TableName +
		" order by " +
		strings.Join(keyFieldIds, ", ")
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

func (repo *tarifsRepo) QueryArgsForInsert(data *tarifs.Tarif) (pointers []any) {
	pointers = make([]any, len(nonKeyFieldIds))
	for idx, id := range nonKeyFieldIds {
		pointers[idx] = repo.fieldAddress(id, data)
	}
	return pointers
}

func (repo *tarifsRepo) QueryArgsForUpdate(data *tarifs.Tarif) (fields []any) {
	updates := make([]any, len(nonKeyFieldIds))
	for idx, id := range nonKeyFieldIds {
		updates[idx] = repo.fieldAddress(id, data)
	}
	keys := make([]any, len(keyFieldIds))
	for idx, id := range keyFieldIds {
		keys[idx] = repo.fieldAddress(id, data)
	}

	return slices.Concat(updates, keys)
}

func (repo *tarifsRepo) QueryArgsForList(filter any) (fields []any) {
	return nil
}

func (repo *tarifsRepo) ScanArgsForAllFields(data *tarifs.Tarif) (pointers []any) {
	pointers = make([]any, len(allFieldIds))
	for idx, id := range allFieldIds {
		pointers[idx] = repo.fieldAddress(id, data)
	}
	return pointers
}

func (repo *tarifsRepo) Create(ctx context.Context, data *tarifs.Tarif) (result common.CRUDResult[*tarifs.Tarif]) {
	const fnName = "Create"
	query := repo.QueryTextForInsert()
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

func (repo *tarifsRepo) Delete(ctx context.Context, keys ...any) (result common.CRUDResult[struct{}]) {
	const fnName = "Delete"
	query := repo.QueryTextForDelete()
	logger.Debug(ctx, fnName, "query", query)

	var execRes sql.Result
	execRes, result.Error = repo.writeHandler.Exec(ctx, query, keys...)
	if result.Error != nil {
		return result
	}
	result.RecordsAffected, result.Error = execRes.RowsAffected()

	return result
}

// Update(ctx context.Context, data DataType) (result CRUDResult[DataType])
func (repo *tarifsRepo) Update(ctx context.Context, data *tarifs.Tarif) (result common.CRUDResult[*tarifs.Tarif]) {
	const fnName = "Update"
	query := repo.QueryTextForUpdate()
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
func (repo *tarifsRepo) Read(ctx context.Context, keys ...any) (result common.CRUDResult[*tarifs.Tarif]) {
	const fnName = "Read"
	data := new(tarifs.Tarif)
	query := repo.QueryTextForRead()
	logger.Debug(ctx, fnName, "query", query)

	var rows *sql.Rows
	rows, result.Error = repo.readHandler.Query(ctx, query, keys...)
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
func (repo *tarifsRepo) List(ctx context.Context, filter any) (result common.CRUDResult[[]*tarifs.Tarif]) {
	const fnName = "List"
	query := repo.QueryTextForList(filter)
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
