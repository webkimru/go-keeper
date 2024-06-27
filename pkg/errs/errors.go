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
	ErrAlreadyExists      = errors.New("already exists")
	ErrBadRequest         = errors.New("bad request")
	ErrInternalServer     = errors.New("internal server error")
	ErrNotFound           = errors.New("not found")
	ErrPermissionDenied   = errors.New("permission denied")
	ErrInvalidCredentials = errors.New("invalid credentials")
)

func MsgInternalServerError(err error) string {
	return fmt.Sprintf("%s: %v", MsgInternalServer, err)
}

func MsgFieldRequiredError(err error) string {
	return fmt.Sprintf("%v", err)
}

func printColorRed(s string) {
	fmt.Printf("%serror: %s%s\n", "\033[31m", s, "\033[0m")
}

func MsgCLI(err error) string {
	return fmt.Sprintf("%serror: %s%s\n", "\033[31m", err, "\033[0m")
}

func CLIMsgAlreadyExists() {
	printColorRed(MsgAlreadyExists)
}

func CLIMsgBadRequest() {
	printColorRed(MsgFieldRequired)
}

func CLIMsgInvalidCredentials() {
	printColorRed(MsgInvalidCred)
}

func CLIMsgPermissionDenied() {
	printColorRed(MsgPermissionDenied)
}

func CLIMsgSeeLog() {
	printColorRed("see go-keeper.log")
}
