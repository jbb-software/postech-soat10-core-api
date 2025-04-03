package entities

import (
	"time"
)

const (
	PaymentTypePixQRCode = "PIX-QRCODE"
)

const (
	PaymentProviderMp = "mercado-pago"
)

type Payment struct {
	Id        string
	OrderId   string
	Provider  string
	Type      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
