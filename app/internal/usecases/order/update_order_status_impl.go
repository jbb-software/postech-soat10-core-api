package order

import (
	"context"
	"errors"
	"fmt"
	"post-tech-challenge-10soat/app/internal/entities"
	"post-tech-challenge-10soat/app/internal/interfaces/gateways"
	"post-tech-challenge-10soat/app/internal/utils"
)

type UpdateOrderStatusUseCaseImpl struct {
	orderGateway interfaces.OrderGateway
}

func NewUpdateOrderStatusUseCaseImpl(orderGateway interfaces.OrderGateway) UpdateOrderStatusUseCase {
	return &UpdateOrderStatusUseCaseImpl{
		orderGateway,
	}
}

func (u UpdateOrderStatusUseCaseImpl) Execute(ctx context.Context, id string, status string) (entities.Order, error) {
	order, err := u.orderGateway.GetOrderById(ctx, id)
	if err != nil {
		return entities.Order{}, err
	}
	validTransitions := map[string][]string{
		"received":  {"preparing"},
		"preparing": {"ready"},
		"ready":     {"completed"},
		"completed": {},
	}
	allowed, exists := validTransitions[string(order.Status)]
	if !exists {
		return entities.Order{}, errors.New("invalid status")
	}
	if !utils.Contains(allowed, status) {
		return entities.Order{}, fmt.Errorf("cannot update order status for '%s' to '%s'", order.Status, status)
	}
	updatedOrder, err := u.orderGateway.UpdateOrderStatus(ctx, id, status)
	if err != nil {
		return entities.Order{}, err
	}
	return updatedOrder, nil
}
