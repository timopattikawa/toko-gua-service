package client

import (
	"context"
	"github.com/timopattikawa/payment-gateway-service/internal/domain"
	pb "github.com/timopattikawa/payment-gateway-service/protos"
)

type CostumerRPC struct {
	client pb.CostumerDataServerClient
}

func (r CostumerRPC) FindCostumerById(ctx context.Context, id int) (domain.Costumer, error) {
	costumer, err := r.client.FindCostumerById(ctx, &pb.IdCostumer{Id: int64(id)})
	if err != nil {
		return domain.Costumer{}, err
	}

	return domain.Costumer{
		Id:            costumer.Id,
		CostumerName:  costumer.CostumerName,
		CostumerEmail: costumer.CostumerEmail,
	}, nil
}

func NewCostumerClientRPC(client pb.CostumerDataServerClient) *CostumerRPC {
	return &CostumerRPC{
		client: client,
	}
}
