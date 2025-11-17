package setup

import (
	handler "github.com/gunawanpras/oc-app-backend/internal/adapter/http/handler/health"
	profileHandler "github.com/gunawanpras/oc-app-backend/internal/adapter/http/handler/profile"
)

type Handler struct {
	HealthHandler  handler.Handler
	ProfileHandler profileHandler.Handler
}

func NewHandler(service Service) *Handler {
	return &Handler{
		HealthHandler: handler.New(handler.InitAttribute{
			Service: handler.ServiceAttribute{
				HealthService: service.HealthService,
			},
		}),
		ProfileHandler: profileHandler.New(profileHandler.InitAttribute{
			Service: profileHandler.ServiceAttribute{
				ProfileService: service.ProfileService,
			},
		}),
	}
}
