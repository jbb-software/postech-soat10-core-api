package order

import (
	"context"
	"post-tech-challenge-10soat/app/internal/entities"
)

type UpdateOrderStatusUseCase interface {
	Execute(ctx context.Context, id string, status string) (entities.Order, error)
}
