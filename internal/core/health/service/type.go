package service

import (
	"github.com/gunawanpras/oc-app-backend/config"
	"github.com/gunawanpras/oc-app-backend/internal/core/health/port"
)

type (
	RepoAttribute struct {
		HealthRepo port.Repository
	}

	ConfigAttribute struct {
		Config *config.Config
	}

	HealthService struct {
		repo   RepoAttribute
		config ConfigAttribute
	}

	InitAttribute struct {
		Repo   RepoAttribute
		Config ConfigAttribute
	}
)
