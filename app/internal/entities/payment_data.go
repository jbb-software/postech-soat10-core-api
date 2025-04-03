package entities

import "time"

type PaymentData struct {
	Id        string
	OrderId   string
	QrCode    string
	Total     float64
	CreatedAt time.Time
}
