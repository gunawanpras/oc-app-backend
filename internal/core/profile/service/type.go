package service

import (
	"github.com/gunawanpras/oc-app-backend/config"
	"github.com/gunawanpras/oc-app-backend/internal/core/profile/port"
)

type (
	RepoAttribute struct {
		ProfileRepo port.Repository
	}

	ConfigAttribute struct {
		Config *config.Config
	}

	ProfileService struct {
		repo   RepoAttribute
		config ConfigAttribute
	}

	InitAttribute struct {
		Repo   RepoAttribute
		Config ConfigAttribute
	}
)
