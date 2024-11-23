package error

type Response struct {
	Code    string `json:"code"`
	Source  string `json:"source"`
	Message string `json:"message"`
}

func NewResponse(detail Detail, source string) *Response {
	return &Response{
		Code:    detail.Code,
		Source:  source,
		Message: detail.Message,
	}
}
