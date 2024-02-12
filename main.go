package main

import (
	"github.com/AaronSaikovski/go-api-starter/api"
	_ "github.com/AaronSaikovski/go-api-starter/docs"
)

// @title go-api-starter
// @version 0.3.1
// @description An example template of a Golang simple backend API using Echo
// @contact.name Aaron Saikovski
// @license.name MIT
// @host localhost:8080
// @BasePath /
func main() {

	// setup and run app
	err := api.SetupAndRunApp()
	if err != nil {
		panic(err)
	}
}
