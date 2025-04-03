package order

import (
	"context"
	"errors"
	gateways "post-tech-challenge-10soat/app/internal/test_utils/mock"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"post-tech-challenge-10soat/app/internal/entities"
)

func TestListOrdersUseCaseImpl_Execute(t *testing.T) {
	ctx := context.Background()

	now := time.Now()

	testCases := []struct {
		name             string
		limit            uint64
		mockOrderGateway func(orderGateway *gateways.MockOrderGateway)
		expectedOrders   []entities.Order
		expectedError    error
	}{
		{
			name:  "Success - No Orders",
			limit: 10,
			mockOrderGateway: func(orderGateway *gateways.MockOrderGateway) {
				orderGateway.On("ListOrders", ctx, uint64(10)).Return([]entities.Order{}, nil)
			},
			expectedOrders: []entities.Order{},
			expectedError:  nil,
		},
		{
			name:  "Success - Multiple Orders",
			limit: 10,
			mockOrderGateway: func(orderGateway *gateways.MockOrderGateway) {
				order1 := entities.Order{
					Id:        "id1",
					Status:    entities.OrderStatusPreparing,
					CreatedAt: now.Add(-2 * time.Minute),
				}
				order2 := entities.Order{
					Id:        "id2",
					Status:    entities.OrderStatusReady,
					CreatedAt: now.Add(-1 * time.Minute),
				}
				order3 := entities.Order{
					Id:        "id3",
					Status:    entities.OrderStatusReceived,
					CreatedAt: now.Add(-3 * time.Minute),
				}
				order4 := entities.Order{
					Id:        "id4",
					Status:    entities.OrderStatusReady,
					CreatedAt: now.Add(-4 * time.Minute),
				}

				expectedOrders := []entities.Order{order4, order2, order1, order3}

				orderGateway.On("ListOrders", ctx, uint64(10)).Return(expectedOrders, nil)

			},
			expectedOrders: func() []entities.Order {
				order1 := entities.Order{
					Id:        "id1",
					Status:    entities.OrderStatusPreparing,
					CreatedAt: now.Add(-2 * time.Minute),
				}
				order2 := entities.Order{
					Id:        "id2",
					Status:    entities.OrderStatusReady,
					CreatedAt: now.Add(-1 * time.Minute),
				}
				order3 := entities.Order{
					Id:        "id3",
					Status:    entities.OrderStatusReceived,
					CreatedAt: now.Add(-3 * time.Minute),
				}
				order4 := entities.Order{
					Id:        "id4",
					Status:    entities.OrderStatusReady,
					CreatedAt: now.Add(-4 * time.Minute),
				}

				return []entities.Order{order4, order2, order1, order3}
			}(),
			expectedError: nil,
		},
		{
			name:  "Error - Gateway Error",
			limit: 5,
			mockOrderGateway: func(orderGateway *gateways.MockOrderGateway) {
				orderGateway.On("ListOrders", ctx, uint64(5)).Return([]entities.Order{}, errors.New("database error"))
			},
			expectedOrders: []entities.Order{},
			expectedError:  errors.New("database error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockOrderGateway := new(gateways.MockOrderGateway)
			if tc.mockOrderGateway != nil {
				tc.mockOrderGateway(mockOrderGateway)
			}

			useCase := NewListOrdersUseCaseImpl(mockOrderGateway)
			orders, err := useCase.Execute(ctx, tc.limit)

			require.Equal(t, tc.expectedError, err)
			require.Equal(t, tc.expectedOrders, orders)

			mockOrderGateway.AssertExpectations(t)
		})
	}
}
