package main

import (
	"log"

	"github.com/micro/go-micro/v2"
	pb "github.com/onrooftop/shippy/shippy-service-user/proto/user"
)

const schema = `
	create table if not exists users (
		id varchar(36) not null,
		name varchar(125) not null,
		email varchar(225) not null unique,
		password varchar(225) not null,
		company varchar(125),
		primary key (id)
	);
`

func main() {

	db, err := NewConnection()
	if err != nil {
		log.Panic(err)
	}

	defer db.Close()

	db.MustExec(schema)

	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	repo := NewPostgresRepository(db)

	service := micro.NewService(
		micro.Name("shippy.service.user"),
	)
	service.Init()

	if err := pb.RegisterUserServiceHandler(service.Server(), &handler{repo}); err != nil {
		log.Panic(err)
	}

	if err := service.Run(); err != nil {
		log.Panic(err)
	}

}
