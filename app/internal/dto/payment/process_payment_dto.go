package dto

type ProcessPaymentDTO struct {
	Provider      string
	TransactionId string
	OrderId       string
	Status        string
}
