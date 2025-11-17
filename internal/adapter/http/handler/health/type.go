package handler

import "github.com/gunawanpras/oc-app-backend/internal/core/health/port"

type (
	ServiceAttribute struct {
		HealthService port.Service
	}

	HealthHandler struct {
		service ServiceAttribute
	}

	InitAttribute struct {
		Service ServiceAttribute
	}
)
