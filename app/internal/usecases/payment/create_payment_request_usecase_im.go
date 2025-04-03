package payment

import (
	"context"
	"fmt"
	"post-tech-challenge-10soat/app/internal/dto/payment"
	"post-tech-challenge-10soat/app/internal/entities"
	"post-tech-challenge-10soat/app/internal/interfaces/gateways"
)

type CreatePaymentRequestUseCaseImpl struct {
	gateway interfaces.PaymentGateway
}

func NewCreatePaymentRequestUsecaseImpl(gateway interfaces.PaymentGateway) CreatePaymentRequestUseCase {
	return &CreatePaymentRequestUseCaseImpl{
		gateway,
	}
}

func (s CreatePaymentRequestUseCaseImpl) Execute(ctx context.Context, createPayment dto.CreatePaymentDTO) (entities.Payment, error) {
	paymentInfo := entities.Payment{
		Type:     createPayment.Type,
		Provider: createPayment.Provider,
	}
	payment, err := s.gateway.CreatePayment(ctx, paymentInfo)
	if err != nil {
		if err == entities.ErrConflictingData {
			return entities.Payment{}, err
		}
		return entities.Payment{}, fmt.Errorf("failed to make payment - %s", err.Error())
	}
	return payment, nil
}
