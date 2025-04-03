package gateways

import (
	"context"
	"post-tech-challenge-10soat/app/internal/dto/order"
	"post-tech-challenge-10soat/app/internal/entities"
	"post-tech-challenge-10soat/app/internal/interfaces/repositories"
)

type OrderProductGatewayImpl struct {
	repository interfaces.OrderProductRepository
}

func NewOrderProductGatewayImpl(repository interfaces.OrderProductRepository) *OrderProductGatewayImpl {
	return &OrderProductGatewayImpl{
		repository,
	}
}

func (og OrderProductGatewayImpl) CreateOrderProduct(ctx context.Context, orderProduct entities.OrderProduct) (entities.OrderProduct, error) {
	orderProductDTO := dto.CreateOrderProductDTO{
		OrderId:     orderProduct.OrderId,
		ProductId:   orderProduct.ProductId,
		Quantity:    orderProduct.Quantity,
		SubTotal:    orderProduct.SubTotal,
		Observation: orderProduct.Observation,
	}
	createdOrderProduct, err := og.repository.CreateOrderProduct(ctx, orderProductDTO)
	if err != nil {
		return entities.OrderProduct{}, nil
	}
	return createdOrderProduct.ToEntity(), nil
}
