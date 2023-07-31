package main

import (
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	greeting "github.com/kaleidoyao/API-GateWay/rpc_server_greeting/kitex_gen/greeting/greetingservice"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"net"
)

func main() {
	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}

	svr := greeting.NewServer(
		new(GreetingServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "GreetingService"}),
		server.WithRegistry(r),
		server.WithServiceAddr(&net.TCPAddr{Port: 9991}),
    )

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
