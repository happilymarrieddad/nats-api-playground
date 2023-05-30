package nats

import "fmt"

type channel string

func (c channel) String() string {
	return string(c)
}

func NewUsersUpdateChannel(id int64) channel {
	return channel(fmt.Sprintf("update.users.%d", id))
}
