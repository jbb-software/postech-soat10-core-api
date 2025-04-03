package mock

import (
	"context"
	"github.com/stretchr/testify/mock"
	"post-tech-challenge-10soat/app/internal/entities"
)

type MockCategoryGateway struct {
	mock.Mock
}

func (m *MockCategoryGateway) GetCategoryById(ctx context.Context, id string) (entities.Category, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(entities.Category), args.Error(1)
}
