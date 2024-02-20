package repository

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/timopattikawa/payment-gateway-service/internal/domain"
)

type OrderRepositoryImpl struct {
	Db *sql.DB
}

func (o OrderRepositoryImpl) SaveRepository(order domain.Order) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	query := `INSERT INTO 
	public.order(id, product_id, costumer_id, total_amount)
	VALUES ($1, $2, $3, $4)`

	stmt, err := o.Db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Fatalf("Error %s expected when close statement", err)
		}
	}(stmt)

	_, err = stmt.ExecContext(
		ctx,
		order.UUID,
		order.ProductId,
		order.CostumerId,
		order.TotalAmount)
	return err
}

func (o OrderRepositoryImpl) FindOrderById(id string) (domain.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	query := `SELECT id, product_id, costumer_id, total_amount
	FROM public.order WHERE id = $1;`

	var order = domain.Order{}

	stmt, err := o.Db.PrepareContext(ctx, query)
	if err != nil {
		return domain.Order{}, err
	}

	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Fatalf("Error %s expected when close statement", err)
		}
	}(stmt)

	err = stmt.QueryRowContext(ctx, id).Scan(
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
		Db: db,
	}
}
