package handler

import "github.com/gunawanpras/oc-app-backend/internal/core/profile/port"

type (
	ServiceAttribute struct {
		ProfileService port.Service
	}

	ProfileHandler struct {
		service ServiceAttribute
	}

	InitAttribute struct {
		Service ServiceAttribute
	}
)
