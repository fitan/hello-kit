package service

import "context"

// Middleware describes a service middleware.
type Middleware func(HelloService) HelloService

type authMiddleware struct {
	next HelloService
}

// AuthMiddleware returns a HelloService Middleware.
func AuthMiddleware() Middleware {
	return func(next HelloService) HelloService {
		return &authMiddleware{next}
	}

}
func (a authMiddleware) Foo(ctx context.Context, s string) (rs string, err error) {
	// Implement your middleware logic here

	return a.next.Foo(ctx, s)
}
func (a authMiddleware) Say(ctx context.Context, req SayReq) (res SayRes, err error) {
	// Implement your middleware logic here

	return a.next.Say(ctx, req)
}
func (a authMiddleware) Say1(ctx context.Context, req SayReq) (res SayRes, err error) {
	// Implement your middleware logic here

	return a.next.Say1(ctx, req)
}
func (a authMiddleware) SayHello(ctx context.Context, req SayReq) (res SayRes, err error) {
	// Implement your middleware logic here

	return a.next.SayHello(ctx, req)
}
func (a authMiddleware) SayHello1(ctx context.Context, s1 string, s2 string) (res SayRes, err error) {
	// Implement your middleware logic here

	return a.next.SayHello1(ctx, s1, s2)
}
