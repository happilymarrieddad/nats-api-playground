package users

import (
	"encoding/json"
	"fmt"

	"github.com/happilymarrieddad/nats-api-playground/api/internal/api/middleware"
	"github.com/happilymarrieddad/nats-api-playground/api/internal/nats"
	"github.com/happilymarrieddad/nats-api-playground/api/internal/repos"
	natspkg "github.com/nats-io/nats.go"
)

type IndexReq struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

func Index(gr repos.GlobalRepo, nc nats.Client) {
	nc.HandleAuthRequest("users.index", "api", func(m *natspkg.Msg) {
		req := IndexReq{}

		if err := json.Unmarshal(m.Data, &req); err != nil {
			fmt.Printf("unable to marshal response err: %s\n", err.Error())
			nc.Respond(m.Reply, middleware.RespondErrMsg("unable to read data"))
			return
		}

		usrs, count, err := gr.Users().Find(req.Limit, req.Offset)
		if err != nil {
			nc.Respond(m.Reply, middleware.RespondError(err))
			return
		}

		nc.Respond(m.Reply, middleware.RespondFind(usrs, count))
	})
}
