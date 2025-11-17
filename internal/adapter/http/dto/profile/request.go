package dto

type (
	GetProfileRequest struct {
		Proc bool `query:"proc" validate:"omitempty"`
	}

	CreateProfileRequest struct {
		Name          string        `json:"name"`
		ProfileDetail ProfileDetail `json:"detail"`
	}

	UpdateProfileIDRequest struct {
		ID int64 `uri:"id"`
	}

	UpdateProfileRequest struct {
		Name          string        `json:"name"`
		ProfileDetail ProfileDetail `json:"detail"`
	}

	DeleteProfileRequest struct {
		ID int64 `uri:"id"`
	}
)
