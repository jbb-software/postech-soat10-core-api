package payment

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	dto2 "post-tech-challenge-10soat/app/internal/dto/payment"
	"post-tech-challenge-10soat/app/internal/entities"
	gateways "post-tech-challenge-10soat/app/internal/test_utils/mock"
)

func TestProcessPaymentResponseUseCaseImpl_Execute(t *testing.T) {
	ctx := context.Background()
	orderID := "123"
	paymentID := "456"

	testCases := []struct {
		name               string
		processPaymentDTO  dto2.ProcessPaymentDTO
		mockPaymentGateway func(paymentGateway *gateways.MockPaymentGateway)
		mockOrderGateway   func(orderGateway *gateways.MockOrderGateway)
		expectedResponse   dto2.ProcessPaymentResponseDTO
		expectedError      error
	}{
		{
			name: "Success - Approved Payment",
			processPaymentDTO: dto2.ProcessPaymentDTO{
				OrderId:  orderID,
				Provider: ValidPaymentProvider,
				Status:   "approved",
			},
			mockOrderGateway: func(orderGateway *gateways.MockOrderGateway) {
				orderGateway.On("GetOrderById", ctx, orderID).Return(entities.Order{
					Id:     orderID,
					Status: entities.OrderStatusPaymentPending,
				}, nil)
				orderGateway.On("UpdateOrderPayment", ctx, orderID, paymentID).Return(entities.Order{}, nil)
			},
			mockPaymentGateway: func(paymentGateway *gateways.MockPaymentGateway) {
				paymentGateway.On("CreatePayment", ctx, entities.Payment{
					Provider: ValidPaymentProvider,
					Type:     entities.PaymentTypePixQRCode,
				}).Return(entities.Payment{
					Id: paymentID,
				}, nil)
			},
			expectedResponse: dto2.ProcessPaymentResponseDTO{
				Status:  dto2.Processed,
				Message: "Payment response processed with success",
			},
			expectedError: nil,
		},
		{
			name: "Denied Payment",
			processPaymentDTO: dto2.ProcessPaymentDTO{
				OrderId:  orderID,
				Provider: ValidPaymentProvider,
				Status:   "denied",
			},
			mockOrderGateway: func(orderGateway *gateways.MockOrderGateway) {
				orderGateway.On("GetOrderById", ctx, orderID).Return(entities.Order{
					Id:     orderID,
					Status: entities.OrderStatusPaymentPending,
				}, nil)
			},
			mockPaymentGateway: func(paymentGateway *gateways.MockPaymentGateway) {},
			expectedResponse: dto2.ProcessPaymentResponseDTO{
				Status:  "denied",
				Message: "Cannot process unaproved payment",
			},
			expectedError: nil,
		},
		{
			name: "Order Not Found",
			processPaymentDTO: dto2.ProcessPaymentDTO{
				OrderId:  orderID,
				Provider: ValidPaymentProvider,
				Status:   "approved",
			},
			mockOrderGateway: func(orderGateway *gateways.MockOrderGateway) {
				orderGateway.On("GetOrderById", ctx, orderID).Return(entities.Order{}, errors.New("order not found"))
			},
			mockPaymentGateway: func(paymentGateway *gateways.MockPaymentGateway) {},
			expectedResponse:   dto2.ProcessPaymentResponseDTO{},
			expectedError:      errors.New("cannot get order"),
		},
		{
			name: "Invalid Order Status",
			processPaymentDTO: dto2.ProcessPaymentDTO{
				OrderId:  orderID,
				Provider: ValidPaymentProvider,
				Status:   "approved",
			},
			mockOrderGateway: func(orderGateway *gateways.MockOrderGateway) {
				orderGateway.On("GetOrderById", ctx, orderID).Return(entities.Order{
					Id:     orderID,
					Status: entities.OrderStatusCompleted,
				}, nil)
			},
			mockPaymentGateway: func(paymentGateway *gateways.MockPaymentGateway) {},
			expectedResponse:   dto2.ProcessPaymentResponseDTO{},
			expectedError:      errors.New("cannot process payment for this order state"),
		},
		{
			name: "Invalid Payment Provider",
			processPaymentDTO: dto2.ProcessPaymentDTO{
				OrderId:  orderID,
				Provider: "invalid-provider",
				Status:   "approved",
			},
			mockOrderGateway: func(orderGateway *gateways.MockOrderGateway) {
				orderGateway.On("GetOrderById", ctx, orderID).Return(entities.Order{
					Id:     orderID,
					Status: entities.OrderStatusPaymentPending,
				}, nil)
			},
			mockPaymentGateway: func(paymentGateway *gateways.MockPaymentGateway) {},
			expectedResponse:   dto2.ProcessPaymentResponseDTO{},
			expectedError:      errors.New("cannot process payment for invalid provider"),
		},
		{
			name: "Payment Creation Error",
			processPaymentDTO: dto2.ProcessPaymentDTO{
				OrderId:  orderID,
				Provider: ValidPaymentProvider,
				Status:   "approved",
			},
			mockOrderGateway: func(orderGateway *gateways.MockOrderGateway) {
				orderGateway.On("GetOrderById", ctx, orderID).Return(entities.Order{
					Id:     orderID,
					Status: entities.OrderStatusPaymentPending,
				}, nil)
			},
			mockPaymentGateway: func(paymentGateway *gateways.MockPaymentGateway) {
				paymentGateway.On("CreatePayment", ctx, entities.Payment{
					Provider: ValidPaymentProvider,
					Type:     entities.PaymentTypePixQRCode,
				}).Return(entities.Payment{}, errors.New("payment creation error"))
			},
			expectedResponse: dto2.ProcessPaymentResponseDTO{},
			expectedError:    errors.New("cannot create payment"),
		},
		{
			name: "Order Update Error",
			processPaymentDTO: dto2.ProcessPaymentDTO{
				OrderId:  orderID,
				Provider: ValidPaymentProvider,
				Status:   "approved",
			},
			mockOrderGateway: func(orderGateway *gateways.MockOrderGateway) {
				orderGateway.On("GetOrderById", ctx, orderID).Return(entities.Order{
					Id:     orderID,
					Status: entities.OrderStatusPaymentPending,
				}, nil)
				orderGateway.On("UpdateOrderPayment", ctx, orderID, paymentID).Return(entities.Order{}, errors.New("order update error"))
			},
			mockPaymentGateway: func(paymentGateway *gateways.MockPaymentGateway) {
				paymentGateway.On("CreatePayment", ctx, entities.Payment{
					Provider: ValidPaymentProvider,
					Type:     entities.PaymentTypePixQRCode,
				}).Return(entities.Payment{
					Id: paymentID,
				}, nil)
			},
			expectedResponse: dto2.ProcessPaymentResponseDTO{},
			expectedError:    errors.New("cannot update order with payment"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockPaymentGateway := new(gateways.MockPaymentGateway)
			mockOrderGateway := new(gateways.MockOrderGateway)

			if tc.mockPaymentGateway != nil {
				tc.mockPaymentGateway(mockPaymentGateway)
			}
			if tc.mockOrderGateway != nil {
				tc.mockOrderGateway(mockOrderGateway)
			}

			useCase := NewProcessPaymentResponseUseCaseImpl(mockPaymentGateway, mockOrderGateway)
			response, err := useCase.Execute(ctx, tc.processPaymentDTO)

			require.Equal(t, tc.expectedError, err)
			require.Equal(t, tc.expectedResponse, response)
		})
	}
}
