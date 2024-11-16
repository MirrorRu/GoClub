package membersrepo

import (
	"context"
	"database/sql"
	"goclub/common/logger"
	"goclub/engine/internal/repository/db"
	"goclub/model/common"
	"goclub/model/members"
)

type (
	membersDbRepo struct {
		writeHandleLoc db.DbHandler
		readHandleLoc  db.DbHandler
	}
)

func NewMembersDbRepo(writeHandle db.DbHandler, readHandle db.DbHandler) *membersDbRepo {
	return &membersDbRepo{
		writeHandleLoc: writeHandle,
		readHandleLoc:  readHandle,
	}
}

func (repo *membersDbRepo) MemberCreate(ctx context.Context, member *members.Member) (result common.CRUDResult[*members.Member]) {
	const fnName = "MemberCreate"
	query := db.QueryTextForInsert(member)
	logger.Debug(ctx, fnName, "query", query)

	queryArgs := db.QueryArgsForInsert(member)
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
			result.Data = new(members.Member)
			scanArgs = db.ScanArgsForAllFields(result.Data)
			logger.Debug(ctx, fnName, "scanArgs", scanArgs)
		}
		result.RecordsAffected++
		rows.Scan(scanArgs...)
	}
	result.Error = rows.Err()

	logger.Debug(ctx, fnName, "member", member)
	return result
}

func (repo *membersDbRepo) MemberUpdate(ctx context.Context, member *members.Member) (result common.CRUDResult[*members.Member]) {
	const fnName = "MemberUpdate"
	query := db.QueryTextForUpdate(member)
	logger.Debug(ctx, fnName, "query", query)

	queryArgs := db.QueryArgsForUpdate(member)
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
			result.Data = new(members.Member)
			scanArgs = db.ScanArgsForAllFields(result.Data)
			logger.Debug(ctx, fnName, "scanArgs", scanArgs)
		}
		result.RecordsAffected++
		rows.Scan(scanArgs...)
	}
	result.Error = rows.Err()

	logger.Debug(ctx, fnName, "member", member)
	return result
}

func (repo *membersDbRepo) MemberRead(ctx context.Context, id members.ID) (result common.CRUDResult[*members.Member]) {
	const fnName = "MemberRead"
	member := &members.Member{
		ID: id,
	}
	query := db.QueryTextForRead(member)
	logger.Debug(ctx, fnName, "query", query)

	queryArgs := db.QueryArgsForRead(member)
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
			result.Data = member
			scanArgs = db.ScanArgsForAllFields(result.Data)
			logger.Debug(ctx, fnName, "scanArgs", scanArgs)
		}
		result.RecordsAffected++
		rows.Scan(scanArgs...)
	}
	result.Error = rows.Err()

	logger.Debug(ctx, fnName, "member", member)
	return result
}

func (repo *membersDbRepo) MemberList(ctx context.Context, filter any) (result common.CRUDResult[[]*members.Member]) {
	const fnName = "MemberList"
	member := new(members.Member)
	query := db.QueryTextForList(member, filter)
	logger.Debug(ctx, fnName, "query", query)

	queryArgs := db.QueryArgsForList(member, filter)
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
			result.Data = make([]*members.Member, 0)
			scanArgs = db.ScanArgsForAllFields(member)
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

func (repo *membersDbRepo) MemberDelete(ctx context.Context, id members.ID) (result common.CRUDResult[struct{}]) {
	const fnName = "MemberDelete"
	member := &members.Member{ID: id}
	query := db.QueryTextForDelete(member)
	logger.Debug(ctx, fnName, "query", query)

	queryArgs := db.QueryArgsForDelete(member)
	logger.Debug(ctx, fnName, "queryArgs", queryArgs)

	var execRes sql.Result
	execRes, result.Error = repo.writeHandleLoc.Exec(ctx, query, queryArgs...)
	if result.Error != nil {
		return result
	}
	result.RecordsAffected, result.Error = execRes.RowsAffected()

	return result
}
