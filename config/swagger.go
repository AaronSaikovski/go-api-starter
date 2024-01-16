// package config

// import (
// 	"github.com/gofiber/fiber/v2"
// 	"github.com/gofiber/swagger"

// 	_ "github.com/AaronSaikovski/go-api-starter/docs"
// )

// func AddSwaggerRoutes(app *fiber.App) {
// 	// setup swagger
// 	app.Get("/swagger/*", swagger.HandlerDefault)
// }

package config

import (
	_ "github.com/AaronSaikovski/go-api-starter/docs"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func AddSwaggerRoutes(app *echo.Echo) {
	// setup swagger
	app.GET("/swagger/*", echoSwagger.WrapHandler)
}
