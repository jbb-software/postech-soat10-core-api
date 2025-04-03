package order

import (
	"context"
	"post-tech-challenge-10soat/app/internal/entities"
	"post-tech-challenge-10soat/app/internal/interfaces/gateways"
	"sort"
)

type ListOrdersUseCaseImpl struct {
	orderGateway interfaces.OrderGateway
}

func NewListOrdersUseCaseImpl(orderGateway interfaces.OrderGateway) ListOrdersUseCase {
	return &ListOrdersUseCaseImpl{
		orderGateway,
	}
}

func (l ListOrdersUseCaseImpl) Execute(ctx context.Context, limit uint64) ([]entities.Order, error) {
	orders, err := l.orderGateway.ListOrders(ctx, limit)
	if err != nil {
		return []entities.Order{}, err
	}
	sortOrdersbyStatus(orders)
	return orders, nil
}

func sortOrdersbyStatus(orders []entities.Order) {
	statusPriority := map[string]int{
		"ready":     1,
		"preparing": 2,
		"received":  3,
	}
	sort.SliceStable(orders, func(i, j int) bool {
		if statusPriority[string(orders[i].Status)] != statusPriority[string(orders[j].Status)] {
			return statusPriority[string(orders[i].Status)] < statusPriority[string(orders[j].Status)]
		}
		return orders[i].CreatedAt.Before(orders[j].CreatedAt)
	})
}
