package config

import (
	_ "github.com/AaronSaikovski/go-api-starter/docs"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// Setup swagger
func AddSwaggerRoutes(app *echo.Echo) {
	app.GET("/swagger/*", echoSwagger.WrapHandler)
}
