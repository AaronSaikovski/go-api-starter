package main

import (
	"github.com/AaronSaikovski/go-api-starter/app"
	_ "github.com/AaronSaikovski/go-api-starter/docs"
)

// @title go-api-starter
// @version 0.1
// @description An example template of a Golang simple backend API using Fiber
// @contact.name Aaron Saikovski
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
