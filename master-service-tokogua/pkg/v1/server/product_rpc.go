package server

import (
	"context"
	"github.com/timopattikawa/master-service-tokogua/internal/domain"
	pb "github.com/timopattikawa/master-service-tokogua/protos"
	"google.golang.org/grpc"
	"sync"
)

type ProductServiceRPC struct {
	useCase domain.ProductUseCase
	pb.UnimplementedDataProductServerServer
	sync sync.Mutex
}

func NewServerProductServiceRPC(server *grpc.Server, useCase domain.ProductUseCase) {
	ProductRPC := &ProductServiceRPC{
		useCase: useCase,
	}
	pb.RegisterDataProductServerServer(server, ProductRPC)
}

func (r *ProductServiceRPC) FindDataProductById(ctx context.Context, id *pb.IdProduct) (*pb.Product, error) {
	product, err := r.useCase.GetProductById(ctx, int(id.Id))
	if err != nil {
		return nil, err
	}

	protoProduct := r.transformToProto(product)
	return protoProduct, nil
}

func (r *ProductServiceRPC) transformToProto(product domain.Product) *pb.Product {
	return &pb.Product{Id: int64(product.Id), Name: product.Name, Price: int64(product.Price)}
}
