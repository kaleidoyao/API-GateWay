package global

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/json"
	kClient "github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/pkg/loadbalance"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"os"
)

func GenerateClient(serviceName string) (genericclient.Client, error) {
	var err error

	lb := loadbalance.NewWeightedBalancer() // 初始化负载均衡

	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"}) // 初始化etcd解析器
	if err != nil {
		log.Fatal("Error: fail to new etcd resolver...")
	}

	wd, _ := os.Getwd() // 获取当前工作路径
	p, err := generic.NewThriftFileProvider(wd + "/../idl/gateway.thrift")
	if err != nil {
		log.Fatal("Error: fail to new thrift file provider...")
	}

	g, err := generic.JSONThriftGeneric(p)
	if err != nil {
		log.Fatal("Error: fail to convert into generic...")
	}

	client, err := genericclient.NewClient(
		serviceName,
		g,
		kClient.WithResolver(r),
		kClient.WithLoadBalancer(lb),
	)
	if err != nil {
		log.Fatal("Error: fail to new generic client...")
	}

	return client, nil
}

func SendRpcRequest(ctx context.Context, genericClient genericclient.Client, methodName string, request interface{}, response interface{}) error {
	strRequest, err := MakeJsonString(request)
	if err != nil {
		log.Fatal(err)
	}

	responseRpc, err := genericClient.GenericCall(ctx, methodName, strRequest)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal([]byte(responseRpc.(string)), response)
	if err != nil {
		return err
	}

	return nil
}

func MakeJsonString(info any) (string, error) {
	jsonData, err := json.Marshal(&info)
	if err != nil {
		log.Fatal(err)
	}
	return string(jsonData), nil
}
