package main

import (
	greeting "github.com/kaleidoyao/rpc_server_greeting/kitex_gen/greeting/greetingservice"
	"log"
)

func main() {
	svr := greeting.NewServer(new(GreetingServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
