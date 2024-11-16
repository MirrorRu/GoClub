package db

import (
	"context"
	"database/sql"
	"goclub/common/errapp"
	"sync"
)

const pkgName = "db."

type DbHandler interface {
	Connection() *sql.DB
	Exec(ctx context.Context, query string, args ...any) (result sql.Result, err error)
	QueryRow(ctx context.Context, query string, args ...any) (result *sql.Row, err error)
	Query(ctx context.Context, query string, args ...any) (result *sql.Rows, err error)
	Close() (err error)
}

// DBHandler - оболочка для работы с SQL-запросами
type dbHandler struct {
	connPool *sql.DB
	stmtMap  sync.Map
}

func NewDbHandler(connPool *sql.DB) DbHandler {
	return &dbHandler{
		connPool: connPool,
		stmtMap:  sync.Map{},
	}
}

func (dbh *dbHandler) Connection() *sql.DB {
	return dbh.connPool
}

func (dbh *dbHandler) shtm(ctx context.Context, query string) (result *sql.Stmt, err error) {
	const fnName = "shmt"
	x, ok := dbh.stmtMap.Load(query)
	if ok {
		return x.(*sql.Stmt), nil
	}

	if result, err = dbh.connPool.PrepareContext(ctx, query); err == nil {
		dbh.stmtMap.Store(query, result)
	}

	return result, errapp.Wrap(pkgName+fnName, err)
}

func (dbh *dbHandler) Exec(ctx context.Context, query string, args ...any) (result sql.Result, err error) {
	const fnName = "Exec"
	stmt, err := dbh.shtm(ctx, query)
	if err == nil {
		result, err = stmt.ExecContext(ctx, args...)
	}
	return result, errapp.Wrap(pkgName+fnName, err)
}

func (dbh *dbHandler) QueryRow(ctx context.Context, query string, args ...any) (result *sql.Row, err error) {
	const fnName = "QueryRow"
	stmt, err := dbh.shtm(ctx, query)
	if err == nil {
		result = stmt.QueryRowContext(ctx, args...)
		err = result.Err()
	}
	return result, errapp.Wrap(pkgName+fnName, err)

}

func (dbh *dbHandler) Query(ctx context.Context, query string, args ...any) (result *sql.Rows, err error) {
	const fnName = "Query"
	stmt, err := dbh.shtm(ctx, query)
	if err != nil {
		return nil, errapp.Wrap(pkgName+fnName, err)
	}
	result, err = stmt.QueryContext(ctx, args...)
	return result, err
}

func (dbh *dbHandler) Close() (err error) {
	return dbh.connPool.Close()
}
