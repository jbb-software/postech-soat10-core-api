package mapper

import (
	"post-tech-challenge-10soat/app/internal/dto/payment"
)

type ProcessedPaymentResponse struct {
	Status  string `json:"status" example:"processed"`
	Message string `json:"message" example:"success"`
}

func NewProcessedPaymentResponse(processedPayment dto.ProcessPaymentResponseDTO) ProcessedPaymentResponse {
	return ProcessedPaymentResponse{
		Status:  string(processedPayment.Status),
		Message: processedPayment.Message,
	}
}
