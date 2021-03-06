package main

import (
	"executor/server"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"path"
	"time"
)

func main() {
	server.RunServer()
}

func init() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if os.Getenv("DEBUG") == "true" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	log.Logger = zerolog.New(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.RFC3339,
		FormatCaller: func(i interface{}) string {
			_, f := path.Split(i.(string))
			return "[" + fmt.Sprintf("%-20s", f) + "]"
		},
	}).With().Caller().Timestamp().Stack().Logger()
}
