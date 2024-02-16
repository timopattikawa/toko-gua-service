package usecase

import (
	"context"
	"github.com/timopattikawa/master-service-tokogua/internal/domain"
	"log"
)

type ProductUseCaseImpl struct {
	repo domain.ProductRepository
}

func (p ProductUseCaseImpl) GetProductById(ctx context.Context, id int) (domain.Product, error) {
	product, err := p.repo.FindProductById(ctx, int(id))
	if err != nil {
		log.Println("UseCase Error: Error cause fail find product in repo")
		return domain.Product{}, err
	}

	return product, nil
}

func NewProductUseCase(repository domain.ProductRepository) domain.ProductUseCase {
	return &ProductUseCaseImpl{
		repo: repository,
	}
}
