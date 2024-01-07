package router

import (
	"github.com/AaronSaikovski/go-api-starter/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/ping", handlers.HandleHealthCheck)
	app.Get("/weatherforecast", handlers.HandleWeatherGet)
}
