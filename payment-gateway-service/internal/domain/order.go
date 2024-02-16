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
	SaveRepository(ctx context.Context, order Order) error
	FindOrderById(ctx context.Context, id int) (Order, error)
}

type OrderUsecase interface {
	OrderPayment(ctx context.Context, req dto.PaymentRequest) (dto.MidtransResponseSnap, error)
	HandlerWebHookPayment(ctx context.Context, req string) (string, error)
	GetDetailOrderById(ctx context.Context, id int) (Order, error)
}
