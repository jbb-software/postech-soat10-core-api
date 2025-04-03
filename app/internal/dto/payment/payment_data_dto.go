package dto

import (
	entity "post-tech-challenge-10soat/app/internal/entities"
	"time"
)

type PaymentDataDTO struct {
	Id        string
	OrderId   string
	QrCode    string
	Total     float64
	CreatedAt time.Time
}

func (d PaymentDataDTO) ToEntity() entity.PaymentData {
	return entity.PaymentData{
		Id:        d.Id,
		OrderId:   d.OrderId,
		QrCode:    d.QrCode,
		Total:     d.Total,
		CreatedAt: d.CreatedAt,
	}
}
