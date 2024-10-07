package grpc

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/webkimru/go-keeper/internal/app/client/config"
	"github.com/webkimru/go-keeper/pkg/grpcserver/client"
	"github.com/webkimru/go-keeper/pkg/logger"
)

// NewClient returns a new gRPC client.
func NewClient(cfg *config.Config, l *logger.Log) *client.Client {
	transportOption := insecure.NewCredentials()
	clientOption := grpc.WithTransportCredentials(transportOption)

	// starting from reusable pkg
	cl, err := client.NewClient(cfg.GRPC.Address, clientOption)
	if err != nil {
		l.Log.Fatal(err)
	}

	return cl
}
