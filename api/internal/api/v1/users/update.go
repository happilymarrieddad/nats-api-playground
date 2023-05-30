package users

import (
	"encoding/json"
	"fmt"

	"github.com/happilymarrieddad/nats-api-playground/api/internal/api/middleware"
	"github.com/happilymarrieddad/nats-api-playground/api/internal/nats"
	"github.com/happilymarrieddad/nats-api-playground/api/internal/repos"
	"github.com/happilymarrieddad/nats-api-playground/api/types"
	natspkg "github.com/nats-io/nats.go"
	"github.com/onsi/ginkgo/v2"
)

type UpdateUser struct {
	FirstName *string `validate:"required_without_all=LastName" json:"first_name"`
	LastName  *string `validate:"required_without_all=FirstName" json:"last_name"`
	// Password        *string `validate:"required_without_all=FirstName LastName,required_with=PasswordConfirm" json:"password"`
	// PasswordConfirm *string `validate:"required_without_all=FirstName LastName,required_with=Password" json:"password_confirm"`
}

func Update(gr repos.GlobalRepo, nc nats.Client) error {
	_, err := nc.HandleAuthRequest("users.update", "api", func(m *natspkg.Msg) ([]byte, string, error) {
		defer ginkgo.GinkgoRecover()
		req := UpdateUser{}

		if err := json.Unmarshal(m.Data, &req); err != nil {
			return nil, fmt.Sprintf("unable to marshal response err: %s\n", err.Error()), err
		}

		if err := types.Validate(req); err != nil {
			return nil, fmt.Sprintf("unable to read data. first_name and/or last_name are required"), err
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
			return nil, "unable to update user", err
		}

		return middleware.Respond(updatedUser), "", nil
	})

	return err
}
