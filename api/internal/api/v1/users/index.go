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
	Offset int `validate:"required" json:"offset"`
}

func Index(gr repos.GlobalRepo, nc nats.Client) error {
	_, err := nc.HandleAuthRequest("users.index", "api", func(m *natspkg.Msg) {
		defer ginkgo.GinkgoRecover()
		req := IndexReq{}

		if err := json.Unmarshal(m.Data, &req); err != nil {
			fmt.Printf("unable to marshal response err: %s\n", err.Error())
			nc.Respond(m.Reply, middleware.RespondErrMsg("unable to read data ['limit','offset'] required in msg request"))
			return
		}

		if err := types.Validate(req); err != nil {
			nc.Respond(m.Reply, middleware.RespondError(err))
			return
		}

		usrs, count, err := gr.Users().Find(req.Limit, req.Offset)
		if err != nil {
			nc.Respond(m.Reply, middleware.RespondError(err))
			return
		}

		nc.Respond(m.Reply, middleware.RespondFind(usrs, count))
	})

	return err
}
