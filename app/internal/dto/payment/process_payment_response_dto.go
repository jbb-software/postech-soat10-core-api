package dto

type PaymentStatus string

const (
	Processed PaymentStatus = "processed"
	Denied    PaymentStatus = "denied"
)

type ProcessPaymentResponseDTO struct {
	Status  PaymentStatus
	Message string
}
