package main

import (
	"github.com/timopattikawa/master-service-tokogua/cmd/config"
	"github.com/timopattikawa/master-service-tokogua/internal/db"
	"github.com/timopattikawa/master-service-tokogua/pkg/v1/repository"
	"github.com/timopattikawa/master-service-tokogua/pkg/v1/server"
	"github.com/timopattikawa/master-service-tokogua/pkg/v1/usecase"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {

	cfg := config.InitConfig()
	sqlCon := db.NewConnectionPsql(cfg)

	// Repository Section
	costumerRepository := repository.NewCostumerRepository(sqlCon)
	productRepository := repository.NewProductRepoImpl(sqlCon)

	//UseCase section
	CostumerUseCase := usecase.NewCostumerUseCase(costumerRepository)
	productUsecase := usecase.NewProductUseCase(productRepository)

	listen, err := net.Listen(cfg.Server.Type, cfg.Server.Port)
	if err != nil {
		log.Fatalf("ERROR STARTING THE SERVER : %v", err)
	}

	grpcServer := grpc.NewServer()
	server.NewCostumerServiceRPC(grpcServer, CostumerUseCase)
	server.NewServerProductServiceRPC(grpcServer, productUsecase)

	log.Println("Listen " + cfg.Server.Port)
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Fail to serve grpc listn")
	}

}
