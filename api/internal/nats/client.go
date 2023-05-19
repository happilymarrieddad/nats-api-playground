package nats

import (
	"fmt"
	"time"

	"github.com/happilymarrieddad/nats-api-playground/api/internal/auth"
	"github.com/nats-io/nats.go"
)

type Client interface {
	HandleRequest(subject string, queueName string, fn func(m *nats.Msg)) (*nats.Subscription, error)
	HandleAuthRequest(subject string, queueName string, fn func(m *nats.Msg)) (*nats.Subscription, error)
	Respond(subj string, data []byte) error
	Request(subj string, data []byte) ([]byte, error)
}

func NewClient(natsUrl, usr, pass string) (Client, error) {
	// Connect to a server
	nc, err := nats.Connect(natsUrl, nats.UserInfo(usr, pass))
	if err != nil {
		return nil, err
	}

	return &client{nc}, nil
}

type client struct {
	nc *nats.Conn
}

// queueName should be the same with all handlers
func (c *client) HandleRequest(subj string, queueName string, fn func(m *nats.Msg)) (*nats.Subscription, error) {
	sub, err := c.nc.QueueSubscribeSync(subj, queueName)
	if err != nil {
		return nil, err
	}

	go func(s *nats.Subscription) {
		for {
			msg, err := s.NextMsg(time.Hour * 24 * 365 * 100)
			if err != nil {
				panic(err)
			}

			fn(msg)
		}
	}(sub)

	return sub, nil
}

func (c *client) HandleAuthRequest(subj string, queueName string, fn func(m *nats.Msg)) (*nats.Subscription, error) {
	sub, err := c.nc.QueueSubscribeSync(subj, queueName)
	if err != nil {
		return nil, err
	}

	go func(s *nats.Subscription) {
		for {
			msg, err := s.NextMsg(time.Hour * 24 * 365 * 100)
			if err != nil {
				panic(err)
			}

			// do some auth
			token := msg.Header.Get("token")

			if _, err := auth.IsTokenValid(token); err != nil {
				fmt.Printf("auth failure err: %s\n", err.Error())
				c.Respond(msg.Reply, []byte(`{"error": "unauthorized"}`))
				continue
			}

			fn(msg)
		}
	}(sub)

	return sub, nil
}

func (c *client) Respond(subj string, data []byte) error {
	return c.nc.Publish(subj, data)
}

func (c *client) Request(subj string, data []byte) ([]byte, error) {
	msg, err := c.nc.Request(subj, data, time.Hour*24)
	if err != nil {
		return nil, err
	}

	return msg.Data, nil
}
