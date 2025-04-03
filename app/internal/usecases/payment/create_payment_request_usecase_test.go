package payment

import (
	"context"
	"errors"
	"fmt"
	"github.com/stretchr/testify/require"
	"post-tech-challenge-10soat/app/internal/dto/payment"
	"post-tech-challenge-10soat/app/internal/entities"
	gateways "post-tech-challenge-10soat/app/internal/test_utils/mock"
	"testing"
)

func TestCreatePaymentRequestUseCaseImpl_Execute(t *testing.T) {
	ctx := context.Background()

	testCases := []struct {
		name               string
		createPaymentDTO   dto.CreatePaymentDTO
		mockPaymentGateway func(paymentGateway *gateways.MockPaymentGateway)
		expectedPayment    entities.Payment
		expectedError      error
	}{
		{
			name: "Success",
			createPaymentDTO: dto.CreatePaymentDTO{
				Type:     "credit_card",
				Provider: "visa",
			},
			mockPaymentGateway: func(paymentGateway *gateways.MockPaymentGateway) {
				paymentGateway.On("CreatePayment", ctx, entities.Payment{
					Type:     "credit_card",
					Provider: "visa",
				}).Return(entities.Payment{
					Id:       "123",
					Type:     "credit_card",
					Provider: "visa",
				}, nil)
			},
			expectedPayment: entities.Payment{
				Id:       "123",
				Type:     "credit_card",
				Provider: "visa",
			},
			expectedError: nil,
		},
		{
			name: "Conflicting Data",
			createPaymentDTO: dto.CreatePaymentDTO{
				Type:     "credit_card",
				Provider: "visa",
			},
			mockPaymentGateway: func(paymentGateway *gateways.MockPaymentGateway) {
				paymentGateway.On("CreatePayment", ctx, entities.Payment{
					Type:     "credit_card",
					Provider: "visa",
				}).Return(entities.Payment{}, entities.ErrConflictingData)
			},
			expectedPayment: entities.Payment{},
			expectedError:   entities.ErrConflictingData,
		},
		{
			name: "Payment Gateway Error",
			createPaymentDTO: dto.CreatePaymentDTO{
				Type:     "credit_card",
				Provider: "visa",
			},
			mockPaymentGateway: func(paymentGateway *gateways.MockPaymentGateway) {
				paymentGateway.On("CreatePayment", ctx, entities.Payment{
					Type:     "credit_card",
					Provider: "visa",
				}).Return(entities.Payment{}, errors.New("database error"))
			},
			expectedPayment: entities.Payment{},
			expectedError:   fmt.Errorf("failed to make payment - database error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockPaymentGateway := new(gateways.MockPaymentGateway)

			if tc.mockPaymentGateway != nil {
				tc.mockPaymentGateway(mockPaymentGateway)
			}

			useCase := NewCreatePaymentRequestUsecaseImpl(mockPaymentGateway)
			payment, err := useCase.Execute(ctx, tc.createPaymentDTO)

			require.Equal(t, tc.expectedError, err)
			require.Equal(t, tc.expectedPayment, payment)

			mockPaymentGateway.AssertExpectations(t)
		})
	}
}
