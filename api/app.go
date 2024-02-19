/*
Package go-api-starter
using the echo web framework
*/

package api

import (
	"os"

	"github.com/AaronSaikovski/go-api-starter/config"
	"github.com/AaronSaikovski/go-api-starter/logging"
	"github.com/AaronSaikovski/go-api-starter/middleware"
	"github.com/AaronSaikovski/go-api-starter/router"

	"github.com/labstack/echo/v4"
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
	errLog := logging.SetupLogging()
	if errLog != nil {
		return errLog
	}

	log.Debug().Msg("calling SetupAndRunApp()")

	// create Echo app -
	app := echo.New()

	//API versioning
	//v1 := app.Group("/v1")
	//router.SetupV1Routes(v1)

	// Uses API key header - 'XApiKey'
	//middleware.AddApiKeyAuth(app)

	// attach swagger
	config.AddSwaggerRoutes(app)

	// attach middleware
	middleware.Recover(app)
	middleware.Logger(app)

	//Use CORS - change AllowOrigins to suit
	middleware.AddCors(app)

	// setup routes
	router.SetupRoutes(app)

	//Add a rate limiter
	//middleware.RateLimiter(app)

	//Add compression
	//middleware.AddCompression(app)

	// get the server port
	port := os.Getenv("PORT")

	// Start the server
	err := app.Start(":" + port)
	if err != nil {
		return err
	}

	return nil
}
