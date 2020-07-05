package main

import (
	"context"
	"log"

	pb "github.com/onrooftop/shippy/shippy-service-consignment/proto/consignment"
	vesselPb "github.com/onrooftop/shippy/shippy-service-vessel/proto/vessel"
)

type handler struct {
	repository
	vesselClient vesselPb.VesselService
}

func (h *handler) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	vesselResponse, err := h.vesselClient.FindAvailable(context.Background(), &vesselPb.Specification{
		Capacity:  int32(len(req.Containers)),
		MaxWeight: req.Weight,
	})

	log.Println("createConsignment Vess: ", req, vesselResponse)

	if err != nil {
		log.Println("vesselResponseErr: ", err)
		return err
	}

	req.VesselId = vesselResponse.Vessel.Id

	if err := h.repository.Create(ctx, MarshalConsignment(req)); err != nil {
		return err
	}

	res.Created = true
	res.Consignment = req
	return nil
}

func (h *handler) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	consignments, err := h.repository.GetAll(ctx)
	if err != nil {
		return err
	}
	res.Consignments = UnmarshalConsignmentCollection(consignments)
	return nil
}
