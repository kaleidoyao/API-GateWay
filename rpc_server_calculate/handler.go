package main

import (
	"context"
	calculate "github.com/kaleidoyao/rpc_server_calculate/kitex_gen/calculate"
)

// CalculateServiceImpl implements the last service interface defined in the IDL.
type CalculateServiceImpl struct{}

// CalculateMethod implements the CalculateServiceImpl interface.
func (s *CalculateServiceImpl) CalculateMethod(ctx context.Context, request *calculate.CalculateRequest) (resp *calculate.CalculateResponse, err error) {
	// TODO: Your code here...
	return
}
