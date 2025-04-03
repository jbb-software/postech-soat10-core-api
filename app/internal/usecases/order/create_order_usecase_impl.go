package order

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"post-tech-challenge-10soat/app/internal/dto/order"
	"post-tech-challenge-10soat/app/internal/entities"
	interfaces2 "post-tech-challenge-10soat/app/internal/interfaces/gateways"
)

type CreateOrderUsecaseImpl struct {
	productGateway      interfaces2.ProductGateway
	clientGateway       interfaces2.ClientGateway
	orderGateway        interfaces2.OrderGateway
	orderProductGateway interfaces2.OrderProductGateway
	paymentGateway      interfaces2.PaymentGateway
}

func NewCreateOrderUsecaseImpl(
	productGateway interfaces2.ProductGateway,
	clientGateway interfaces2.ClientGateway,
	orderGateway interfaces2.OrderGateway,
	orderProductGateway interfaces2.OrderProductGateway,
	paymentGateway interfaces2.PaymentGateway,
) CreateOrderUseCase {
	return &CreateOrderUsecaseImpl{
		productGateway,
		clientGateway,
		orderGateway,
		orderProductGateway,
		paymentGateway,
	}
}

func (s CreateOrderUsecaseImpl) Execute(ctx context.Context, createOrder dto.CreateOrderDTO) (entities.Order, error) {
	var totalValue float64
	for _, orderProduct := range createOrder.Products {
		product, err := s.productGateway.GetProductById(ctx, orderProduct.ProductId)
		if err != nil {
			if err == entities.ErrDataNotFound {
				return entities.Order{}, err
			}
			return entities.Order{}, fmt.Errorf("cannot create order because has invalid product - %s", err.Error())
		}
		subTotal := product.Value * float64(orderProduct.Quantity)
		totalValue += subTotal
	}

	orderInfo := entities.Order{
		Status: entities.OrderStatusPaymentPending,
		Total:  totalValue,
	}
	if createOrder.ClientId != "" && uuid.Validate(createOrder.ClientId) == nil {
		client, err := s.clientGateway.GetClientById(ctx, createOrder.ClientId)
		if err != nil {
			if err == entities.ErrDataNotFound {
				return entities.Order{}, err
			}
			return entities.Order{}, fmt.Errorf("cannot create order because has invalid client - %s", err.Error())
		}
		orderInfo.ClientId = client.Id
	} else {
		orderInfo.ClientId = ""
	}
	order, err := s.orderGateway.CreateOrder(ctx, orderInfo)
	if err != nil {
		if err == entities.ErrDataNotFound {
			return entities.Order{}, err
		}
		return entities.Order{}, fmt.Errorf("cannot create order - %s", err.Error())
	}
	for _, orderProduct := range createOrder.Products {
		product, err := s.productGateway.GetProductById(ctx, orderProduct.ProductId)
		if err != nil {
			if err == entities.ErrDataNotFound {
				return entities.Order{}, err
			}
			return entities.Order{}, fmt.Errorf("cannot create order because has invalid product - %s", err.Error())
		}
		subTotal := product.Value * float64(orderProduct.Quantity)
		orderProductInfo := entities.OrderProduct{
			OrderId:     order.Id,
			ProductId:   product.Id,
			Quantity:    orderProduct.Quantity,
			SubTotal:    subTotal,
			Observation: orderProduct.Observation,
		}
		_, err = s.orderProductGateway.CreateOrderProduct(ctx, orderProductInfo)
		if err != nil {
			if err == entities.ErrDataNotFound {
				return entities.Order{}, err
			}
			err := s.orderGateway.DeleteOrder(ctx, order.Id)
			if err != nil {
				return entities.Order{}, err
			}
			return entities.Order{}, fmt.Errorf("cannot complete order - %s", err.Error())
		}
	}
	paymentData := entities.PaymentData{
		OrderId: order.Id,
		Total:   order.Total,
	}
	createdPaymentData, err := s.paymentGateway.CreatePaymentData(ctx, paymentData)
	if err == nil {
		order.PaymentData = createdPaymentData
	}
	return order, nil
}
