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

type IndexReq struct {
	Limit  int `validate:"required" json:"limit"`
	Offset int `json:"offset"`
}

func Index(gr repos.GlobalRepo, nc nats.Client) error {
	_, err := nc.HandleAuthRequest("users.index", "api", func(m *natspkg.Msg) ([]byte, string, error) {
		fmt.Println("Inside users.index")
		defer ginkgo.GinkgoRecover()
		req := IndexReq{}

		if err := json.Unmarshal(m.Data, &req); err != nil {
			fmt.Println("users.index Unmarshal err: ", err.Error())
			return nil, "unable to read data ['limit','offset'] required in msg request", err
		}

		if err := types.Validate(req); err != nil {
			fmt.Println("users.index Validate err: ", err.Error())
			return nil, "unable to read data ['limit','offset'] required in msg request", err
		}

		usrs, count, err := gr.Users().Find(req.Limit, req.Offset)
		if err != nil {
			fmt.Println("users.index Users().Find err: ", err.Error())
			return nil, "unable to get users", err
		}

		fmt.Println("Finished users.index")
		return middleware.RespondFind(usrs, count), "", nil
	})

	return err
}
