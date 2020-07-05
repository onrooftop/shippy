package main

import (
	"context"

	pb "github.com/onrooftop/shippy/shippy-service-vessel/proto/vessel"
)

type handler struct {
	repository
}

func (h *handler) FindAvailable(ctx context.Context, req *pb.Specification, res *pb.Response) error {
	vessel, err := h.repository.FindAvailable(ctx, MarshalSpecification(req))
	if err != nil {
		return err
	}

	res.Vessel = UnmarshalVessel(vessel)
	return nil
}
func (h *handler) Create(ctx context.Context, req *pb.Vessel, res *pb.Response) error {
	if err := h.repository.Create(ctx, MarshalVessel(req)); err != nil {
		return err
	}
	res.Vessel = req
	return nil
}
