package nats

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/happilymarrieddad/nats-api-playground/api/internal/api/middleware"
	"github.com/happilymarrieddad/nats-api-playground/api/internal/auth"
	"github.com/nats-io/nats.go"
)

//go:generate mockgen -destination=./mocks/Client.go -package=mock_nats github.com/happilymarrieddad/nats-api-playground/api/internal/nats Client
type Client interface {
	HandleRequest(subject string, queueName string, fn func(m *nats.Msg) (data []byte, fnRespMsg string, err error)) (*nats.Subscription, error)
	HandleAuthRequest(subject string, queueName string, fn func(m *nats.Msg) (data []byte, fnRespMsg string, err error)) (*nats.Subscription, error)
	Respond(subj string, data []byte) error
	Request(subj string, data []byte, headers map[string]string) ([]byte, error)
	SetDebug(val bool)
}

func NewClient(natsUrl, usr, pass string) (Client, error) {
	// Connect to a server
	nc, err := nats.Connect(natsUrl, nats.UserInfo(usr, pass))
	if err != nil {
		return nil, err
	}

	return &client{nc, false}, nil
}

type client struct {
	nc    *nats.Conn
	debug bool
}

func (c *client) SetDebug(val bool) {
	c.debug = val
}

// queueName should be the same with all handlers
func (c *client) HandleRequest(subj string, queueName string, fn func(m *nats.Msg) (data []byte, fnRespMsg string, err error)) (*nats.Subscription, error) {
	sub, err := c.nc.QueueSubscribeSync(subj, queueName)
	if err != nil {
		return nil, err
	}

	go func(s *nats.Subscription) {
		for {
			msgID := uuid.New()
			msg, err := s.NextMsg(time.Hour * 24 * 365 * 100)
			if err != nil {
				c.log("HandleRequest unable to get NextMsg err: %s", err.Error())
				c.Respond(msg.Reply, []byte(`{"error": "unable to read message"}`))
				return
			}

			c.handleReq(msgID, "HandleAuthRequest", msg, fn)
		}
	}(sub)

	return sub, nil
}

func (c *client) HandleAuthRequest(subj string, queueName string, fn func(m *nats.Msg) (data []byte, fnRespMsg string, err error)) (*nats.Subscription, error) {
	sub, err := c.nc.QueueSubscribeSync(subj, queueName)
	if err != nil {
		return nil, err
	}

	go func(s *nats.Subscription) {
		for {
			msgID := uuid.New()
			msg, err := s.NextMsg(time.Hour * 24 * 365 * 100)
			if err != nil {
				c.log("HandleAuthRequest unable to get NextMsg err: %s", err.Error())
				c.Respond(msg.Reply, []byte(`{"error": "unable to read message"}`))
				return
			}

			// do some auth
			token := msg.Header.Get("token")

			if _, err := auth.IsTokenValid(token); err != nil {
				c.log("HandleAuthRequest auth err: %s", err.Error())
				c.Respond(msg.Reply, []byte(`{"error": "unauthorized"}`))
				continue
			}

			c.handleReq(msgID, "HandleAuthRequest", msg, fn)
		}
	}(sub)

	return sub, nil
}

func (c *client) Respond(subj string, data []byte) error {
	return c.nc.Publish(subj, data)
}

func (c *client) Request(subj string, data []byte, headers map[string]string) ([]byte, error) {
	msg := nats.NewMsg(subj)
	msg.Data = data
	if headers != nil {
		for key, val := range headers {
			msg.Header.Add(key, val)
		}
	}

	msg, err := c.nc.RequestMsg(msg, time.Hour*24)
	if err != nil {
		return nil, err
	}

	type errStruct struct {
		Error string `json:"error"`
	}

	es := new(errStruct)
	if err = json.Unmarshal(msg.Data, es); err != nil {
		return nil, err
	}
	if len(es.Error) > 0 {
		return nil, errors.New(es.Error)
	}

	return msg.Data, nil
}

func (c *client) handleReq(id uuid.UUID, from string, msg *nats.Msg, fn func(m *nats.Msg) (data []byte, fnRespMsg string, err error)) {
	c.log("'%s': msg '%s' received data: '%s'", from, id.String(), string(msg.Data))

	data, resMsg, err := fn(msg)
	if err != nil {
		c.log("'%s': msg '%s' err: %s", err.Error())
		if err = c.Respond(msg.Reply, middleware.RespondErrMsg(resMsg)); err != nil {
			c.log("'%s': msg '%s' unable to respond err: %s", err.Error())
		}
		return
	}

	// Log After
	c.log("'%s': msg '%s' sending data: '%s'", from, id.String(), string(data))
	if err = c.Respond(msg.Reply, data); err != nil {
		c.log("'%s': msg '%s' unable to respond err: %s", err.Error())
	}
}

func (c *client) log(format string, v ...any) {
	if c.debug {
		fmt.Printf(format+"\n", v...)
	}
}
