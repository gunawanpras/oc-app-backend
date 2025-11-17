package setup

import (
	"github.com/gunawanpras/oc-app-backend/config"
	setupClient "github.com/gunawanpras/oc-app-backend/internal/setup/client"
	"github.com/jmoiron/sqlx"
)

type ExternalServices struct {
	Oracle *sqlx.DB
}

type CoreServices struct {
	Handler Handler
}

func InitExternalServices(conf *config.Config) *ExternalServices {
	oracle := setupClient.InitOracle(conf)

	return &ExternalServices{
		Oracle: oracle,
	}
}

func InitCoreServices(conf *config.Config, externalService *ExternalServices) *CoreServices {
	repo := NewRepository(externalService.Oracle)
	service := NewService(conf, repo)
	handler := NewHandler(service)

	return &CoreServices{
		Handler: *handler,
	}
}
