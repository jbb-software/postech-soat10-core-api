package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	router "post-tech-challenge-10soat/app/internal/delivery/http"
	"post-tech-challenge-10soat/app/internal/external/database/postgres"
	"post-tech-challenge-10soat/app/internal/infrastructure/config"
	dependency "post-tech-challenge-10soat/app/internal/infrastructure/di"
	"post-tech-challenge-10soat/app/internal/infrastructure/logger"

	_ "post-tech-challenge-10soat/app/docs"
)

//	@title			POS-Tech API
//	@version		1.0
//	@description	API em Go para o desafio na pos-tech fiap de Software Architecture.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@BasePath	/v1

//	@securityDefinitions.basic	BasicAuth

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	conf, err := config.New()
	if err != nil {
		slog.Error("Error loading environment variables", "error", err)
		os.Exit(1)
	}
	logger.Set(conf.App)
	slog.Info("Starting the application", "app", conf.App.Name, "env", conf.App.Env)

	ctx := context.Background()
	db, err := postgres.New(ctx, conf.DB)
	if err != nil {
		slog.Error("Error initializing database connection", "error", err)
		os.Exit(1)
	}

	slog.Info("Successfully connected to the database", "db", conf.DB.Connection)

	err = db.Migrate()
	if err != nil {
		slog.Error("Error migrating database", "error", err)
		os.Exit(1)
	}

	defer db.Close()

	// di
	healthHandler, clientHandler, productHandler, orderHandler, paymentHandler := dependency.Setup(conf.App, db)

	router, err := router.NewRouter(
		conf.HTTP,
		healthHandler,
		clientHandler,
		productHandler,
		orderHandler,
		paymentHandler,
	)
	if err != nil {
		slog.Error("Error initializing router", "error", err)
		os.Exit(1)
	}

	listenAddress := fmt.Sprintf("%s:%s", conf.HTTP.URL, conf.HTTP.Port)
	slog.Info("Starting the HTTP server", "listen_address", listenAddress)
	err = router.Run(listenAddress)
	if err != nil {
		slog.Error("Error starting the HTTP server", "error", err)
		os.Exit(1)
	}
}
