package category

import (
	"context"
	"errors"
	"fmt"
	gateways "post-tech-challenge-10soat/app/internal/test_utils/mock"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"post-tech-challenge-10soat/app/internal/entities"
)

func TestGetCategoryUsecaseImpl_Execute(t *testing.T) {
	ctx := context.Background()
	categoryID := uuid.New().String()

	testCases := []struct {
		name             string
		categoryID       string
		mockGateway      func(gateway *gateways.MockCategoryGateway)
		expectedCategory entities.Category
		expectedError    error
	}{
		{
			name:       "Success",
			categoryID: categoryID,
			mockGateway: func(gateway *gateways.MockCategoryGateway) {
				gateway.On("GetCategoryById", ctx, categoryID).Return(entities.Category{
					Id:   categoryID,
					Name: "Test Category",
				}, nil)
			},
			expectedCategory: entities.Category{
				Id:   categoryID,
				Name: "Test Category",
			},
			expectedError: nil,
		},
		{
			name:       "Category Not Found",
			categoryID: "non-existent-id",
			mockGateway: func(gateway *gateways.MockCategoryGateway) {
				gateway.On("GetCategoryById", ctx, "non-existent-id").Return(entities.Category{}, entities.ErrDataNotFound)
			},
			expectedCategory: entities.Category{},
			expectedError:    fmt.Errorf("failed to get category by id - %s", entities.ErrDataNotFound.Error()),
		},
		{
			name:       "Gateway Error",
			categoryID: categoryID,
			mockGateway: func(gateway *gateways.MockCategoryGateway) {
				gateway.On("GetCategoryById", ctx, categoryID).Return(entities.Category{}, errors.New("database error"))
			},
			expectedCategory: entities.Category{},
			expectedError:    fmt.Errorf("failed to get category by id - database error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockGateway := new(gateways.MockCategoryGateway)
			if tc.mockGateway != nil {
				tc.mockGateway(mockGateway)
			}

			usecase := NewGetCategoryUsecase(mockGateway)
			category, err := usecase.Execute(ctx, tc.categoryID)

			require.Equal(t, tc.expectedError, err)
			require.Equal(t, tc.expectedCategory, category)

			mockGateway.AssertExpectations(t)
		})
	}
}
