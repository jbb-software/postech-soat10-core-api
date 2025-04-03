package clients

import (
	"context"
	dto2 "post-tech-challenge-10soat/app/internal/dto/payment"
)

type PaymentClient interface {
	CreatePaymentData(ctx context.Context, paymentData dto2.CreatePaymentDataDTO) (dto2.PaymentDataDTO, error)
}
