package handler

import (
	"fmt"
	"log"
)

func New(attr InitAttribute) *HealthHandler {
	if err := attr.validate(); err != nil {
		log.Panic(err)
	}
	return &HealthHandler{
		service: attr.Service,
	}
}

func (attr InitAttribute) validate() error {
	if !attr.Service.validate() {
		return fmt.Errorf("missing health service : %+v", attr.Service.HealthService)
	}

	return nil
}

func (service ServiceAttribute) validate() bool {
	return service.HealthService != nil
}
