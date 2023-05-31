package users

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/happilymarrieddad/nats-api-playground/api/internal/api/middleware"
	"github.com/happilymarrieddad/nats-api-playground/api/internal/auth"
	"github.com/happilymarrieddad/nats-api-playground/api/internal/nats"
	"github.com/happilymarrieddad/nats-api-playground/api/internal/repos"
	"github.com/happilymarrieddad/nats-api-playground/api/types"
	natspkg "github.com/nats-io/nats.go"
	"github.com/onsi/ginkgo/v2"
)

type UpdateUser struct {
	FirstName       *string `validate:"required_without_all=LastName" json:"first_name"`
	LastName        *string `validate:"required_without_all=FirstName" json:"last_name"`
	Password        *string `json:"password"`
	PasswordConfirm *string `json:"password_confirm"`
}

func Update(gr repos.GlobalRepo, nc nats.Client) error {
	_, err := nc.HandleAuthRequest("users.update.*", "api", func(m *natspkg.Msg) ([]byte, string, error) {
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

		req := UpdateUser{}

		if err := json.Unmarshal(m.Data, &req); err != nil {
			return nil, fmt.Sprintf("unable to marshal response err: %s\n", err.Error()), err
		}

		if err := types.Validate(req); err != nil {
			return nil, fmt.Sprintf("unable to read data. first_name and/or last_name are required"), err
		}

		token := m.Header.Get("token")
		tokenUser, err := auth.GetUserFromToken(token)
		if err != nil {
			return nil, "unauthorized", err
		} else if tokenUser.ID != id {
			return nil, "unauthorized", fmt.Errorf("unauthorized id: %d neq %d", id, tokenUser.ID)
		}

		userToUpdate := types.UserUpdate{
			ID:        tokenUser.ID,
			FirstName: req.FirstName,
			LastName:  req.LastName,
		}

		if req.Password != nil && req.PasswordConfirm != nil {
			if len(*req.Password) == 0 || *req.Password != *req.PasswordConfirm {
				return nil, "passwords must match", fmt.Errorf("passwords must match")
			}

			// TODO: update the user with password
		}

		updatedUser, err := gr.Users().Update(ctx, userToUpdate)
		if err != nil {
			return nil, "unable to update user", err
		}

		resBts := middleware.Respond(updatedUser)
		ch := nats.NewUsersUpdateChannel(updatedUser.ID).String()
		if err = nc.Publish(ch, resBts, nil); err != nil {
			log.Printf("unable to publish to channel '%s' err: '%s'\n", ch, err.Error())
		}

		return resBts, "", nil
	})

	return err
}
