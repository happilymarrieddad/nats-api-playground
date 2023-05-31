package middleware

import (
	"strconv"

	"github.com/nats-io/nats.go"
)

func GetIntKeyFromHeaders(m *nats.Msg, key string) (int64, error) {
	return strconv.ParseInt(m.Header.Get(key), 10, 64)
}
