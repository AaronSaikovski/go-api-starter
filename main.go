package main

import (
	"github.com/AaronSaikovski/go-api-starter/app"
	_ "github.com/AaronSaikovski/go-api-starter/docs"
)

// @title The Better Backend Template
// @version 0.1
// @description An example template of a Golang backend API using Fiber and MongoDB
// @contact.name Ben Davis
// @license.name MIT
// @host localhost:8080
// @BasePath /
func main() {
	// setup and run app
	err := app.SetupAndRunApp()
	if err != nil {
		panic(err)
	}
}
