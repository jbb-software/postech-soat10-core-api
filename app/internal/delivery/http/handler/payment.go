package handler

import (
	"github.com/gin-gonic/gin"
	"post-tech-challenge-10soat/app/internal/controllers"
	pm "post-tech-challenge-10soat/app/internal/delivery/http/mapper"
	"post-tech-challenge-10soat/app/internal/dto/payment"
)

type PaymentHandler struct {
	paymentController controllers.PaymentController
}

func NewPaymentHandler(paymentController controllers.PaymentController) PaymentHandler {
	return PaymentHandler{
		paymentController: paymentController,
	}
}

type processPaymentRequest struct {
	Provider      string `json:"provider" binding:"required" example:"mercado-pago"`
	TransactionId string `json:"transactionId" binding:"required" example:"ed6ac028-8016-4cbd-aeee-c3a155cdb2a4"`
	OrderId       string `json:"orderId" binding:"required" example:"ed6ac028-8016-4cbd-aeee-c3a155cdb2a4"`
	Status        string `json:"status" binding:"required" example:"approved"`
}

// ProccesPaymentResponse godoc
//
//	@Summary     Webhook para processar confirmação de um pagamento
//	@Description Webhook que deve receber a confirmação do pagamento se foi aprovado ou recusado
//	@Tags        Payments
//	@Accept      json
//	@Produce		json
//	@Param	    processPaymentRequest	body processPaymentRequest true "Processsar confirmação de pagamento"
//	@Success		200	{object} pm.ProcessedPaymentResponse	"Pagamento processado ou recusado"
//	@Failure		400	{object} ErrorResponse	"Erro de validação"
//	@Router		/payments/webhook/process [post]
func (h *PaymentHandler) ProccesPaymentResponse(ctx *gin.Context) {
	var request processPaymentRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		validationError(ctx, err)
		return
	}
	processPayment := dto.ProcessPaymentDTO{
		Provider:      request.Provider,
		OrderId:       request.OrderId,
		TransactionId: request.TransactionId,
		Status:        request.Status,
	}
	processedPayment, err := h.paymentController.ProcessPaymentResponse(ctx, processPayment)
	if err != nil {
		handleError(ctx, err)
		return
	}
	response := pm.NewProcessedPaymentResponse(processedPayment)
	handleSuccess(ctx, response)
}
