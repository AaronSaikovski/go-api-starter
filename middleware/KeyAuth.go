package middleware

// import (
// 	"github.com/gofiber/fiber/v2"
// 	"github.com/gofiber/fiber/v2/middleware/keyauth"
// )

// // Attach API key auth to the app header - 'XApiKey'
// func AddApiKeyAuth(app *fiber.App) {
// 	app.Use(keyauth.New(keyauth.Config{
// 		KeyLookup: "header:XApiKey",
// 		Validator: ValidateAPIKey,
// 	}))
// }

import (
	"github.com/labstack/echo/v4"
)

// Attach API key auth to the app header - 'XApiKey'
func AddApiKeyAuth(app *echo.Echo) {

	// app.Use(keyauth.New(keyauth.Config{
	// 	KeyLookup: "header:XApiKey",
	// 	Validator: ValidateAPIKey,
	// }))
}
