package controllers

import (
	"context"
	"post-tech-challenge-10soat/app/internal/dto/client"
	entity "post-tech-challenge-10soat/app/internal/entities"
	client2 "post-tech-challenge-10soat/app/internal/usecases/client"
)

type ClientController struct {
	getClientByCpf client2.GetClientByCpfUseCase
	getClientById  client2.GetClientByIdUseCase
	createClient   client2.CreateClientUseCase
}

func NewClientController(
	getClientByCpf client2.GetClientByCpfUseCase,
	getClientById client2.GetClientByIdUseCase,
	createClient client2.CreateClientUseCase,
) *ClientController {
	return &ClientController{
		getClientByCpf,
		getClientById,
		createClient,
	}
}

func (c *ClientController) GetClientByCpf(ctx context.Context, cpf string) (entity.Client, error) {
	client, err := c.getClientByCpf.Execute(ctx, cpf)
	if err != nil {
		return entity.Client{}, err
	}
	return client, nil
}

func (c *ClientController) GetClientById(ctx context.Context, id string) (entity.Client, error) {
	client, err := c.getClientById.Execute(ctx, id)
	if err != nil {
		return entity.Client{}, err
	}
	return client, nil
}

func (c *ClientController) CreateClient(ctx context.Context, createClient dto.CreateClientDTO) (entity.Client, error) {
	client, err := c.createClient.Execute(ctx, createClient)
	if err != nil {
		return entity.Client{}, err
	}
	return client, nil
}
