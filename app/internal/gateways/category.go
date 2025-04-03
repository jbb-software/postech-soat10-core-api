package gateways

import (
	"context"
	"post-tech-challenge-10soat/app/internal/entities"
	"post-tech-challenge-10soat/app/internal/interfaces/repositories"
)

type CategoryGatewayImpl struct {
	repository interfaces.CategoryRepository
}

func NewCategoryGatewayImpl(repository interfaces.CategoryRepository) *CategoryGatewayImpl {
	return &CategoryGatewayImpl{
		repository,
	}
}

func (cg CategoryGatewayImpl) GetCategoryById(ctx context.Context, categoryId string) (entities.Category, error) {
	category, err := cg.repository.GetCategoryById(ctx, categoryId)
	if err != nil {
		return entities.Category{}, err
	}
	return category.ToEntity(), nil
}
