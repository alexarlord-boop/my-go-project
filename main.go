package main

import (
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
)

func main() {
	log := hclog.Default()

	grpcServer := grpc.NewServer()

	protos.RegisterCurrencyServiceServer(grpcServer, &CurrencyService{})
}
