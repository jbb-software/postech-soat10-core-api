package interfaces

import (
	"context"
	"post-tech-challenge-10soat/app/internal/dto/category"
)

type CategoryRepository interface {
	GetCategoryById(ctx context.Context, categoryId string) (dto.CategoryDTO, error)
}
