package category

import (
	"context"
	"fmt"
	"post-tech-challenge-10soat/app/internal/entities"
	"post-tech-challenge-10soat/app/internal/interfaces/gateways"
)

type GetCategoryUsecaseImpl struct {
	gateway interfaces.CategoryGateway
}

func NewGetCategoryUsecase(gateway interfaces.CategoryGateway) GetCategoryUseCase {
	return &GetCategoryUsecaseImpl{
		gateway,
	}
}

func (s *GetCategoryUsecaseImpl) Execute(ctx context.Context, id string) (entities.Category, error) {
	category, err := s.gateway.GetCategoryById(ctx, id)
	if err != nil {
		return entities.Category{}, fmt.Errorf("failed to get category by id - %s", err.Error())
	}
	return category, nil
}
