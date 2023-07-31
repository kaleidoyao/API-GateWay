# API-GateWay

# 1. 项目背景
随着微服务架构的流行，许多企业开始将其业务拆分成多个小型服务。这种架构可以提高开发效率、扩展性和灵活性。然而，随着服务数量的增加，也带来了新的挑战，如请求转发、认证、授权、限流、监控等问题。为了有效地管理和保护这些后端服务，我们决定开发一个高性能、可扩展的API网关。

# 2. 项目概览
## 2.1 项目概述
API网关是一个用于转发请求、认证、授权、限流、监控等功能的中间层服务。本项目使用GO语言开发，目标是为了集中管理和保护后端服务，并提供更好的开发和维护体验。

本项目实现了API网关的基本功能，包括服务注册、服务发现、路由服务、负载均衡等功能。

- 接受和响应HTTP请求：API网关作为中间层，能够接受来自客户端的HTTP请求，并将这些请求转发到后端的服务。同时，API网关还能够接收后端服务返回的响应，并将其转发回客户端，实现客户端和后端服务之间的通信。
- 服务注册和服务发现：本项目使用etcd作为服务注册和服务发现的组件。
- 构造Kitex泛化调用客户端：本项目实现了管理多个由Kitex生成的RPC服务所需要的处理逻辑。API网关采用Kitex泛化调用客户端，实现了对后端服务的泛化调用。这意味着API网关可以以一种通用的方式调用不同后端服务的不同API接口，而无需事先了解后端服务的具体接口定义。通过泛化调用，API网关可以动态地构造请求参数和解析响应，实现灵活的请求转发和响应处理。
- Client池化：为了提高性能和资源利用率，本项目实现了Client池化功能。当收到请求时，API网关从Client池中获取一个可用的Client，使用它来转发请求到后端服务。通过重用Client，API网关避免了频繁创建和销毁连接的开销，提高了处理请求的效率。

## 2.2 技术栈
- **go**：Go语言是一种由Google开发的开源编程语言。它具有高性能、强大的并发支持、垃圾回收机制等特点，适用于构建高性能的网络服务和分布式系统。
- **etcd**：etcd是一个分布式的键值存储系统，由CoreOS开发。它基于Raft算法实现了强一致性，并提供了高可用性和可靠的数据存储服务。etcd通常用于服务发现、配置管理和分布式锁等场景，是构建分布式系统的重要基础组件。
- **Hertz**：Hertz[həːts]是一个Golang微服务HTTP框架，在设计之初参考了其他开源框架fasthttp、gin、echo的优势，并结合字节跳动内部的需求，使其具有高易用性、高性能、高扩展性等特点，目前在字节跳动内部已广泛使用。如今越来越多的微服务选择使用Golang，如果对微服务性能有要求，又希望框架能够充分满足内部的可定制化需求，Hertz会是一个不错的选择。
- **Kitex**：Kitex[kaɪt’eks]字节跳动内部的Golang微服务RPC框架，具有高性能、强可扩展的特点，在字节内部已广泛使用。如果对微服务性能有要求，又希望定制扩展融入自己的治理体系，Kitex会是一个不错的选择。

## 2.3 团队成员
|学号|姓名|
|:-:|:-:|
|211250029|姚周珩|
|211250030|石锐婷|
|211250083|周鑫|

# 3. 开发指南
## 3.1 项目构建
### 3.1.1 初始化项目
```
go mod init github.com/kaleidoyao/API-GateWay
```
### 3.1.2 IDL文件编写
``` thrift
// reverse.thrift
struct ReverseRequest {
    1: required string inputString;
}

struct ReverseResponse {
    1: required string outputString;
}

service ReverseService {
    ReverseResponse reverseMethod(1: ReverseRequest request) (api.post="/reverse")
}
```
``` thrift
include "calculate.thrift"
include "greeting.thrift"
include "reverse.thrift"

service GatewayService {
    calculate.CalculateResponse calculateMethod(1: calculate.CalculateRequest request) (api.post="/calculate");
    greeting.GreetingResponse greetingMethod(1: greeting.GreetingRequest request) (api.get="/greeting")
    reverse.ReverseResponse reverseMethod(1: reverse.ReverseRequest request) (api.post="/reverse")
}
```
### 3.1.3 生成RPC_SERVER
```
kitex -module github.com/kaleidoyao/API-GateWay -service calculate ../idl/calculate.thrift
```
```
kitex -module github.com/kaleidoyao/API-GateWay -service greeting ../idl/greeting.thrift
```
```
kitex -module github.com/kaleidoyao/API-GateWay -service reverse ../idl/reverse.thrift
```
修改`rpc_server_reverse`中的`handler.go`，使其能够处理相应的业务逻辑。
``` go
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
```
修改`rpc_server_reverse`中的`main.go`，使其将服务注册到相应端口。
``` go
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
```
### 3.1.4 生成HTTP_SERVER
```
hz new -idl ../idl/gateway.thrift
```
处理HTTP请求并与远程RPC服务进行通信。
``` go
func ReverseMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req reverse.ReverseRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	clientPool := global.ClientPool{}
	reverseClient, err := clientPool.GetClient("ReverseService")
	if err != nil {
		log.Fatal(err)
	}

	requestRpc := &reverse.ReverseRequest{
		InputString: req.InputString,
	}

	var responseRpc reverse.ReverseResponse
	err = global.SendRpcRequest(ctx, reverseClient, "ReverseMethod", requestRpc, &responseRpc)
	if err != nil {
		log.Fatal(err)
	}

	resp := &reverse.ReverseResponse{
		OutputString: responseRpc.OutputString,
	}

	c.JSON(consts.StatusOK, resp)
}
```

## 3.2 项目结构
```
.
├── default.etcd
├── global
├── go.mod
├── go.sum
├── http_server
├── idl
├── README.md
├── rpc_server_calculate
├── rpc_server_greeting
└── rpc_server_reverse
```

## 3.3 项目运行
### 3.3.1 启动etcd注册中心
```
etcd --log-level debug
```
### 3.3.2 启动http_server
in directory `http_server`
```
go build
./http_server
```
### 3.3.3 启动rpc_server
- in directory `rpc_server_calculate`
```
sh ./build.sh && sh ./output/bootstrap.sh
```
- in directory `rpc_server_greeting`
```
sh ./build.sh && sh ./output/bootstrap.sh
```
- in directory `rpc_server_reverse`
```
sh ./build.sh && sh ./output/bootstrap.sh
```
> 执行：`sh ./build.sh && sh ./output/bootstrap.sh`<br/>
> 报错：`error obtaining VCS status: exit status 128  Use -buildvcs=false to disable VCS stamping.`<br/>
> 解决方法：`go env -w GOFLAGS=-buildvcs=false`
#### 3.3.4 发送HTTP请求
- CalculateService
```
curl --location 'http://127.0.0.1:8888/calculate' \
--header 'Content-Type: application/json' \
--data '{
    "operand_1": 1, 
    "operand_2": 2
}'
```
```
{"outcome":3}
```
- GreetingService
```
curl 127.0.0.1:8888/greeting?name=TestName
```
```
{"ResponseBody":"Hello, TestName"}
```
- ReverseService
```
curl --location 'http://127.0.0.1:8888/reverse' \
--header 'Content-Type: application/json' \
--data '{
    "inputString": "ThisIsATestString"
}'
```
```
{"outputString":"gnirtStseTAsIsihT"}
```
