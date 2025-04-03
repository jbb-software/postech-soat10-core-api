package mock

import (
	"context"
	"github.com/stretchr/testify/mock"
	"post-tech-challenge-10soat/app/internal/entities"
)

type MockClientGateway struct {
	mock.Mock
}

func (m *MockClientGateway) GetClientById(ctx context.Context, id string) (entities.Client, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(entities.Client), args.Error(1)
}

func (m *MockClientGateway) CreateClient(ctx context.Context, client entities.Client) (entities.Client, error) {
	args := m.Called(ctx, client)
	return args.Get(0).(entities.Client), args.Error(1)
}

func (m *MockClientGateway) GetClientByCpf(ctx context.Context, cpf string) (entities.Client, error) {
	args := m.Called(ctx, cpf)
	return args.Get(0).(entities.Client), args.Error(1)
}
