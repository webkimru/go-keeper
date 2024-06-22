// Package errs consists of the base errors and messages including wrapped errors.
package errs

import (
	"errors"
	"fmt"
)

const (
	MsgAlreadyExists    = "already exists"
	MsgInternalServer   = "internal server error"
	MsgInvalidCred      = "invalid credentials"
	MsgFieldRequired    = "field is required"
	MsgNotFound         = "not found"
	MsgPermissionDenied = "permission denied"
)

var (
	ErrAlreadyExists    = errors.New("already exists")
	ErrBadRequest       = errors.New("bad request")
	ErrInternalServer   = errors.New("internal server error")
	ErrNotFound         = errors.New("not found")
	ErrPermissionDenied = errors.New("permission denied")
)

func MsgInternalServerError(err error) string {
	return fmt.Sprintf("%s: %v", MsgInternalServer, err)
}

func MsgFieldRequiredError(err error) string {
	return fmt.Sprintf("%v", err)
}
