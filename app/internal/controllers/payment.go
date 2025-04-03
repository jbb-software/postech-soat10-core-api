package controllers

import (
	"context"
	dto2 "post-tech-challenge-10soat/app/internal/dto/payment"
	"post-tech-challenge-10soat/app/internal/usecases/payment"
)

type PaymentController struct {
	processPayment payment.ProcessPaymentResponseUseCase
}

func NewPaymentController(
	processPayment payment.ProcessPaymentResponseUseCase,
) *PaymentController {
	return &PaymentController{
		processPayment,
	}
}

func (c *PaymentController) ProcessPaymentResponse(ctx context.Context, processPaymentDto dto2.ProcessPaymentDTO) (dto2.ProcessPaymentResponseDTO, error) {
	processPayment, err := c.processPayment.Execute(ctx, processPaymentDto)
	if err != nil {
		return dto2.ProcessPaymentResponseDTO{}, err
	}
	return processPayment, nil
}
