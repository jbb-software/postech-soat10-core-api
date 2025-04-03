package interfaces

import (
	"context"
	dto2 "post-tech-challenge-10soat/app/internal/dto/payment"
)

type PaymentRepository interface {
	CreatePayment(ctx context.Context, payment dto2.CreatePaymentDTO) (dto2.PaymentDTO, error)
}
