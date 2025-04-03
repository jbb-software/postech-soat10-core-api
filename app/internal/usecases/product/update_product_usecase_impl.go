package product

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"post-tech-challenge-10soat/app/internal/dto/product"
	"post-tech-challenge-10soat/app/internal/entities"
	interfaces2 "post-tech-challenge-10soat/app/internal/interfaces/gateways"
)

type UpdateProductUsecaseImpl struct {
	productGateway  interfaces2.ProductGateway
	categoryGateway interfaces2.CategoryGateway
}

func NewUpdateProductUsecaseImpl(productGateway interfaces2.ProductGateway, categoryGateway interfaces2.CategoryGateway) UpdateProductUseCase {
	return &UpdateProductUsecaseImpl{
		productGateway,
		categoryGateway,
	}
}

func (s UpdateProductUsecaseImpl) Execute(ctx context.Context, updateProductDTO dto.UpdateProductDTO) (entities.Product, error) {
	existingProduct, err := s.productGateway.GetProductById(ctx, updateProductDTO.Id)
	if err != nil {
		if err == entities.ErrDataNotFound {
			return entities.Product{}, err
		}
		return entities.Product{}, fmt.Errorf("cannot find product to update - %s", err.Error())
	}
	emptyData := uuid.Validate(updateProductDTO.CategoryId) != nil &&
		updateProductDTO.Name == "" &&
		updateProductDTO.Value == 0
	sameData := existingProduct.CategoryId == updateProductDTO.CategoryId &&
		existingProduct.Name == updateProductDTO.Name &&
		existingProduct.Value == updateProductDTO.Value &&
		existingProduct.Description == updateProductDTO.Description
	if emptyData || sameData {
		return entities.Product{}, entities.ErrNoUpdatedData
	}
	if uuid.Validate(updateProductDTO.CategoryId) != nil {
		updateProductDTO.CategoryId = existingProduct.CategoryId
	}
	category, err := s.categoryGateway.GetCategoryById(ctx, updateProductDTO.CategoryId)
	if err != nil {
		if err == entities.ErrDataNotFound {
			return entities.Product{}, err
		}
		return entities.Product{}, fmt.Errorf("cannot update product for this category - %s", err.Error())
	}
	newUpdateProduct := entities.Product{
		Name:        updateProductDTO.Name,
		Description: updateProductDTO.Description,
		Image:       updateProductDTO.Image,
		Value:       updateProductDTO.Value,
		CategoryId:  updateProductDTO.CategoryId,
		Category:    category,
	}
	_, err = s.productGateway.UpdateProduct(ctx, newUpdateProduct)
	if err != nil {
		if err == entities.ErrConflictingData {
			return entities.Product{}, err
		}
		return entities.Product{}, entities.ErrInternal
	}
	return newUpdateProduct, nil
}
