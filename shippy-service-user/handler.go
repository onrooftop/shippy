package main

import (
	"context"

	pb "github.com/onrooftop/shippy/shippy-service-user/proto/user"
)

type handler struct {
	repository
}

// Create -
func (h *handler) Create(ctx context.Context, req *pb.User, res *pb.Response) error {
	err := h.repository.Create(ctx, MarshalUser(req))
	if err != nil {
		return err
	}
	req.Password = ""
	res.User = req
	return nil
}

// Get -
func (h *handler) Get(ctx context.Context, req *pb.User, res *pb.Response) error {
	result, err := h.repository.Get(ctx, req.Id)
	if err != nil {
		return err
	}

	user := UnmarshalUser(result)
	res.User = user
	return nil
}

// GetAll -
func (h *handler) GetAll(ctx context.Context, req *pb.Request, res *pb.Response) error {
	result, err := h.repository.GetAll(ctx)
	if err != nil {
		return err
	}

	users := UnmarshalUserCollection(result)
	res.Users = users
	return nil
}

func (h *handler) Auth(ctx context.Context, req *pb.User, res *pb.Token) error {
	return nil
}

func (h *handler) ValidateToken(ctx context.Context, req *pb.Token, res *pb.Token) error {
	return nil
}
