package order

import (
	"context"
	"post-tech-challenge-10soat/app/internal/dto/order"
	"post-tech-challenge-10soat/app/internal/entities"
)

type CreateOrderUseCase interface {
	Execute(ctx context.Context, createOrder dto.CreateOrderDTO) (entities.Order, error)
}
