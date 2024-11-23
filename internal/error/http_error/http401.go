package http_error

import (
	error2 "blue-archive-damage/internal/error"
	"net/http"
)

func newHttp401(code, message string) *error2.Detail {
	return &error2.Detail{
		Status:  http.StatusUnauthorized,
		Code:    code,
		Message: message,
	}
}

var (
	Http401Generic = newHttp401("UNAUTHORIZED", "Unauthorized")
)
