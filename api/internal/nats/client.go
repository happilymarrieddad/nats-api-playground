package nats

import (
	"encoding/json"
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/happilymarrieddad/nats-api-playground/api/internal/api/middleware"
	"github.com/happilymarrieddad/nats-api-playground/api/internal/auth"
	"github.com/nats-io/nats.go"
)

//go:generate mockgen -destination=./mocks/Client.go -package=mock_nats github.com/happilymarrieddad/nats-api-playground/api/internal/nats Client
type Client interface {
	Publish(subj string, data []byte, headers map[string]string) error
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

func (c *client) Publish(subj string, data []byte, headers map[string]string) error {
	msg := nats.NewMsg(subj)
	msg.Data = data
	if headers != nil {
		for key, val := range headers {
			msg.Header.Add(key, val)
		}
	}

	return c.nc.PublishMsg(msg)
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

func (c *client) HandleAuthRequest(subj string, queueName string, fn func(m *nats.Msg) ([]byte, string, error)) (*nats.Subscription, error) {
	sub, err := c.nc.QueueSubscribeSync(subj, queueName)
	if err != nil {
		return nil, err
	}

	go func(s *nats.Subscription) {
		for {
			msgID := uuid.New()
			msg, subErr := s.NextMsg(time.Hour * 24 * 365 * 5)
			if subErr != nil {
				c.log("HandleAuthRequest unable to get NextMsg err: %s", subErr.Error())
				c.Respond(msg.Reply, []byte(`{"error": "unable to read message"}`))
				return
			}

			// do some auth
			token := msg.Header.Get("token")
			if tokenErr := auth.IsTokenValid(token); tokenErr != nil {
				c.log("HandleAuthRequest auth err: '%s' token: '%s'", tokenErr.Error(), token)
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

	resMsg, err := c.nc.RequestMsg(msg, time.Hour)
	if err != nil {
		return nil, err
	}

	type errStruct struct {
		Error string `json:"error"`
	}

	es := new(errStruct)
	if err = json.Unmarshal(resMsg.Data, es); err != nil {
		return nil, err
	}
	if len(es.Error) > 0 {
		return nil, errors.New(es.Error)
	}

	return resMsg.Data, nil
}

func (c *client) handleReq(id uuid.UUID, from string, msg *nats.Msg, fn func(m *nats.Msg) ([]byte, string, error)) {
	c.log("'%s': msgId: '%s' subject: '%s' received data: '%s'", from, id.String(), msg.Subject, string(msg.Data))
	reply := msg.Reply

	data, resMsg, err := fn(msg)
	if err != nil {
		c.log("'%s': msgId: '%s' err: '%s'", from, id.String(), err.Error())
		if err = c.Respond(reply, middleware.RespondErrMsg(resMsg)); err != nil {
			c.log("'%s': msgId: '%s' unable to respond err: '%s'", from, id.String(), err.Error())
		}
		return
	}

	// Log After
	c.log("'%s': msgId: '%s' sending data: '%s'", from, id.String(), string(data))
	if err = c.Respond(reply, data); err != nil {
		c.log("'%s': msgId: '%s' unable to respond err: '%s'", from, id.String(), err.Error())
	}
}

func (c *client) log(format string, v ...any) {
	if c.debug {
		log.Printf(format+"\n", v...)
	}
}
