package usecase_test

import (
	"context"
	"github.com/stretchr/testify/mock"
	pb "github.com/timopattikawa/payment-gateway-service/protos"
)

type ProductServiceClient struct {
	mock.Mock
}

func (p *ProductServiceClient) FindCostumerById(ctx context.Context, id *pb.IdCostumer) (bool, error) {
	args := p.Called(&id.Id)
	return args.Is(pb.Costumer{}), args.Error(1)
}
