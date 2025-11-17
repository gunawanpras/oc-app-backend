package oracle

import (
	"github.com/jmoiron/sqlx"
)

type (
	ProfileRepository struct {
		db        DB
		statement StatementList
	}

	DB struct {
		Db *sqlx.DB
	}

	StatementList struct {
		GetProfile             *sqlx.Stmt
		CreateProfile          *sqlx.Stmt
		UpdateProfile          *sqlx.Stmt
		DeleteProfile          *sqlx.Stmt
		GetProfileViaProcedure *sqlx.Stmt
	}

	InitAttribute struct {
		DB DB
	}
)
