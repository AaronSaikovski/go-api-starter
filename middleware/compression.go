package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// add compression: https://echo.labstack.com/docs/middleware/gzip
func AddCompression(app *echo.Echo) {

	app.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))

}
