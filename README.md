# API-GateWay

# 1. 项目背景

# 2. 项目概览

# 3. 开发指南
## 3.1 开发准备

```
hz new -idl ../idl/gateway.thrift
```

```
kitex -module github.com/kaleidoyao/API-GateWay -service calculate ../idl/calculate.thrift
```

```
kitex -module github.com/kaleidoyao/API-GateWay -service greeting ../idl/greeting.thrift
```

```
kitex -module github.com/kaleidoyao/API-GateWay -service reverse ../idl/reverse.thrift
```

启动etcd注册中心
etcd --log-level debug


执行：`sh ./build.sh && sh ./output/bootstrap.sh`
报错：`error obtaining VCS status: exit status 128  Use -buildvcs=false to disable VCS stamping.`
解决方法：`go env -w GOFLAGS=-buildvcs=false`
