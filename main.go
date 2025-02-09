package main

import (
	"fmt"
	server "my-go-project/grpc_server"
	protos "my-go-project/protos/currency"
	"net"
	"os"

	hclog "github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log := hclog.Default()

	grpcServer := grpc.NewServer()
	currencyService := server.NewCurrencyService(log)

	protos.RegisterCurrencyServiceServer(grpcServer, currencyService)

	l, err := net.Listen("tcp", ":9092")
	if err != nil {
		log.Error("Unable to listen", "error", err)
		os.Exit(1)
	}

	reflection.Register(grpcServer)
	fmt.Println("⚡ Reflection Enabled (Development Mode)")
	grpcServer.Serve(l)
}
