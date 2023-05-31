package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/happilymarrieddad/nats-api-playground/api/internal/api"
	"github.com/happilymarrieddad/nats-api-playground/api/internal/nats"
	"github.com/happilymarrieddad/nats-api-playground/api/internal/repos"
	"github.com/happilymarrieddad/nats-api-playground/api/types"
	_ "github.com/jackc/pgx/v4/stdlib" //postgres driver for sqlx
	natspkg "github.com/nats-io/nats.go"
	"xorm.io/xorm"
)

func main() {
	var (
		dbHost     = os.Getenv("NATS_PLAYGROUND_DB_HOST")
		dbPort     = os.Getenv("NATS_PLAYGROUND_DB_PORT")
		dbUser     = os.Getenv("NATS_PLAYGROUND_DB_USER")
		dbPass     = os.Getenv("NATS_PLAYGROUND_DB_PASS")
		dbDatabase = os.Getenv("NATS_PLAYGROUND_DB_DATABASE") + "_test" // force a test database for now
		natsUrl    = os.Getenv("NATS_PLAYGROUND_NATS_URL")
		natsUser   = os.Getenv("NATS_PLAYGROUND_NATS_USER")
		natsPass   = os.Getenv("NATS_PLAYGROUND_NATS_PASS")
	)

	conn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?connect_timeout=180&sslmode=disable",
		dbUser, dbPass, dbHost, dbPort, dbDatabase,
	)

	db, err := xorm.NewEngine("pgx", conn)
	if err != nil {
		panic(err)
	}

	// Just default to localhost
	if len(natsUrl) == 0 {
		natsUrl = natspkg.DefaultURL
	}
	if len(natsUser) == 0 {
		natsUser = "usr"
	}
	if len(natsPass) == 0 {
		natsPass = "pass"
	}

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?connect_timeout=180&sslmode=disable",
		dbUser, dbPass, dbHost, dbPort, dbDatabase,
	)

	db, err := xorm.NewEngine("pgx", connStr)
	if err != nil {
		panic(err)
	}

	// TODO: remove when migrations are up and running
	if err = db.Sync(
		&types.User{},
	); err != nil {
		panic(err)
	}

	gr, err := repos.NewGlobalRepo(db)
	if err != nil {
		panic(err)
	}

	nc, err := nats.NewClient(natsUrl, natsUser, natsPass)
	if err != nil {
		panic(err)
	}

	nc.SetDebug(true)

	api.Start(gin.Default(), 4000, gr, nc)
}
