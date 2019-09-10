package main

import (
	"context"
	pb "github.com/itswcg/micro-demo/consignment-srv/proto/consignment"
	vesselPb "github.com/itswcg/micro-demo/vessel-srv/proto/vessel"
	"gopkg.in/mgo.v2"
	"log"
)

type handler struct {
	session      *mgo.Session
	vesselClient vesselPb.VesselService
}

func (h *handler) GetRepo() Repository {
	return &ConsignmentRepository{h.session.Clone()}
}

func (h *handler) CreateConsignment(ctx context.Context, req *pb.Consignment, resp *pb.Response) error {
	defer h.GetRepo().Close()

	vReq := &vesselPb.Specification{
		Capacity:  int32(len(req.Containers)),
		MaxWeight: req.Weight,
	}

	vResp, err := h.vesselClient.FindAvailable(context.Background(), vReq)
	if err != nil {
		return err
	}

	log.Printf("found vessel: %s\n", vResp.Vessel.Name)
	req.VesselId = vResp.Vessel.Id
	err = h.GetRepo().Create(req)
	if err != nil {
		return err
	}
	resp.Created = true
	resp.Consignment = req
	return nil
}

func (h *handler) GetConsignments(ctx context.Context, req *pb.GetRequest, resp *pb.Response) error {
	defer h.GetRepo().Close()
	consignments, err := h.GetRepo().GetAll()
	if err != nil {
		return err
	}

	resp.Consignments = consignments
	return nil
}
