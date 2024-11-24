package dbext

import (
	"context"
	"database/sql"
	"goclub/common/logger"
	"goclub/model/common"
	"slices"
)

type (
	crudRepo[T comparable] struct {
		tableInfo    TableInfoProvider[T]
		writeHandler DbHandler
		readHandler  DbHandler
	}
)

// func (ti *TableInfoData[T]) GetFieldsAddr(data *T, fieldIds []int) (pointers []any) {
// 	panic("needs to override GetFieldsAddr")
// }

func NewCRUDRepo[T comparable](tableInfo TableInfoProvider[T], writeHandler DbHandler, readHandler DbHandler) *crudRepo[T] {
	return &crudRepo[T]{
		tableInfo:    tableInfo,
		writeHandler: writeHandler,
		readHandler:  readHandler,
	}
}

func (repo *crudRepo[T]) Create(ctx context.Context, data *T) (result common.CRUDResult[*T]) {
	const fnName = "Create"
	query := QueryTextForInsert(repo.tableInfo.Name(), repo.tableInfo.NonKeyFieldNames(), repo.tableInfo.AllFieldNames())
	logger.Debug(ctx, fnName, "query", query)

	queryArgs := repo.tableInfo.GetFieldsAddr(data, repo.tableInfo.NonKeyFieldIds())
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
			result.Data = new(T)
			scanArgs = repo.tableInfo.GetFieldsAddr(result.Data, repo.tableInfo.AllFieldIds())
			logger.Debug(ctx, fnName, "scanArgs", scanArgs)
		}
		result.RecordsAffected++
		rows.Scan(scanArgs...)
	}
	result.Error = rows.Err()

	logger.Debug(ctx, fnName, "data", data)
	return result
}

func (repo *crudRepo[T]) Delete(ctx context.Context, keys *T) (result common.CRUDResult[struct{}]) {
	const fnName = "Delete"
	query := QueryTextForDelete(repo.tableInfo.Name(), repo.tableInfo.KeyFieldNames())
	logger.Debug(ctx, fnName, "query", query)

	queryArgs := repo.tableInfo.GetFieldsAddr(keys, repo.tableInfo.KeyFieldIds())

	var execRes sql.Result
	execRes, result.Error = repo.writeHandler.Exec(ctx, query, queryArgs...)
	if result.Error != nil {
		return result
	}
	result.RecordsAffected, result.Error = execRes.RowsAffected()

	return result
}

// Update(ctx context.Context, data DataType) (result CRUDResult[DataType])
func (repo *crudRepo[T]) Update(ctx context.Context, data *T) (result common.CRUDResult[*T]) {
	const fnName = "Update"
	query := QueryTextForUpdate(repo.tableInfo.Name(), repo.tableInfo.NonKeyFieldNames(), repo.tableInfo.KeyFieldNames(), repo.tableInfo.AllFieldNames())
	logger.Debug(ctx, fnName, "query", query)

	queryArgs := repo.tableInfo.GetFieldsAddr(data, slices.Concat(repo.tableInfo.NonKeyFieldIds(), repo.tableInfo.KeyFieldIds()))
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
			result.Data = new(T)
			scanArgs = repo.tableInfo.GetFieldsAddr(result.Data, repo.tableInfo.AllFieldIds())
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
func (repo *crudRepo[T]) Read(ctx context.Context, keys *T) (result common.CRUDResult[*T]) {
	const fnName = "Read"
	data := new(T)
	query := QueryTextForRead(repo.tableInfo.Name(), repo.tableInfo.KeyFieldNames(), repo.tableInfo.AllFieldNames())
	logger.Debug(ctx, fnName, "query", query)

	queryArgs := repo.tableInfo.GetFieldsAddr(keys, repo.tableInfo.KeyFieldIds())
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
			scanArgs = repo.tableInfo.GetFieldsAddr(result.Data, repo.tableInfo.AllFieldIds())
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
func (repo *crudRepo[T]) List(ctx context.Context, filter common.ListFilter) (result common.CRUDResult[[]*T]) {
	const fnName = "List"
	query := QueryTextForList(repo.tableInfo.Name(), repo.tableInfo.KeyFieldNames(), repo.tableInfo.AllFieldNames(), filter)
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
	var data *T
	for rows.Next() {
		if result.Data == nil {
			result.Data = make([]*T, 0)
			data = new(T)
			scanArgs = repo.tableInfo.GetFieldsAddr(data, repo.tableInfo.AllFieldIds())
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
