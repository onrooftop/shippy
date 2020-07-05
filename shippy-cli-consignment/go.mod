module github.com/onrooftop/shippy/shippy-cli-consignment

go 1.13

// replace github.com/onrooftop/shippy/shippy-service-consignment => ../shippy-service-consignment

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/micro/go-micro/v2 v2.9.1
	github.com/onrooftop/shippy/shippy-service-consignment v0.0.0-20200705163235-b56576c0909e
	google.golang.org/grpc/examples v0.0.0-20200630190442-3de8449f8555 // indirect
)
