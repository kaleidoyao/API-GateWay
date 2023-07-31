package unit_test

import (
	"context"
	"testing"

	reverse "github.com/kaleidoyao/API-GateWay/rpc_server_reverse/kitex_gen/reverse"
	handler "github.com/kaleidoyao/API-GateWay/test/unit_test/utils"
)

func Test_reverseCase1(t *testing.T) {
	testString := "kjadkjfkjdksjfdnskdj"
	expectedResp := "jdksndfjskdjkfjkdajk"
	csi := new(handler.ReverseServiceImpl)
	req := &reverse.ReverseRequest{
		InputString: testString,
	}
	resp, err := csi.ReverseMethod(context.Background(), req)
	if err != nil {
		t.Error(err)
	}
	if resp.OutputString != expectedResp {
		t.Errorf("expected resp is %s but resp in fact is %s", expectedResp, resp.OutputString)
	}
}

func Test_reverseCase2(t *testing.T) {
	testString := "a"
	expectedResp := "a"
	csi := new(handler.ReverseServiceImpl)
	req := &reverse.ReverseRequest{
		InputString: testString,
	}
	resp, err := csi.ReverseMethod(context.Background(), req)
	if err != nil {
		t.Error(err)
	}
	if resp.OutputString != expectedResp {
		t.Errorf("expected resp is %s but resp in fact is %s", expectedResp, resp.OutputString)
	}
}

func Test_reverseCase3(t *testing.T) {
	testString := ""
	expectedResp := ""
	csi := new(handler.ReverseServiceImpl)
	req := &reverse.ReverseRequest{
		InputString: testString,
	}
	resp, err := csi.ReverseMethod(context.Background(), req)
	if err != nil {
		t.Error(err)
	}
	if resp.OutputString != expectedResp {
		t.Errorf("expected resp is %s but resp in fact is %s", expectedResp, resp.OutputString)
	}
}
