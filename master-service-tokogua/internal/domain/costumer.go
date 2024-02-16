package domain

import "context"

type Costumer struct {
	Id            int64  `db:"id" json:"id"`
	CostumerName  string `db:"costumer_name" json:"costumer_name"`
	CostumerEmail string `db:"costumer_email" json:"costumer_email"`
}

type CostumerRepository interface {
	FindCostumerById(ctx context.Context, id int64) (Costumer, error)
}

type CostumerUseCase interface {
	FindCostumerById(ctx context.Context, id int64) (Costumer, error)
}
