package handler

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"

	dto "github.com/gunawanpras/oc-app-backend/internal/adapter/http/dto/profile"
	"github.com/gunawanpras/oc-app-backend/internal/core/profile/domain"
	"github.com/gunawanpras/oc-app-backend/pkg/response"
	"github.com/gunawanpras/oc-app-backend/pkg/util/constant"
	"github.com/gunawanpras/oc-app-backend/pkg/validator"
)

func (handler *ProfileHandler) GetProfile(c *fiber.Ctx) error {
	ctx := c.UserContext()

	var req dto.GetProfileRequest
	if err := c.QueryParser(&req); err != nil {
		return response.Error(c, constant.BindingParameterFailed, err, constant.GenericHttpStatusMappings)
	}

	options := domain.GetProfileOptions{
		Proc: req.Proc,
	}

	resp, err := handler.service.ProfileService.GetProfile(ctx, options)
	if err != nil {
		return response.Error(c, constant.GetProfileFailed, err, constant.HealthHttpStatusMappings)
	}

	respData, err := handler.buildGetProfile(resp)
	if err != nil {
		return response.Error(c, constant.GetProfileFailed, err, constant.HealthHttpStatusMappings)
	}

	return response.OK(c, constant.GetProfileSuccess, respData, constant.HealthHttpStatusMappings)
}

func (handler *ProfileHandler) buildGetProfile(resp domain.Profiles) (respData dto.GetProfileResponses, err error) {
	for _, v := range resp {
		detail, err := handler.buildProfileDetail(v.Detail)
		if err != nil {
			return nil, err
		}

		respData = append(respData, dto.GetProfileResponse{
			ID:        v.ID,
			Name:      v.Name,
			Detail:    detail,
			CreatedAt: v.CreatedAt.String(),
		})
	}

	return respData, nil
}

func (handler *ProfileHandler) buildProfileDetail(data []byte) (res *dto.ProfileDetail, err error) {
	err = json.Unmarshal(data, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (handler *ProfileHandler) CreateProfile(c *fiber.Ctx) error {
	ctx := c.UserContext()

	var req dto.CreateProfileRequest
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, constant.BindingParameterFailed, err, constant.GenericHttpStatusMappings)
	}

	errv := validator.Validate(req)
	if errv != nil {
		return response.ErrorValidator(c, errv)
	}

	data, err := json.Marshal(req.ProfileDetail)
	if err != nil {
		return response.Error(c, constant.CreateProfileFailed, err, constant.HealthHttpStatusMappings)
	}

	args := domain.Profile{
		Name:   req.Name,
		Detail: data,
	}

	res, err := handler.service.ProfileService.CreateProfile(ctx, args)
	if err != nil {
		return response.Error(c, constant.CreateProfileFailed, err, constant.HealthHttpStatusMappings)
	}

	respData := dto.CreateProfileResponse{
		ID: res.ID,
	}

	return response.OK(c, constant.CreateProfileSuccess, respData, constant.HealthHttpStatusMappings)
}

func (handler *ProfileHandler) UpdateProfile(c *fiber.Ctx) error {
	var (
		id  dto.UpdateProfileIDRequest
		req dto.UpdateProfileRequest
	)

	ctx := c.UserContext()

	if err := c.ParamsParser(&id); err != nil {
		return response.Error(c, constant.BindingParameterFailed, err, constant.GenericHttpStatusMappings)
	}

	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, constant.BindingParameterFailed, err, constant.GenericHttpStatusMappings)
	}

	data, err := json.Marshal(req.ProfileDetail)
	if err != nil {
		return response.Error(c, constant.CreateProfileFailed, err, constant.HealthHttpStatusMappings)
	}

	args := domain.Profile{
		ID:     id.ID,
		Name:   req.Name,
		Detail: data,
	}

	errv := validator.Validate(req)
	if errv != nil {
		return response.ErrorValidator(c, errv)
	}

	resp, err := handler.service.ProfileService.UpdateProfile(ctx, args)
	if err != nil {
		return response.Error(c, constant.UpdateProfileFailed, err, constant.HealthHttpStatusMappings)
	}

	respData := dto.UpdateProfileResponse{
		ID: resp.ID,
	}

	return response.OK(c, constant.UpdateProfileSuccess, respData, constant.HealthHttpStatusMappings)
}

func (handler *ProfileHandler) DeleteProfile(c *fiber.Ctx) error {
	ctx := c.UserContext()

	var req dto.DeleteProfileRequest
	if err := c.ParamsParser(&req); err != nil {
		return response.Error(c, constant.BindingParameterFailed, err, constant.GenericHttpStatusMappings)
	}

	args := domain.Profile{
		ID: req.ID,
	}

	errv := validator.Validate(req)
	if errv != nil {
		return response.ErrorValidator(c, errv)
	}

	err := handler.service.ProfileService.DeleteProfile(ctx, args.ID)
	if err != nil {
		return response.Error(c, constant.DeleteProfileFailed, err, constant.HealthHttpStatusMappings)
	}

	return response.OK(c, constant.DeleteProfileSuccess, nil, constant.HealthHttpStatusMappings)
}
