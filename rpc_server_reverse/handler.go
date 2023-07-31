package main

import (
	"context"
	reverse "github.com/kaleidoyao/API-GateWay/rpc_server_reverse/kitex_gen/reverse"
)

// ReverseServiceImpl implements the last service interface defined in the IDL.
type ReverseServiceImpl struct{}

// ReverseMethod implements the ReverseServiceImpl interface.
func (s *ReverseServiceImpl) ReverseMethod(ctx context.Context, request *reverse.ReverseRequest) (resp *reverse.ReverseResponse, err error) {
	resp = &reverse.ReverseResponse{
		OutputString: reverseString(request.InputString),
	}
	return
}

func reverseString(s string) string {
	runes := []rune(s) // 将字符串转换为一个rune切片，rune表示UTF-8字符
	n := len(runes)

	for i, j := 0, n-1; i < j; i, j = i+1, j-1 { // // 使用双指针法反转字符串
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes) // 将rune切片转换回字符串
}
