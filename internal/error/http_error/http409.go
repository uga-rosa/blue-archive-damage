package http_error

import (
	error2 "blue-archive-damage/internal/error"
	"net/http"
)

func newHttp409(code, message string) *error2.Detail {
	return &error2.Detail{
		Status:  http.StatusConflict,
		Code:    code,
		Message: message,
	}
}

var (
	Http409Generic = newHttp409("CONFLICT", "Conflict")
)
