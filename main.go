package main

import (
	server "my-go-project/grpc_server"
	protos "my-go-project/protos/currency"

	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
)

func main() {
	log := hclog.Default()

	grpcServer := grpc.NewServer()

	protos.RegisterCurrencyServiceServer(grpcServer, &server.CurrencyService{})
}
