package users

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/happilymarrieddad/nats-api-playground/api/internal/api/middleware"
	"github.com/happilymarrieddad/nats-api-playground/api/internal/nats"
	"github.com/happilymarrieddad/nats-api-playground/api/internal/repos"
	"github.com/happilymarrieddad/nats-api-playground/api/types"
	natspkg "github.com/nats-io/nats.go"
	"github.com/onsi/ginkgo/v2"
)

type NewUser struct {
	FirstName       string `validate:"required" json:"first_name"`
	LastName        string `validate:"required" json:"last_name"`
	Email           string `validate:"required" json:"email"`
	Password        string `validate:"required" json:"password"`
	PasswordConfirm string `validate:"required" json:"password_confirm"`
}

func Create(gr repos.GlobalRepo, nc nats.Client) error {
	_, err := nc.HandleRequest("users.create", "api", func(m *natspkg.Msg) ([]byte, string, error) {
		ctx := context.Background()
		defer ginkgo.GinkgoRecover()
		req := NewUser{}

		if err := json.Unmarshal(m.Data, &req); err != nil {
			return nil, "unable to read data", err
		}

		if err := types.Validate(req); err != nil {
			return nil, "unable to read data", err
		}

		if len(req.Password) == 0 || req.Password != req.PasswordConfirm {
			return nil, "password must match", fmt.Errorf("password must match")
		}

		usr := &types.User{
			FirstName: req.FirstName,
			LastName:  req.LastName,
			Email:     req.Email,
		}
		usr.SetPassword(req.Password)
		if err := gr.Users().Create(ctx, usr); err != nil {
			return nil, "unable to create user", err
		}

		return middleware.Respond(usr), "", nil
	})

	return err
}
