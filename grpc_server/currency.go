package grpcserver

import (
	"context"
	"log"
	protos "my-go-project/protos/currency"
)

type CurrencyService struct {
	log *log.Logger
}

func (c *CurrencyService) GetCurrency(ctx context.Context, in *protos.CurrencyRequest) (*protos.CurrencyResponse, error) {
	return &protos.CurrencyResponse{}, nil
}
