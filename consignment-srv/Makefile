GOPATH:=$(shell go env GOPATH)

.PHONY: proto
proto:

	protoc --proto_path=. --micro_out=. --go_out=. proto/consignment/consignment.proto

.PHONY: build
build:

	CGO_ENABLED=0 GOOS=linux go build
	docker build -t consignment-srv .

.PHONY: run
run:

	docker run -p 50051:50051 -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=mdns consignment-srv
