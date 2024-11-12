package db

import (
	"context"
	"database/sql"
	"goclub/common/errapp"
	"sync"
)

const pkgName = "db."

type DbHandler interface {
	Exec(ctx context.Context, query string, args ...any) (result sql.Result, err error)
	Close() (err error)
}

// DBHandler - оболочка для работы с SQL-запросами
type dbHandler struct {
	connPool *sql.DB
	stmtMap  sync.Map
}

func NewDbHandler(connPool *sql.DB) *dbHandler {
	return &dbHandler{
		connPool: connPool,
		stmtMap:  sync.Map{},
	}
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

func (dbh *dbHandler) Close() (err error) {
	return dbh.connPool.Close()
}
