package order

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"post-tech-challenge-10soat/app/internal/entities"
	gateways "post-tech-challenge-10soat/app/internal/test_utils/mock"
)

func TestGetOrderPaymentStatusUseCaseImpl_Execute(t *testing.T) {
	ctx := context.Background()
	orderID := uuid.New().String()
	paymentID := uuid.New().String()

	testCases := []struct {
		name             string
		orderID          string
		mockOrderGateway func(orderGateway *gateways.MockOrderGateway)
		expectedStatus   OrderPaymentStatus
		expectedError    error
	}{
		{
			name:    "Payment Pending",
			orderID: orderID,
			mockOrderGateway: func(orderGateway *gateways.MockOrderGateway) {
				orderGateway.On("GetOrderById", ctx, orderID).Return(entities.Order{
					Id: orderID,
				}, nil)
			},
			expectedStatus: OrderPaymentStatus{
				PaymentStatus: PaymentPending,
			},
			expectedError: nil,
		},
		{
			name:    "Payment Approved",
			orderID: orderID,
			mockOrderGateway: func(orderGateway *gateways.MockOrderGateway) {
				orderGateway.On("GetOrderById", ctx, orderID).Return(entities.Order{
					Id:        orderID,
					PaymentId: paymentID,
				}, nil)
			},
			expectedStatus: OrderPaymentStatus{
				PaymentStatus: PaymentApproved,
			},
			expectedError: nil,
		},
		{
			name:    "Order Not Found",
			orderID: orderID,
			mockOrderGateway: func(orderGateway *gateways.MockOrderGateway) {
				orderGateway.On("GetOrderById", ctx, orderID).Return(entities.Order{}, errors.New("order not found"))
			},
			expectedStatus: OrderPaymentStatus{},
			expectedError:  errors.New("order not found"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockOrderGateway := new(gateways.MockOrderGateway)
			if tc.mockOrderGateway != nil {
				tc.mockOrderGateway(mockOrderGateway)
			}

			useCase := NewGetOrderPaymentStatusUseCaseImpl(mockOrderGateway)
			status, err := useCase.Execute(ctx, tc.orderID)

			require.Equal(t, tc.expectedError, err)
			require.Equal(t, tc.expectedStatus, status)

			mockOrderGateway.AssertExpectations(t)
		})
	}
}
