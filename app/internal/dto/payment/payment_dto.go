package dto

import (
	entity "post-tech-challenge-10soat/app/internal/entities"
	"time"
)

type PaymentDTO struct {
	Id        string
	OrderId   string
	Provider  string
	Type      string
	Total     float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (d PaymentDTO) ToEntity() entity.Payment {
	return entity.Payment{
		Id:        d.Id,
		OrderId:   d.OrderId,
		Provider:  d.Provider,
		Type:      d.Type,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}
}
