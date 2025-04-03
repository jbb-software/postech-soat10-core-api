package interfaces

import (
	"context"
	dto2 "post-tech-challenge-10soat/app/internal/dto/client"
)

type ClientRepository interface {
	CreateClient(ctx context.Context, client dto2.CreateClientDTO) (dto2.ClientDTO, error)
	GetClientByCpf(ctx context.Context, cpf string) (dto2.ClientDTO, error)
	GetClientById(ctx context.Context, id string) (dto2.ClientDTO, error)
}
