package app

import (
	"flag"
	"os"
	"strconv"

	"github.com/AaronSaikovski/go-api-starter/config"
	"github.com/AaronSaikovski/go-api-starter/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Setup logging from .env file
func setupLogging() error {
	//get debug flag
	debug_flag, err := strconv.ParseBool(os.Getenv("DEBUG"))
	if err != nil {
		return err
	}

	//setup logging
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	debug := flag.Bool("debug", debug_flag, "sets log level to debug")

	flag.Parse()

	// Default level for this example is info, unless debug flag is present
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	return nil
}

// Setup and run
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

	// create app
	app := fiber.New()

	// attach middleware
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path} ${latency}\n",
	}))

	//Use CORS - change AllowOrigins to suit
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// setup routes
	router.SetupRoutes(app)

	// attach swagger
	config.AddSwaggerRoutes(app)

	// get the port and start
	port := os.Getenv("PORT")
	app.Listen(":" + port)

	return nil
}
