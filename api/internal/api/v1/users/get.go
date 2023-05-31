package users

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/happilymarrieddad/nats-api-playground/api/internal/api/middleware"
	"github.com/happilymarrieddad/nats-api-playground/api/internal/auth"
	"github.com/happilymarrieddad/nats-api-playground/api/internal/nats"
	"github.com/happilymarrieddad/nats-api-playground/api/internal/repos"
	natspkg "github.com/nats-io/nats.go"
	"github.com/onsi/ginkgo/v2"
)

func Get(gr repos.GlobalRepo, nc nats.Client) error {
	_, err := nc.HandleAuthRequest("users.get.*", "api", func(m *natspkg.Msg) ([]byte, string, error) {
		ctx := context.Background()
		defer ginkgo.GinkgoRecover()

		strs := strings.Split(m.Subject, ".")
		if len(strs) != 3 {
			return nil, "id required and it must be an integer", fmt.Errorf("id required and it must be an integer")
		}

		id, err := strconv.ParseInt(strs[2], 10, 64)
		if err != nil {
			return nil, "id required and it must be an integer", err
		}

		token := m.Header.Get("token")

		tokenUser, err := auth.GetUserFromToken(token)
		if err != nil {
			return nil, "unauthorized", err
		} else if tokenUser.ID != id {
			return nil, "unauthorized", fmt.Errorf("unauthorized id: %d neq %d", id, tokenUser.ID)
		}

		usr, exists, err := gr.Users().Get(ctx, id)
		if err != nil {
			return nil, "user not found", err
		} else if !exists {
			return nil, "user not found", fmt.Errorf("user not found")
		}

		return middleware.Respond(usr), "", nil
	})

	return err
}
