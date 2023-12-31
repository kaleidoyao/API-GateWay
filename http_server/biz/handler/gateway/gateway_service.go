// Code generated by hertz generator.

package gateway

import (
	"context"
	"github.com/kaleidoyao/API-GateWay/global"
	"log"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	calculate "github.com/kaleidoyao/API-GateWay/http_server/biz/model/calculate"
	greeting "github.com/kaleidoyao/API-GateWay/http_server/biz/model/greeting"
	reverse "github.com/kaleidoyao/API-GateWay/http_server/biz/model/reverse"
)

// CalculateMethod .
// @router /calculate [POST]
func CalculateMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req calculate.CalculateRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	clientPool := global.ClientPool{}
	calculateClient, err := clientPool.GetClient("CalculateService")
	if err != nil {
		log.Fatal(err)
	}

	requestRpc := &calculate.CalculateRequest{
		Operand1: req.Operand1,
		Operand2: req.Operand2,
	}

	var responseRpc calculate.CalculateResponse
	err = global.SendRpcRequest(ctx, calculateClient, "calculateMethod", requestRpc, &responseRpc)
	if err != nil {
		log.Fatal(err)
	}

	resp := &calculate.CalculateResponse{
		Outcome: responseRpc.Outcome,
	}

	c.JSON(consts.StatusOK, resp)
}

// GreetingMethod .
// @router /greeting [GET]
func GreetingMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req greeting.GreetingRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	clientPool := global.ClientPool{}
	greetingClient, err := clientPool.GetClient("GreetingService")
	if err != nil {
		log.Fatal(err)
	}

	requestRpc := &greeting.GreetingRequest{
		Name: req.Name,
	}

	var responseRpc greeting.GreetingResponse
	err = global.SendRpcRequest(ctx, greetingClient, "greetingMethod", requestRpc, &responseRpc)
	if err != nil {
		log.Fatal(err)
	}

	resp := &greeting.GreetingResponse{
		ResponseBody: responseRpc.ResponseBody,
	}

	c.JSON(consts.StatusOK, resp)
}

// ReverseMethod .
// @router /reverse [POST]
func ReverseMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req reverse.ReverseRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	clientPool := global.ClientPool{}
	reverseClient, err := clientPool.GetClient("ReverseService")
	if err != nil {
		log.Fatal(err)
	}

	requestRpc := &reverse.ReverseRequest{
		InputString: req.InputString,
	}

	var responseRpc reverse.ReverseResponse
	err = global.SendRpcRequest(ctx, reverseClient, "reverseMethod", requestRpc, &responseRpc)
	if err != nil {
		log.Fatal(err)
	}

	resp := &reverse.ReverseResponse{
		OutputString: responseRpc.OutputString,
	}

	c.JSON(consts.StatusOK, resp)
}
