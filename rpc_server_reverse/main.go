package main

import (
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	reverse "github.com/kaleidoyao/API-GateWay/rpc_server_reverse/kitex_gen/reverse/reverseservice"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"net"
)

func main() {
	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}

	svr := reverse.NewServer(
		new(ReverseServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "ReverseService"}),
		server.WithRegistry(r),
		server.WithServiceAddr(&net.TCPAddr{Port: 9991}),
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
