package client

import (
	"context"
	"post-tech-challenge-10soat/app/internal/entities"
)

type GetClientByCpfUseCase interface {
	Execute(ctx context.Context, cpf string) (entities.Client, error)
}
