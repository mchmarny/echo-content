package main

import (
	"log"
	"net"
	"os"

	"github.com/gin-gonic/gin"
	ev "github.com/mchmarny/gcputil/env"
)

const (
	defaultPort         = "8080"
	portVariableName    = "PORT"
	messageVariableName = "MSG"
)

var (
	logger = log.New(os.Stdout, "", 0)
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.NoRoute(func(c *gin.Context) {
		c.String(200, "%s", ev.MustGetEnvVar(messageVariableName,
			"MSG env var not set"))
	})
	port := ev.MustGetEnvVar(portVariableName, defaultPort)
	hostPost := net.JoinHostPort("0.0.0.0", port)
	logger.Printf("Server starting: %s \n", hostPost)
	if err := r.Run(hostPost); err != nil {
		logger.Fatal(err)
	}
}
