package gateways

import (
	"context"
	"post-tech-challenge-10soat/app/internal/dto/order"
	"post-tech-challenge-10soat/app/internal/entities"
	"post-tech-challenge-10soat/app/internal/interfaces/repositories"
)

type OrderGatewayImpl struct {
	repository interfaces.OrderRepository
}

func NewOrderGatewayImpl(repository interfaces.OrderRepository) *OrderGatewayImpl {
	return &OrderGatewayImpl{
		repository,
	}
}

func (og OrderGatewayImpl) CreateOrder(ctx context.Context, order entities.Order) (entities.Order, error) {
	createOrderDTO := dto.CreateOrderDTO{
		Status:   string(order.Status),
		ClientId: order.ClientId,
		Total:    order.Total,
	}
	createdOrder, err := og.repository.CreateOrder(ctx, createOrderDTO)
	if err != nil {
		return entities.Order{}, err
	}
	return createdOrder.ToEntity(), nil
}

func (og OrderGatewayImpl) DeleteOrder(ctx context.Context, id string) error {
	err := og.repository.DeleteOrder(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (og OrderGatewayImpl) ListOrders(ctx context.Context, limit uint64) ([]entities.Order, error) {
	orders, err := og.repository.ListOrders(ctx, limit)
	var ordersRes []entities.Order
	if err != nil {
		return []entities.Order{}, err
	}
	for _, order := range orders {
		ordersRes = append(ordersRes, order.ToEntity())
	}
	return ordersRes, nil
}

func (og OrderGatewayImpl) GetOrderById(ctx context.Context, id string) (entities.Order, error) {
	order, err := og.repository.GetOrderById(ctx, id)
	if err != nil {
		return entities.Order{}, err
	}
	return order.ToEntity(), nil
}

func (og OrderGatewayImpl) UpdateOrderStatus(ctx context.Context, id string, status string) (entities.Order, error) {
	order, err := og.repository.UpdateOrderStatus(ctx, id, status)
	if err != nil {
		return entities.Order{}, err
	}
	return order.ToEntity(), nil
}

func (og OrderGatewayImpl) UpdateOrderPayment(ctx context.Context, id string, paymentId string) (entities.Order, error) {
	order, err := og.repository.UpdateOrderPayment(ctx, id, paymentId)
	if err != nil {
		return entities.Order{}, err
	}
	return order.ToEntity(), nil
}
