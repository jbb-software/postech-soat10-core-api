package mercadopago

import (
	"context"
	dto2 "post-tech-challenge-10soat/app/internal/dto/payment"
	"post-tech-challenge-10soat/app/internal/interfaces/clients"
	"time"
)

type MercadoPagoClientImpl struct {
	//Poderiamos passar configs de ambiente aqui
}

func NewMercadoPagoClientImpl() clients.PaymentClient {
	return MercadoPagoClientImpl{}
}

func (m MercadoPagoClientImpl) CreatePaymentData(ctx context.Context, paymentData dto2.CreatePaymentDataDTO) (dto2.PaymentDataDTO, error) {
	// TODO: Mock, trocaria pela implementacao real
	return dto2.PaymentDataDTO{
		Id:        "123",
		OrderId:   paymentData.OrderId,
		QrCode:    "0020101021243650016COM.MERCADOLIBRE02013063638f1192a-5fd1-4180-a180-8bcae3556bc35204000053039865802BR5925IZABEL AAAA DE MELO6007BARUERI62070503***63040B6D",
		Total:     paymentData.Total,
		CreatedAt: time.Now(),
	}, nil
}
