package service

import (
	"fmt"
	"log"
)

func New(attr InitAttribute) *ProfileService {
	if err := attr.validate(); err != nil {
		log.Panic(err)
	}

	return &ProfileService{
		repo:   attr.Repo,
		config: attr.Config,
	}
}

func (attr InitAttribute) validate() error {
	if !attr.Repo.validate() {
		return fmt.Errorf("missing profile repo : %+v", attr.Repo.ProfileRepo)
	}

	return nil
}

func (repo RepoAttribute) validate() bool {
	return repo.ProfileRepo != nil
}
