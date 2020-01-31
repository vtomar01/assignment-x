package apis

import (
	"github.com/vtomar01/user-service/src/main/context"
	"net/http"
)

func CorrelationIdMiddleWare(next RequestHandler) RequestHandler {
	return func(req *http.Request) *Response {
		context.AddCorrelationId(req)
		return next(req)
	}
}
