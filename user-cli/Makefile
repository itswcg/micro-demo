.PHONY: build
build:

	CGO_ENABLED=0 GOOS=linux go build
	docker build -t user-cli .

.PHONY: run
run:

	docker run -e MICRO_REGISTRY=mdns user-cli
