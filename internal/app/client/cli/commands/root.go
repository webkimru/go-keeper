package commands

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/webkimru/go-keeper/internal/app/client/cli/commands/user"
	"github.com/webkimru/go-keeper/internal/app/client/config"
	"github.com/webkimru/go-keeper/internal/app/client/service"
	"github.com/webkimru/go-keeper/pkg/logger"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(userService *service.UserService, cfg *config.Config, log *logger.Log) {
	// rootCmd represents the base command when called without any subcommands
	var rootCmd = &cobra.Command{
		Use:   "go-keeper",
		Short: "Go-keeper basic CLI",
		Long:  "Go-keeper is a friendly command line application for safe keeping key-value data",
	}
	rootCmd.AddCommand(user.NewUserCommand(userService, log))
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
