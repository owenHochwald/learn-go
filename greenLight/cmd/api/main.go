package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
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

	r := gin.Default()

	SetupRoutes(r, &app)

	r.Run(fmt.Sprintf(":%d", cfg.port))
}
