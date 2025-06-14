package repository

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	dto2 "post-tech-challenge-10soat/app/internal/dto/client"
	"post-tech-challenge-10soat/app/internal/external/database/postgres"
	"post-tech-challenge-10soat/app/internal/external/database/postgres/model"
)

type ClientRepositoryImpl struct {
	db *postgres.DB
}

func NewClientRepositoryImpl(db *postgres.DB) ClientRepositoryImpl {
	return ClientRepositoryImpl{
		db,
	}
}

func (repository ClientRepositoryImpl) CreateClient(ctx context.Context, client dto2.CreateClientDTO) (dto2.ClientDTO, error) {
	var clientModel model.ClientModel
	query := repository.db.QueryBuilder.Insert("clients").
		Columns("name", "email").
		Values(client.Name, client.Email).
		Suffix("RETURNING id, name, email, created_at, updated_at")
	sql, args, err := query.ToSql()
	if err != nil {
		return dto2.ClientDTO{}, err
	}
	err = repository.db.QueryRow(ctx, sql, args...).Scan(
		&clientModel.Id,
		&clientModel.Name,
		&clientModel.Email,
		&clientModel.CreatedAt,
		&clientModel.UpdatedAt,
	)
	if err != nil {
		return dto2.ClientDTO{}, err
	}
	return clientModel.ToDTO(), nil
}

func (repository ClientRepositoryImpl) GetClientByCpf(ctx context.Context, cpf string) (dto2.ClientDTO, error) {
	var clientModel model.ClientModel
	query := repository.db.QueryBuilder.Select("*").
		From("clients").
		Where(sq.Eq{"cpf": cpf}).
		Limit(1)
	sql, args, err := query.ToSql()
	if err != nil {
		return dto2.ClientDTO{}, err
	}
	err = repository.db.QueryRow(ctx, sql, args...).Scan(
		&clientModel.Id,
		&clientModel.Cpf,
		&clientModel.Name,
		&clientModel.Email,
		&clientModel.CreatedAt,
		&clientModel.UpdatedAt,
	)
	if err != nil {
		return dto2.ClientDTO{}, err
	}
	return clientModel.ToDTO(), nil
}

func (repository ClientRepositoryImpl) GetClientById(ctx context.Context, id string) (dto2.ClientDTO, error) {
	var clientModel model.ClientModel
	query := repository.db.QueryBuilder.Select("*").
		From("clients").
		Where(sq.Eq{"id": id}).
		Limit(1)
	sql, args, err := query.ToSql()
	if err != nil {
		return dto2.ClientDTO{}, err
	}
	err = repository.db.QueryRow(ctx, sql, args...).Scan(
		&clientModel.Id,
		&clientModel.Cpf,
		&clientModel.Name,
		&clientModel.Email,
		&clientModel.CreatedAt,
		&clientModel.UpdatedAt,
	)
	if err != nil {
		return dto2.ClientDTO{}, err
	}
	return clientModel.ToDTO(), nil
}
