package constant

import (
	"net/http"
)

const (
	SUCCESS = "success"
	ERROR   = "error"
)

const (
	SYSTEM = "SYSTEM"
)

const (
	BindingParameterFailed = "failed to bind parameter"
)

const (
	CheckAppHealthSuccess         = "app health check successful"
	CheckAppHealthFailed          = "app health check failed"
	CheckDatabaseHealthSuccess    = "database health check successful"
	CheckDatabaseHealthFailed     = "database health check failed"
	GetProfileSuccess             = "get profile successful"
	GetProfileFailed              = "get profile failed"
	GetProfileViaProcedureSuccess = "get profile via procedure successful"
	GetProfileViaProcedureFailed  = "get profile via procedure failed"
	CreateProfileSuccess          = "create profile successful"
	CreateProfileFailed           = "create profile failed"
	UpdateProfileSuccess          = "update profile successful"
	UpdateProfileFailed           = "update profile failed"
	DeleteProfileSuccess          = "delete profile successful"
	DeleteProfileFailed           = "delete profile failed"
)

var (
	GenericHttpStatusMappings = map[string]int{
		BindingParameterFailed: http.StatusBadRequest,
	}

	HealthHttpStatusMappings = map[string]int{
		CheckAppHealthSuccess:      http.StatusOK,
		CheckAppHealthFailed:       http.StatusInternalServerError,
		CheckDatabaseHealthSuccess: http.StatusOK,
		CheckDatabaseHealthFailed:  http.StatusInternalServerError,
	}

	ProfileHttpStatusMappings = map[string]int{
		GetProfileSuccess:    http.StatusOK,
		GetProfileFailed:     http.StatusInternalServerError,
		CreateProfileSuccess: http.StatusCreated,
		CreateProfileFailed:  http.StatusInternalServerError,
		UpdateProfileSuccess: http.StatusOK,
		UpdateProfileFailed:  http.StatusInternalServerError,
		DeleteProfileSuccess: http.StatusOK,
		DeleteProfileFailed:  http.StatusInternalServerError,
	}
)
