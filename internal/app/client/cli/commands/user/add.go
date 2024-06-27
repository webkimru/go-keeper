package user

import (
	"context"
	"fmt"
	"io"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/webkimru/go-keeper/internal/app/client/models"
	"github.com/webkimru/go-keeper/internal/app/client/service"
	"github.com/webkimru/go-keeper/pkg/errs"
	"github.com/webkimru/go-keeper/pkg/logger"
)

// NewUserAddCommand represents the initialized create user command
func NewUserAddCommand(in io.Reader, userService *service.UserService, l *logger.Log) *cobra.Command {
	return &cobra.Command{
		Use:   "add",
		Short: "Creator new user",
		Long:  "Allows creating a new user with login and password",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("To register a user, enter login and password.")

			login, err := readString(in, "Login: ")
			CLIog(l, "commands - NewUserAddCommand - login - readString(): %w", err)

			password, err := readString(in, "Password: ")
			CLIog(l, "commands - NewUserAddCommand - password - readString(): %w", err)

			user := models.User{Login: login, Password: password}
			err = userService.Add(context.Background(), &user)
			if err != nil {
				if errors.Is(err, errs.ErrAlreadyExists) {
					errs.CLIMsgAlreadyExists()
					return
				}

				l.Log.Errorf("commands - NewUserAddCommand - userService.Add(): %w", err)
				errs.CLIMsgSeeLog()
				return
			}
		},
	}
}

// readString works as prompt UI
func readString(in io.Reader, s string) (string, error) {
	var input string
	for {
		fmt.Print(s)
		_, err := fmt.Fscanln(in, &input)
		if err != nil {
			return "", fmt.Errorf("commands - readString - fmt.Fscanln(): %w", err)
		}
		if input != "" {
			break
		}
	}

	return input, nil
}
