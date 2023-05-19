package types

import (
	"fmt"
	"strings"
)

const (
	notFoundErr string = "Not Found:"
)

func IsNotFoundError(err error) bool {
	return strings.Contains(err.Error(), notFoundErr)
}

func NewNotFoundError(msg string) error {
	return fmt.Errorf("%s %s", notFoundErr, msg)
}
