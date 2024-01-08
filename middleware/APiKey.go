package middleware

import (
	"crypto/sha256"
	"crypto/subtle"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/keyauth"
	"os"
)

// Validate with API key
func ValidateAPIKey(c *fiber.Ctx, key string) (bool, error) {

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
	return false, keyauth.ErrMissingOrMalformedAPIKey
}
