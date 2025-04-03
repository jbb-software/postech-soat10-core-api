package product

import (
	"context"
	"post-tech-challenge-10soat/app/internal/entities"
)

type ListProductsUseCase interface {
	Execute(ctx context.Context, categoryId string) ([]entities.Product, error)
}
