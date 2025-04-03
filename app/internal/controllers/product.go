package controllers

import (
	"context"
	dto2 "post-tech-challenge-10soat/app/internal/dto/product"
	entity "post-tech-challenge-10soat/app/internal/entities"
	product2 "post-tech-challenge-10soat/app/internal/usecases/product"
)

type ProductController struct {
	createProduct product2.CreateProductUseCase
	deleteProduct product2.DeleteProductUseCase
	updateProduct product2.UpdateProductUseCase
	listProducts  product2.ListProductsUseCase
}

func NewProductController(
	createProduct product2.CreateProductUseCase,
	deleteProduct product2.DeleteProductUseCase,
	updateProduct product2.UpdateProductUseCase,
	listProducts product2.ListProductsUseCase,
) *ProductController {
	return &ProductController{
		createProduct,
		deleteProduct,
		updateProduct,
		listProducts,
	}
}

func (c *ProductController) CreateProduct(ctx context.Context, createProductDTO dto2.CreateProductDTO) (entity.Product, error) {
	product, err := c.createProduct.Execute(ctx, createProductDTO)
	if err != nil {
		return entity.Product{}, err
	}
	return product, nil
}

func (c *ProductController) DeleteProduct(ctx context.Context, id string) error {
	err := c.deleteProduct.Execute(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (c *ProductController) UpdateProduct(ctx context.Context, updateProductDTO dto2.UpdateProductDTO) (entity.Product, error) {
	product, err := c.updateProduct.Execute(ctx, updateProductDTO)
	if err != nil {
		return entity.Product{}, err
	}
	return product, nil
}

func (c *ProductController) ListProducts(ctx context.Context, categoryId string) ([]entity.Product, error) {
	products, err := c.listProducts.Execute(ctx, categoryId)
	if err != nil {
		return []entity.Product{}, err
	}
	return products, nil
}
