.PHONY: build
build:

	CGO_ENABLED=0 GOOS=linux go build
	docker build -t consignment-cli .

.PHONY: run
run:

	docker run -e MICRO_REGISTRY=mdns consignment-cli
