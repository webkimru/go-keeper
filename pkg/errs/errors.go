package errs

import (
	"errors"
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
