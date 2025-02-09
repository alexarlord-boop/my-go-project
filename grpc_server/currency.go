package grpcserver

// We want to implement server from grpc generated code.

import (
	"context"
	protos "my-go-project/protos/currency"

	"github.com/hashicorp/go-hclog"
)

type CurrencyService struct {
	log hclog.Logger
	protos.UnimplementedCurrencyServiceServer
}

func NewCurrencyService(l hclog.Logger) *CurrencyService {
	return &CurrencyService{l, protos.UnimplementedCurrencyServiceServer{}}
}

func (c *CurrencyService) GetRate(ctx context.Context, in *protos.RateRequest) (*protos.RateResponse, error) {
	c.log.Info("Handle Get", "base", in.GetBase(), "destination", in.GetDestination())
	return &protos.RateResponse{Rate: 0.5}, nil
}
