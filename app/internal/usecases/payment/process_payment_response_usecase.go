package payment

import (
	"context"
	dto2 "post-tech-challenge-10soat/app/internal/dto/payment"
)

type ProcessPaymentResponseUseCase interface {
	Execute(ctx context.Context, processPayment dto2.ProcessPaymentDTO) (dto2.ProcessPaymentResponseDTO, error)
}
