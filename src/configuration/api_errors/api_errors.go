package api_errors

import "net/http"

type ApiError struct {
	Code    int     `json:"code"`
	Message string  `json:"message"`
	Err     string  `json:"error"`
	Causes  []Cause `json:"causes"`
}

type Cause struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (e *ApiError) Error() string {
	return e.Message
}

func NewApiError(message, err string, code int, causes []Cause) *ApiError {
	return &ApiError{
		Message: message,
		Err:     err,
		Code:    code,
	}
}

func NewBadRequestError(message string) *ApiError {
	return &ApiError{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
	}
}

func NewBadRequestValidationError(message string, causes []Cause) *ApiError {
	return &ApiError{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func NewInternalServerError(message string) *ApiError {
	return &ApiError{
		Message: message,
		Err:     "internal_server_error",
		Code:    http.StatusInternalServerError,
	}
}

func NewNotFoundError(message string) *ApiError {
	return &ApiError{
		Message: message,
		Err:     "not_found",
		Code:    http.StatusNotFound,
	}
}

func NewUnauthorizedError(message string) *ApiError {
	return &ApiError{
		Message: message,
		Err:     "unauthorized",
		Code:    http.StatusUnauthorized,
	}
}
