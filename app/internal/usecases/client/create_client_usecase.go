package client

import (
	"context"
	"post-tech-challenge-10soat/app/internal/dto/client"
	"post-tech-challenge-10soat/app/internal/entities"
)

type CreateClientUseCase interface {
	Execute(ctx context.Context, client dto.CreateClientDTO) (entities.Client, error)
}
