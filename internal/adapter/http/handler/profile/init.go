package handler

import (
	"fmt"
	"log"
)

func New(attr InitAttribute) *ProfileHandler {
	if err := attr.validate(); err != nil {
		log.Panic(err)
	}
	return &ProfileHandler{
		service: attr.Service,
	}
}

func (attr InitAttribute) validate() error {
	if !attr.Service.validate() {
		return fmt.Errorf("missing profile service : %+v", attr.Service.ProfileService)
	}

	return nil
}

func (service ServiceAttribute) validate() bool {
	return service.ProfileService != nil
}
