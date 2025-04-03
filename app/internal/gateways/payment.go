package gateways

import (
	"context"
	"errors"
	dto2 "post-tech-challenge-10soat/app/internal/dto/payment"
	"post-tech-challenge-10soat/app/internal/entities"
	"post-tech-challenge-10soat/app/internal/interfaces/clients"
	"post-tech-challenge-10soat/app/internal/interfaces/repositories"
)

type PaymentGatewayImpl struct {
	repository    interfaces.PaymentRepository
	paymentClient clients.PaymentClient
}

func NewPaymentGatewayImpl(repository interfaces.PaymentRepository, paymentClient clients.PaymentClient) *PaymentGatewayImpl {
	return &PaymentGatewayImpl{
		repository,
		paymentClient,
	}
}

func (pg PaymentGatewayImpl) CreatePayment(ctx context.Context, payment entities.Payment) (entities.Payment, error) {
	createPaymentDTO := dto2.CreatePaymentDTO{
		Provider: payment.Provider,
		Type:     payment.Type,
	}
	createdPayment, err := pg.repository.CreatePayment(ctx, createPaymentDTO)
	if err != nil {
		return entities.Payment{}, errors.New("cannot create payment")
	}
	return createdPayment.ToEntity(), nil
}

func (pg PaymentGatewayImpl) CreatePaymentData(ctx context.Context, paymentData entities.PaymentData) (entities.PaymentData, error) {
	createPaymentDataDTO := dto2.CreatePaymentDataDTO{
		OrderId: paymentData.OrderId,
		Total:   paymentData.Total,
	}
	createdPaymentData, err := pg.paymentClient.CreatePaymentData(ctx, createPaymentDataDTO)
	if err != nil {
		return entities.PaymentData{}, errors.New("cannot create payment data")
	}
	return createdPaymentData.ToEntity(), nil
}
