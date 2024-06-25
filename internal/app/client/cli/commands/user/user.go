package user

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/webkimru/go-keeper/internal/app/client/service"
	"github.com/webkimru/go-keeper/pkg/logger"
)

// NewUserCommand represents the user command init
func NewUserCommand(userService *service.UserService, log *logger.Log) *cobra.Command {
	var userCmd = cobra.Command{
		Use:   "user",
		Short: "Manage users",
		Long:  "Use subcommands to manage the users",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("to control use this command: user new")

		},
	}

	userCmd.AddCommand(NewUserNewCommand(userService, log))
	userCmd.AddCommand(NewUserLoginCommand(userService, log))

	return &userCmd
}
