package repository

import (
	"context"
	"database/sql"
	"github.com/timopattikawa/payment-gateway-service/internal/domain"
	"log"
	"time"
)

type OrderRepositoryImpl struct {
	db *sql.DB
}

func (o OrderRepositoryImpl) SaveRepository(ctx context.Context, order domain.Order) error {
	query := `INSERT INTO 
	"order"(id, product_id, costumer_id, total_amount)
	VALUES ($1, $2, $3, $4)`

	_, err := o.db.ExecContext(ctx, query,
		order.UUID,
		order.ProductId,
		order.CostumerId,
		order.TotalAmount)

	if err != nil {
		log.Printf("Fail to query: %v\n", query)
		log.Printf("Error Message: %v\n", err.Error())
		return err
	}

	return nil
}

func (o OrderRepositoryImpl) FindOrderById(ctx context.Context, id int) (domain.Order, error) {
	context.WithTimeout(ctx, 10*time.Second)
	query := `SELECT id, product_id, costumer_id, total_amount
	FROM public.order WHERE id = $1;`

	var order = domain.Order{}

	err := o.db.QueryRowContext(ctx, query, id).Scan(
		&order.UUID,
		&order.ProductId,
		&order.CostumerId,
		&order.TotalAmount,
	)
	if err != nil {
		log.Printf("Fail to query: %v\n", query)
		log.Printf("Error Message: %v\n", err.Error())
		return domain.Order{}, err
	}

	return order, nil
}

func NewOrderRepository(db *sql.DB) domain.OrderRepository {
	return &OrderRepositoryImpl{
		db: db,
	}
}
