package interfaces

import (
	"context"
	dto2 "post-tech-challenge-10soat/app/internal/dto/product"
)

type ProductRepository interface {
	ListProducts(ctx context.Context, categoryId string) ([]dto2.ProductDTO, error)
	GetProductById(ctx context.Context, id string) (dto2.ProductDTO, error)
	CreateProduct(ctx context.Context, product dto2.CreateProductDTO) (dto2.ProductDTO, error)
	UpdateProduct(ctx context.Context, product dto2.UpdateProductDTO) (dto2.ProductDTO, error)
	DeleteProduct(ctx context.Context, id string) error
}
