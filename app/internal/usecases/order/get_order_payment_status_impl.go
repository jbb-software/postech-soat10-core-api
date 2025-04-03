package order

import (
	"context"
	"post-tech-challenge-10soat/app/internal/interfaces/gateways"
)

type GetOrderPaymentStatusUseCaseImpl struct {
	orderGateway interfaces.OrderGateway
}

func NewGetOrderPaymentStatusUseCaseImpl(orderGateway interfaces.OrderGateway) GetOrderPaymentStatusUseCase {
	return &GetOrderPaymentStatusUseCaseImpl{
		orderGateway,
	}
}

func (u GetOrderPaymentStatusUseCaseImpl) Execute(ctx context.Context, id string) (OrderPaymentStatus, error) {
	order, err := u.orderGateway.GetOrderById(ctx, id)
	if err != nil {
		return OrderPaymentStatus{}, err
	}
	var paymentStatus PaymentStatus = PaymentPending
	if order.PaymentId != "" {
		paymentStatus = PaymentApproved
	}
	return OrderPaymentStatus{
		PaymentStatus: paymentStatus,
	}, nil
}
