package db

import (
	"context"
	"database/sql"
	"goclub/common/logger"
	"goclub/model/common"
)

type (
	dictBaseDbRepo[T comparable] struct {
		writeHandleLoc DbHandler
		readHandleLoc  DbHandler
	}
)

func NewDictBaseDbRepo[T comparable](writeHandle DbHandler, readHandle DbHandler) *dictBaseDbRepo[T] {
	return &dictBaseDbRepo[T]{
		writeHandleLoc: writeHandle,
		readHandleLoc:  readHandle,
	}
}

func (repo *dictBaseDbRepo[T]) Create(ctx context.Context, member *T) (result common.CRUDResult[*T]) {
	const fnName = "Create"
	query := QueryTextForInsert(member)
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

func (repo *dictBaseDbRepo[T]) Update(ctx context.Context, member *T) (result common.CRUDResult[*T]) {
	const fnName = "Update"
	query := QueryTextForUpdate(member)
	logger.Debug(ctx, fnName, "query", query)

	queryArgs := QueryArgsForUpdate(member)
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

func (repo *dictBaseDbRepo[T]) Read(ctx context.Context, keys ...any) (result common.CRUDResult[*T]) {
	const fnName = "Read"
	member := new(T)
	query := QueryTextForRead(member)
	logger.Debug(ctx, fnName, "query", query)

	// queryArgs := QueryArgsForRead(member)
	// logger.Debug(ctx, fnName, "queryArgs", queryArgs)

	var rows *sql.Rows
	rows, result.Error = repo.readHandleLoc.Query(ctx, query, keys...)
	if result.Error != nil {
		return result
	}
	logger.Debug(ctx, fnName, "rows", rows)

	defer rows.Close()
	var scanArgs []any
	for rows.Next() {
		if result.Data == nil {
			result.Data = member
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

func (repo *dictBaseDbRepo[T]) List(ctx context.Context, filter any) (result common.CRUDResult[[]*T]) {
	const fnName = "List"
	member := new(T)
	query := QueryTextForList(member, filter)
	logger.Debug(ctx, fnName, "query", query)

	queryArgs := QueryArgsForList(member, filter)
	logger.Debug(ctx, fnName, "queryArgs", queryArgs)

	var rows *sql.Rows
	rows, result.Error = repo.readHandleLoc.Query(ctx, query, queryArgs...)
	if result.Error != nil {
		return result
	}
	logger.Debug(ctx, fnName, "rows", rows)

	defer rows.Close()
	var scanArgs []any
	for rows.Next() {
		if result.Data == nil {
			result.Data = make([]*T, 0)
			scanArgs = ScanArgsForAllFields(member)
			logger.Debug(ctx, fnName, "scanArgs", scanArgs)
		}
		result.RecordsAffected++
		rows.Scan(scanArgs...)
		memberCopy := *member
		result.Data = append(result.Data, &memberCopy) // cloning
	}
	result.Error = rows.Err()

	logger.Debug(ctx, fnName, "member", member)
	return result
}

func (repo *dictBaseDbRepo[T]) Delete(ctx context.Context, keys ...any) (result common.CRUDResult[struct{}]) {
	const fnName = "Delete"
	member := new(T)
	query := QueryTextForDelete(member)
	logger.Debug(ctx, fnName, "query", query)

	// queryArgs := QueryArgsForDelete(member)
	// logger.Debug(ctx, fnName, "queryArgs", queryArgs)

	var execRes sql.Result
	execRes, result.Error = repo.writeHandleLoc.Exec(ctx, query, keys...)
	if result.Error != nil {
		return result
	}
	result.RecordsAffected, result.Error = execRes.RowsAffected()

	return result
}
