package client

import (
	"context"
	"fmt"
	"post-tech-challenge-10soat/app/internal/entities"
	"post-tech-challenge-10soat/app/internal/interfaces/gateways"
)

type GetClientByIdUseCaseImpl struct {
	gateway interfaces.ClientGateway
}

func NewGetClientByIdUseCaseImpl(gateway interfaces.ClientGateway) GetClientByIdUseCase {
	return &GetClientByIdUseCaseImpl{
		gateway,
	}
}

func (s GetClientByIdUseCaseImpl) Execute(ctx context.Context, id string) (entities.Client, error) {
	client, err := s.gateway.GetClientById(ctx, id)
	if err != nil {
		return entities.Client{}, fmt.Errorf("failed to get client by id - %s", err.Error())
	}
	return client, nil
}
