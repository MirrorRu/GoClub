package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"goclub/engine/internal/config"
	"goclub/engine/internal/repository/db"
	"goclub/engine/internal/repository/member_repo"
	"goclub/engine/internal/repository/room_repo"
	"goclub/engine/internal/repository/tarif_repo"
	"goclub/model/members"
	"goclub/model/rooms"

	_ "github.com/jackc/pgx/v5/stdlib"
)

const dbDriver = "pgx"

type pgDBRepo struct {
	memberrepo.MembersRepo
	roomrepo.RoomsRepo
	tarifrepo.TarifsRepo
	dbDSN  config.DBDSN
	writer db.DbHandler
	reader db.DbHandler
}

func NewPgDBRepo(dbdsn config.DBDSN) *pgDBRepo {
	return &pgDBRepo{
		dbDSN: dbdsn,
	}
}

func (repo *pgDBRepo) Open() (err error) {
	var dbMaster, dbSlave *sql.DB
	if dbMaster, err = sql.Open(dbDriver, repo.dbDSN.Master); err != nil {
		return fmt.Errorf("MasterDSN Open failed: %w", err)
	}
	if err = dbMaster.Ping(); err != nil {
		return fmt.Errorf("MasterDNS Ping failed: %w", err)
	}

	if dbSlave, err = sql.Open(dbDriver, repo.dbDSN.Slave); err != nil {
		return errors.Join(
			fmt.Errorf("SlaveDSN Open failed: %w", err),
			dbMaster.Close(),
		)
	}

	if err = dbSlave.Ping(); err != nil {
		return errors.Join(
			fmt.Errorf("SlaveDSN Open failed: %w", err),
			dbMaster.Close(),
			dbSlave.Close(),
		)
	}

	repo.writer = db.NewDbHandler(dbMaster)
	repo.reader = db.NewDbHandler(dbSlave)

	repo.MembersRepo = db.NewDictBaseDbRepo[members.Member](repo.writer, repo.reader)
	repo.RoomsRepo = db.NewDictBaseDbRepo[rooms.Room](repo.writer, repo.reader)
	repo.TarifsRepo = tarifrepo.NewTarifsRepo(repo.writer, repo.reader)
	return nil
}

func (repo *pgDBRepo) Close() (err error) {
	var errWriter, errReader error
	if repo.writer != nil {
		if errWriter = repo.writer.Close(); err == nil {
			repo.writer = nil
		}
	}

	if repo.reader != nil {
		if errReader = repo.reader.Close(); err == nil {
			repo.writer = nil
		}
	}

	return errors.Join(errWriter, errReader)

}

func (repo *pgDBRepo) Members() memberrepo.MembersRepo {
	return repo.MembersRepo
}

func (repo *pgDBRepo) Rooms() roomrepo.RoomsRepo {
	return repo.RoomsRepo
}

func (repo *pgDBRepo) Tarifs() tarifrepo.TarifsRepo {
	return repo.TarifsRepo
}
