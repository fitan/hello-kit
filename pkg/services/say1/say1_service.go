package say1

// Say1Service describes the service.
// if use kit. go:generate gowrap gen -g -p ./ -i Say1Service -bt "prometheus:say1_with_prometheus_gen.go log:say1_with_log_gen.go opentracing:say1_with_trace_gen.go http-gin:say1_http_gen.go endpoint:say1_endpoint_gen.go"
//go:generate gowrap gen -ps=false -g -p ./ -i Say1Service -bt "prometheus:say1_with_prometheus_gen.go log:say1_with_log_gen.go opentracing:say1_with_trace_gen.go"
type Say1Service interface {
}

type basicSay1Service struct {
}

type BaseService Say1Service

// NewBasicSay1Service returns a naive, stateless implementation of Say1Service.
func NewBasicService() BaseService {
	return &basicSay1Service{}
}
