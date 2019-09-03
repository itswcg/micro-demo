GOPATH:=$(shell go env GOPATH)

.PHONY: proto
proto:

	protoc -I. --go_out=plugins=grpc:$(GOPATH)/src/github.com/itswcg/micro-demo/consignment-srv \
    proto/consignment/consignment.proto

.PHONY: build
build:

	docker build -t consignment-srv .

.PHONY: run
run:

	docker run -p 50051:50051 -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=mdns consignment-srv
