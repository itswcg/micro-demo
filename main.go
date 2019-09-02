package main

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	"github.com/itswcg/micro-demo/handler"
	"github.com/itswcg/micro-demo/subscriber"

	micro "github.com/itswcg/micro-demo/proto/micro"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.micro"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	micro.RegisterMicroHandler(service.Server(), new(handler.Micro))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.micro", service.Server(), new(subscriber.Micro))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.micro", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
