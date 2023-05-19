package repos_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	. "github.com/happilymarrieddad/nats-api-playground/api/internal/repos"
	"github.com/happilymarrieddad/nats-api-playground/api/types"

	// TODO: upgrade to pgx5
	_ "github.com/jackc/pgx/v4/stdlib"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"xorm.io/xorm"
)

var (
	db  *xorm.Engine
	gr  GlobalRepo
	ctx context.Context
)

var _ = BeforeSuite(func() {
	defer GinkgoRecover()
	ctx = context.Background()
	var err error
	var (
		dbHost     = os.Getenv("NATS_PLAYGROUND_DB_HOST")
		dbPort     = os.Getenv("NATS_PLAYGROUND_DB_PORT")
		dbUser     = os.Getenv("NATS_PLAYGROUND_DB_USER")
		dbPass     = os.Getenv("NATS_PLAYGROUND_DB_PASS")
		dbDatabase = "nats_api_test" // force a test database
	)

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?connect_timeout=180&sslmode=disable",
		dbUser, dbPass, dbHost, dbPort, dbDatabase,
	)

	db, err = xorm.NewEngine("pgx", connStr)
	Expect(err).To(BeNil())

	db.Sync(&types.User{})

	gr, err = NewGlobalRepo(db)
	Expect(err).To(BeNil())
})

func clearDatabase(tables ...string) {
	for _, table := range tables {
		_, err := db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", table))
		Expect(err).To(Succeed())
	}
}

func TestRepos(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Repos Suite")
}
