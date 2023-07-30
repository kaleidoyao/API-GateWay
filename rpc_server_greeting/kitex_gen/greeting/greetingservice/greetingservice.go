// Code generated by Kitex v0.6.1. DO NOT EDIT.

package greetingservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	greeting "github.com/kaleidoyao/rpc_server_greeting/kitex_gen/greeting"
)

func serviceInfo() *kitex.ServiceInfo {
	return greetingServiceServiceInfo
}

var greetingServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "GreetingService"
	handlerType := (*greeting.GreetingService)(nil)
	methods := map[string]kitex.MethodInfo{
		"greetingMethod": kitex.NewMethodInfo(greetingMethodHandler, newGreetingServiceGreetingMethodArgs, newGreetingServiceGreetingMethodResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "greeting",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.6.1",
		Extra:           extra,
	}
	return svcInfo
}

func greetingMethodHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*greeting.GreetingServiceGreetingMethodArgs)
	realResult := result.(*greeting.GreetingServiceGreetingMethodResult)
	success, err := handler.(greeting.GreetingService).GreetingMethod(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newGreetingServiceGreetingMethodArgs() interface{} {
	return greeting.NewGreetingServiceGreetingMethodArgs()
}

func newGreetingServiceGreetingMethodResult() interface{} {
	return greeting.NewGreetingServiceGreetingMethodResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) GreetingMethod(ctx context.Context, request *greeting.GreetingRequest) (r *greeting.GreetingResponse, err error) {
	var _args greeting.GreetingServiceGreetingMethodArgs
	_args.Request = request
	var _result greeting.GreetingServiceGreetingMethodResult
	if err = p.c.Call(ctx, "greetingMethod", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
