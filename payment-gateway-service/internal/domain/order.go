package domain

import (
	"context"

	"github.com/timopattikawa/payment-gateway-service/internal/dto"
)

type Order struct {
	UUID        string `db:"id"`
	ProductId   int    `db:"product_id"`
	CostumerId  int    `db:"costumer_id"`
	TotalAmount int    `db:"total_amount"`
}

type OrderRepository interface {
	SaveRepository(order Order) error
	FindOrderById(id string) (Order, error)
}

type OrderUsecase interface {
	OrderPayment(ctx context.Context, req dto.PaymentRequest) (dto.MidtransResponseSnap, error)
	HandlerWebHookPayment(eq map[string]string) (string, error)
	GetDetailOrderById(ctx context.Context, id int) (Order, error)
}
