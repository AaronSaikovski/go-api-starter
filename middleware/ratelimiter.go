package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Set a Rate limiter with a sliding window  - https://echo.labstack.com/docs/middleware/rate-limiter
func RateLimiter(app *echo.Echo) {

	//The example below will limit the application to 50 requests/sec using the default in-memory store
	app.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(50)))

}
