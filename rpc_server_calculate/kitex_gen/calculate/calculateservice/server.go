// Code generated by Kitex v0.6.1. DO NOT EDIT.
package calculateservice

import (
	server "github.com/cloudwego/kitex/server"
	calculate "github.com/kaleidoyao/API-GateWay/rpc_server_calculate/kitex_gen/calculate"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler calculate.CalculateService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
