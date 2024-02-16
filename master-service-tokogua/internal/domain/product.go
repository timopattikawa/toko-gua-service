package domain

import "context"

type Product struct {
	Id    int    `db:"id"  json:"id"`
	Name  string `db:"name" json:"name"`
	Price int    `db:"price" json:"price"`
}

type ProductRepository interface {
	FindProductById(ctx context.Context, id int) (Product, error)
}

type ProductUseCase interface {
	GetProductById(ctx context.Context, id int) (Product, error)
}
