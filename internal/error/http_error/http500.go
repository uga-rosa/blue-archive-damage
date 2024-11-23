package http_error

import (
	error2 "blue-archive-damage/internal/error"
	"net/http"
)

func newHttp500(code, message string) *error2.Detail {
	return &error2.Detail{
		Status:  http.StatusInternalServerError,
		Code:    code,
		Message: message,
	}
}

var (
	Http500Generic = newHttp500("INTERNAL_SERVER_ERROR", "Internal Server Error")
)
