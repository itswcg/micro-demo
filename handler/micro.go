package handler

import (
	"context"

	"github.com/micro/go-micro/util/log"

	micro "github.com/itswcg/micro-demo/proto/micro"
)

type Micro struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Micro) Call(ctx context.Context, req *micro.Request, rsp *micro.Response) error {
	log.Log("Received Micro.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Micro) Stream(ctx context.Context, req *micro.StreamingRequest, stream micro.Micro_StreamStream) error {
	log.Logf("Received Micro.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&micro.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Micro) PingPong(ctx context.Context, stream micro.Micro_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&micro.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
