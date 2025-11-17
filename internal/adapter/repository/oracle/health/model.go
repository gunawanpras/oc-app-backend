package oracle

import "github.com/gunawanpras/oc-app-backend/internal/core/health/domain"

type (
	HealthStatus struct {
		Status string
	}
)

func (h HealthStatus) ToDomain() domain.HealthStatus {
	return domain.HealthStatus{
		Status: h.Status,
	}
}

func (h HealthStatus) Validate() bool {
	return h.Status != ""
}
