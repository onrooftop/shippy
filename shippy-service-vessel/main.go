package main

import (
	"context"
	"log"
	"os"

	"github.com/micro/go-micro/v2"
	pb "github.com/onrooftop/shippy/shippy-service-vessel/proto/vessel"
)

func main() {
	service := micro.NewService(micro.Name("shippy.service.vessel"))
	service.Init()

	uri := os.Getenv("DB_HOST")

	client, err := CreateClient(context.Background(), uri, 0)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.Background())

	vesselCollection := client.Database("shippy").Collection("vessels")
	repository := &MongoRepository{vesselCollection}
	repository.Create(context.Background(), &Vessel{ID: "vessel001", Name: "Boaty McBoatface", MaxWeight: 200000, Capacity: 500})

	h := &handler{repository}

	if err := pb.RegisterVesselServiceHandler(service.Server(), h); err != nil {
		log.Panic(err)
	}

	if err := service.Run(); err != nil {
		log.Panic(err)
	}
}
