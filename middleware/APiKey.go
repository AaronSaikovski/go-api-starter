package middleware

// import (
// 	"crypto/sha256"
// 	"crypto/subtle"
// 	"errors"
// 	"os"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/gofiber/fiber/v2/middleware/keyauth"
// )

// // Validate with API key
// func ValidateAPIKey(c *fiber.Ctx, key string) (bool, error) {

// 	//get API key
// 	apiKey := os.Getenv("APIKEY")

// 	//return if API is empty
// 	if apiKey == "" {
// 		return false, errors.New("APIKEY is empty")
// 	}

// 	hashedAPIKey := sha256.Sum256([]byte(apiKey))
// 	hashedKey := sha256.Sum256([]byte(key))

// 	if subtle.ConstantTimeCompare(hashedAPIKey[:], hashedKey[:]) == 1 {
// 		return true, nil
// 	}
// 	return false, keyauth.ErrMissingOrMalformedAPIKey
// }

// ref: https://echo.labstack.com/docs/middleware/key-auth
import (
	"crypto/sha256"
	"crypto/subtle"
	"errors"
	"os"

	"github.com/labstack/echo/v4"
)

// Validate with API key
func ValidateAPIKey(c echo.Context, key string) (bool, error) {

	//get API key
	apiKey := os.Getenv("APIKEY")

	//return if API is empty
	if apiKey == "" {
		return false, errors.New("APIKEY is empty")
	}

	hashedAPIKey := sha256.Sum256([]byte(apiKey))
	hashedKey := sha256.Sum256([]byte(key))

	if subtle.ConstantTimeCompare(hashedAPIKey[:], hashedKey[:]) == 1 {
		return true, nil
	}
	return false, errors.New("APIKEY is missing or invalid.")
}
