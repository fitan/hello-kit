package httpclient

import (
	"context"
	//"github.com/fitan/magic/pkg/types"
	"github.com/go-resty/resty/v2"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/registry/cache"
	"go-micro.dev/v4/selector"
)

type microNext struct{}

func BeforeMicroSelect(serviceName string, r registry.Registry, options ...selector.SelectOption) resty.RequestMiddleware {
	s := selector.NewSelector(selector.Registry(cache.New(r)), selector.SetStrategy(selector.RoundRobin))
	return func(client *resty.Client, request *resty.Request) error {

		v := request.Context().Value(microNext{})
		if v != nil {
			next := v.(selector.Next)
			node, err := next()
			if err != nil {
				return err
			}
			client.SetHostURL("http://" + node.Address)
			return nil
		}

		next, err := s.Select(serviceName, options...)
		if err != nil {
			return err
		}
		node, err := next()
		if err != nil {
			return err
		}
		client.SetHostURL("http://" + node.Address)
		context.WithValue(request.Context(), microNext{}, next)
		return nil
	}
}
