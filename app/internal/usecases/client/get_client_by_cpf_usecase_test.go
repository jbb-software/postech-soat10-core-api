package client

import (
	"context"
	"errors"
	"fmt"
	gateways "post-tech-challenge-10soat/app/internal/test_utils/mock"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"post-tech-challenge-10soat/app/internal/entities"
)

func TestGetClientByCpfUseCaseImpl_Execute(t *testing.T) {
	ctx := context.Background()
	cpf := "12345678900"
	clientID := uuid.New().String()

	testCases := []struct {
		name           string
		cpf            string
		mockGateway    func(gateway *gateways.MockClientGateway)
		expectedClient entities.Client
		expectedError  error
	}{
		{
			name: "Success",
			cpf:  cpf,
			mockGateway: func(gateway *gateways.MockClientGateway) {
				gateway.On("GetClientByCpf", ctx, cpf).Return(entities.Client{
					Id:    clientID,
					Cpf:   cpf,
					Name:  "John Doe",
					Email: "john.doe@example.com",
				}, nil)
			},
			expectedClient: entities.Client{
				Id:    clientID,
				Cpf:   cpf,
				Name:  "John Doe",
				Email: "john.doe@example.com",
			},
			expectedError: nil,
		},
		{
			name: "Client Not Found",
			cpf:  "99999999999",
			mockGateway: func(gateway *gateways.MockClientGateway) {
				gateway.On("GetClientByCpf", ctx, "99999999999").Return(entities.Client{}, entities.ErrDataNotFound)
			},
			expectedClient: entities.Client{},
			expectedError:  fmt.Errorf("failed to get client by cpf - %s", entities.ErrDataNotFound.Error()),
		},
		{
			name: "Gateway Error",
			cpf:  cpf,
			mockGateway: func(gateway *gateways.MockClientGateway) {
				gateway.On("GetClientByCpf", ctx, cpf).Return(entities.Client{}, errors.New("database error"))
			},
			expectedClient: entities.Client{},
			expectedError:  fmt.Errorf("failed to get client by cpf - database error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockGateway := new(gateways.MockClientGateway)
			if tc.mockGateway != nil {
				tc.mockGateway(mockGateway)
			}

			usecase := NewGetClientByCpfUseCaseImpl(mockGateway)
			client, err := usecase.Execute(ctx, tc.cpf)

			require.Equal(t, tc.expectedError, err)
			require.Equal(t, tc.expectedClient, client)

			mockGateway.AssertExpectations(t)
		})
	}
}
