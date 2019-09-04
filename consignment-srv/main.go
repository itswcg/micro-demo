package main

import (
	"context"
	pb "github.com/itswcg/micro-demo/consignment-srv/proto/consignment"
	"github.com/micro/go-micro"
	"log"
)

type IRepository interface {
	Create(consignment *pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

type Repository struct {
	consignments []*pb.Consignment
}

func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	updated := append(repo.consignments, consignment)
	repo.consignments = updated
	return consignment, nil
}

func (repo *Repository) GetAll() []*pb.Consignment {
	return repo.consignments
}

type service struct {
	repo Repository
}

func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, resp *pb.Response) error {
	consignment, err := s.repo.Create(req)
	if err != nil {
		return err
	}

	resp = &pb.Response{}
	resp.Created = true
	resp.Consignment = consignment
	return nil
}

func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest, resp *pb.Response) error {
	consignments := s.repo.GetAll()
	resp = &pb.Response{}
	resp.Consignments = consignments
	return nil
}

func main() {
	server := micro.NewService(
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)

	server.Init()
	repo := Repository{}
	_ = pb.RegisterShippingServiceHandler(server.Server(), &service{repo})
	if err := server.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
