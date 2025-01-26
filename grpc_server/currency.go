package grpcserver

// We want to implement server from grpc generated code.

import (
	"context"
	"log"
	protos "my-go-project/protos/currency"
)

type CurrencyService struct {
	log *log.Logger
	protos.UnimplementedCurrencyServiceServer
}

func (c *CurrencyService) Get(ctx context.Context, in *protos.CurrencyRequest) (*protos.CurrencyResponse, error) {
	return &protos.CurrencyResponse{}, nil
}
