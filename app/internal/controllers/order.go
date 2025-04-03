package controllers

import (
	"context"
	"post-tech-challenge-10soat/app/internal/dto/order"
	entity "post-tech-challenge-10soat/app/internal/entities"
	order2 "post-tech-challenge-10soat/app/internal/usecases/order"
)

type OrderController struct {
	createOrder           order2.CreateOrderUseCase
	listOrders            order2.ListOrdersUseCase
	getOrderPaymentStatus order2.GetOrderPaymentStatusUseCase
	updateOrderStatus     order2.UpdateOrderStatusUseCase
}

func NewOrderController(
	createOrder order2.CreateOrderUseCase,
	listOrders order2.ListOrdersUseCase,
	getOrderPaymentStatus order2.GetOrderPaymentStatusUseCase,
	updateOrderStatus order2.UpdateOrderStatusUseCase,
) *OrderController {
	return &OrderController{
		createOrder,
		listOrders,
		getOrderPaymentStatus,
		updateOrderStatus,
	}
}

func (c *OrderController) CreateOrder(ctx context.Context, createOrderDTO dto.CreateOrderDTO) (entity.Order, error) {
	order, err := c.createOrder.Execute(ctx, createOrderDTO)
	if err != nil {
		return entity.Order{}, err
	}
	return order, nil
}

func (c *OrderController) ListOrders(ctx context.Context, limit uint64) ([]entity.Order, error) {
	orders, err := c.listOrders.Execute(ctx, limit)
	if err != nil {
		return []entity.Order{}, err
	}
	return orders, nil
}

func (c *OrderController) GetOrderPaymentStatus(ctx context.Context, id string) (order2.OrderPaymentStatus, error) {
	orderPaymentStatus, err := c.getOrderPaymentStatus.Execute(ctx, id)
	if err != nil {
		return order2.OrderPaymentStatus{}, err
	}
	return orderPaymentStatus, nil
}

func (c *OrderController) UpdateOrderStatus(ctx context.Context, id string, status string) (entity.Order, error) {
	order, err := c.updateOrderStatus.Execute(ctx, id, status)
	if err != nil {
		return entity.Order{}, err
	}
	return order, nil
}
