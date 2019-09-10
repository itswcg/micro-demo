package main

import (
	pb "github.com/itswcg/micro-demo/vessel-srv/proto/vessel"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	DB_NAME        = "shippy"
	CON_COLLECTION = "vessels"
)

type Repository interface {
	Create(vessel *pb.Vessel) error
	FindAvailable(specification *pb.Specification) (*pb.Vessel, error)
	Close()
}

type VesselRepository struct {
	session *mgo.Session
}

func (repo *VesselRepository) collection() *mgo.Collection {
	return repo.session.DB(DB_NAME).C(CON_COLLECTION)
}

func (repo *VesselRepository) Create(v *pb.Vessel) error {
	return repo.collection().Insert(v)
}

func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	var vessel *pb.Vessel
	err := repo.collection().Find(bson.M{
		"capacity":  bson.M{"$gte": spec.Capacity},
		"maxweight": bson.M{"$gte": spec.MaxWeight},
	}).One(&vessel)
	if err != nil {
		return nil, err
	}

	return vessel, nil
}

func (repo *VesselRepository) Close() {
	repo.session.Close()
}
