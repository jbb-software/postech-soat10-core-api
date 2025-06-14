package mapper

import (
	"post-tech-challenge-10soat/app/internal/entities"
	"post-tech-challenge-10soat/app/internal/usecases/order"
	"post-tech-challenge-10soat/app/internal/utils"
	"time"

	"github.com/google/uuid"
)

type PaymentDataResponse struct {
	Id        string    `json:"id" example:"ed6ac028-8016-4cbd-aeee-c3a155cdb2a4"`
	OrderId   string    `json:"orderId" example:"ed6ac028-8016-4cbd-aeee-c3a155cdb2a4"`
	QrCode    string    `json:"qrCode" example:""`
	Total     float64   `json:"total" example:"100"`
	CreatedAt time.Time `json:"created_at" example:"1970-01-01T00:00:00Z"`
}

type OrderResponse struct {
	Id          uuid.UUID            `json:"id" example:"ed6ac028-8016-4cbd-aeee-c3a155cdb2a4"`
	Number      int                  `json:"number" example:"123"`
	ClientId    uuid.UUID            `json:"client_id" example:"ed6ac028-8016-4cbd-aeee-c3a155cdb2a4"`
	Total       float64              `json:"total" example:"100.90"`
	Status      entities.OrderStatus `json:"status" example:"received"`
	PaymentData PaymentDataResponse  `json:"payment" swaggertype:"object"`
	CreatedAt   time.Time            `json:"created_at" example:"1970-01-01T00:00:00Z"`
	UpdatedAt   time.Time            `json:"updated_at" example:"1970-01-01T00:00:00Z"`
}

func NewPaymentDataResponse(payment entities.PaymentData) PaymentDataResponse {
	return PaymentDataResponse{
		Id:        payment.Id,
		OrderId:   payment.OrderId,
		QrCode:    payment.QrCode,
		Total:     payment.Total,
		CreatedAt: payment.CreatedAt,
	}
}

func NewOrderResponse(order entities.Order) OrderResponse {
	orderResponse := OrderResponse{
		Id:          utils.StringToUuid(order.Id),
		Number:      order.Number,
		Total:       order.Total,
		Status:      order.Status,
		PaymentData: NewPaymentDataResponse(order.PaymentData),
		CreatedAt:   order.CreatedAt,
		UpdatedAt:   order.UpdatedAt,
	}
	if order.ClientId != "" {
		orderResponse.ClientId = utils.StringToUuid(order.ClientId)
	}
	return orderResponse
}

type ListOrdersResponse struct {
	Orders []OrderResponse `json:"orders"`
}

func NewListOrdersResponse(orders []entities.Order) []OrderResponse {
	var ordersResponse []OrderResponse
	for _, order := range orders {
		ordersResponse = append(ordersResponse, NewOrderResponse(order))
	}
	return ordersResponse
}

type OrderPaymentStatusResponse struct {
	PaymentStatus string `json:"paymentStatus" example:"payment_approved"`
}

func NewOrderPaymentStatusResponse(order order.OrderPaymentStatus) OrderPaymentStatusResponse {
	orderPaymentStatusResponse := OrderPaymentStatusResponse{
		PaymentStatus: string(order.PaymentStatus),
	}
	return orderPaymentStatusResponse
}

type UpdateOrderStatusResponse struct {
	Id        uuid.UUID            `json:"id" example:"ed6ac028-8016-4cbd-aeee-c3a155cdb2a4"`
	Status    entities.OrderStatus `json:"status" example:"received"`
	CreatedAt time.Time            `json:"created_at" example:"1970-01-01T00:00:00Z"`
	UpdatedAt time.Time            `json:"updated_at" example:"1970-01-01T00:00:00Z"`
}

func NewOrderUpdateStatusResponse(order entities.Order) UpdateOrderStatusResponse {
	orderResponse := UpdateOrderStatusResponse{
		Id:        utils.StringToUuid(order.Id),
		Status:    order.Status,
		CreatedAt: order.CreatedAt,
		UpdatedAt: order.UpdatedAt,
	}
	return orderResponse
}
