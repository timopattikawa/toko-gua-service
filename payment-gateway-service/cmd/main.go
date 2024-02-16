package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/timopattikawa/payment-gateway-service/cmd/config"
	"github.com/timopattikawa/payment-gateway-service/internal/db"
	error2 "github.com/timopattikawa/payment-gateway-service/internal/exception"
	"github.com/timopattikawa/payment-gateway-service/internal/repository"
	"github.com/timopattikawa/payment-gateway-service/pkg/v1/client"
	"github.com/timopattikawa/payment-gateway-service/pkg/v1/handler"
	"github.com/timopattikawa/payment-gateway-service/pkg/v1/usecase"
	pb "github.com/timopattikawa/payment-gateway-service/protos"
)

func main() {

	cfg := config.InitConfiguration()
	postgresDBCon := db.InitPostgresDB(cfg)

	orderRepo := repository.NewOrderRepository(postgresDBCon)

	rpcClient := client.NewClientMasterService(cfg)
	productClient := pb.NewDataProductServerClient(rpcClient)
	costumerClient := pb.NewCostumerDataServerClient(rpcClient)

	orderUsecase := usecase.NewOrderUsecase(orderRepo, costumerClient, productClient, cfg)

	orderHandler := handler.NewOrderHandler(orderUsecase, cfg)

	app := fiber.New(fiber.Config{ErrorHandler: error2.CustomErrorHandler})

	orderHandler.OrderRoute(app)

	log.Println("Listen 3000 :)")
	err := app.Listen(":3000")
	if err != nil {
		log.Fatal("fail to listen")
	}
}
