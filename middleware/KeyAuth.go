package middleware

import (
	"crypto/sha256"
	"crypto/subtle"
	"errors"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Attach API key auth to the app header - 'XApiKey' - https://echo.labstack.com/docs/middleware/key-auth
func AddApiKeyAuth(app *echo.Echo) {

	app.Use(middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		KeyLookup: "header:XApiKey",
		Validator: func(key string, c echo.Context) (bool, error) {
				  
				//get API key
				apiKey := os.Getenv("APIKEY")
				hashedAPIKey := sha256.Sum256([]byte(apiKey))
				hashedKey := sha256.Sum256([]byte(key))
				if subtle.ConstantTimeCompare(hashedAPIKey[:], hashedKey[:]) == 1 {
					return true, nil
				}
				return false, errors.New("apikey is missing or invalid.")
		},
	  }))
}
