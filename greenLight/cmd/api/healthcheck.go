package main

import (
	"fmt"
	"net/http"
)

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "status available\n")
	fmt.Fprintf(w, "port %d\n", app.config.port)
	fmt.Fprintf(w, "version %s\n", version)
	w.WriteHeader(http.StatusOK)
}
