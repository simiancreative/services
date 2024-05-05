package errors

import (
	"net/http"
)

var Common = Available{Details: DefaultErrors}

const (
	NotAuthorized        = http.StatusUnauthorized
	NotFound             = http.StatusNotFound
	ValidationFailed     = http.StatusUnprocessableEntity
	StatusNotImplemented = http.StatusNotImplemented
	InternalServerError  = http.StatusInternalServerError
)

var DefaultErrors = map[int]Details{
	NotAuthorized: {
		Status:  http.StatusUnauthorized,
		Key:     "not_authorized",
		Message: "Not authorized",
	},

	NotFound: {
		Status:  http.StatusNotFound,
		Key:     "not_found",
		Message: "Not found",
	},

	ValidationFailed: {
		Status:  http.StatusUnprocessableEntity,
		Key:     "validation_failed",
		Message: "Validation failed",
	},

	StatusNotImplemented: {
		Status:  http.StatusNotImplemented,
		Key:     "not_implemented",
		Message: "Not implemented",
	},

	InternalServerError: {
		Status:  http.StatusInternalServerError,
		Key:     "internal_server_error",
		Message: "Internal server error",
	},
}
