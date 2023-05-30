package authroutes

import (
	"encoding/json"

	"github.com/happilymarrieddad/nats-api-playground/api/internal/api/middleware"
	"github.com/happilymarrieddad/nats-api-playground/api/internal/auth"
	"github.com/happilymarrieddad/nats-api-playground/api/internal/nats"
	"github.com/happilymarrieddad/nats-api-playground/api/internal/repos"
	natspkg "github.com/nats-io/nats.go"
)

func Login(gr repos.GlobalRepo, nc nats.Client) {
	nc.HandleRequest("login", "api", func(m *natspkg.Msg) ([]byte, string, error) {
		type Login struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		var login Login
		if err := json.Unmarshal(m.Data, &login); err != nil {
			return nil, "unable to read data", err
		}

		usr, exists, err := gr.Users().GetByEmail(login.Email)
		if err != nil {
			return nil, "unable to get user", err
		} else if !exists || !usr.PasswordMatches(login.Password) {
			return nil, "unauthorized", err
		}

		// For now just return a token
		token, err := auth.CreateToken(map[string]interface{}{
			"user": usr,
		})
		if err != nil {
			return nil, "unauthorized", err
		}

		return middleware.Respond(struct {
			Token string `json:"token"`
		}{
			Token: token,
		}), "", nil
	})
}
