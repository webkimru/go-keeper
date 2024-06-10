package errs

import "errors"

const (
	MsgInternalServer = "internal server error"
	MsgAlreadyExists  = "already exists"
	MsgNotFound       = "not found"
	MsgInvalidCred    = "invalid credentials"
)

var (
	ErrNotFound       = errors.New("not found")
	ErrAlreadyExists  = errors.New("already exists")
	ErrInternalServer = errors.New("internal server error")
)
