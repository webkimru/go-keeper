package errs

import (
	"errors"
	"fmt"
)

const (
	MsgAlreadyExists  = "already exists"
	MsgInternalServer = "internal server error"
	MsgInvalidCred    = "invalid credentials"
	MsgFieldRequired  = "field is required"
	MsgNotFound       = "not found"
)

var (
	ErrAlreadyExists = errors.New("already exists")
	ErrBadRequest    = errors.New("bad request")
	ErrNotFound      = errors.New("not found")
)

func MsgInternalServerError(err error) string {
	return fmt.Sprintf("%s: %v", MsgInternalServer, err)
}
