package di

import (
	controllers2 "post-tech-challenge-10soat/app/internal/controllers"
	handler2 "post-tech-challenge-10soat/app/internal/delivery/http/handler"
	"post-tech-challenge-10soat/app/internal/external/clients/mercadopago"
	"post-tech-challenge-10soat/app/internal/external/database/postgres"
	repository2 "post-tech-challenge-10soat/app/internal/external/database/postgres/repositories"
	gateways2 "post-tech-challenge-10soat/app/internal/gateways"
	"post-tech-challenge-10soat/app/internal/infrastructure/config"
	"post-tech-challenge-10soat/app/internal/infrastructure/logger"
	client2 "post-tech-challenge-10soat/app/internal/usecases/client"
	order2 "post-tech-challenge-10soat/app/internal/usecases/order"
	"post-tech-challenge-10soat/app/internal/usecases/payment"
	product2 "post-tech-challenge-10soat/app/internal/usecases/product"
)

func Setup(config *config.App, db *postgres.DB) (
	handler2.HealthHandler,
	handler2.ClientHandler,
	handler2.ProductHandler,
	handler2.OrderHandler,
	handler2.PaymentHandler,
) {
	logger.Set(config)

	// Repositories
	clientRepo := repository2.NewClientRepositoryImpl(db)
	productRepo := repository2.NewProductRepositoryImpl(db)
	categoryRepo := repository2.NewCategoryRepositoryImpl(db)
	orderRepo := repository2.NewOrderRepositoryImpl(db)
	orderProductRepo := repository2.NewOrderProductRepositoryImpl(db)
	paymentRepo := repository2.NewPaymentRepositoryImpl(db)

	// Clients
	mercadoPagoClient := mercadopago.NewMercadoPagoClientImpl()

	// Gateways
	clientGateway := gateways2.NewClientGatewayImpl(
		clientRepo,
	)
	productGateway := gateways2.NewProductGatewayImpl(
		productRepo,
	)
	categoryGateway := gateways2.NewCategoryGatewayImpl(
		categoryRepo,
	)
	orderGateway := gateways2.NewOrderGatewayImpl(
		orderRepo,
	)
	orderProductGateway := gateways2.NewOrderProductGatewayImpl(
		orderProductRepo,
	)
	paymentGateway := gateways2.NewPaymentGatewayImpl(
		paymentRepo,
		mercadoPagoClient,
	)

	// Usecases
	getClientByCpf := client2.NewGetClientByCpfUseCaseImpl(
		clientGateway,
	)
	getClientById := client2.NewGetClientByCpfUseCaseImpl(
		clientGateway,
	)
	createClient := client2.NewCreateClientUsecaseImpl(
		clientGateway,
	)
	createProduct := product2.NewCreateProductUsecaseImpl(
		productGateway,
		categoryGateway,
	)
	updateProduct := product2.NewUpdateProductUsecaseImpl(
		productGateway,
		categoryGateway,
	)
	deleteProduct := product2.NewDeleteProductUsecaseImpl(
		productGateway,
	)
	listProducts := product2.NewListProductsUsecaseImpl(
		productGateway,
		categoryGateway,
	)
	// paymentUseCase := payment.NewPaymentCheckoutUsecaseImpl(
	// 	paymentGateway,
	// )
	createOrder := order2.NewCreateOrderUsecaseImpl(
		productGateway,
		clientGateway,
		orderGateway,
		orderProductGateway,
		paymentGateway,
	)
	listOrders := order2.NewListOrdersUseCaseImpl(
		orderGateway,
	)
	getOrderPaymentStatus := order2.NewGetOrderPaymentStatusUseCaseImpl(
		orderGateway,
	)
	updateOrderStatus := order2.NewUpdateOrderStatusUseCaseImpl(
		orderGateway,
	)
	processPayment := payment.NewProcessPaymentResponseUseCaseImpl(
		paymentGateway,
		orderGateway,
	)

	// Controllers
	clientController := controllers2.NewClientController(
		getClientByCpf,
		getClientById,
		createClient,
	)
	productController := controllers2.NewProductController(
		createProduct,
		deleteProduct,
		updateProduct,
		listProducts,
	)
	orderController := controllers2.NewOrderController(
		createOrder,
		listOrders,
		getOrderPaymentStatus,
		updateOrderStatus,
	)
	paymentController := controllers2.NewPaymentController(
		processPayment,
	)

	// Handlers
	healthHandler := handler2.NewHealthHandler()
	clientHandler := handler2.NewClientHandler(*clientController)
	productHandler := handler2.NewProductHandler(*productController)
	orderHandler := handler2.NewOrderHandler(*orderController)
	paymentHandler := handler2.NewPaymentHandler(*paymentController)

	return healthHandler, clientHandler, productHandler, orderHandler, paymentHandler
}
