package main

import (
	pb "github.com/itswcg/micro-demo/consignment-srv/proto/consignment"
	vesselPb "github.com/itswcg/micro-demo/vessel-srv/proto/vessel"
	"github.com/micro/go-micro"
	"log"
	"os"
)

const (
	DEFAULT_HOST = "mongodb://@212.64.50.167:2225"
)

func main() {
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = DEFAULT_HOST
	}

	session, err := CreateSession(dbHost)
	defer session.Close()
	if err != nil {
		log.Fatalf("create session error:%v\n", err)
	}

	server := micro.NewService(
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)

	server.Init()

	vClient := vesselPb.NewVesselService("go.micro.srv.vessel", server.Client())
	_ = pb.RegisterShippingServiceHandler(server.Server(), &handler{session, vClient})

	if err := server.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
