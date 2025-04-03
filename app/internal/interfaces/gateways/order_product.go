package interfaces

import (
	"context"
	"post-tech-challenge-10soat/app/internal/entities"
)

type OrderProductGateway interface {
	CreateOrderProduct(ctx context.Context, orderProduct entities.OrderProduct) (entities.OrderProduct, error)
}
