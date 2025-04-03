package product

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"post-tech-challenge-10soat/app/internal/dto/product"
	"post-tech-challenge-10soat/app/internal/entities"
	gateways "post-tech-challenge-10soat/app/internal/test_utils/mock"
	"testing"
)

func TestUpdateProductUsecaseImpl_Execute(t *testing.T) {
	ctx := context.Background()
	validUUID := uuid.New().String()
	categoryUUID := uuid.New().String()

	testCases := []struct {
		name                string
		updateProductDTO    dto.UpdateProductDTO
		mockProductGateway  func(productGateway *gateways.MockProductGateway)
		mockCategoryGateway func(categoryGateway *gateways.MockCategoryGateway)
		expectedProduct     entities.Product
		expectedError       error
	}{
		{
			name: "Success",
			updateProductDTO: dto.UpdateProductDTO{
				Id:          validUUID,
				Name:        "Updated Product",
				Description: "Updated Description",
				Value:       20.0,
				CategoryId:  categoryUUID,
			},
			mockProductGateway: func(productGateway *gateways.MockProductGateway) {
				productGateway.On("GetProductById", ctx, validUUID).Return(entities.Product{
					Id:          validUUID,
					Name:        "Original Product",
					Description: "Original Description",
					Value:       10.0,
					CategoryId:  categoryUUID,
				}, nil)
				productGateway.On("UpdateProduct", ctx, mock.AnythingOfType("entities.Product")).Return(entities.Product{}, nil)
			},
			mockCategoryGateway: func(categoryGateway *gateways.MockCategoryGateway) {
				categoryGateway.On("GetCategoryById", ctx, categoryUUID).Return(entities.Category{Id: categoryUUID, Name: "Test Category"}, nil)
			},
			expectedProduct: entities.Product{
				Name:        "Updated Product",
				Description: "Updated Description",
				Value:       20.0,
				CategoryId:  categoryUUID,
				Category:    entities.Category{Id: categoryUUID, Name: "Test Category"},
			},
			expectedError: nil,
		},
		{
			name: "Product Not Found",
			updateProductDTO: dto.UpdateProductDTO{
				Id: "non-existent-id",
			},
			mockProductGateway: func(productGateway *gateways.MockProductGateway) {
				productGateway.On("GetProductById", ctx, "non-existent-id").Return(entities.Product{}, entities.ErrDataNotFound)
			},
			mockCategoryGateway: func(categoryGateway *gateways.MockCategoryGateway) {},
			expectedProduct:     entities.Product{},
			expectedError:       entities.ErrDataNotFound,
		},
		{
			name: "No Updated Data",
			updateProductDTO: dto.UpdateProductDTO{
				Id:          validUUID,
				Name:        "Original Product",
				Description: "Original Description",
				Value:       10.0,
				CategoryId:  categoryUUID,
			},
			mockProductGateway: func(productGateway *gateways.MockProductGateway) {
				productGateway.On("GetProductById", ctx, validUUID).Return(entities.Product{
					Id:          validUUID,
					Name:        "Original Product",
					Description: "Original Description",
					Value:       10.0,
					CategoryId:  categoryUUID,
				}, nil)
			},
			mockCategoryGateway: func(categoryGateway *gateways.MockCategoryGateway) {},
			expectedProduct:     entities.Product{},
			expectedError:       entities.ErrNoUpdatedData,
		},
		{
			name: "Category Not Found",
			updateProductDTO: dto.UpdateProductDTO{
				Id:          validUUID,
				Name:        "Updated Product",
				Description: "Updated Description",
				Value:       20.0,
				CategoryId:  "non-existent-category",
			},
			mockProductGateway: func(productGateway *gateways.MockProductGateway) {
				productGateway.On("GetProductById", ctx, validUUID).Return(entities.Product{
					Id:         validUUID,
					CategoryId: categoryUUID,
				}, entities.ErrDataNotFound)
			},
			mockCategoryGateway: func(categoryGateway *gateways.MockCategoryGateway) {
				categoryGateway.On("GetCategoryById", ctx, "non-existent-category").Return(entities.Category{}, entities.ErrDataNotFound)
			},
			expectedProduct: entities.Product{},
			expectedError:   entities.ErrDataNotFound,
		},
		{
			name: "Update Product Error",
			updateProductDTO: dto.UpdateProductDTO{
				Id:          validUUID,
				Name:        "Updated Product",
				Description: "Updated Description",
				Value:       20.0,
				CategoryId:  categoryUUID,
			},
			mockProductGateway: func(productGateway *gateways.MockProductGateway) {
				productGateway.On("GetProductById", ctx, validUUID).Return(entities.Product{
					Id:         validUUID,
					CategoryId: categoryUUID,
				}, nil)
				productGateway.On("UpdateProduct", ctx, mock.AnythingOfType("entities.Product")).Return(entities.Product{}, errors.New("database error"))
			},
			mockCategoryGateway: func(categoryGateway *gateways.MockCategoryGateway) {
				categoryGateway.On("GetCategoryById", ctx, categoryUUID).Return(entities.Category{Id: categoryUUID, Name: "Test Category"}, nil)
			},
			expectedProduct: entities.Product{},
			expectedError:   entities.ErrInternal,
		},
		{
			name: "Conflicting Data Error",
			updateProductDTO: dto.UpdateProductDTO{
				Id:          validUUID,
				Name:        "Updated Product",
				Description: "Updated Description",
				Value:       20.0,
				CategoryId:  categoryUUID,
			},
			mockProductGateway: func(productGateway *gateways.MockProductGateway) {
				productGateway.On("GetProductById", ctx, validUUID).Return(entities.Product{
					Id:         validUUID,
					CategoryId: categoryUUID,
				}, nil)
				productGateway.On("UpdateProduct", ctx, mock.AnythingOfType("entities.Product")).Return(entities.Product{}, entities.ErrConflictingData)
			},
			mockCategoryGateway: func(categoryGateway *gateways.MockCategoryGateway) {
				categoryGateway.On("GetCategoryById", ctx, categoryUUID).Return(entities.Category{Id: categoryUUID, Name: "Test Category"}, nil)
			},
			expectedProduct: entities.Product{},
			expectedError:   entities.ErrConflictingData,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockProductGateway := new(gateways.MockProductGateway)
			mockCategoryGateway := new(gateways.MockCategoryGateway)

			if tc.mockProductGateway != nil {
				tc.mockProductGateway(mockProductGateway)
			}
			if tc.mockCategoryGateway != nil {
				tc.mockCategoryGateway(mockCategoryGateway)
			}

			useCase := NewUpdateProductUsecaseImpl(mockProductGateway, mockCategoryGateway)
			product, err := useCase.Execute(ctx, tc.updateProductDTO)

			require.Equal(t, tc.expectedError, err)
			require.Equal(t, tc.expectedProduct, product)
		})
	}
}
