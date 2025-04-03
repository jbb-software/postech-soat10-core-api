package order

import (
	"context"
	"post-tech-challenge-10soat/app/internal/entities"
)

type ListOrdersUseCase interface {
	Execute(ctx context.Context, limit uint64) ([]entities.Order, error)
}
