package mock

import (
	"context"
	"github.com/stretchr/testify/mock"
	"post-tech-challenge-10soat/app/internal/entities"
)

type MockPaymentGateway struct {
	mock.Mock
}

func (m *MockPaymentGateway) CreatePayment(ctx context.Context, payment entities.Payment) (entities.Payment, error) {
	args := m.Called(ctx, payment)
	return args.Get(0).(entities.Payment), args.Error(1)
}

func (m *MockPaymentGateway) CreatePaymentData(ctx context.Context, paymentData entities.PaymentData) (entities.PaymentData, error) {
	args := m.Called(ctx, paymentData)
	return args.Get(0).(entities.PaymentData), args.Error(1)
}
