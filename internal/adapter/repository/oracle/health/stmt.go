package oracle

import (
	"log"

	"github.com/jmoiron/sqlx"
)

func (repo *HealthRepository) prepareStatements() {
	repo.statement = StatementList{}
}

func (repo *HealthRepository) prepareCheckDatabaseHealth() {
	var (
		err  error
		stmt *sqlx.Stmt
		db   = repo.db.Db
	)

	if stmt, err = db.Preparex(queryCheckHealth); err != nil {
		log.Panic("[prepareCheckHealth] error:", err)
	}
	repo.statement.CheckHealth = stmt
}
