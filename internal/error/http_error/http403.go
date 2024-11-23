package http_error

import (
	error2 "blue-archive-damage/internal/error"
	"net/http"
)

func newHttp403(code, message string) *error2.Detail {
	return &error2.Detail{
		Status:  http.StatusUnauthorized,
		Code:    code,
		Message: message,
	}
}

var (
	Http403Generic = newHttp403("FORBIDDEN", "Forbidden")
)
