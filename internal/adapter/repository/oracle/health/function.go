package oracle

import (
	"context"
	"database/sql"
	"errors"

	"github.com/gunawanpras/oc-app-backend/internal/core/health/domain"
)

func (repo *HealthRepository) CheckOracleHealth(ctx context.Context) (res domain.HealthStatus, err error) {
	var healthStatus HealthStatus

	repo.prepareCheckDatabaseHealth()
	_, err = repo.statement.CheckHealth.QueryxContext(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return res, errors.New("error when select 1 to db")
		}

		return res, err
	}

	healthStatus.Status = "OK"

	if !healthStatus.Validate() {
		return res, errors.New("invalid health status data")
	}

	res = healthStatus.ToDomain()

	return res, nil
}
