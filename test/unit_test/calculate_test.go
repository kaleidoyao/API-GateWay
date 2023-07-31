package unit_test

import (
	"context"
	"testing"

	calculate "github.com/kaleidoyao/API-GateWay/rpc_server_calculate/kitex_gen/calculate"
	handler "github.com/kaleidoyao/API-GateWay/test/unit_test/utils"
)

func Test_calculateCase1(t *testing.T) {
	operand_1 := int32(1)
	operand_2 := int32(2)
	outcome := int32(3)
	csi := new(handler.CalculateServiceImpl)
	req := &calculate.CalculateRequest{
		Operand_1: operand_1,
		Operand_2: operand_2,
	}
	resp, err := csi.CalculateMethod(context.Background(), req)
	if err != nil {
		t.Error(err)
	}
	if resp.Outcome != outcome {
		t.Errorf("expected resp is %d but resp in fact is %d", outcome, resp.Outcome)
	}
}

func Test_calculateCase2(t *testing.T) {
	operand_1 := int32(2147483647)
	operand_2 := int32(0)
	outcome := int32(2147483647)
	csi := new(handler.CalculateServiceImpl)
	req := &calculate.CalculateRequest{
		Operand_1: operand_1,
		Operand_2: operand_2,
	}
	resp, err := csi.CalculateMethod(context.Background(), req)
	if err != nil {
		t.Error(err)
	}
	if resp.Outcome != outcome {
		t.Errorf("expected resp is %d but resp in fact is %d", outcome, resp.Outcome)
	}
}

func Test_calculateCase3(t *testing.T) {
	operand_1 := int32(-2147483648)
	operand_2 := int32(0)
	outcome := int32(-2147483648)
	csi := new(handler.CalculateServiceImpl)
	req := &calculate.CalculateRequest{
		Operand_1: operand_1,
		Operand_2: operand_2,
	}
	resp, err := csi.CalculateMethod(context.Background(), req)
	if err != nil {
		t.Error(err)
	}
	if resp.Outcome != outcome {
		t.Errorf("expected resp is %d but resp in fact is %d", outcome, resp.Outcome)
	}
}
