package subscriber

import (
	"context"
	"github.com/micro/go-micro/util/log"

	micro "github.com/itswcg/micro-demo/proto/micro"
)

type Micro struct{}

func (e *Micro) Handle(ctx context.Context, msg *micro.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *micro.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
