package oracle

import (
	"fmt"
	"log"

	"github.com/gunawanpras/oc-app-backend/internal/core/profile/port"
)

func New(attr InitAttribute) port.Repository {
	if err := attr.validate(); err != nil {
		log.Panic(err)
	}

	repo := &ProfileRepository{
		db: attr.DB,
	}

	repo.prepareStatements()

	return repo
}

func (init InitAttribute) validate() error {
	if !init.DB.validate() {
		return fmt.Errorf("missing DB driver : %+v", init.DB)
	}

	return nil
}

func (db DB) validate() bool {
	return db.Db != nil
}
