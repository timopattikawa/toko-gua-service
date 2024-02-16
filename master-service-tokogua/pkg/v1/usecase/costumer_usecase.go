package usecase

import (
	"context"
	"github.com/timopattikawa/master-service-tokogua/internal/domain"
)

type CostumerUseCaseImpl struct {
	repo domain.CostumerRepository
}

func (c CostumerUseCaseImpl) FindCostumerById(ctx context.Context, id int64) (domain.Costumer, error) {
	costumer, err := c.repo.FindCostumerById(ctx, id)
	if err != nil {
		return domain.Costumer{}, err
	}

	return costumer, nil
}

func NewCostumerUseCase(repository domain.CostumerRepository) domain.CostumerUseCase {
	return &CostumerUseCaseImpl{
		repo: repository,
	}
}
