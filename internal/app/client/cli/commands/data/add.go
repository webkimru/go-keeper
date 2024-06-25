package data

import (
	"context"
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/webkimru/go-keeper/internal/app/client/models"
	"github.com/webkimru/go-keeper/internal/app/client/service"
	"github.com/webkimru/go-keeper/pkg/errs"
	"github.com/webkimru/go-keeper/pkg/logger"
)

// NewKeyValueAddCommand represents the initialized create key-value command
func NewKeyValueAddCommand(ctx context.Context, keyValueService *service.KeyValueService, l *logger.Log) *cobra.Command {
	return &cobra.Command{
		Use:   "add",
		Short: "Creator new key-value",
		Long:  "Allows creating a new ke-value data",
		Run: func(cmd *cobra.Command, args []string) {
			title, err := readString("Title: ")
			CLIog(l, "commands - NewKeyValueAddCommand - readString(title): %w", err)

			key, err := readString("Key: ")
			CLIog(l, "commands - NewKeyValueAddCommand - readString(key): %w", err)

			value, err := readString("Value: ")
			CLIog(l, "commands - NewKeyValueAddCommand - readString(value): %w", err)

			data := models.KeyValue{Title: title, Key: key, Value: value}
			err = keyValueService.Add(ctx, data)
			if err != nil {
				if errors.Is(err, errs.ErrPermissionDenied) {
					errs.CLIMsgPermissionDenied()
					return
				}
				if errors.Is(err, errs.ErrBadRequest) {
					errs.CLIMsgBadRequest()
					return
				}

				l.Log.Errorf("commands - NewKeyValueAddCommand - keyValueService.Add(): %w", err)
				errs.CLIMsgSeeLog()
				return
			}
		},
	}
}

// CLIog wrapped logger and CLI message.
func CLIog(l *logger.Log, s string, err error) {
	if err != nil {
		l.Log.Errorf(s, err)
		errs.CLIMsgSeeLog()
		os.Exit(1)
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
