package service

import (
	"fmt"
	"log"
)

func New(attr InitAttribute) *HealthService {
	if err := attr.validate(); err != nil {
		log.Panic(err)
	}

	return &HealthService{
		repo:   attr.Repo,
		config: attr.Config,
	}
}

func (attr InitAttribute) validate() error {
	if !attr.Repo.validate() {
		return fmt.Errorf("missing health repo : %+v", attr.Repo.HealthRepo)
	}

	return nil
}

func (repo RepoAttribute) validate() bool {
	return repo.HealthRepo != nil
}
