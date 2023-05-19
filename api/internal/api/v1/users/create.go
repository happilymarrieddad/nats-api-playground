package users

import (
	"encoding/json"
	"fmt"

	"github.com/happilymarrieddad/nats-api-playground/api/internal/api/middleware"
	"github.com/happilymarrieddad/nats-api-playground/api/internal/nats"
	"github.com/happilymarrieddad/nats-api-playground/api/internal/repos"
	"github.com/happilymarrieddad/nats-api-playground/api/types"
	natspkg "github.com/nats-io/nats.go"
)

type NewUser struct {
	FirstName       string `validate:"required" json:"first_name"`
	LastName        string `validate:"required" json:"last_name"`
	Email           string `validate:"required" json:"email"`
	Password        string `validate:"required" json:"password"`
	PasswordConfirm string `validate:"required" json:"password_confirm"`
}

func Create(gr repos.GlobalRepo, nc nats.Client) {
	nc.HandleRequest("users.create", "api", func(m *natspkg.Msg) {
		req := NewUser{}

		if err := json.Unmarshal(m.Data, &req); err != nil {
			fmt.Printf("unable to marshal response err: %s\n", err.Error())
			nc.Respond(m.Reply, middleware.RespondErrMsg("unable to read data"))
			return
		}

		if err := types.Validate(req); err != nil {
			nc.Respond(m.Reply, middleware.RespondError(err))
			return
		}

		if len(req.Password) == 0 || req.Password != req.PasswordConfirm {
			nc.Respond(m.Reply, middleware.RespondErrMsg("password must match"))
			return
		}

		usr := &types.User{
			FirstName: req.FirstName,
			LastName:  req.LastName,
			Email:     req.Email,
		}
		usr.SetPassword(req.Password)
		if err := gr.Users().Create(usr); err != nil {
			nc.Respond(m.Reply, middleware.RespondError(err))
			return
		}

		nc.Respond(m.Reply, middleware.Respond(usr))
	})
}
