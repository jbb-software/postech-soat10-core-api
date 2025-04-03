package client

import (
	"context"
	"fmt"
	"post-tech-challenge-10soat/app/internal/dto/client"
	"post-tech-challenge-10soat/app/internal/entities"
	"post-tech-challenge-10soat/app/internal/interfaces/gateways"
)

type CreateClientUseCaseImpl struct {
	gateway interfaces.ClientGateway
}

func NewCreateClientUsecaseImpl(gateway interfaces.ClientGateway) CreateClientUseCase {
	return &CreateClientUseCaseImpl{
		gateway,
	}
}

func (s CreateClientUseCaseImpl) Execute(ctx context.Context, createClientDTO dto.CreateClientDTO) (entities.Client, error) {
	newClient := entities.Client{
		Cpf:   createClientDTO.Cpf,
		Name:  createClientDTO.Name,
		Email: createClientDTO.Email,
	}
	client, err := s.gateway.CreateClient(ctx, newClient)
	if err != nil {
		return entities.Client{}, fmt.Errorf("failed to create client - %s", err.Error())
	}
	return client, nil
}
