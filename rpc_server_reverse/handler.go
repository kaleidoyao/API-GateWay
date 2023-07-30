package main

import (
	"context"
	reverse "github.com/kaleidoyao/rpc_server_reverse/kitex_gen/reverse"
)

// ReverseServiceImpl implements the last service interface defined in the IDL.
type ReverseServiceImpl struct{}

// ReverseMethod implements the ReverseServiceImpl interface.
func (s *ReverseServiceImpl) ReverseMethod(ctx context.Context, request *reverse.ReverseRequest) (resp *reverse.ReverseResponse, err error) {
	// TODO: Your code here...
	return
}
