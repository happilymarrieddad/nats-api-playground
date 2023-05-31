package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/happilymarrieddad/nats-api-playground/api/internal/api"
	"github.com/happilymarrieddad/nats-api-playground/api/internal/repos"
	"github.com/happilymarrieddad/nats-api-playground/api/types"
	_ "github.com/jackc/pgx/v4/stdlib" //postgres driver for sqlx
	"xorm.io/xorm"
)

func main() {
	var (
		dbHost     = os.Getenv("DBHOST")
		dbPort     = os.Getenv("DBPORT")
		dbUser     = os.Getenv("DBUSER")
		dbPass     = os.Getenv("DBPASS")
		dbDatabase = os.Getenv("DBDATABASE")
	)

	conn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?connect_timeout=180&sslmode=disable",
		dbUser, dbPass, dbHost, dbPort, dbDatabase,
	)

	db, err := xorm.NewEngine("pgx", conn)
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

	api.Start(gin.Default(), 3000, gr)
}
