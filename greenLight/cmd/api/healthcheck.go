package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (app *application) healthCheckHandler(c *gin.Context) {
	fmt.Fprint(c.Writer, "version %s", version)
	fmt.Fprint(c.Writer, "port %d", app.config.port)
	c.Writer.WriteHeader(200)
}
