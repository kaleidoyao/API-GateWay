package unit_test

import (
	"context"
	"testing"

	greeting "github.com/kaleidoyao/API-GateWay/rpc_server_greeting/kitex_gen/greeting"
	handler "github.com/kaleidoyao/API-GateWay/test/unit_test/utils"
)

func Test_greetingCase1(t *testing.T) {
	testString := "Barbie"
	expectedResp := "Hello,Barbie"
	csi := new(handler.GreetingServiceImpl)
	req := &greeting.GreetingRequest{
		Name: testString,
	}
	resp, err := csi.GreetingMethod(context.Background(), req)
	if err != nil {
		t.Error(err)
	}
	if resp.ResponseBody != expectedResp {
		t.Errorf("expected resp is %s but resp in fact is %s", expectedResp, resp.ResponseBody)
	}
}
