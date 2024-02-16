package server

import (
	"context"
	"github.com/timopattikawa/master-service-tokogua/internal/domain"
	pb "github.com/timopattikawa/master-service-tokogua/protos"
	"google.golang.org/grpc"
	"log"
	"sync"
)

type CostumerServiceRPC struct {
	useCase domain.CostumerUseCase
	pb.UnimplementedCostumerDataServerServer
	sync sync.Mutex
}

func NewCostumerServiceRPC(server *grpc.Server, useCase domain.CostumerUseCase) {
	costumerGRPC := &CostumerServiceRPC{
		useCase: useCase,
	}
	pb.RegisterCostumerDataServerServer(server, costumerGRPC)
}

func (cs *CostumerServiceRPC) FindCostumerById(ctx context.Context, costumerId *pb.IdCostumer) (*pb.Costumer, error) {
	log.Printf("Find costumer by id %v\n", costumerId.Id)
	costumer, err := cs.useCase.FindCostumerById(ctx, costumerId.Id)
	if err != nil {
		log.Printf("Error when get costumer from usecase err{%v}", err.Error())
		return nil, err
	}
	return cs.transformToProto(costumer), nil
}

func (cs *CostumerServiceRPC) transformToProto(costumer domain.Costumer) *pb.Costumer {
	return &pb.Costumer{Id: costumer.Id, CostumerName: costumer.CostumerName, CostumerEmail: costumer.CostumerEmail}
}
