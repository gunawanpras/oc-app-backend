package port

import (
	"context"

	"github.com/gunawanpras/oc-app-backend/internal/core/profile/domain"
)

type Service interface {
	GetProfile(ctx context.Context, options domain.GetProfileOptions) (res domain.Profiles, err error)
	CreateProfile(ctx context.Context, profile domain.Profile) (res domain.Profile, err error)
	UpdateProfile(ctx context.Context, profile domain.Profile) (res domain.Profile, err error)
	DeleteProfile(ctx context.Context, id int64) (err error)
}
