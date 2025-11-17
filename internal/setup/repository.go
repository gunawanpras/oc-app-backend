package setup

import (
	healthRepoOc "github.com/gunawanpras/oc-app-backend/internal/adapter/repository/oracle/health"
	profileRepoOc "github.com/gunawanpras/oc-app-backend/internal/adapter/repository/oracle/profile"
	healthRepo "github.com/gunawanpras/oc-app-backend/internal/core/health/port"
	profileRepo "github.com/gunawanpras/oc-app-backend/internal/core/profile/port"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	HealthRepo  healthRepo.Repository
	ProfileRepo profileRepo.Repository
}

func NewRepository(db *sqlx.DB) Repository {
	healthRepo := healthRepoOc.New(healthRepoOc.InitAttribute{
		DB: healthRepoOc.DB{
			Db: db,
		},
	})

	profileRepo := profileRepoOc.New(profileRepoOc.InitAttribute{
		DB: profileRepoOc.DB{
			Db: db,
		},
	})

	return Repository{
		HealthRepo:  healthRepo,
		ProfileRepo: profileRepo,
	}
}
