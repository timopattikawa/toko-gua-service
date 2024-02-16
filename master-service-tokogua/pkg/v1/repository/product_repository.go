package repository

import (
	"context"
	"database/sql"
	"github.com/timopattikawa/master-service-tokogua/internal/domain"
	"log"
)

type ProductRepoImpl struct {
	db *sql.DB
}

func (p ProductRepoImpl) FindProductById(ctx context.Context, id int) (domain.Product, error) {
	log.Printf("Query to get product from DB by id : %v", id)
	var product = domain.Product{}

	query := `SELECT id, name, price
	FROM product WHERE id = $1`

	err := p.db.QueryRowContext(ctx, query, id).Scan(
		&product.Id,
		&product.Name,
		&product.Price)

	if err != nil {
		log.Println("Error DB: ", err.Error())
		return domain.Product{}, err
	}

	return product, nil
}

func NewProductRepoImpl(db *sql.DB) domain.ProductRepository {
	return &ProductRepoImpl{
		db: db,
	}
}
