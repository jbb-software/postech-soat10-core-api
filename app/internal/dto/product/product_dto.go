package dto

import (
	"post-tech-challenge-10soat/app/internal/dto/category"
	entity "post-tech-challenge-10soat/app/internal/entities"
	"time"
)

type ProductDTO struct {
	Id          string
	Name        string
	Description string
	Image       string
	Value       float64
	CategoryId  string
	CategoryDTO dto.CategoryDTO
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (d ProductDTO) ToEntity() entity.Product {
	return entity.Product{
		Id:          d.Id,
		Name:        d.Name,
		Description: d.Description,
		Image:       d.Image,
		Value:       d.Value,
		CategoryId:  d.CategoryId,
		Category:    d.CategoryDTO.ToEntity(),
		CreatedAt:   d.CreatedAt,
		UpdatedAt:   d.UpdatedAt,
	}
}
