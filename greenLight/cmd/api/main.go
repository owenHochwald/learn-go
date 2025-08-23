package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type config struct {
	port int
	env  string
}

type application struct {
	config config
	logger *log.Logger
}

const version = "1.0.0"

func main() {
	var cfg config
	flag.IntVar(&cfg.port, "port", 8080, "port to run the server on")
	flag.StringVar(&cfg.env, "env", "dev", "environment to run the server in")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	app := application{
		config: cfg,
		logger: logger,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthcheck", app.healthCheckHandler)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.config.port),
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	logger.Printf("Starting server on port %d", app.config.port)
	log.Fatal(server.ListenAndServe())

}
