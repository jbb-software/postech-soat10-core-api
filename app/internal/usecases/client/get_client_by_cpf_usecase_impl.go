package client

import (
	"context"
	"fmt"
	"post-tech-challenge-10soat/app/internal/entities"
	"post-tech-challenge-10soat/app/internal/interfaces/gateways"
)

type GetClientByCpfUseCaseImpl struct {
	gateway interfaces.ClientGateway
}

func NewGetClientByCpfUseCaseImpl(gateway interfaces.ClientGateway) GetClientByCpfUseCase {
	return &GetClientByCpfUseCaseImpl{
		gateway,
	}
}

func (s GetClientByCpfUseCaseImpl) Execute(ctx context.Context, cpf string) (entities.Client, error) {
	client, err := s.gateway.GetClientByCpf(ctx, cpf)
	if err != nil {
		return entities.Client{}, fmt.Errorf("failed to get client by cpf - %s", err.Error())
	}
	return client, nil
}
