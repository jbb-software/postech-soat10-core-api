package interfaces

import (
	"context"
	"post-tech-challenge-10soat/app/internal/entities"
)

type OrderGateway interface {
	CreateOrder(ctx context.Context, order entities.Order) (entities.Order, error)
	DeleteOrder(ctx context.Context, id string) error
	ListOrders(ctx context.Context, limit uint64) ([]entities.Order, error)
	GetOrderById(ctx context.Context, id string) (entities.Order, error)
	UpdateOrderStatus(ctx context.Context, id string, status string) (entities.Order, error)
	UpdateOrderPayment(ctx context.Context, id string, paymentId string) (entities.Order, error)
}
