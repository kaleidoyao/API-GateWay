package main

import (
	reverse "github.com/kaleidoyao/rpc_server_reverse/kitex_gen/reverse/reverseservice"
	"log"
)

func main() {
	svr := reverse.NewServer(new(ReverseServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
