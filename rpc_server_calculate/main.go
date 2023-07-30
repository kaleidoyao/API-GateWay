package main

import (
	calculate "github.com/kaleidoyao/rpc_server_calculate/kitex_gen/calculate/calculateservice"
	"log"
)

func main() {
	svr := calculate.NewServer(new(CalculateServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
