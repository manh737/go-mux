package model

type (
	// Response is the http json response schema
	Response struct {
		Status  int         `json:"status"`
		Message string      `json:"message"`
		Content interface{} `json:"content"`
	}
)

// NewResponse is the Response struct factory function.
func NewResponse(status int, message string, content interface{}) *Response {
	return &Response{
		Status:  status,
		Message: message,
		Content: content,
	}
}
