package say

import "github.com/google/wire"

var SayKitSet = wire.NewSet(NewBasicService, NewService, NewEndpointMiddleware, NewServiceMiddleware, NewEndpoints, NewServiceOption, NewHTTPHandler)