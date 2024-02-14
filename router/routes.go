package router

import (
	"net/http"
	"github.com/AaronSaikovski/go-api-starter/handlers"
)

// Setup up API routes
func SetupRoutes(app *http.ServeMux) {
	app.HandleFunc("GET /api/health", handlers.HandleHealthCheck)
	app.HandleFunc("GET /api/weatherforecast", handlers.HandleWeatherGet)
}
