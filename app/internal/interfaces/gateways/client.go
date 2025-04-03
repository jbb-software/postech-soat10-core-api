package interfaces

import (
	"context"
	"post-tech-challenge-10soat/app/internal/entities"
)

type ClientGateway interface {
	CreateClient(ctx context.Context, client entities.Client) (entities.Client, error)
	GetClientByCpf(ctx context.Context, cpf string) (entities.Client, error)
	GetClientById(ctx context.Context, id string) (entities.Client, error)
}
