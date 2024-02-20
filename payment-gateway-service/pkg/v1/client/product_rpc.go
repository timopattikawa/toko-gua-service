package client

import (
	"context"
	"github.com/timopattikawa/payment-gateway-service/internal/domain"
	pb "github.com/timopattikawa/payment-gateway-service/protos"
)

type ProductRPC struct {
	client pb.DataProductServerClient
}

func (r ProductRPC) FindCostumerById(ctx context.Context, id int) (domain.Product, error) {
	product, err := r.client.FindDataProductById(ctx, &pb.IdProduct{Id: int64(id)})
	if err != nil {
		return domain.Product{}, err
	}

	return domain.Product{
		Id:    int(product.Id),
		Name:  product.Name,
		Price: int(product.Price),
	}, nil
}

func NewProductClientRPC(client pb.DataProductServerClient) *ProductRPC {
	return &ProductRPC{
		client: client,
	}
}
