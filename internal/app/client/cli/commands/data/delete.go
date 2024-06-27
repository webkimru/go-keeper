package data

import (
	"context"
	"io"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/webkimru/go-keeper/internal/app/client/service"
	"github.com/webkimru/go-keeper/pkg/errs"
	"github.com/webkimru/go-keeper/pkg/logger"
)

// NewKeyValueDelCommand represents the initialized delete key-value command
func NewKeyValueDelCommand(ctx context.Context, in io.Reader, keyValueService *service.KeyValueService, l *logger.Log) *cobra.Command {
	return &cobra.Command{
		Use:   "del",
		Short: "Deleter key-value data",
		Long:  "Allows delete key-value data",
		Run: func(cmd *cobra.Command, args []string) {
			id, err := readInt(in, "ID: ")
			CLIog(l, "commands - NewKeyValueDelCommand - readString(id): %w", err)

			err = keyValueService.Delete(ctx, int64(id))
			if err != nil {
				if errors.Is(err, errs.ErrPermissionDenied) {
					errs.CLIMsgPermissionDenied()
					return
				}
				if errors.Is(err, errs.ErrBadRequest) {
					errs.CLIMsgBadRequest()
					return
				}

				l.Log.Errorf("commands - NewKeyValueDelCommand - keyValueService.Delete(): %w", err)
				errs.CLIMsgSeeLog()
				return
			}
		},
	}
}
