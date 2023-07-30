package main

import (
	"context"
	greeting "github.com/kaleidoyao/rpc_server_greeting/kitex_gen/greeting"
)

// GreetingServiceImpl implements the last service interface defined in the IDL.
type GreetingServiceImpl struct{}

// GreetingMethod implements the GreetingServiceImpl interface.
func (s *GreetingServiceImpl) GreetingMethod(ctx context.Context, request *greeting.GreetingRequest) (resp *greeting.GreetingResponse, err error) {
	// TODO: Your code here...
	return
}
