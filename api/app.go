package api

import (
	"fmt"
	"os"

	"github.com/AaronSaikovski/go-api-starter/config"
	"github.com/AaronSaikovski/go-api-starter/logging"
	"github.com/AaronSaikovski/go-api-starter/router"

	"github.com/rs/zerolog/log"

	"net/http"
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
	mux := http.NewServeMux()

	// create Echo app -
	//app := echo.New()

	// Uses API key header - 'XApiKey'
	//middleware.AddApiKeyAuth(app)

	// attach swagger
	//config.AddSwaggerRoutes(mux)

	// attach middleware
	//middleware.Recover(mux)
	//middleware.Logger(mux)

	//Use CORS - change AllowOrigins to suit
	//middleware.AddCors(mux)

	// setup routes
	router.SetupRoutes(mux)

	//Add a rate limiter
	//middleware.RateLimiter(app)

	//Add compression
	//middleware.AddCompression(app)

	// get the server port
	port := os.Getenv("PORT")

	fmt.Println("Starting http server....")

	// Start the server
	err := http.ListenAndServe(":"+port, mux)

	//err := app.Start(":" + port)
	if err != nil {
		return err
	}

	return nil
}
