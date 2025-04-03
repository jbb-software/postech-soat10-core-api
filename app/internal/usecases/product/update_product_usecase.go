package product

import (
	"context"
	"post-tech-challenge-10soat/app/internal/dto/product"
	"post-tech-challenge-10soat/app/internal/entities"
)

type UpdateProductUseCase interface {
	Execute(ctx context.Context, product dto.UpdateProductDTO) (entities.Product, error)
}
