module github.com/onrooftop/shippy/shippy-service-consignment

go 1.13

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/golang/protobuf v1.4.2
	github.com/micro/go-micro/v2 v2.9.1
	github.com/onrooftop/shippy/shippy-service-vessel v0.0.0-20200710165438-08c52678b68a
	github.com/stretchr/testify v1.5.1 // indirect
	go.mongodb.org/mongo-driver v1.3.5
	golang.org/x/crypto v0.0.0-20200709230013-948cd5f35899 // indirect
	google.golang.org/genproto v0.0.0-20200710124503-20a17af7bd0e // indirect
	google.golang.org/protobuf v1.25.0
)
