package router

// import (
// 	"github.com/AaronSaikovski/go-api-starter/handlers"
// 	"github.com/gofiber/fiber/v2"
// )

// func SetupRoutes(app *fiber.App) {
// 	app.Get("/api/ping", handlers.HandlePingCheck)
// 	app.Get("/api/weatherforecast", handlers.HandleWeatherGet)
// }

import (
	"github.com/AaronSaikovski/go-api-starter/handlers"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(app *echo.Echo) {
	app.GET("/api/ping", handlers.HandlePingCheck)
	app.GET("/api/weatherforecast", handlers.HandleWeatherGet)
}
