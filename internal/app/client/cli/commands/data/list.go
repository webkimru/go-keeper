package data

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/webkimru/go-keeper/internal/app/client/service"
	"github.com/webkimru/go-keeper/pkg/errs"
	"github.com/webkimru/go-keeper/pkg/logger"
)

// NewKeyValueListCommand represents the initialized list key-value command
func NewKeyValueListCommand(ctx context.Context, keyValueService *service.KeyValueService, l *logger.Log) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "Explorer key-value data",
		Long:  "Allows to show the list of the key-value data",
		Run: func(cmd *cobra.Command, args []string) {
			data, err := keyValueService.List(ctx)
			if err != nil {
				if errors.Is(err, errs.ErrPermissionDenied) {
					errs.CLIMsgPermissionDenied()
					return
				}
				if errors.Is(err, errs.ErrBadRequest) {
					errs.CLIMsgBadRequest()
					return
				}

				l.Log.Errorf("commands - NewKeyValueListCommand - keyValueService.List(): %w", err)
				errs.CLIMsgSeeLog()
				return
			}

			for _, item := range data {
				fmt.Println(
					"ID:", item.ID, "|",
					"Title:", item.Title, "|",
					"Key:", item.Key, "|",
					"Value:", item.Value,
				)
			}
		},
	}
}
