package product

import (
	"context"
	"errors"
	dto "post-tech-challenge-10soat/app/internal/dto/product"
	"post-tech-challenge-10soat/app/internal/entities"
	_ "post-tech-challenge-10soat/app/internal/interfaces/gateways"
	gateways "post-tech-challenge-10soat/app/internal/test_utils/mock"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCreateProductUseCaseImpl_Execute(t *testing.T) {
	ctx := context.Background()

	testCases := []struct {
		name                string
		createProductDTO    dto.CreateProductDTO
		mockProductGateway  func(productGateway *gateways.MockProductGateway)
		mockCategoryGateway func(categoryGateway *gateways.MockCategoryGateway)
		expectedProduct     entities.Product
		expectedError       error
	}{
		{
			name: "Success",
			createProductDTO: dto.CreateProductDTO{
				Name:        "Test Product",
				Description: "Test Description",
				Image:       "Test Image",
				Value:       10.0,
				CategoryId:  "1",
			},
			mockCategoryGateway: func(categoryGateway *gateways.MockCategoryGateway) {
				categoryGateway.On("GetCategoryById", ctx, "1").Return(entities.Category{Id: "1", Name: "Test Category"}, nil)
			},
			mockProductGateway: func(productGateway *gateways.MockProductGateway) {
				productGateway.On("CreateProduct", ctx, mock.AnythingOfType("entities.Product")).Return(entities.Product{
					Id:          "1",
					Name:        "Test Product",
					Description: "Test Description",
					Image:       "Test Image",
					Value:       10.0,
					CategoryId:  "1",
					Category:    entities.Category{Id: "1", Name: "Test Category"},
				}, nil)
			},
			expectedProduct: entities.Product{
				Id:          "1",
				Name:        "Test Product",
				Description: "Test Description",
				Image:       "Test Image",
				Value:       10.0,
				CategoryId:  "1",
				Category:    entities.Category{Id: "1", Name: "Test Category"},
			},
			expectedError: nil,
		},
		{
			name: "Category Not Found",
			createProductDTO: dto.CreateProductDTO{
				CategoryId: "1",
			},
			mockCategoryGateway: func(categoryGateway *gateways.MockCategoryGateway) {
				categoryGateway.On("GetCategoryById", ctx, "1").Return(entities.Category{}, entities.ErrDataNotFound)
			},
			mockProductGateway: func(productGateway *gateways.MockProductGateway) {},
			expectedProduct:    entities.Product{},
			expectedError:      entities.ErrDataNotFound,
		},
		{
			name: "Create Product Error",
			createProductDTO: dto.CreateProductDTO{
				Name:        "Test Product",
				Description: "Test Description",
				Image:       "Test Image",
				Value:       10.0,
				CategoryId:  "1",
			},
			mockCategoryGateway: func(categoryGateway *gateways.MockCategoryGateway) {
				categoryGateway.On("GetCategoryById", ctx, "1").Return(entities.Category{Id: "1", Name: "Test Category"}, nil)
			},
			mockProductGateway: func(productGateway *gateways.MockProductGateway) {
				productGateway.On("CreateProduct", ctx, mock.AnythingOfType("entities.Product")).Return(entities.Product{}, errors.New("database error"))
			},
			expectedProduct: entities.Product{},
			expectedError:   entities.ErrInternal,
		},
		{
			name: "Conflicting Data Error",
			createProductDTO: dto.CreateProductDTO{
				Name:        "Test Product",
				Description: "Test Description",
				Image:       "Test Image",
				Value:       10.0,
				CategoryId:  "1",
			},
			mockCategoryGateway: func(categoryGateway *gateways.MockCategoryGateway) {
				categoryGateway.On("GetCategoryById", ctx, "1").Return(entities.Category{Id: "1", Name: "Test Category"}, nil)
			},
			mockProductGateway: func(productGateway *gateways.MockProductGateway) {
				productGateway.On("CreateProduct", ctx, mock.AnythingOfType("entities.Product")).Return(entities.Product{}, entities.ErrConflictingData)
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

			useCase := NewCreateProductUsecaseImpl(mockProductGateway, mockCategoryGateway)
			product, err := useCase.Execute(ctx, tc.createProductDTO)

			require.Equal(t, tc.expectedError, err)
			require.Equal(t, tc.expectedProduct, product)

			mockProductGateway.AssertExpectations(t)
			mockCategoryGateway.AssertExpectations(t)
		})
	}
}
