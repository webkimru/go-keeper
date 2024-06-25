package user

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/webkimru/go-keeper/internal/app/client/models"
	"github.com/webkimru/go-keeper/internal/app/client/service"
	"github.com/webkimru/go-keeper/pkg/errs"
	"github.com/webkimru/go-keeper/pkg/logger"
)

// NewUserNewCommand represents the initialized create user command
func NewUserNewCommand(userService *service.UserService, l *logger.Log) *cobra.Command {
	return &cobra.Command{
		Use:   "add",
		Short: "Creator new user",
		Long:  "Allows creating a new user with login and password",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("To register a user, enter login and password.")

			login, err := readString("Login: ")
			CLIog(l, "commands - NewUserNewCommand - login - readString(): %w", err)

			password, err := readString("Password: ")
			CLIog(l, "commands - NewUserNewCommand - password - readString(): %w", err)

			user := models.User{Login: login, Password: password}
			err = userService.Add(context.Background(), &user)
			if err != nil {
				if errors.Is(err, errs.ErrAlreadyExists) {
					errs.CLIMsgAlreadyExists()
					return
				}

				l.Log.Errorf("commands - NewUserNewCommand - userService.Add(): %w", err)
				errs.CLIMsgSeeLog()
				return
			}
		},
	}
}

// readString works as prompt UI
func readString(s string) (string, error) {
	var input string
	for {
		fmt.Print(s)
		_, err := fmt.Scan(&input)
		if err != nil {
			return "", fmt.Errorf("commands - readString - fmt.Scan(): %w", err)
		}
		if input != "" {
			break
		}
	}

	return input, nil
}
