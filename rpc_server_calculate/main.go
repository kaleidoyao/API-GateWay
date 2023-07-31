package main

import (
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	calculate "github.com/kaleidoyao/API-GateWay/rpc_server_calculate/kitex_gen/calculate/calculateservice"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"net"
)

func main() {
	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}

	svr := calculate.NewServer(
		new(CalculateServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "CalculateService"}),
		server.WithRegistry(r),
		server.WithServiceAddr(&net.TCPAddr{Port: 9990}),
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
