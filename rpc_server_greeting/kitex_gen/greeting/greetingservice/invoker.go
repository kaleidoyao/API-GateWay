// Code generated by Kitex v0.6.1. DO NOT EDIT.

package greetingservice

import (
	server "github.com/cloudwego/kitex/server"
	greeting "github.com/kaleidoyao/API-GateWay/rpc_server_greeting/kitex_gen/greeting"
)

// NewInvoker creates a server.Invoker with the given handler and options.
func NewInvoker(handler greeting.GreetingService, opts ...server.Option) server.Invoker {
	var options []server.Option

	options = append(options, opts...)

	s := server.NewInvoker(options...)
	if err := s.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	return s
}