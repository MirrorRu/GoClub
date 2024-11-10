package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"goclub/engine/internal/config"
	membersrepo "goclub/engine/internal/repository/members_repo"

	_ "github.com/jackc/pgx/v5/stdlib"
)

const dbDriver = "pgx"

type pgDBRepo struct {
	membersrepo.MembersRepo
	dbDSN  config.DBDSN
	writer *sql.DB
	reader *sql.DB
}

func NewPgDBRepo(dbdsn config.DBDSN) *pgDBRepo {
	return &pgDBRepo{
		dbDSN: dbdsn,
	}
}

func (repo *pgDBRepo) Open() (err error) {
	if repo.writer, err = sql.Open(dbDriver, repo.dbDSN.Master); err != nil {
		return fmt.Errorf("MasterDSN Open failed: %w", err)
	}

	if err = repo.writer.Ping(); err != nil {
		return fmt.Errorf("MasterDNS Ping failed: %w", err)
	}

	if repo.reader, err = sql.Open(dbDriver, repo.dbDSN.Slave); err != nil {
		return errors.Join(
			fmt.Errorf("SlaveDSN Open failed: %w", err),
			repo.writer.Close(),
		)
	}

	if err = repo.reader.Ping(); err != nil {
		return errors.Join(
			fmt.Errorf("SlaveDSN Open failed: %w", err),
			repo.writer.Close(),
		)
	}

	repo.MembersRepo = membersrepo.NewMembersDbRepo(repo.writer, repo.reader)
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
