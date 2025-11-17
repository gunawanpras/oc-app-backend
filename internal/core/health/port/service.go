package port

import (
	"context"

	"github.com/gunawanpras/oc-app-backend/internal/core/health/domain"
)

type Service interface {
	CheckAppHealth(ctx context.Context) (status domain.HealthStatus, err error)
	CheckDatabaseHealth(ctx context.Context) (status domain.HealthStatus, err error)
}
