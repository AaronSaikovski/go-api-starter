package router

import (
	"github.com/AaronSaikovski/go-api-starter/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/api/ping", handlers.HandlePingCheck)
	app.Get("/api/weatherforecast", handlers.HandleWeatherGet)
}
