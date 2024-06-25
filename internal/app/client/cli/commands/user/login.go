package user

import (
	"context"
	"os"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/webkimru/go-keeper/internal/app/client/service"
	"github.com/webkimru/go-keeper/pkg/errs"
	"github.com/webkimru/go-keeper/pkg/logger"
)

// NewUserLoginCommand represents the initialization for the login user command
func NewUserLoginCommand(userService *service.UserService, l *logger.Log) *cobra.Command {
	return &cobra.Command{
		Use:   "login",
		Short: "Login the user",
		Long:  "Allows logging the user by login and password",
		Run: func(cmd *cobra.Command, args []string) {
			login, err := readString("Login: ")
			CLIog(l, "commands - NewUserLoginCommand - login - readString(): %w", err)

			password, err := readString("Password: ")
			CLIog(l, "commands - NewUserLoginCommand - password - readString(): %w", err)

			token, err := userService.Auth(context.Background(), login, password)
			if err != nil {
				if errors.Is(err, errs.ErrInvalidCredentials) {
					errs.CLIMsgInvalidCredentials()
					return
				}

				l.Log.Errorf("commands - NewUserLoginCommand - userService.Auth(): %w", err)
				errs.CLIMsgSeeLog()
				return
			}

			// write token to the config for some commands required authentication by token
			file, err := os.Create("config.json")
			CLIog(l, "commands - NewUserLoginCommand  - os.Create(): %w", err)
			defer file.Close()
			config := `{
	"app": {
		"token": "{{$token}}"
	}
}`
			config = strings.Replace(config, "{{$token}}", token, 1)
			_, err = file.Write([]byte(config))
			CLIog(l, "commands - NewUserLoginCommand - os.Write(): %w", err)
		},
	}
}

func CLIog(l *logger.Log, s string, err error) {
	if err != nil {
		l.Log.Errorf(s, err)
		errs.CLIMsgSeeLog()
		os.Exit(1)
	}
}
