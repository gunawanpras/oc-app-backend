package oracle

import (
	"log"

	"github.com/jmoiron/sqlx"
)

func (repo *ProfileRepository) prepareStatements() {
	repo.statement = StatementList{}
}

func (repo *ProfileRepository) prepareGetProfile() {
	var (
		err  error
		stmt *sqlx.Stmt
		db   = repo.db.Db
	)

	if stmt, err = db.Preparex(queryGetProfile); err != nil {
		log.Panic("[prepareGetProfile] error:", err)
	}
	repo.statement.GetProfile = stmt
}

func (repo *ProfileRepository) prepareGetProfileViaProcedure() {
	var (
		err  error
		stmt *sqlx.Stmt
		db   = repo.db.Db
	)

	if stmt, err = db.Preparex(queryGetProfileViaProcedure); err != nil {
		log.Panic("[prepareGetProfileViaProcedure] error:", err)
	}
	repo.statement.GetProfileViaProcedure = stmt
}

func (repo *ProfileRepository) prepareCreateProfile() {
	var (
		err  error
		stmt *sqlx.Stmt
		db   = repo.db.Db
	)

	if stmt, err = db.Preparex(queryCreateProfile); err != nil {
		log.Panic("[prepareCreateProfile] error:", err)
	}
	repo.statement.CreateProfile = stmt
}

func (repo *ProfileRepository) prepareUpdateProfile() {
	var (
		err  error
		stmt *sqlx.Stmt
		db   = repo.db.Db
	)

	if stmt, err = db.Preparex(queryUpdateProfile); err != nil {
		log.Panic("[prepareUpdateProfile] error:", err)
	}
	repo.statement.UpdateProfile = stmt
}

func (repo *ProfileRepository) prepareDeleteProfile() {
	var (
		err  error
		stmt *sqlx.Stmt
		db   = repo.db.Db
	)

	if stmt, err = db.Preparex(queryDeleteProfile); err != nil {
		log.Panic("[prepareDeleteProfile] error:", err)
	}
	repo.statement.DeleteProfile = stmt
}
