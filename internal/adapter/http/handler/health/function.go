package handler

import (
	"github.com/gofiber/fiber/v2"

	dto "github.com/gunawanpras/oc-app-backend/internal/adapter/http/dto/health"
	"github.com/gunawanpras/oc-app-backend/pkg/response"
	"github.com/gunawanpras/oc-app-backend/pkg/util/constant"
)

func (handler *HealthHandler) CheckAppHealth(c *fiber.Ctx) error {
	ctx := c.UserContext()

	resp, err := handler.service.HealthService.CheckAppHealth(ctx)
	if err != nil {
		return response.Error(c, constant.CheckAppHealthFailed, err, constant.HealthHttpStatusMappings)
	}

	respData := dto.HealthResponse{
		Status: resp.Status,
	}

	return response.OK(c, constant.CheckAppHealthSuccess, respData, constant.HealthHttpStatusMappings)
}

func (handler *HealthHandler) CheckDatabaseHealth(c *fiber.Ctx) error {
	ctx := c.UserContext()

	resp, err := handler.service.HealthService.CheckDatabaseHealth(ctx)
	if err != nil {
		return response.Error(c, constant.CheckDatabaseHealthFailed, err, constant.HealthHttpStatusMappings)
	}

	respData := dto.HealthResponse{
		Status: resp.Status,
	}

	return response.OK(c, constant.CheckDatabaseHealthSuccess, respData, constant.HealthHttpStatusMappings)
}
