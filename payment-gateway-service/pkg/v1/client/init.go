package client

import (
	"log"

	"github.com/timopattikawa/payment-gateway-service/cmd/config"
	"google.golang.org/grpc"
)

func NewClientMasterService(config *config.Config) *grpc.ClientConn {

	log.Println("Open Dial GRPC")
	dial, err := grpc.Dial(config.GRPCMaster.Target, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err.Error())
	}

	return dial
}
