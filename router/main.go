package router

import (
	"github.com/AaronSaikovski/go-api-starter/handlers"
	"github.com/labstack/echo/v4"
)

// Setup up API routes
func SetupRoutes(app *echo.Echo) {
	app.GET("/api/ping", handlers.HandlePingCheck)
	app.GET("/api/weatherforecast", handlers.HandleWeatherGet)
}
