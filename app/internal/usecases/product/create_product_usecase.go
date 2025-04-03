package product

import (
	"context"
	"post-tech-challenge-10soat/app/internal/dto/product"
	"post-tech-challenge-10soat/app/internal/entities"
)

type CreateProductUseCase interface {
	Execute(ctx context.Context, product dto.CreateProductDTO) (entities.Product, error)
}
