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

type Err struct {
	Message string  `json:"message"`
	Code    int     `json:"code"`
	Err     string  `json:"error"`
	Causes  []Cause `json:"causes,omitempty"`
}

type Cause struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (r *Err) Error() string {
	return r.Message
}

func NewRestErr(message, err string, code int, causes []Cause) *Err {
	return &Err{
		Message: message,
		Code:    code,
		Err:     err,
		Causes:  causes,
	}
}

func NewBadRequestErr(message string, causes []Cause) *Err {
	return &Err{
		Message: message,
		Code:    http.StatusBadRequest,
		Err:     formatStatus(http.StatusBadRequest),
		Causes:  causes,
	}
}

func NewInternalServerErr(message string) *Err {
	return &Err{
		Message: message,
		Code:    http.StatusInternalServerError,
		Err:     formatStatus(http.StatusInternalServerError),
		Causes:  nil,
	}
}

func NewNotFoundErr(message string) *Err {
	return &Err{
		Message: message,
		Code:    http.StatusNotFound,
		Err:     formatStatus(http.StatusNotFound),
		Causes:  nil,
	}
}

func NewForbiddenErr(message string) *Err {
	return &Err{
		Message: message,
		Code:    http.StatusForbidden,
		Err:     formatStatus(http.StatusForbidden),
		Causes:  nil,
	}
}

func NewUnauthorizedErr() *Err {
	return &Err{
		Message: http.StatusText(http.StatusUnauthorized),
		Code:    http.StatusUnauthorized,
		Err:     formatStatus(http.StatusUnauthorized),
		Causes:  nil,
	}
}
