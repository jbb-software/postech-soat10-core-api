package mock

import (
	"context"
	"github.com/stretchr/testify/mock"
	"post-tech-challenge-10soat/app/internal/entities"
)

type MockProductGateway struct {
	mock.Mock
}

func (m *MockProductGateway) ListProducts(ctx context.Context, categoryId string) ([]entities.Product, error) {
	args := m.Called(ctx, categoryId)
	return args.Get(0).([]entities.Product), args.Error(1)
}

func (m *MockProductGateway) CreateProduct(ctx context.Context, product entities.Product) (entities.Product, error) {
	args := m.Called(ctx, product)
	return args.Get(0).(entities.Product), args.Error(1)
}

func (m *MockProductGateway) GetProductById(ctx context.Context, id string) (entities.Product, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(entities.Product), args.Error(1)
}

func (m *MockProductGateway) UpdateProduct(ctx context.Context, product entities.Product) (entities.Product, error) {
	args := m.Called(ctx, product)
	return args.Get(0).(entities.Product), args.Error(1)
}

func (m *MockProductGateway) DeleteProduct(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
