package config

import (
	"net/http"
	//_ "github.com/AaronSaikovski/go-api-starter/docs"
	//echoSwagger "github.com/swaggo/echo-swagger"
)

// Setup swagger
func AddSwaggerRoutes(app *http.ServeMux) {
	//app.HandleFunc("GET /swagger/*", echoSwagger.WrapHandler)
}

// func AddSwaggerRoutes(mux *http.ServeMux) {
// 	mux.HandleFunc("GET /swagger/*", echoSwagger.WrapHandler)
// }
