package v1

import (
	"github.com/happilymarrieddad/nats-api-playground/api/internal/api/v1/users"
	"github.com/happilymarrieddad/nats-api-playground/api/internal/nats"
	"github.com/happilymarrieddad/nats-api-playground/api/internal/repos"
)

func SetupRoutes(gr repos.GlobalRepo, nc nats.Client) {
	// Users
	users.Index(gr, nc)
	users.Create(gr, nc)
	users.Update(gr, nc)
}
