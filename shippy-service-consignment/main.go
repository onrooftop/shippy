// shippy-service-consignment/main.go
package main

import (
	"context"
	"log"
	"os"

	"github.com/micro/go-micro/v2"
	pb "github.com/onrooftop/shippy/shippy-service-consignment/proto/consignment"
	vesselPb "github.com/onrooftop/shippy/shippy-service-vessel/proto/vessel"
)

const (
	defaultHost = "datastore:27017"
)

func main() {
	service := micro.NewService(
		micro.Name("shippy.service.consignment"),
	)
	service.Init()

	uri := os.Getenv("DB_HOST")

	client, err := CreateClient(context.Background(), uri, 0)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.Background())
	consignmentCollection := client.Database("shippy").Collection("consignments")

	repository := &MongoRepository{consignmentCollection}
	vesselClient := vesselPb.NewVesselService("shippy.service.vessel", service.Client())

	h := &handler{repository, vesselClient}
	if err := pb.RegisterShippingServiceHandler(service.Server(), h); err != nil {
		log.Panic(err)
	}

	if err := service.Run(); err != nil {
		log.Panic(err)
	}
}
