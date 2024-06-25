package data

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

// NewKeyValueUpdCommand represents the initialized update key-value command
func NewKeyValueUpdCommand(ctx context.Context, keyValueService *service.KeyValueService, l *logger.Log) *cobra.Command {
	return &cobra.Command{
		Use:   "upd",
		Short: "Updater key-value data",
		Long:  "Allows update key-value data",
		Run: func(cmd *cobra.Command, args []string) {
			id, err := readInt("ID: ")
			CLIog(l, "commands - NewKeyValueUpdCommand - readString(id): %w", err)

			title, err := readString("Title: ")
			CLIog(l, "commands - NewKeyValueUpdCommand - readString(title): %w", err)

			key, err := readString("Key: ")
			CLIog(l, "commands - NewKeyValueUpdCommand - readString(key): %w", err)

			value, err := readString("Value: ")
			CLIog(l, "commands - NewKeyValueUpdCommand - readString(value): %w", err)

			data := models.KeyValue{ID: int64(id), Title: title, Key: key, Value: value}
			err = keyValueService.Update(ctx, data)
			if err != nil {
				if errors.Is(err, errs.ErrPermissionDenied) {
					errs.CLIMsgPermissionDenied()
					return
				}
				if errors.Is(err, errs.ErrBadRequest) {
					errs.CLIMsgBadRequest()
					return
				}

				l.Log.Errorf("commands - NewKeyValueUpdCommand - keyValueService.Update(): %w", err)
				errs.CLIMsgSeeLog()
				return
			}
		},
	}
}

// readInt works as prompt UI
func readInt(s string) (int, error) {
	var input int
	for {
		fmt.Print(s)
		_, err := fmt.Scan(&input)
		if err != nil {
			return -1, fmt.Errorf("commands - readString - fmt.Scan(): %w", err)
		}
		if input != 0 {
			break
		}
	}

	return input, nil
}
