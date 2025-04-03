package interfaces

import (
	"context"
	dto2 "post-tech-challenge-10soat/app/internal/dto/order"
)

type OrderProductRepository interface {
	CreateOrderProduct(ctx context.Context, orderProduct dto2.CreateOrderProductDTO) (dto2.OrderProductDTO, error)
}
