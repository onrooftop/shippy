package main

import (
	"context"
	"log"

	pb "github.com/onrooftop/shippy/shippy-service-vessel/proto/vessel"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Vessel -
type Vessel struct {
	ID        string `bson:"id"`
	Capacity  int32  `bson:"capacity"`
	MaxWeight int32  `bson:"maxWeight"`
	Name      string `bson:"name"`
	Available bool   `bson:"available"`
	OwnerID   string `bson:"ownerID "`
}

// Specification -
type Specification struct {
	Capacity  int32 `bson:"capacity"`
	MaxWeight int32 `bson:"maxWeight"`
}

// UnmarshalSpecification -
func UnmarshalSpecification(spec *Specification) *pb.Specification {
	return &pb.Specification{
		Capacity:  spec.Capacity,
		MaxWeight: spec.MaxWeight,
	}
}

// UnmarshalVessel -
func UnmarshalVessel(vessel *Vessel) *pb.Vessel {
	return &pb.Vessel{
		Available: vessel.Available,
		Capacity:  vessel.Capacity,
		Id:        vessel.ID,
		MaxWeight: vessel.MaxWeight,
		Name:      vessel.Name,
		OwnerId:   vessel.OwnerID,
	}
}

// MarshalSpecification -
func MarshalSpecification(spec *pb.Specification) *Specification {
	return &Specification{
		Capacity:  spec.Capacity,
		MaxWeight: spec.MaxWeight,
	}
}

// MarshalVessel -
func MarshalVessel(vessel *pb.Vessel) *Vessel {
	return &Vessel{
		ID:        vessel.Id,
		Available: vessel.Available,
		Capacity:  vessel.Capacity,
		MaxWeight: vessel.MaxWeight,
		Name:      vessel.Name,
		OwnerID:   vessel.OwnerId,
	}
}

// repository -
type repository interface {
	FindAvailable(ctx context.Context, spec *Specification) (*Vessel, error)
	Create(ctx context.Context, vessel *Vessel) error
}

// MongoRepository -
type MongoRepository struct {
	collection *mongo.Collection
}

// FindAvailable -
func (repository *MongoRepository) FindAvailable(ctx context.Context, spec *Specification) (*Vessel, error) {
	filter := bson.D{
		{
			"capacity", bson.D{
				{"$gte", spec.Capacity},
			},
		},
	}

	log.Println("FindAvailable: ", spec, filter)
	vessel := &Vessel{}
	if err := repository.collection.FindOne(ctx, filter).Decode(vessel); err != nil {
		return nil, err
	}
	return vessel, nil
}

// Create -
func (repository *MongoRepository) Create(ctx context.Context, vessel *Vessel) error {
	_, err := repository.collection.InsertOne(ctx, vessel)
	log.Println("Created", err, vessel)
	return err
}
