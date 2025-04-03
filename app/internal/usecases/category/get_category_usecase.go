package category

import (
	"context"
	"post-tech-challenge-10soat/app/internal/entities"
)

type GetCategoryUseCase interface {
	Execute(ctx context.Context, id string) (entities.Category, error)
}
