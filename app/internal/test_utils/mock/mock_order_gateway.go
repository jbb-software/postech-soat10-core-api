package mock

import (
	"context"
	"github.com/stretchr/testify/mock"
	"post-tech-challenge-10soat/app/internal/entities"
)

type MockOrderGateway struct {
	mock.Mock
}

func (m *MockOrderGateway) GetOrderById(ctx context.Context, id string) (entities.Order, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(entities.Order), args.Error(1)
}

func (m *MockOrderGateway) UpdateOrderPayment(ctx context.Context, orderId string, paymentId string) (entities.Order, error) {
	args := m.Called(ctx, orderId, paymentId)
	return args.Get(0).(entities.Order), args.Error(1)
}

func (m *MockOrderGateway) CreateOrder(ctx context.Context, order entities.Order) (entities.Order, error) {
	args := m.Called(ctx, order)
	return args.Get(0).(entities.Order), args.Error(1)
}
func (m *MockOrderGateway) DeleteOrder(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
func (m *MockOrderGateway) ListOrders(ctx context.Context, limit uint64) ([]entities.Order, error) {
	args := m.Called(ctx, limit)
	return args.Get(0).([]entities.Order), args.Error(1)
}
func (m *MockOrderGateway) UpdateOrderStatus(ctx context.Context, id string, status string) (entities.Order, error) {
	args := m.Called(ctx, id, status)
	return args.Get(0).(entities.Order), args.Error(1)
}
