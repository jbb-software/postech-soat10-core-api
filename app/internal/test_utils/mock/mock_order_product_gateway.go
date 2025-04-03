package mock

import (
	"context"
	"github.com/stretchr/testify/mock"
	"post-tech-challenge-10soat/app/internal/entities"
)

type MockOrderProductGateway struct {
	mock.Mock
}

func (m *MockOrderProductGateway) CreateOrderProduct(ctx context.Context, orderProduct entities.OrderProduct) (entities.OrderProduct, error) {
	args := m.Called(ctx, orderProduct)
	return args.Get(0).(entities.OrderProduct), args.Error(1)
}
