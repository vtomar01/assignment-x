package context

import (
	"github.com/vtomar01/user-service/src/main/utils/uuid"
	"go.uber.org/zap"
	"net/http"
)

const CorrelationId = "Correlation-ID"

type Context struct {
	CorrelationId string
	Data          map[string]interface{}
	Logger        *zap.SugaredLogger
}

func CreateLoggableContext(r *http.Request, log *zap.SugaredLogger) *Context {
	correlationId := GetCorrelationId(r)
	log = log.With(CorrelationId, correlationId)
	return &Context{CorrelationId: correlationId, Logger: log}
}

func GetCorrelationId(r *http.Request) string {
	return r.Header.Get(CorrelationId)
}

func AddCorrelationId(r *http.Request) {
	correlationId := GetCorrelationId(r)
	if correlationId == "" {
		correlationId = uuid.V4()
		r.Header.Set(CorrelationId, correlationId)
	}
}

func AttachCorrelationIdFromContext(r *http.Request, ctx *Context) {
	r.Header.Set(CorrelationId, ctx.CorrelationId)
}
