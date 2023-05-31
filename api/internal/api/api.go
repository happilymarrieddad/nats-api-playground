package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	authroutes "github.com/happilymarrieddad/nats-api-playground/api/internal/api/auth"
	"github.com/happilymarrieddad/nats-api-playground/api/internal/api/middleware"
	v1 "github.com/happilymarrieddad/nats-api-playground/api/internal/api/v1"
	"github.com/happilymarrieddad/nats-api-playground/api/internal/nats"
	"github.com/happilymarrieddad/nats-api-playground/api/internal/repos"
)

func Start(r *gin.Engine, port int, gr repos.GlobalRepo, nc nats.Client) {
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.SetGlobalRepoToContext(gr))

	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "okay"})
	})

	authroutes.Login(gr, nc)

	v1.SetupRoutes(gr, nc)

	fmt.Printf("Running server on port: %d\n", port)
	r.Run(fmt.Sprintf(":%d", port))
}
