package product

import (
	"context"
	"fmt"
	"post-tech-challenge-10soat/app/internal/dto/product"
	"post-tech-challenge-10soat/app/internal/entities"
	interfaces2 "post-tech-challenge-10soat/app/internal/interfaces/gateways"
)

type CreateProductUseCaseImpl struct {
	productGateway  interfaces2.ProductGateway
	categoryGateway interfaces2.CategoryGateway
}

func NewCreateProductUsecaseImpl(productGateway interfaces2.ProductGateway, categoryGateway interfaces2.CategoryGateway) CreateProductUseCase {
	return &CreateProductUseCaseImpl{
		productGateway,
		categoryGateway,
	}
}

func (s CreateProductUseCaseImpl) Execute(ctx context.Context, createProductDTO dto.CreateProductDTO) (entities.Product, error) {
	category, err := s.categoryGateway.GetCategoryById(ctx, createProductDTO.CategoryId)
	if err != nil {
		if err == entities.ErrDataNotFound {
			return entities.Product{}, err
		}
		return entities.Product{}, fmt.Errorf("cannot create product for this category - %s", err.Error())
	}
	newProduct := entities.Product{
		Name:        createProductDTO.Name,
		Description: createProductDTO.Description,
		Image:       createProductDTO.Image,
		Value:       createProductDTO.Value,
		CategoryId:  createProductDTO.CategoryId,
		Category:    category,
	}
	product, err := s.productGateway.CreateProduct(ctx, newProduct)
	if err != nil {
		if err == entities.ErrConflictingData {
			return entities.Product{}, err
		}
		return entities.Product{}, entities.ErrInternal
	}
	return product, nil
}
