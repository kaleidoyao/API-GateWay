// Code generated by Kitex v0.6.1. DO NOT EDIT.

package reverseservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	reverse "github.com/kaleidoyao/API-GateWay/rpc_server_reverse/kitex_gen/reverse"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	ReverseMethod(ctx context.Context, request *reverse.ReverseRequest, callOptions ...callopt.Option) (r *reverse.ReverseResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kReverseServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kReverseServiceClient struct {
	*kClient
}

func (p *kReverseServiceClient) ReverseMethod(ctx context.Context, request *reverse.ReverseRequest, callOptions ...callopt.Option) (r *reverse.ReverseResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ReverseMethod(ctx, request)
}
