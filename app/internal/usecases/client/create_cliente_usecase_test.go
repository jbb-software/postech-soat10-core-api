package client

import (
	"context"
	"errors"
	"fmt"
	gateways "post-tech-challenge-10soat/app/internal/test_utils/mock"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"post-tech-challenge-10soat/app/internal/dto/client"
	"post-tech-challenge-10soat/app/internal/entities"
)

func TestCreateClientUseCaseImpl_Execute(t *testing.T) {
	ctx := context.Background()

	testCases := []struct {
		name            string
		createClientDTO dto.CreateClientDTO
		mockGateway     func(gateway *gateways.MockClientGateway)
		expectedClient  entities.Client
		expectedError   error
	}{
		{
			name: "Success",
			createClientDTO: dto.CreateClientDTO{
				Cpf:   "12345678900",
				Name:  "John Doe",
				Email: "john.doe@example.com",
			},
			mockGateway: func(gateway *gateways.MockClientGateway) {
				gateway.On("CreateClient", ctx, entities.Client{
					Cpf:   "12345678900",
					Name:  "John Doe",
					Email: "john.doe@example.com",
				}).Return(entities.Client{
					Id:    uuid.New().String(),
					Cpf:   "12345678900",
					Name:  "John Doe",
					Email: "john.doe@example.com",
				}, nil)
			},
			expectedClient: entities.Client{
				Cpf:   "12345678900",
				Name:  "John Doe",
				Email: "john.doe@example.com",
			},
			expectedError: nil,
		},
		{
			name: "Gateway Error",
			createClientDTO: dto.CreateClientDTO{
				Cpf:   "12345678900",
				Name:  "John Doe",
				Email: "john.doe@example.com",
			},
			mockGateway: func(gateway *gateways.MockClientGateway) {
				gateway.On("CreateClient", ctx, entities.Client{
					Cpf:   "12345678900",
					Name:  "John Doe",
					Email: "john.doe@example.com",
				}).Return(entities.Client{}, errors.New("database error"))
			},
			expectedClient: entities.Client{},
			expectedError:  fmt.Errorf("failed to create client - database error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockGateway := new(gateways.MockClientGateway)
			if tc.mockGateway != nil {
				tc.mockGateway(mockGateway)
			}

			usecase := NewCreateClientUsecaseImpl(mockGateway)
			client, err := usecase.Execute(ctx, tc.createClientDTO)

			if err == nil {
				client.Id = "" // ID is generated by gateway, so it is not possible to test
			}

			require.Equal(t, tc.expectedError, err)
			require.Equal(t, tc.expectedClient, client)

			mockGateway.AssertExpectations(t)
		})
	}
}
