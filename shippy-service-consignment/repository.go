package main

import (
	"context"
	"log"

	pb "github.com/onrooftop/shippy/shippy-service-consignment/proto/consignment"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Consignment -
type Consignment struct {
	ID          string       `bson:"id"`
	Description string       `bson:"description"`
	Weight      int32        `bson:"weight"`
	Containers  []*Container `bson:"containers"`
	VesselID    string       `bson:"vessel_id"`
}

// Container -
type Container struct {
	ID         string `bson:"id"`
	CustomerID string `bson:"customer_id"`
	UserID     string `bson:"user_id"`
}

// MarshalContainerCollection -
func MarshalContainerCollection(containers []*pb.Container) []*Container {
	collection := make([]*Container, 0)
	for _, container := range containers {
		collection = append(collection, MarshalContainer(container))
	}
	return collection
}

// MarshalContainer -
func MarshalContainer(container *pb.Container) *Container {
	return &Container{
		ID:         container.Id,
		CustomerID: container.CustomerId,
		UserID:     container.UserId,
	}
}

// MarshalConsignment -
func MarshalConsignment(consignment *pb.Consignment) *Consignment {
	return &Consignment{
		ID:          consignment.Id,
		Containers:  MarshalContainerCollection(consignment.Containers),
		Description: consignment.Description,
		VesselID:    consignment.VesselId,
		Weight:      consignment.Weight,
	}
}

// UnmarshalContainerCollection -
func UnmarshalContainerCollection(containers []*Container) []*pb.Container {
	collection := make([]*pb.Container, 0)
	for _, container := range containers {
		collection = append(collection, UnmarshalContainer(container))
	}
	return collection
}

// UnmarshalContainer -
func UnmarshalContainer(container *Container) *pb.Container {
	return &pb.Container{
		Id:         container.ID,
		CustomerId: container.CustomerID,
		UserId:     container.UserID,
	}
}

// UnmarshalConsignmentCollection -
func UnmarshalConsignmentCollection(consignments []*Consignment) []*pb.Consignment {
	collection := make([]*pb.Consignment, 0)
	for _, consignment := range consignments {
		collection = append(collection, UnmarshalConsignment(consignment))
	}
	return collection
}

// UnmarshalConsignment -
func UnmarshalConsignment(consignment *Consignment) *pb.Consignment {
	return &pb.Consignment{
		Id:          consignment.ID,
		Description: consignment.Description,
		VesselId:    consignment.VesselID,
		Weight:      consignment.Weight,
		Containers:  UnmarshalContainerCollection(consignment.Containers),
	}
}

type repository interface {
	Create(ctx context.Context, consignment *Consignment) error
	GetAll(ctx context.Context) ([]*Consignment, error)
}

// MongoRepository -
type MongoRepository struct {
	collection *mongo.Collection
}

// Create -
func (repository *MongoRepository) Create(ctx context.Context, consignment *Consignment) error {
	_, err := repository.collection.InsertOne(ctx, consignment)
	return err
}

// GetAll -
func (repository *MongoRepository) GetAll(ctx context.Context) ([]*Consignment, error) {
	cur, err := repository.collection.Find(ctx, bson.M{})
	log.Println("GetAll: ", cur, err)
	if err != nil {
		return nil, err
	}
	var consignments []*Consignment
	for cur.Next(ctx) {
		consignment := &Consignment{}
		if err := cur.Decode(consignment); err != nil {
			return nil, err
		}
		consignments = append(consignments, consignment)
	}
	return consignments, nil
}
