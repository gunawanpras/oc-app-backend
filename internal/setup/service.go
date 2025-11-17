package setup

import (
	"github.com/gunawanpras/oc-app-backend/config"
	healthPort "github.com/gunawanpras/oc-app-backend/internal/core/health/port"
	healthService "github.com/gunawanpras/oc-app-backend/internal/core/health/service"
	profilePort "github.com/gunawanpras/oc-app-backend/internal/core/profile/port"
	profileService "github.com/gunawanpras/oc-app-backend/internal/core/profile/service"
)

type Service struct {
	HealthService  healthPort.Service
	ProfileService profilePort.Service
}

func NewService(conf *config.Config, repo Repository) Service {
	return Service{
		HealthService: healthService.New(healthService.InitAttribute{
			Repo: healthService.RepoAttribute{
				HealthRepo: repo.HealthRepo,
			},
			Config: healthService.ConfigAttribute{
				Config: conf,
			},
		}),
		ProfileService: profileService.New(profileService.InitAttribute{
			Repo: profileService.RepoAttribute{
				ProfileRepo: repo.ProfileRepo,
			},
			Config: profileService.ConfigAttribute{
				Config: conf,
			},
		}),
	}
}
