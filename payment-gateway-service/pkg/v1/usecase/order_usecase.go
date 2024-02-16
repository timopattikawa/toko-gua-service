package usecase

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	uuid2 "github.com/google/uuid"
	"github.com/timopattikawa/payment-gateway-service/cmd/config"
	"github.com/timopattikawa/payment-gateway-service/internal/domain"
	"github.com/timopattikawa/payment-gateway-service/internal/dto"
	pb "github.com/timopattikawa/payment-gateway-service/protos"
)

type OrderUsecaseImpl struct {
	cfg         *config.Config
	costumerRPC pb.CostumerDataServerClient
	productRPC  pb.DataProductServerClient
	repo        domain.OrderRepository
}

func (o OrderUsecaseImpl) OrderPayment(ctx context.Context, req dto.PaymentRequest) (dto.MidtransResponseSnap, error) {
	log.Printf("Payment Proccess started for costumer witGet Request From Midtrans h id %v\n", req.CostumerId)
	costumer, err := o.costumerRPC.FindCostumerById(ctx, &pb.IdCostumer{
		Id: int64(req.CostumerId),
	})
	if err != nil {
		log.Println("Not found costumer with id: ", req.CostumerId)
		return dto.MidtransResponseSnap{}, err
	}
	log.Printf("Get costumer {%v}\n", costumer)

	product, err := o.productRPC.FindDataProductById(ctx, &pb.IdProduct{Id: int64(req.ProductId)})
	if err != nil {
		return dto.MidtransResponseSnap{}, err
	}
	log.Printf("Get product {%v}\n", product)

	uuid, err := uuid2.NewUUID()
	if err != nil {
		log.Println("Failed to generate uuid")
		return dto.MidtransResponseSnap{}, err
	}

	var newOrder = domain.Order{
		UUID:        uuid.String(),
		ProductId:   int(product.Id),
		CostumerId:  int(costumer.Id),
		TotalAmount: int(req.Qty * int(product.Price)),
	}

	client := &http.Client{}

	var payloadMidtransRequest = fmt.Sprintf(`{
		"transaction_details": {
			"order_id": "%s",
			"gross_amount": %d
		},
		"credit_card":{
			"secure" : %v
		},
		"customer_details": {
			"name": "%s",
			"email": "%s",
			"phone": "08111222333"
		}
	}`, newOrder.UUID, newOrder.TotalAmount, true, costumer.CostumerName, costumer.CostumerEmail)

	log.Println("Payload : ", payloadMidtransRequest)

	reqToMidtrans, err := http.NewRequest("POST", "https://app.sandbox.midtrans.com/snap/v1/transactions",
		bytes.NewBuffer([]byte(payloadMidtransRequest)))

	if err != nil {
		log.Println("Fail prepare req for midtrans")
		return dto.MidtransResponseSnap{}, err
	}

	authString := base64.StdEncoding.EncodeToString([]byte(o.cfg.Midtrans.ServerKey + ":"))
	log.Println(authString)
	reqToMidtrans.Header.Add("Accept", "application/json")
	reqToMidtrans.Header.Add("Content-Type", "application/json")
	reqToMidtrans.Header.Add("Authorization",
		"Basic "+authString)

	resp, err := client.Do(reqToMidtrans)
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

	if err := o.repo.SaveRepository(ctx, newOrder); err != nil {
		log.Println("Fail to save order with exception: ", err.Error())
		return dto.MidtransResponseSnap{}, err
	}

	return midtransResponse, nil
}

func (o OrderUsecaseImpl) HandlerWebHookPayment(ctx context.Context, req string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (o OrderUsecaseImpl) GetDetailOrderById(ctx context.Context, id int) (domain.Order, error) {
	//TODO implement me
	panic("implement me")
}

func NewOrderUsecase(repository domain.OrderRepository,
	costumer pb.CostumerDataServerClient,
	product pb.DataProductServerClient,
	cfg *config.Config) domain.OrderUsecase {
	return &OrderUsecaseImpl{
		repo:        repository,
		productRPC:  product,
		costumerRPC: costumer,
		cfg:         cfg,
	}
}
