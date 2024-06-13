package errs

import "errors"

const (
	MsgInternalServer = "internal server error"
	MsgAlreadyExists  = "already exists"
	MsgNotFound       = "not found"
	MsgInvalidCred    = "invalid credentials"
	MsgFieldRequired  = "field is required"
)

var (
	ErrNotFound       = errors.New("not found")
	ErrAlreadyExists  = errors.New("already exists")
	ErrInternalServer = errors.New("internal server error")
	ErrBadRequest     = errors.New("bad request")
)
