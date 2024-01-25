package rest

import (
	"net/http"
	"strings"
)

func formatStatus(code int) string {
	status := http.StatusText(code)
	status = strings.Join(strings.Split(status, " "), "_")

	return strings.ToLower(status)
}

// Default rest error type.
type Err struct {
	Message string  `json:"message"`
	Code    int     `json:"code"`
	Err     string  `json:"error"`
	Causes  []Cause `json:"causes,omitempty"`
}

// Describes the cause of the error occurrence.
type Cause struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// Returns error message on string format.
func (r *Err) Error() string {
	return r.Message
}

// Creates a new rest.Err.
func NewRestErr(message, err string, code int, causes []Cause) *Err {
	return &Err{
		Message: message,
		Code:    code,
		Err:     err,
		Causes:  causes,
	}
}

// Creates a new Bad Request Error.
func NewBadRequestErr(message string, causes []Cause) *Err {
	return &Err{
		Message: message,
		Code:    http.StatusBadRequest,
		Err:     formatStatus(http.StatusBadRequest),
		Causes:  causes,
	}
}

// Creates a new Internal Server Error.
func NewInternalServerErr(message string) *Err {
	return &Err{
		Message: message,
		Code:    http.StatusInternalServerError,
		Err:     formatStatus(http.StatusInternalServerError),
		Causes:  nil,
	}
}

// Creates a new Not Found Error.
func NewNotFoundErr(message string) *Err {
	return &Err{
		Message: message,
		Code:    http.StatusNotFound,
		Err:     formatStatus(http.StatusNotFound),
		Causes:  nil,
	}
}

// Creates a new Forbidden Error.
func NewForbiddenErr(message string) *Err {
	return &Err{
		Message: message,
		Code:    http.StatusForbidden,
		Err:     formatStatus(http.StatusForbidden),
		Causes:  nil,
	}
}

// Creates a new Unauthorized Error.
func NewUnauthorizedErr() *Err {
	return &Err{
		Message: http.StatusText(http.StatusUnauthorized),
		Code:    http.StatusUnauthorized,
		Err:     formatStatus(http.StatusUnauthorized),
		Causes:  nil,
	}
}
