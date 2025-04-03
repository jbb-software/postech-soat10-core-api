package gateways

import (
	"context"
	"post-tech-challenge-10soat/app/internal/dto/client"
	"post-tech-challenge-10soat/app/internal/entities"
	"post-tech-challenge-10soat/app/internal/interfaces/repositories"
)

type ClientGatewayImpl struct {
	repository interfaces.ClientRepository
}

func NewClientGatewayImpl(repository interfaces.ClientRepository) *ClientGatewayImpl {
	return &ClientGatewayImpl{
		repository,
	}
}

func (cg ClientGatewayImpl) CreateClient(ctx context.Context, client entities.Client) (entities.Client, error) {
	createClientDTO := dto.CreateClientDTO{
		Cpf:   client.Cpf,
		Name:  client.Name,
		Email: client.Email,
	}
	createdClient, err := cg.repository.CreateClient(ctx, createClientDTO)
	if err != nil {
		return entities.Client{}, err
	}
	return createdClient.ToEntity(), nil
}

func (cg ClientGatewayImpl) GetClientByCpf(ctx context.Context, cpf string) (entities.Client, error) {
	client, err := cg.repository.GetClientByCpf(ctx, cpf)
	if err != nil {
		return entities.Client{}, err
	}
	return client.ToEntity(), nil
}

func (cg ClientGatewayImpl) GetClientById(ctx context.Context, id string) (entities.Client, error) {
	client, err := cg.repository.GetClientById(ctx, id)
	if err != nil {
		return entities.Client{}, err
	}
	return client.ToEntity(), nil
}
