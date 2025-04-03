package product

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"post-tech-challenge-10soat/app/internal/entities"
	gateways "post-tech-challenge-10soat/app/internal/test_utils/mock"
	"testing"
)

func TestDeleteProductUsecaseImpl_Execute(t *testing.T) {
	ctx := context.Background()
	validUUID := uuid.New().String()
	invalidUUID := "invalid-uuid"

	testCases := []struct {
		name               string
		id                 string
		mockProductGateway func(productGateway *gateways.MockProductGateway)
		expectedError      error
	}{
		{
			name: "Success",
			id:   validUUID,
			mockProductGateway: func(productGateway *gateways.MockProductGateway) {
				productGateway.On("GetProductById", ctx, validUUID).Return(entities.Product{}, nil)
				productGateway.On("DeleteProduct", ctx, validUUID).Return(nil)
			},
			expectedError: nil,
		},
		{
			name:               "Invalid UUID",
			id:                 invalidUUID,
			mockProductGateway: func(productGateway *gateways.MockProductGateway) {},
			expectedError:      fmt.Errorf("invalid product id"),
		},
		{
			name: "Product Not Found",
			id:   validUUID,
			mockProductGateway: func(productGateway *gateways.MockProductGateway) {
				productGateway.On("GetProductById", ctx, validUUID).Return(entities.Product{}, entities.ErrDataNotFound)
			},
			expectedError: entities.ErrDataNotFound,
		},
		{
			name: "Delete Product Error",
			id:   validUUID,
			mockProductGateway: func(productGateway *gateways.MockProductGateway) {
				productGateway.On("GetProductById", ctx, validUUID).Return(entities.Product{}, nil)
				productGateway.On("DeleteProduct", ctx, validUUID).Return(errors.New("database error"))
			},
			expectedError: fmt.Errorf("database error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockProductGateway := new(gateways.MockProductGateway)
			if tc.mockProductGateway != nil {
				tc.mockProductGateway(mockProductGateway)
			}

			useCase := NewDeleteProductUsecaseImpl(mockProductGateway)
			err := useCase.Execute(ctx, tc.id)

			require.Equal(t, tc.expectedError, err)
			mockProductGateway.AssertExpectations(t)
		})
	}
}
