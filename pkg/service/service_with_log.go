package service

import "context"

// Code generated by gowrap. DO NOT EDIT.
// template: ../gowrap/templates/log
// gowrap: http://github.com/fitan/gowrap

//go:generate gowrap gen -p hello/pkg/service -i HelloService -t ../gowrap/templates/log -o service_with_log.go -l ""

type Log interface {
	Log(keyvals ...interface{}) error
}

// HelloServiceWithLog implements HelloService that is instrumented with logging
type HelloServiceWithLog struct {
	_log  Log
	_base HelloService
}

// NewHelloServiceWithLog instruments an implementation of the HelloService with simple logging
func NewHelloServiceWithLog(base HelloService, log Log) HelloServiceWithLog {
	return HelloServiceWithLog{
		_base: base,
		_log:  log,
	}
}

// Foo implements HelloService
func (_d HelloServiceWithLog) Foo(ctx context.Context, s string) (rs string, err error) {
	_params := []interface{}{"HelloServiceWithLog calling", "Foo", "params", map[string]interface{}{
		"s": s}}
	_d._log.Log(_params...)
	defer func() {
		_results := []interface{}{"HelloServiceWithLog calling", "Foo", "results", map[string]interface{}{
			"ctx": ctx,
			"s":   s}}
		if err != nil {
			_d._log.Log(_results...)
		} else {
			_d._log.Log(_results...)
		}
	}()
	return _d._base.Foo(ctx, s)
}

// Say implements HelloService
func (_d HelloServiceWithLog) Say(ctx context.Context, req SayReq) (res SayRes, err error) {
	_params := []interface{}{"HelloServiceWithLog calling", "Say", "params", map[string]interface{}{
		"req": req}}
	_d._log.Log(_params...)
	defer func() {
		_results := []interface{}{"HelloServiceWithLog calling", "Say", "results", map[string]interface{}{
			"ctx": ctx,
			"req": req}}
		if err != nil {
			_d._log.Log(_results...)
		} else {
			_d._log.Log(_results...)
		}
	}()
	return _d._base.Say(ctx, req)
}

// Say1 implements HelloService
func (_d HelloServiceWithLog) Say1(ctx context.Context, req SayReq) (res SayRes, err error) {
	_params := []interface{}{"HelloServiceWithLog calling", "Say1", "params", map[string]interface{}{
		"req": req}}
	_d._log.Log(_params...)
	defer func() {
		_results := []interface{}{"HelloServiceWithLog calling", "Say1", "results", map[string]interface{}{
			"ctx": ctx,
			"req": req}}
		if err != nil {
			_d._log.Log(_results...)
		} else {
			_d._log.Log(_results...)
		}
	}()
	return _d._base.Say1(ctx, req)
}

// SayHello implements HelloService
func (_d HelloServiceWithLog) SayHello(ctx context.Context, req SayReq) (res SayRes, err error) {
	_params := []interface{}{"HelloServiceWithLog calling", "SayHello", "params", map[string]interface{}{
		"req": req}}
	_d._log.Log(_params...)
	defer func() {
		_results := []interface{}{"HelloServiceWithLog calling", "SayHello", "results", map[string]interface{}{
			"ctx": ctx,
			"req": req}}
		if err != nil {
			_d._log.Log(_results...)
		} else {
			_d._log.Log(_results...)
		}
	}()
	return _d._base.SayHello(ctx, req)
}

// SayHello1 implements HelloService
func (_d HelloServiceWithLog) SayHello1(ctx context.Context, s1 string, s2 string) (res SayRes, err error) {
	_params := []interface{}{"HelloServiceWithLog calling", "SayHello1", "params", map[string]interface{}{
		"s1": s1,
		"s2": s2}}
	_d._log.Log(_params...)
	defer func() {
		_results := []interface{}{"HelloServiceWithLog calling", "SayHello1", "results", map[string]interface{}{
			"ctx": ctx,
			"s1":  s1,
			"s2":  s2}}
		if err != nil {
			_d._log.Log(_results...)
		} else {
			_d._log.Log(_results...)
		}
	}()
	return _d._base.SayHello1(ctx, s1, s2)
}
