package order

//
//import (
//	"context"
//	"github.com/google/uuid"
//	"github.com/stretchr/testify/mock"
//	"github.com/stretchr/testify/require"
//	dto "post-tech-challenge-10soat/app/internal/dto/order"
//	"post-tech-challenge-10soat/app/internal/entities"
//	gateways "post-tech-challenge-10soat/app/internal/test_utils/mock"
//	"testing"
//)
//
//func TestCreateOrderUsecaseImpl_Execute(t *testing.T) {
//	ctx := context.Background()
//	clientID := uuid.New().String()
//	productID1 := uuid.New().String()
//	productID2 := uuid.New().String()
//	orderID := uuid.New().String()
//	paymentDataID := uuid.New().String()
//
//	testCases := []struct {
//		name                    string
//		createOrderDTO          dto.CreateOrderDTO
//		mockProductGateway      func(productGateway *gateways.MockProductGateway)
//		mockClientGateway       func(clientGateway *gateways.MockClientGateway)
//		mockOrderGateway        func(orderGateway *gateways.MockOrderGateway)
//		mockOrderProductGateway func(orderProductGateway *gateways.MockOrderProductGateway)
//		mockPaymentGateway      func(paymentGateway *gateways.MockPaymentGateway)
//		expectedOrder           entities.Order
//		expectedError           error
//	}{
//		{
//			name: "Success with client",
//			createOrderDTO: dto.CreateOrderDTO{
//				ClientId: clientID,
//				Products: []dto.CreateOrderProduct{
//					{ProductId: productID1, Quantity: 1},
//					{ProductId: productID2, Quantity: 2},
//				},
//			},
//			mockProductGateway: func(productGateway *gateways.MockProductGateway) {
//				productGateway.On("GetProductById", ctx, productID1).Return(entities.Product{Id: productID1, Value: 10.0}, nil)
//				productGateway.On("GetProductById", ctx, productID2).Return(entities.Product{Id: productID2, Value: 20.0}, nil)
//			},
//			mockClientGateway: func(clientGateway *gateways.MockClientGateway) {
//				clientGateway.On("GetClientById", ctx, clientID).Return(entities.Client{Id: clientID}, nil)
//			},
//			mockOrderGateway: func(orderGateway *gateways.MockOrderGateway) {
//				orderGateway.On("CreateOrder", ctx, mock.AnythingOfType("entities.Order")).Return(entities.Order{Id: orderID, ClientId: clientID, Total: 50.0}, nil)
//				orderGateway.On("DeleteOrder", ctx, orderID).Return(nil)
//			},
//			mockOrderProductGateway: func(orderProductGateway *gateways.MockOrderProductGateway) {
//				orderProductGateway.On("CreateOrderProduct", ctx, mock.AnythingOfType("entities.OrderProduct")).Return(entities.OrderProduct{}, nil).Times(2)
//			},
//			mockPaymentGateway: func(paymentGateway *gateways.MockPaymentGateway) {
//				paymentGateway.On("CreatePaymentData", ctx, mock.AnythingOfType("entities.PaymentData")).Return(entities.PaymentData{Id: paymentDataID}, nil)
//			},
//			expectedOrder: entities.Order{
//				Id:          orderID,
//				ClientId:    clientID,
//				Total:       50.0,
//				Status:      entities.OrderStatusPaymentPending,
//				PaymentData: entities.PaymentData{Id: paymentDataID},
//			},
//			expectedError: nil,
//		},
//		{
//			name: "Success without client",
//			createOrderDTO: dto.CreateOrderDTO{
//				Products: []dto.CreateOrderProduct{
//					{ProductId: productID1, Quantity: 1},
//				},
//			},
//			mockProductGateway: func(productGateway *gateways.MockProductGateway) {
//				productGateway.On("GetProductById", ctx, productID1).Return(entities.Product{Id: productID1, Value: 10.0}, nil).Times(2)
//			},
//			mockClientGateway: func(clientGateway *gateways.MockClientGateway) {},
//			mockOrderGateway: func(orderGateway *gateways.MockOrderGateway) {
//				orderGateway.On("CreateOrder", ctx, mock.AnythingOfType("entities.Order")).Return(entities.Order{Id: orderID, Total: 10.0}, nil)
//				orderGateway.On("DeleteOrder", ctx, orderID).Return(nil)
//			},
//			mockOrderProductGateway: func(orderProductGateway *gateways.MockOrderProductGateway) {
//				orderProductGateway.On("CreateOrderProduct", ctx, mock.AnythingOfType("entities.OrderProduct")).Return(entities.OrderProduct{}, nil)
//			},
//			mockPaymentGateway: func(paymentGateway *gateways.MockPaymentGateway) {
//				paymentGateway.On("CreatePaymentData", ctx, mock.AnythingOfType("entities.PaymentData")).Return(entities.PaymentData{Id: paymentDataID}, nil)
//			},
//			expectedOrder: entities.Order{
//				Id:          orderID,
//				Total:       10.0,
//				Status:      entities.OrderStatusPaymentPending,
//				PaymentData: entities.PaymentData{Id: paymentDataID},
//			},
//			expectedError: nil,
//		},
//		{
//			name: "Product Not Found",
//			createOrderDTO: dto.CreateOrderDTO{
//				Products: []dto.CreateOrderProduct{},
//			},
//			mockProductGateway: func(productGateway *gateways.MockProductGateway) {
//				productGateway.On("GetProductById", ctx, "invalid-product").Return(entities.Product{}, entities.ErrDataNotFound)
//			},
//			mockClientGateway:       func(clientGateway *gateways.MockClientGateway) {},
//			mockOrderGateway:        func(orderGateway *gateways.MockOrderGateway) {},
//			mockOrderProductGateway: func(orderProductGateway *gateways.MockOrderProductGateway) {},
//			mockPaymentGateway:      func(paymentGateway *gateways.MockPaymentGateway) {},
//			expectedOrder:           entities.Order{},
//			expectedError:           entities.ErrDataNotFound,
//		},
//		{
//			name: "Client Not Found",
//			createOrderDTO: dto.CreateOrderDTO{
//				ClientId: "invalid-client",
//				Products: []dto.CreateOrderProduct{},
//			},
//			mockProductGateway: func(productGateway *gateways.MockProductGateway) {
//				productGateway.On("GetProductById", ctx, productID1).Return(entities.Product{Id: productID1, Value: 10.0}, nil)
//			},
//			mockClientGateway: func(clientGateway *gateways.MockClientGateway) {
//				clientGateway.On("GetClientById", ctx, "invalid-client").Return(entities.Client{}, entities.ErrDataNotFound)
//			},
//			expectedOrder: entities.Order{},
//			expectedError: entities.ErrDataNotFound,
//		},
//	}
//
//	for _, tc := range testCases {
//		t.Run(tc.name, func(t *testing.T) {
//			mockProductGateway := new(gateways.MockProductGateway)
//			mockClientGateway := new(gateways.MockClientGateway)
//			mockOrderGateway := new(gateways.MockOrderGateway)
//			mockOrderProductGateway := new(gateways.MockOrderProductGateway)
//			mockPaymentGateway := new(gateways.MockPaymentGateway)
//
//			useCase := NewCreateOrderUsecaseImpl(mockProductGateway, mockClientGateway, mockOrderGateway, mockOrderProductGateway, mockPaymentGateway)
//			status, err := useCase.Execute(ctx, tc.createOrderDTO)
//
//			require.Equal(t, tc.expectedError, err)
//			require.Equal(t, tc.expectedOrder, status)
//
//			mockOrderGateway.AssertExpectations(t)
//		})
//	}
//
//}
