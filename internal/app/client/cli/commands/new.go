package commands

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/webkimru/go-keeper/internal/app/client/models"
	"github.com/webkimru/go-keeper/internal/app/client/service"
	"github.com/webkimru/go-keeper/pkg/logger"
)

// NewUserNewCommand represents the initialized create user command
func NewUserNewCommand(userService *service.UserService, l *logger.Log) *cobra.Command {
	return &cobra.Command{
		Use:   "new",
		Short: "Creator new user",
		Long:  "Allows creating a new user with login and password",
		Run: func(cmd *cobra.Command, args []string) {
			login, err := readString("Login: ")
			if err != nil {
				l.Log.Errorf("commands - NewUserNewCommand - login - readString(): %w", err)
				return
			}

			password, err := readString("Password: ")
			if err != nil {
				l.Log.Errorf("commands - NewUserNewCommand - password - readString(): %w", err)
				return
			}

			user := models.User{Login: login, Password: password}
			if err = userService.Add(context.Background(), &user); err != nil {
				l.Log.Errorf("commands - NewUserNewCommand - userService.Add(): %w", err)
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
