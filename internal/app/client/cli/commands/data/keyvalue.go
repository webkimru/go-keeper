package data

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/webkimru/go-keeper/internal/app/client/service"
	"github.com/webkimru/go-keeper/pkg/logger"
)

// NewKeyValueCommand represents the key-value data command init
func NewKeyValueCommand(ctx context.Context, keyValueService *service.KeyValueService, log *logger.Log) *cobra.Command {
	var dataCmd = cobra.Command{
		Use:   "keyvalue",
		Short: "Manage key-value",
		Long:  "Use subcommands to manage the key-value data",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("to control use this subcommand:")
			fmt.Println("keyvalue [add] or [get] or [upd] or [del]")

		},
	}

	dataCmd.AddCommand(NewKeyValueAddCommand(ctx, keyValueService, log))

	return &dataCmd
}
