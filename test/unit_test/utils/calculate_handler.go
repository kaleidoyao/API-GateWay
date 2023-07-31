package utils

import (
	"context"

	calculate "github.com/kaleidoyao/API-GateWay/rpc_server_calculate/kitex_gen/calculate"
)

// CalculateServiceImpl implements the last service interface defined in the IDL.
type CalculateServiceImpl struct{}

// CalculateMethod implements the CalculateServiceImpl interface.
func (s *CalculateServiceImpl) CalculateMethod(ctx context.Context, request *calculate.CalculateRequest) (resp *calculate.CalculateResponse, err error) {
	op1 := request.Operand_1
	op2 := request.Operand_2

	resp = &calculate.CalculateResponse{
		Outcome: op1 + op2,
	}
	return
}
