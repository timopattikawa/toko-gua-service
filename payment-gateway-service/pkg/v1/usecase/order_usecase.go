package usecase

import (
	"context"
	"encoding/json"
	"github.com/timopattikawa/payment-gateway-service/pkg/v1/client"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/timopattikawa/payment-gateway-service/cmd/config"
	"github.com/timopattikawa/payment-gateway-service/internal/dao"
	"github.com/timopattikawa/payment-gateway-service/internal/domain"
	"github.com/timopattikawa/payment-gateway-service/internal/dto"
)

type OrderUsecaseImpl struct {
	cfg         *config.Config
	costumerRPC *client.CostumerRPC
	productRPC  *client.ProductRPC
	repo        domain.OrderRepository
	midtrans    dao.MidtransDao
}

func (o OrderUsecaseImpl) OrderPayment(ctx context.Context, req dto.PaymentRequest) (dto.MidtransResponseSnap, error) {
	log.Printf("Payment Proccess started for costumer witGet Request From Midtrans h id %v\n", req.CostumerId)
	costumer, err := o.costumerRPC.FindCostumerById(ctx, req.CostumerId)
	if err != nil {
		log.Println("Not found costumer with id: ", req.CostumerId)
		return dto.MidtransResponseSnap{}, err
	}

	product, err := o.productRPC.FindCostumerById(ctx, req.ProductId)
	if err != nil {
		return dto.MidtransResponseSnap{}, err
	}

	if err != nil {
		log.Println("Failed to generate uuid")
		return dto.MidtransResponseSnap{}, err
	}

	var newOrder = domain.Order{
		ProductId:   int(product.Id),
		CostumerId:  int(costumer.Id),
		TotalAmount: req.Qty * int(product.Price),
	}

	resp, err := o.midtrans.DoRequestMidtransSnap(newOrder, costumer)
	if err != nil {
		return dto.MidtransResponseSnap{}, err
	}

	if resp.StatusCode != fiber.StatusCreated {
		log.Printf("Response Status Code Midtrans : {%v}\n", resp.StatusCode)
		var errorResp interface{}
		err := json.NewDecoder(resp.Body).Decode(&errorResp)
		log.Printf("Body Response From Midtrans : {%v}\n", errorResp)
		return dto.MidtransResponseSnap{}, err
	}

	midtransResponse := dto.MidtransResponseSnap{}

	if err := json.NewDecoder(resp.Body).Decode(&midtransResponse); err != nil {
		log.Printf("Midtrans Response : {%v}", midtransResponse)
		return dto.MidtransResponseSnap{}, nil
	}

	if err := o.repo.SaveRepository(newOrder); err != nil {
		log.Println("Fail to save order with exception: ", err.Error())
		return dto.MidtransResponseSnap{}, err
	}

	return midtransResponse, nil
}

func (o OrderUsecaseImpl) HandlerWebHookPayment(req map[string]string) (string, error) {

	params := req["order_id"] + req["status_code"] + req["gross_amount"]

	if err := o.midtrans.DoCheckingMidtransWebhook(params, req["signature_key"]); err != nil {
		log.Println("Failed to checking midtrans webhook")
		return "Fail", err
	}

	return "OK", nil
}

func (o OrderUsecaseImpl) GetDetailOrderById(ctx context.Context, id int) (domain.Order, error) {
	//TODO implement me
	panic("implement me")
}

func NewOrderUsecase(repository domain.OrderRepository,
	costumer *client.CostumerRPC,
	product *client.ProductRPC,
	cfg *config.Config,
	midtrans dao.MidtransDao) domain.OrderUsecase {
	return &OrderUsecaseImpl{
		repo:        repository,
		productRPC:  product,
		costumerRPC: costumer,
		cfg:         cfg,
		midtrans:    midtrans,
	}
}
