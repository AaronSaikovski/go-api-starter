/*
Package go-api-starter
using the new  servemux features in Go 1.22
*/

package api

// import (
// 	"fmt"
// 	"os"

// 	"github.com/AaronSaikovski/go-api-starter/config"
// 	"github.com/AaronSaikovski/go-api-starter/logging"
// 	"github.com/AaronSaikovski/go-api-starter/router"

// 	"github.com/rs/zerolog/log"

// 	"net/http"
// )

// // MAIN - Setup and run
// func SetupAndRunApp() error {

// 	// load env
// 	errEnv := config.LoadENV()
// 	if errEnv != nil {
// 		return errEnv
// 	}

// 	//setup logging
// 	errLog := logging.SetupLogging()
// 	if errLog != nil {
// 		return errLog
// 	}

// 	log.Debug().Msg("calling SetupAndRunApp()")
// 	mux := http.NewServeMux()

// 	// setup routes
// 	router.SetupRoutes(mux)

// 	// get the server port
// 	port := os.Getenv("PORT")

// 	fmt.Println("Starting http server....")
// 	fmt.Printf("⇨ http server started on %s\n", ":"+port)

// 	// Start the server
// 	err := http.ListenAndServe(":"+port, mux)

// 	//err := app.Start(":" + port)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
