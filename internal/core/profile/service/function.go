package service

import (
	"context"
	"errors"

	"github.com/gunawanpras/oc-app-backend/internal/core/profile/domain"
)

func (service *ProfileService) GetProfile(ctx context.Context, options domain.GetProfileOptions) (res domain.Profiles, err error) {
	if options.Proc {
		res, err = service.repo.ProfileRepo.GetProfileViaProcedure(ctx)

	} else {
		res, err = service.repo.ProfileRepo.GetProfile(ctx)
	}

	if err != nil {
		if err.Error() == "data not found" {
			return res, errors.New("profile not found")
		}

		return res, err
	}

	return res, nil
}

func (service *ProfileService) CreateProfile(ctx context.Context, profile domain.Profile) (res domain.Profile, err error) {
	res, err = service.repo.ProfileRepo.CreateProfile(ctx, profile)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (service *ProfileService) UpdateProfile(ctx context.Context, profile domain.Profile) (res domain.Profile, err error) {
	res, err = service.repo.ProfileRepo.UpdateProfile(ctx, profile)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (service *ProfileService) DeleteProfile(ctx context.Context, id int64) (err error) {
	err = service.repo.ProfileRepo.DeleteProfile(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
