package data

import (
	"context"
	"fmt"
	"io"

	"github.com/spf13/cobra"

	"github.com/webkimru/go-keeper/internal/app/client/service"
	"github.com/webkimru/go-keeper/pkg/logger"
)

// NewKeyValueCommand represents the key-value data command init
func NewKeyValueCommand(ctx context.Context, in io.Reader, keyValueService *service.KeyValueService, log *logger.Log) *cobra.Command {
	var dataCmd = cobra.Command{
		Use:   "keyvalue",
		Short: "Manage key-value",
		Long:  "Use subcommands to manage the key-value data",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("to control use this subcommand:")
			fmt.Println("keyvalue [add] or [get] or [upd] or [del]")

		},
	}

	dataCmd.AddCommand(NewKeyValueAddCommand(ctx, in, keyValueService, log))
	dataCmd.AddCommand(NewKeyValueUpdCommand(ctx, in, keyValueService, log))
	dataCmd.AddCommand(NewKeyValueDelCommand(ctx, in, keyValueService, log))
	dataCmd.AddCommand(NewKeyValueListCommand(ctx, keyValueService, log))

	return &dataCmd
}
