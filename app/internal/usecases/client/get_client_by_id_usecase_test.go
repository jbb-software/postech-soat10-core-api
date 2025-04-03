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

func TestGetClientByIdUseCaseImpl_Execute(t *testing.T) {
	ctx := context.Background()
	clientID := uuid.New().String()

	testCases := []struct {
		name           string
		clientID       string
		mockGateway    func(gateway *gateways.MockClientGateway)
		expectedClient entities.Client
		expectedError  error
	}{
		{
			name:     "Success",
			clientID: clientID,
			mockGateway: func(gateway *gateways.MockClientGateway) {
				gateway.On("GetClientById", ctx, clientID).Return(entities.Client{
					Id:    clientID,
					Cpf:   "12345678900",
					Name:  "John Doe",
					Email: "john.doe@example.com",
				}, nil)
			},
			expectedClient: entities.Client{
				Id:    clientID,
				Cpf:   "12345678900",
				Name:  "John Doe",
				Email: "john.doe@example.com",
			},
			expectedError: nil,
		},
		{
			name:     "Client Not Found",
			clientID: "non-existent-id",
			mockGateway: func(gateway *gateways.MockClientGateway) {
				gateway.On("GetClientById", ctx, "non-existent-id").Return(entities.Client{}, entities.ErrDataNotFound)
			},
			expectedClient: entities.Client{},
			expectedError:  fmt.Errorf("failed to get client by id - %s", entities.ErrDataNotFound.Error()),
		},
		{
			name:     "Gateway Error",
			clientID: clientID,
			mockGateway: func(gateway *gateways.MockClientGateway) {
				gateway.On("GetClientById", ctx, clientID).Return(entities.Client{}, errors.New("database error"))
			},
			expectedClient: entities.Client{},
			expectedError:  fmt.Errorf("failed to get client by id - database error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockGateway := new(gateways.MockClientGateway)
			if tc.mockGateway != nil {
				tc.mockGateway(mockGateway)
			}

			usecase := NewGetClientByIdUseCaseImpl(mockGateway)
			client, err := usecase.Execute(ctx, tc.clientID)

			require.Equal(t, tc.expectedError, err)
			require.Equal(t, tc.expectedClient, client)

			mockGateway.AssertExpectations(t)
		})
	}
}
