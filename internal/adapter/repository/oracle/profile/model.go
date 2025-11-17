package oracle

import (
	"time"

	"github.com/gunawanpras/oc-app-backend/internal/core/profile/domain"
)

type (
	Profile struct {
		ID        int64     `db:"ID"`
		Name      string    `db:"NAME"`
		Detail    []byte    `db:"DETAIL"`
		CreatedAt time.Time `db:"CREATED_AT"`
	}

	Profiles []Profile
)

func (h Profile) ToDomain() domain.Profile {
	return domain.Profile{
		ID:        h.ID,
		Name:      h.Name,
		Detail:    h.Detail,
		CreatedAt: h.CreatedAt,
	}
}

func (ps Profiles) ToDomain() domain.Profiles {
	var profiles domain.Profiles

	for _, profile := range ps {
		p := domain.Profile{
			ID:        profile.ID,
			Name:      profile.Name,
			Detail:    profile.Detail,
			CreatedAt: profile.CreatedAt,
		}

		profiles = append(profiles, p)
	}

	return profiles
}

func (h Profile) Validate() bool {
	if h.ID == 0 {
		return false
	}

	if h.Name == "" {
		return false
	}

	if h.Detail == nil {
		return false
	}

	if h.CreatedAt.IsZero() {
		return false
	}

	return true
}

func (ps Profiles) Validate() bool {
	for _, profile := range ps {
		if !profile.Validate() {
			return false
		}
	}

	return true
}
