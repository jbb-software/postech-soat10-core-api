package product

import (
	"context"
	"post-tech-challenge-10soat/app/internal/entities"
	interfaces2 "post-tech-challenge-10soat/app/internal/interfaces/gateways"
)

type ListProductsUseCaseImpl struct {
	productGateway  interfaces2.ProductGateway
	categoryGateway interfaces2.CategoryGateway
}

func NewListProductsUsecaseImpl(productGateway interfaces2.ProductGateway, categoryGateway interfaces2.CategoryGateway) ListProductsUseCase {
	return &ListProductsUseCaseImpl{
		productGateway,
		categoryGateway,
	}
}

func (s ListProductsUseCaseImpl) Execute(ctx context.Context, categoryId string) ([]entities.Product, error) {
	var products []entities.Product
	products, err := s.productGateway.ListProducts(ctx, categoryId)
	if err != nil {
		return nil, err
	}
	for i, product := range products {
		category, err := s.categoryGateway.GetCategoryById(ctx, product.CategoryId)
		if err != nil {
			if err == entities.ErrDataNotFound {
				return nil, err
			}
			return nil, err
		}

		products[i].Category = category
	}
	return products, nil
}
