package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// api version number
const version = "1.0.0"

type config struct {
	port int
	env  string
}

type application struct {
	config config
	logger *log.Logger
}

func main() {
	// instance of config struct
	var cfg config

	// define the port number and env
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	// needs to be called after seting setting flags
	flag.Parse()

	// intialize a new logger which writes messages to the standard stream with prefixed with current date and time
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// instance of application struct containing config struct and logger
	app := &application{
		config: cfg,
		logger: logger,
	}

	// declare multiplexer
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthcheck", app.healthcheckHandler)

	srv := &http.Server{
		Addr:         fmt.Sprintf(": %d", cfg.port),
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	logger.Printf("starting %s server on %d", cfg.env, cfg.port)
	err := srv.ListenAndServe()
	logger.Fatal(err)
}
