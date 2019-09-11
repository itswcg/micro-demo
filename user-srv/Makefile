GOPATH:=$(shell go env GOPATH)

.PHONY: proto
proto:

	protoc --proto_path=. --micro_out=. --go_out=. proto/user/user.proto

.PHONY: build
build:

	CGO_ENABLED=0 GOOS=linux go build
	docker build -t user-srv .

.PHONY: run
run:

	docker run -p 50053:50051 -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=mdns consignment-srv
