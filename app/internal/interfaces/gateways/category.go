package interfaces

import (
	"context"
	"post-tech-challenge-10soat/app/internal/entities"
)

type CategoryGateway interface {
	GetCategoryById(ctx context.Context, categoryId string) (entities.Category, error)
}
