package service

import (
	"context"

	"github.com/gunawanpras/oc-app-backend/internal/core/health/domain"
)

func (service *HealthService) CheckAppHealth(ctx context.Context) (res domain.HealthStatus, err error) {
	res = domain.HealthStatus{
		Status: "OK",
	}

	return res, nil
}

func (service *HealthService) CheckDatabaseHealth(ctx context.Context) (res domain.HealthStatus, err error) {
	res, err = service.repo.HealthRepo.CheckOracleHealth(ctx)
	if err != nil {
		return res, err
	}

	return res, nil
}
