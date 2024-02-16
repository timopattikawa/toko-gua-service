package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/timopattikawa/master-service-tokogua/internal/domain"
)

type CostumerRepoImpl struct {
	db *sql.DB
}

func (c CostumerRepoImpl) FindCostumerById(ctx context.Context, id int64) (domain.Costumer, error) {
	log.Printf("Query to get costumer from DB by id : %v", id)
	var costumer = domain.Costumer{}

	query := "SELECT id, costumer_name, costumer_email FROM costumer WHERE id = $1"

	err := c.db.QueryRowContext(ctx, query, id).Scan(
		&costumer.Id,
		&costumer.CostumerName,
		&costumer.CostumerEmail)

	if err != nil {
		log.Println("Error DB: ", err.Error())
		return domain.Costumer{}, err
	}

	return costumer, nil
}

func NewCostumerRepository(db *sql.DB) domain.CostumerRepository {
	return &CostumerRepoImpl{
		db: db,
	}
}
