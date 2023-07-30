# API-GateWay

# 1. 项目背景

# 2. 项目概览

# 3. 开发指南
## 3.1 开发准备

```
hz new -mod github.com/kaleidoyao/http_server -idl ../idl/gateway.thrift
```

```
kitex -module github.com/kaleidoyao/rpc_server_calculate -service calculate ../idl/calculate.thrift
```

```
kitex -module github.com/kaleidoyao/rpc_server_greeting -service greeting ../idl/greeting.thrift
```

```
kitex -module github.com/kaleidoyao/rpc_server_reverse -service reverse ../idl/reverse.thrift
```