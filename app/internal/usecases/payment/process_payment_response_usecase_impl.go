package payment

import (
	"context"
	"errors"
	dto2 "post-tech-challenge-10soat/app/internal/dto/payment"
	"post-tech-challenge-10soat/app/internal/entities"
	interfaces2 "post-tech-challenge-10soat/app/internal/interfaces/gateways"
)

const (
	ValidPaymentProvider = "mercado-pago"
)

type ProcessPaymentResponseUseCaseImpl struct {
	paymentGateway interfaces2.PaymentGateway
	orderGateway   interfaces2.OrderGateway
}

func NewProcessPaymentResponseUseCaseImpl(paymentGateway interfaces2.PaymentGateway, orderGateway interfaces2.OrderGateway) ProcessPaymentResponseUseCase {
	return &ProcessPaymentResponseUseCaseImpl{
		paymentGateway,
		orderGateway,
	}
}

func (p *ProcessPaymentResponseUseCaseImpl) Execute(ctx context.Context, processPayment dto2.ProcessPaymentDTO) (dto2.ProcessPaymentResponseDTO, error) {
	order, err := p.orderGateway.GetOrderById(ctx, processPayment.OrderId)
	if err != nil {
		return dto2.ProcessPaymentResponseDTO{}, errors.New("cannot get order")
	}
	if order.Status != entities.OrderStatusPaymentPending {
		return dto2.ProcessPaymentResponseDTO{}, errors.New("cannot process payment for this order state")
	}
	if processPayment.Provider != ValidPaymentProvider {
		return dto2.ProcessPaymentResponseDTO{}, errors.New("cannot process payment for invalid provider")
	}

	if processPayment.Status == "approved" {
		payment := entities.Payment{
			Provider: processPayment.Provider,
			Type:     entities.PaymentTypePixQRCode,
		}
		createdPayment, err := p.paymentGateway.CreatePayment(ctx, payment)
		if err != nil {
			return dto2.ProcessPaymentResponseDTO{}, errors.New("cannot create payment")
		}
		_, err = p.orderGateway.UpdateOrderPayment(ctx, order.Id, createdPayment.Id)
		if err != nil {
			return dto2.ProcessPaymentResponseDTO{}, errors.New("cannot update order with payment")
		}
		return dto2.ProcessPaymentResponseDTO{
			Status:  dto2.Processed,
			Message: "Payment response processed with success",
		}, nil
	}
	return dto2.ProcessPaymentResponseDTO{
		Status:  dto2.Denied,
		Message: "Cannot process unaproved payment",
	}, nil
}
