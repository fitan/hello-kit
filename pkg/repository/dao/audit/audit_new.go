package audit

import (
	"github.com/google/wire"
	"go.uber.org/zap"
)

type Middleware func(AuditService) AuditService

func NewServiceMiddleware(logger *zap.SugaredLogger) (mw []Middleware) {
	mw = []Middleware{}
	// Append your middleware here
	mw = append(mw, func(AuditService AuditService) AuditService {
		return NewAuditServiceWithPrometheus(AuditService)
	})
	mw = append(mw, func(AuditService AuditService) AuditService {
		return NewAuditServiceWithLog(AuditService, logger)
	})
	mw = append(mw, func(AuditService AuditService) AuditService {
		return NewAuditServiceWithTracing(AuditService)
	})
	return
}

// New returns a AuditService with all of the expected middleware wired in.
func NewService(svc BaseService, middleware []Middleware) AuditService {

	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}

var AuditServiceSet = wire.NewSet(NewBasicService, NewService, NewServiceMiddleware)
