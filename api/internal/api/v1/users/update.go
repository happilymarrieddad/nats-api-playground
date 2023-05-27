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

type UpdateUser struct {
	FirstName *string `validate:"required_without_all=LastName" json:"first_name"`
	LastName  *string `validate:"required_without_all=FirstName" json:"last_name"`
	// Password        *string `validate:"required_without_all=FirstName LastName,required_with=PasswordConfirm" json:"password"`
	// PasswordConfirm *string `validate:"required_without_all=FirstName LastName,required_with=Password" json:"password_confirm"`
}

func Update(gr repos.GlobalRepo, nc nats.Client) error {
	_, err := nc.HandleAuthRequest("users.update", "api", func(m *natspkg.Msg) {
		req := UpdateUser{}

		if err := json.Unmarshal(m.Data, &req); err != nil {
			fmt.Printf("unable to marshal response err: %s\n", err.Error())
			nc.Respond(m.Reply, middleware.RespondErrMsg("unable to read data"))
			return
		}

		if err := types.Validate(req); err != nil {
			nc.Respond(m.Reply, middleware.RespondError(err))
			return
		}

		// TODO: add ability to change password
		// if req.Password != nil && req.PasswordConfirm != nil {
		// 	if len(*req.Password) == 0 || *req.Password != *req.PasswordConfirm {
		// 		nc.Respond(m.Reply, middleware.RespondErrMsg("password must match"))
		// 		return
		// 	}

		// }

		updatedUser, err := gr.Users().Update(types.UserUpdate{
			FirstName: req.FirstName,
			LastName:  req.LastName,
		})
		if err != nil {
			nc.Respond(m.Reply, middleware.RespondError(err))
			return
		}

		nc.Respond(m.Reply, middleware.Respond(updatedUser))
	})

	return err
}
