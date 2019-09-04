# micro-demo

This is the Micro service

Generated with

```
micro new github.com/itswcg/micro-demo --namespace=go.micro --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.srv.micro
- Type: srv
- Alias: micro

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend consul.

```
# install consul
wget https://releases.hashicorp.com/consul/1.6.0/consul_1.6.0_linux_amd64.zip
unzip consul_1.6.0_linux_amd64.zip
cp consul /usr/bin
# server and client
consul agent -dev -config-dir /etc/consul.d
# client

```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./micro-srv
```

Build a docker image
```
make docker
```

## Micro command
```
# 所有服务
micro list services
# 单个服务
micro get service
# 调用服务 
micro call
# 执行api
micro api
# 交互
micro cli
> micro health
> micro register service
> micro deregister service
# 代理
micro proxy
```

package
```
# 本地打包
go build -i -o micro ./main.go ./plugins.go

# 打包成docker镜像
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' -i -o micro ./main.go ./plugins.go
```

```
go get -u google.golang.org/grpc
go get -u github.com/golang/protobuf/protoc-gen-go

https://yinzige.com/2018/05/10/microservices-part-1-introduction-and-consignment-service/
http://www.ruanyifeng.com/blog/2016/07/google-monolithic-source-repository.html

protoc --proto_path=. --micro_out=. --go_out=. proto/consignment/consignment.proto
--proto_path= -I 表示指定搜索的目录
--micro_out=输出目录
--go_out=输出目录
表示所在proto文件
```
