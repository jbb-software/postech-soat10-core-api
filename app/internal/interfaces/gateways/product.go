package interfaces

import (
	"context"
	"post-tech-challenge-10soat/app/internal/entities"
)

type ProductGateway interface {
	ListProducts(ctx context.Context, categoryId string) ([]entities.Product, error)
	GetProductById(ctx context.Context, id string) (entities.Product, error)
	CreateProduct(ctx context.Context, product entities.Product) (entities.Product, error)
	UpdateProduct(ctx context.Context, product entities.Product) (entities.Product, error)
	DeleteProduct(ctx context.Context, id string) error
}
