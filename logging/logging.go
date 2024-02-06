package logging

import (
	"flag"
	"os"
	"strconv"

	"github.com/rs/zerolog"
)

// Setup Zerolog logging from the .env file
func SetupLogging() error {
	//get debug flag
	debug_flag, err := strconv.ParseBool(os.Getenv("DEBUG"))
	if err != nil {
		return err
	}

	//setup logging
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	debug := flag.Bool("debug", debug_flag, "sets log level to debug")

	flag.Parse()

	// Default level for this example is info, unless debug flag is present
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	return nil
}
