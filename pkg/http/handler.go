package http

import (
	"context"
	"encoding/json"
	"errors"
	endpoint "hello/pkg/endpoint"
	"net/http"

	http1 "github.com/go-kit/kit/transport/http"
)

// makeFooHandler creates the handler logic
func makeFooHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/foo", http1.NewServer(endpoints.FooEndpoint, decodeFooRequest, encodeFooResponse, options...))
}

// decodeFooRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeFooRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.FooRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeFooResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeFooResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
func ErrorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	w.WriteHeader(err2code(err))
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}
func ErrorDecoder(r *http.Response) error {
	var w errorWrapper
	if err := json.NewDecoder(r.Body).Decode(&w); err != nil {
		return err
	}
	return errors.New(w.Error)
}

// This is used to set the http status, see an example here :
// https://github.com/go-kit/kit/blob/master/examples/addsvc/pkg/addtransport/http.go#L133
func err2code(err error) int {
	return http.StatusInternalServerError
}

type errorWrapper struct {
	Error string `json:"error"`
}

// makeSayHandler creates the handler logic
func makeSayHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/say", http1.NewServer(endpoints.SayEndpoint, decodeSayRequest, encodeSayResponse, options...))
}

// decodeSayRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeSayRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.SayRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeSayResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeSayResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeSay1Handler creates the handler logic
func makeSay1Handler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/say1", http1.NewServer(endpoints.Say1Endpoint, decodeSay1Request, encodeSay1Response, options...))
}

// decodeSay1Request is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeSay1Request(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.Say1Request{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeSay1Response is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeSay1Response(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeSayHelloHandler creates the handler logic
func makeSayHelloHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/say-hello", http1.NewServer(endpoints.SayHelloEndpoint, decodeSayHelloRequest, encodeSayHelloResponse, options...))
}

// decodeSayHelloRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeSayHelloRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.SayHelloRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeSayHelloResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeSayHelloResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeSayHello1Handler creates the handler logic
func makeSayHello1Handler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/say-hello1", http1.NewServer(endpoints.SayHello1Endpoint, decodeSayHello1Request, encodeSayHello1Response, options...))
}

// decodeSayHello1Request is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeSayHello1Request(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.SayHello1Request{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeSayHello1Response is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeSayHello1Response(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
