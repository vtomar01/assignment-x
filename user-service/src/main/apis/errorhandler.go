package apis

import "net/http"

type ErrorResponse struct {
	Error string `json:"error"`
	Code  string `json:"code"`
}

func HandleError(err error) *Response {
	return &Response{
		Code: http.StatusInternalServerError,
		Payload: &ErrorResponse{
			Error: err.Error(),
		},
	}
}
