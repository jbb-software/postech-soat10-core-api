package interfaces

import (
	"context"
	"post-tech-challenge-10soat/app/internal/entities"
)

type PaymentGateway interface {
	CreatePayment(ctx context.Context, payment entities.Payment) (entities.Payment, error)
	CreatePaymentData(ctx context.Context, paymentData entities.PaymentData) (entities.PaymentData, error)
}
