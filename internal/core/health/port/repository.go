package port

import (
	"context"

	"github.com/gunawanpras/oc-app-backend/internal/core/health/domain"
)

type Repository interface {
	CheckOracleHealth(ctx context.Context) (status domain.HealthStatus, err error)
}
