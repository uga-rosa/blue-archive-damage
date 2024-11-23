package http_error

import (
	error2 "blue-archive-damage/internal/error"
	"net/http"
)

func newHttp404(code, message string) *error2.Detail {
	return &error2.Detail{
		Status:  http.StatusNotFound,
		Code:    code,
		Message: message,
	}
}

var (
	Http404Generic = newHttp404("NOT_FOUND", "Not Found")
)
