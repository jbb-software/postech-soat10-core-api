package order

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

func TestUpdateOrderStatusUseCaseImpl_Execute(t *testing.T) {
	ctx := context.Background()
	orderID := uuid.New().String()

	testCases := []struct {
		name             string
		orderID          string
		initialStatus    entities.OrderStatus
		targetStatus     string
		mockOrderGateway func(orderGateway *gateways.MockOrderGateway)
		expectedOrder    entities.Order
		expectedError    error
	}{
		{
			name:          "Valid Transition - received to preparing",
			orderID:       orderID,
			initialStatus: entities.OrderStatusReceived,
			targetStatus:  "preparing",
			mockOrderGateway: func(orderGateway *gateways.MockOrderGateway) {
				orderGateway.On("GetOrderById", ctx, orderID).Return(entities.Order{Id: orderID, Status: entities.OrderStatusReceived}, nil)
				orderGateway.On("UpdateOrderStatus", ctx, orderID, "preparing").Return(entities.Order{Id: orderID, Status: entities.OrderStatusPreparing}, nil)
			},
			expectedOrder: entities.Order{Id: orderID, Status: entities.OrderStatusPreparing},
			expectedError: nil,
		},
		{
			name:          "Valid Transition - preparing to ready",
			orderID:       orderID,
			initialStatus: entities.OrderStatusPreparing,
			targetStatus:  "ready",
			mockOrderGateway: func(orderGateway *gateways.MockOrderGateway) {
				orderGateway.On("GetOrderById", ctx, orderID).Return(entities.Order{Id: orderID, Status: entities.OrderStatusPreparing}, nil)
				orderGateway.On("UpdateOrderStatus", ctx, orderID, "ready").Return(entities.Order{Id: orderID, Status: entities.OrderStatusReady}, nil)
			},
			expectedOrder: entities.Order{Id: orderID, Status: entities.OrderStatusReady},
			expectedError: nil,
		},
		{
			name:          "Valid Transition - ready to completed",
			orderID:       orderID,
			initialStatus: entities.OrderStatusReady,
			targetStatus:  "completed",
			mockOrderGateway: func(orderGateway *gateways.MockOrderGateway) {
				orderGateway.On("GetOrderById", ctx, orderID).Return(entities.Order{Id: orderID, Status: entities.OrderStatusReady}, nil)
				orderGateway.On("UpdateOrderStatus", ctx, orderID, "completed").Return(entities.Order{Id: orderID, Status: entities.OrderStatusCompleted}, nil)
			},
			expectedOrder: entities.Order{Id: orderID, Status: entities.OrderStatusCompleted},
			expectedError: nil,
		},
		{
			name:          "Invalid Initial Status",
			orderID:       orderID,
			initialStatus: "invalid",
			targetStatus:  "preparing",
			mockOrderGateway: func(orderGateway *gateways.MockOrderGateway) {
				orderGateway.On("GetOrderById", ctx, orderID).Return(entities.Order{Id: orderID, Status: "invalid"}, nil)
			},
			expectedOrder: entities.Order{},
			expectedError: errors.New("invalid status"),
		},
		{
			name:          "Invalid Transition",
			orderID:       orderID,
			initialStatus: entities.OrderStatusReceived,
			targetStatus:  "completed",
			mockOrderGateway: func(orderGateway *gateways.MockOrderGateway) {
				orderGateway.On("GetOrderById", ctx, orderID).Return(entities.Order{Id: orderID, Status: entities.OrderStatusReceived}, nil)
			},
			expectedOrder: entities.Order{},
			expectedError: fmt.Errorf("cannot update order status for '%s' to '%s'", entities.OrderStatusReceived, "completed"),
		},
		{
			name:          "Order Not Found",
			orderID:       orderID,
			initialStatus: entities.OrderStatusReceived,
			targetStatus:  "preparing",
			mockOrderGateway: func(orderGateway *gateways.MockOrderGateway) {
				orderGateway.On("GetOrderById", ctx, orderID).Return(entities.Order{}, errors.New("order not found"))
			},
			expectedOrder: entities.Order{},
			expectedError: errors.New("order not found"),
		},
		{
			name:          "Update Order Status Error",
			orderID:       orderID,
			initialStatus: entities.OrderStatusReceived,
			targetStatus:  "preparing",
			mockOrderGateway: func(orderGateway *gateways.MockOrderGateway) {
				orderGateway.On("GetOrderById", ctx, orderID).Return(entities.Order{Id: orderID, Status: entities.OrderStatusReceived}, nil)
				orderGateway.On("UpdateOrderStatus", ctx, orderID, "preparing").Return(entities.Order{}, errors.New("database error"))
			},
			expectedOrder: entities.Order{},
			expectedError: errors.New("database error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockOrderGateway := new(gateways.MockOrderGateway)
			if tc.mockOrderGateway != nil {
				tc.mockOrderGateway(mockOrderGateway)
			}

			useCase := NewUpdateOrderStatusUseCaseImpl(mockOrderGateway)
			order, err := useCase.Execute(ctx, tc.orderID, tc.targetStatus)

			require.Equal(t, tc.expectedError, err)
			require.Equal(t, tc.expectedOrder, order)

			mockOrderGateway.AssertExpectations(t)
		})
	}
}
