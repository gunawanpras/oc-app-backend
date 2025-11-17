package dto

type (
	ProfileDetail struct {
		Name     string  `json:"name"`
		Age      int     `json:"age"`
		Contacts Contact `json:"contacts"`
		Address  Address `json:"address"`
	}

	Contact struct {
		Email string `json:"email"`
		Phone string `json:"phone"`
	}

	Address struct {
		Street     string `json:"street"`
		City       string `json:"city"`
		PostalCode string `json:"postal_code"`
	}

	GetProfileResponse struct {
		ID        int64          `json:"id"`
		Name      string         `json:"name"`
		Detail    *ProfileDetail `json:"detail"`
		CreatedAt string         `json:"created_at"`
	}

	GetProfileResponses []GetProfileResponse

	CreateProfileResponse struct {
		ID int64 `json:"id"`
	}

	UpdateProfileResponse struct {
		ID int64 `json:"id"`
	}

	ProfilesResponse struct {
		Profiles []GetProfileResponse `json:"profiles"`
	}
)
