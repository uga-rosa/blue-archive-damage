package http_error

import (
	error2 "blue-archive-damage/internal/error"
	"net/http"
)

func newHttp400(code, message string) *error2.Detail {
	return &error2.Detail{
		Status:  http.StatusBadRequest,
		Code:    code,
		Message: message,
	}
}

var (
	Http400Generic = newHttp400("BAD_REQUEST", "Bad Request")
)
