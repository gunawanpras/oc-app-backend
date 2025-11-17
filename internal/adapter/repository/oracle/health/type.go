package oracle

import (
	"github.com/jmoiron/sqlx"
)

type (
	HealthRepository struct {
		db        DB
		statement StatementList
	}

	DB struct {
		Db *sqlx.DB
	}

	StatementList struct {
		CheckHealth *sqlx.Stmt
	}

	InitAttribute struct {
		DB DB
	}
)
