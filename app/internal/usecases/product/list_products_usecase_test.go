package product

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"post-tech-challenge-10soat/app/internal/entities"
	gateways "post-tech-challenge-10soat/app/internal/test_utils/mock"
)

func TestListProductsUseCaseImpl_Execute(t *testing.T) {
	ctx := context.Background()
	categoryId := "1"

	testCases := []struct {
		name                string
		categoryId          string
		mockProductGateway  func(productGateway *gateways.MockProductGateway)
		mockCategoryGateway func(categoryGateway *gateways.MockCategoryGateway)
		expectedProducts    []entities.Product
		expectedError       error
	}{
		{
			name:       "Success",
			categoryId: categoryId,
			mockProductGateway: func(productGateway *gateways.MockProductGateway) {
				productGateway.On("ListProducts", ctx, categoryId).Return([]entities.Product{
					{Id: "1", Name: "Product 1", CategoryId: "1"},
					{Id: "2", Name: "Product 2", CategoryId: "1"},
				}, nil)
			},
			mockCategoryGateway: func(categoryGateway *gateways.MockCategoryGateway) {
				categoryGateway.On("GetCategoryById", ctx, "1").Return(entities.Category{Id: "1", Name: "Category 1"}, nil).Times(2)
			},
			expectedProducts: []entities.Product{
				{Id: "1", Name: "Product 1", CategoryId: "1", Category: entities.Category{Id: "1", Name: "Category 1"}},
				{Id: "2", Name: "Product 2", CategoryId: "1", Category: entities.Category{Id: "1", Name: "Category 1"}},
			},
			expectedError: nil,
		},
		{
			name:       "Product Gateway Error",
			categoryId: categoryId,
			mockProductGateway: func(productGateway *gateways.MockProductGateway) {
				productGateway.On("ListProducts", ctx, categoryId).Return([]entities.Product{
					{Id: "1", Name: "Product 1", CategoryId: "1"},
				}, errors.New("database error"))
			},
			mockCategoryGateway: func(categoryGateway *gateways.MockCategoryGateway) {},
			expectedProducts:    nil,
			expectedError:       errors.New("database error"),
		},
		{
			name:       "Category Not Found",
			categoryId: categoryId,
			mockProductGateway: func(productGateway *gateways.MockProductGateway) {
				productGateway.On("ListProducts", ctx, categoryId).Return([]entities.Product{
					{Id: "1", Name: "Product 1", CategoryId: "1"},
				}, nil)
			},
			mockCategoryGateway: func(categoryGateway *gateways.MockCategoryGateway) {
				categoryGateway.On("GetCategoryById", ctx, "1").Return(entities.Category{}, entities.ErrDataNotFound)
			},
			expectedProducts: nil,
			expectedError:    entities.ErrDataNotFound,
		},
		{
			name:       "Category Gateway Error",
			categoryId: categoryId,
			mockProductGateway: func(productGateway *gateways.MockProductGateway) {
				productGateway.On("ListProducts", ctx, categoryId).Return([]entities.Product{
					{Id: "1", Name: "Product 1", CategoryId: "1"},
				}, nil)
			},
			mockCategoryGateway: func(categoryGateway *gateways.MockCategoryGateway) {
				categoryGateway.On("GetCategoryById", ctx, "1").Return(entities.Category{}, errors.New("category error"))
			},
			expectedProducts: nil,
			expectedError:    errors.New("category error"),
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

			useCase := NewListProductsUsecaseImpl(mockProductGateway, mockCategoryGateway)
			products, err := useCase.Execute(ctx, tc.categoryId)

			require.Equal(t, tc.expectedError, err)
			require.Equal(t, tc.expectedProducts, products)

			mockProductGateway.AssertExpectations(t)
			mockCategoryGateway.AssertExpectations(t)
		})
	}
}
