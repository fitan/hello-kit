package endpoint

import (
	"context"
	service "hello/pkg/service"

	endpoint "github.com/go-kit/kit/endpoint"
)

// FooRequest collects the request parameters for the Foo method.
type FooRequest struct {
	S string `json:"s"`
}

// FooResponse collects the response parameters for the Foo method.
type FooResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeFooEndpoint returns an endpoint that invokes Foo on the service.
func MakeFooEndpoint(s service.HelloService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(FooRequest)
		rs, err := s.Foo(ctx, req.S)
		return FooResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r FooResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Foo implements Service. Primarily useful in a client.
func (e Endpoints) Foo(ctx context.Context, s string) (rs string, err error) {
	request := FooRequest{S: s}
	response, err := e.FooEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(FooResponse).Rs, response.(FooResponse).Err
}

// SayRequest collects the request parameters for the Say method.
type SayRequest struct {
	Req service.SayReq `json:"req"`
}

// SayResponse collects the response parameters for the Say method.
type SayResponse struct {
	Res service.SayRes `json:"res"`
	Err error          `json:"err"`
}

// MakeSayEndpoint returns an endpoint that invokes Say on the service.
func MakeSayEndpoint(s service.HelloService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SayRequest)
		res, err := s.Say(ctx, req.Req)
		return SayResponse{
			Err: err,
			Res: res,
		}, nil
	}
}

// Failed implements Failer.
func (r SayResponse) Failed() error {
	return r.Err
}

// Say implements Service. Primarily useful in a client.
func (e Endpoints) Say(ctx context.Context, req service.SayReq) (res service.SayRes, err error) {
	request := SayRequest{Req: req}
	response, err := e.SayEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(SayResponse).Res, response.(SayResponse).Err
}

// Say1Request collects the request parameters for the Say1 method.
type Say1Request struct {
	Req service.SayReq `json:"req"`
}

// Say1Response collects the response parameters for the Say1 method.
type Say1Response struct {
	Res service.SayRes `json:"res"`
	Err error          `json:"err"`
}

// MakeSay1Endpoint returns an endpoint that invokes Say1 on the service.
func MakeSay1Endpoint(s service.HelloService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(Say1Request)
		res, err := s.Say1(ctx, req.Req)
		return Say1Response{
			Err: err,
			Res: res,
		}, nil
	}
}

// Failed implements Failer.
func (r Say1Response) Failed() error {
	return r.Err
}

// Say1 implements Service. Primarily useful in a client.
func (e Endpoints) Say1(ctx context.Context, req service.SayReq) (res service.SayRes, err error) {
	request := Say1Request{Req: req}
	response, err := e.Say1Endpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(Say1Response).Res, response.(Say1Response).Err
}

// SayHelloRequest collects the request parameters for the SayHello method.
type SayHelloRequest struct {
	Req service.SayReq `json:"req"`
}

// SayHelloResponse collects the response parameters for the SayHello method.
type SayHelloResponse struct {
	Res service.SayRes `json:"res"`
	Err error          `json:"err"`
}

// MakeSayHelloEndpoint returns an endpoint that invokes SayHello on the service.
func MakeSayHelloEndpoint(s service.HelloService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SayHelloRequest)
		res, err := s.SayHello(ctx, req.Req)
		return SayHelloResponse{
			Err: err,
			Res: res,
		}, nil
	}
}

// Failed implements Failer.
func (r SayHelloResponse) Failed() error {
	return r.Err
}

// SayHello implements Service. Primarily useful in a client.
func (e Endpoints) SayHello(ctx context.Context, req service.SayReq) (res service.SayRes, err error) {
	request := SayHelloRequest{Req: req}
	response, err := e.SayHelloEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(SayHelloResponse).Res, response.(SayHelloResponse).Err
}

// SayHello1Request collects the request parameters for the SayHello1 method.
type SayHello1Request struct {
	S1 string `json:"s1"`
	S2 string `json:"s2"`
}

// SayHello1Response collects the response parameters for the SayHello1 method.
type SayHello1Response struct {
	Res service.SayRes `json:"res"`
	Err error          `json:"err"`
}

// MakeSayHello1Endpoint returns an endpoint that invokes SayHello1 on the service.
func MakeSayHello1Endpoint(s service.HelloService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SayHello1Request)
		res, err := s.SayHello1(ctx, req.S1, req.S2)
		return SayHello1Response{
			Err: err,
			Res: res,
		}, nil
	}
}

// Failed implements Failer.
func (r SayHello1Response) Failed() error {
	return r.Err
}

// SayHello1 implements Service. Primarily useful in a client.
func (e Endpoints) SayHello1(ctx context.Context, s1 string, s2 string) (res service.SayRes, err error) {
	request := SayHello1Request{
		S1: s1,
		S2: s2,
	}
	response, err := e.SayHello1Endpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(SayHello1Response).Res, response.(SayHello1Response).Err
}
