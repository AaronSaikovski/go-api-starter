package app

import (
	"os"

	"github.com/AaronSaikovski/go-api-starter/config"
	"github.com/AaronSaikovski/go-api-starter/middleware"
	"github.com/AaronSaikovski/go-api-starter/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/rs/zerolog/log"
)

// MAIN - Setup and run
func SetupAndRunApp() error {

	// load env
	errEnv := config.LoadENV()
	if errEnv != nil {
		return errEnv
	}

	//setup logging
	errLog := setupLogging()
	if errLog != nil {
		return errLog
	}

	log.Debug().Msg("calling SetupAndRunApp()")

	// create Fiber app
	app := fiber.New()

	// Uses API key header - 'XApiKey'
	middleware.AddApiKeyAuth(app)

	// attach middleware
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path} ${latency}\n",
	}))

	//Use CORS - change AllowOrigins to suit
	middleware.AddCors(app)

	// setup routes
	router.SetupRoutes(app)

	// attach swagger
	config.AddSwaggerRoutes(app)

	// Provide a minimal healthcheck - Default Endpoint: /livez
	app.Use(healthcheck.New())

	// get the port and start
	port := os.Getenv("PORT")
	app.Listen(":" + port)

	return nil
}
