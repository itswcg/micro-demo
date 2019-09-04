package main

import (
	"context"
	pb "github.com/itswcg/micro-demo/consignment-srv/proto/consignment"
	vesselPb "github.com/itswcg/micro-demo/vessel-srv/proto/vessel"
	"github.com/micro/go-micro"
	"log"
)

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
	repo         Repository
	vesselClient vesselPb.VesselService
}

func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, resp *pb.Response) error {
	vReq := &vesselPb.Specification{
		Capacity:  int32(len(req.Containers)),
		MaxWeight: req.Weight,
	}

	vResp, err := s.vesselClient.FindAvailable(context.Background(), vReq)
	if err != nil {
		return err
	}

	log.Printf("found vessel: %s\n", vResp.Vessel.Name)
	req.VesselId = vResp.Vessel.Id
	consignment, err := s.repo.Create(req)
	if err != nil {
		return err
	}

	resp.Created = true
	resp.Consignment = consignment

	return nil
}

func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest, resp *pb.Response) error {
	consignments := s.repo.GetAll()

	resp.Consignments = consignments
	//log.Println(resp)
	return nil
}

func main() {
	server := micro.NewService(
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)

	server.Init()
	repo := Repository{}
	vClient := vesselPb.NewVesselService("go.micro.srv.vessel", server.Client())
	_ = pb.RegisterShippingServiceHandler(server.Server(), &service{repo, vClient})
	if err := server.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
