package port

import (
	"context"

	"github.com/gunawanpras/oc-app-backend/internal/core/profile/domain"
)

type Repository interface {
	GetProfile(ctx context.Context) (res domain.Profiles, err error)
	GetProfileViaProcedure(ctx context.Context) (res domain.Profiles, err error)
	CreateProfile(ctx context.Context, newProfile domain.Profile) (res domain.Profile, err error)
	UpdateProfile(ctx context.Context, updatedProfile domain.Profile) (res domain.Profile, err error)
	DeleteProfile(ctx context.Context, id int64) (err error)
}
