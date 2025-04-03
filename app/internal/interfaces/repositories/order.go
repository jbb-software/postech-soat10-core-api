package interfaces

import (
	"context"
	dto2 "post-tech-challenge-10soat/app/internal/dto/order"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, order dto2.CreateOrderDTO) (dto2.OrderDTO, error)
	DeleteOrder(ctx context.Context, id string) error
	ListOrders(ctx context.Context, limit uint64) ([]dto2.OrderDTO, error)
	GetOrderById(ctx context.Context, id string) (dto2.OrderDTO, error)
	UpdateOrderStatus(ctx context.Context, id string, status string) (dto2.OrderDTO, error)
	UpdateOrderPayment(ctx context.Context, id string, paymentId string) (dto2.OrderDTO, error)
}
