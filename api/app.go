package api

import (
	"os"

	"github.com/AaronSaikovski/go-api-starter/config"
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
	errLog := setupLogging()
	if errLog != nil {
		return errLog
	}

	log.Debug().Msg("calling SetupAndRunApp()")

	// create Echo app
	app := echo.New()

	// Uses API key header - 'XApiKey'
	//middleware.AddApiKeyAuth(app)

	// attach middleware
	middleware.Recover(app)
	middleware.Logger(app)

	//Use CORS - change AllowOrigins to suit
	middleware.AddCors(app)

	// setup routes
	router.SetupRoutes(app)

	// attach swagger
	config.AddSwaggerRoutes(app)

	//Add a rate limiter
	//middleware.RateLimiter(app)

	//Add compression
	//middleware.AddCompression(app)

	// get the port and start
	port := os.Getenv("PORT")
	app.Logger.Fatal(app.Start(":" + port))

	return nil
}
