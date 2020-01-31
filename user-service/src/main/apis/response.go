package apis

import (
	"encoding/json"
	"github.com/vtomar01/user-service/src/main/context"
	"github.com/vtomar01/user-service/src/main/logging"
	"net/http"
	"runtime/debug"
)

func Process(ctx *context.Context, statusCode int,
	processor func() (interface{}, error)) *Response {

	return responseWithStatus(ctx, statusCode, processor)
}

func responseWithStatus(ctx *context.Context, statusCode int,
	processor func() (interface{}, error)) *Response {

	res, err := processor()

	if err != nil {
		return HandleError(err)
	}
	return &Response{
		Code:    statusCode,
		Payload: res,
	}
}

// Response is written to http.ResponseWriter
type Response struct {
	Code    int
	Payload interface{}
}

// Make creates a http handler from a request handler func
func Make(r RequestHandler) func(w http.ResponseWriter, req *http.Request) {

	handler := func(w http.ResponseWriter, req *http.Request) {

		defer func() {
			if r := recover(); r != nil {
				logging.Log.Errorf("Panic Recovered!: %v \n %v", r, string(debug.Stack()))
				response := &Response{
					Code:    http.StatusInternalServerError,
					Payload: r,
				}
				writeToResponseWriter(w, response)
			}
		}()

		res := r(req)
		writeToResponseWriter(w, res)
		req.Body.Close()
	}

	return handler
}

func writeToResponseWriter(w http.ResponseWriter, response *Response) {
	responseBody, _ := json.Marshal(response.Payload)
	w.WriteHeader(response.Code)
	w.Write(responseBody)
	w.Header().Set("Content-Type", "application/json")
}
