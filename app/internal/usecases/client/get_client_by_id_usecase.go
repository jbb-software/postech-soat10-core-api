package client

import (
	"context"
	"post-tech-challenge-10soat/app/internal/entities"
)

type GetClientByIdUseCase interface {
	Execute(ctx context.Context, id string) (entities.Client, error)
}
