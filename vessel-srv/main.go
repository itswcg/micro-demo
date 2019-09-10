package main

import (
	pb "github.com/itswcg/micro-demo/vessel-srv/proto/vessel"
	"github.com/micro/go-micro"
	"log"
	"os"
)

const (
	DEFAULT_HOST = "mongodb://@212.64.50.167:2225"
)

func createDummyData(repo Repository) {
	defer repo.Close()
	vessels := []*pb.Vessel{
		{Id: "vessel001", Name: "Kane's Salty Secret", MaxWeight: 200000, Capacity: 500},
	}
	for _, v := range vessels {
		_ = repo.Create(v)
	}
}

func main() {

	host := os.Getenv("DB_HOST")
	if host == "" {
		host = DEFAULT_HOST
	}

	session, err := CreateSession(host)
	defer session.Close()

	if err != nil {
		log.Fatalf("Error connectiong to datastore: %v", err)
	}

	repo := &VesselRepository{session.Copy()}

	createDummyData(repo)

	server := micro.NewService(
		micro.Name("go.micro.srv.vessel"),
		micro.Version("latest"),
	)
	server.Init()

	_ = pb.RegisterVesselServiceHandler(server.Server(), &handler{session})

	if err := server.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
