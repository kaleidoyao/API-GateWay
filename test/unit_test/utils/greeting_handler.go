package utils

import (
	"context"

	greeting "github.com/kaleidoyao/API-GateWay/rpc_server_greeting/kitex_gen/greeting"
)

// GreetingServiceImpl implements the last service interface defined in the IDL.
type GreetingServiceImpl struct{}

// GreetingMethod implements the GreetingServiceImpl interface.
func (s *GreetingServiceImpl) GreetingMethod(ctx context.Context, request *greeting.GreetingRequest) (resp *greeting.GreetingResponse, err error) {
	greetingMessage := "Hello, " + request.Name

	resp = &greeting.GreetingResponse{
		ResponseBody: greetingMessage,
	}
	return
}
