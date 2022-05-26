package audit

// Code generated by gowrap. DO NOT EDIT.
// template:
// gowrap: http://github.com/fitan/gowrap

import (
	"context"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"hello/pkg/ent"
)

// AuditServiceWithPrometheus implements AuditService interface with all methods wrapped
// with Prometheus metrics
type AuditServiceWithPrometheus struct {
	base         AuditService
	instanceName string
}

var auditserviceDurationSummaryVec = promauto.NewSummaryVec(
	prometheus.SummaryOpts{
		Name:       "dao_audit_duration_seconds",
		Help:       "auditservice runtime duration and result",
		MaxAge:     time.Minute,
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	},
	[]string{"instance_name", "method", "result"})

// NewAuditServiceWithPrometheus returns an instance of the AuditService decorated with prometheus summary metric
func NewAuditServiceWithPrometheus(base AuditService) AuditServiceWithPrometheus {
	return AuditServiceWithPrometheus{
		base:         base,
		instanceName: "auditservice",
	}
}

// ByQueriesAll implements AuditService
func (_d AuditServiceWithPrometheus) ByQueriesAll(ctx context.Context, i interface{}) (res []ent.AuditBaseGetRes, count int, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		auditserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "ByQueriesAll", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.ByQueriesAll(ctx, i)
}

// ByQueriesOne implements AuditService
func (_d AuditServiceWithPrometheus) ByQueriesOne(ctx context.Context, i interface{}) (res ent.AuditBaseGetRes, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		auditserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "ByQueriesOne", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.ByQueriesOne(ctx, i)
}

// Create implements AuditService
func (_d AuditServiceWithPrometheus) Create(ctx context.Context, v ent.AuditBaseCreateReq) (res *ent.Audit, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		auditserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "Create", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.Create(ctx, v)
}

// CreateMany implements AuditService
func (_d AuditServiceWithPrometheus) CreateMany(ctx context.Context, vs []ent.AuditBaseCreateReq) (a1 ent.Audits, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		auditserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "CreateMany", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.CreateMany(ctx, vs)
}

// DeleteById implements AuditService
func (_d AuditServiceWithPrometheus) DeleteById(ctx context.Context, id int) (err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		auditserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "DeleteById", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.DeleteById(ctx, id)
}

// DeleteMany implements AuditService
func (_d AuditServiceWithPrometheus) DeleteMany(ctx context.Context, ids []int) (err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		auditserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "DeleteMany", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.DeleteMany(ctx, ids)
}

// GetById implements AuditService
func (_d AuditServiceWithPrometheus) GetById(ctx context.Context, id int) (res ent.AuditBaseGetRes, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		auditserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "GetById", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.GetById(ctx, id)
}

// UpdateById implements AuditService
func (_d AuditServiceWithPrometheus) UpdateById(ctx context.Context, id int, v ent.AuditBaseUpdateReq) (ap1 *ent.Audit, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		auditserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "UpdateById", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.UpdateById(ctx, id, v)
}

// UpdateMany implements AuditService
func (_d AuditServiceWithPrometheus) UpdateMany(ctx context.Context, vs []ent.AuditBaseUpdateReq) (err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		auditserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "UpdateMany", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.UpdateMany(ctx, vs)
}
