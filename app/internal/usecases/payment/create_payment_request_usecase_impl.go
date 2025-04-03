package payment

import (
	"context"
	"post-tech-challenge-10soat/app/internal/dto/payment"
	"post-tech-challenge-10soat/app/internal/entities"
)

type CreatePaymentRequestUseCase interface {
	Execute(ctx context.Context, createPayment dto.CreatePaymentDTO) (entities.Payment, error)
}
