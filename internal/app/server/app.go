package server

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/webkimru/go-keeper/internal/app/server/api/grpc"
	"github.com/webkimru/go-keeper/internal/app/server/config"
	"github.com/webkimru/go-keeper/internal/app/server/repository/store/pg"
	"github.com/webkimru/go-keeper/internal/app/server/service"
	"github.com/webkimru/go-keeper/pkg/crypt"
	"github.com/webkimru/go-keeper/pkg/jwtmanager"
	"github.com/webkimru/go-keeper/pkg/logger"
	"github.com/webkimru/go-keeper/pkg/postgres"
)

// Run runs application.
func Run(cfg *config.Config) {
	l, err := logger.NewZap(cfg.Log.Level)
	if err != nil {
		log.Fatal(err)
	}

	l.Log.Infoln("Starting configuration:",
		"APP_SECRET_KEY", cfg.SecretKey,
		"APP_TOKEN_EXP", cfg.TokenExp,
		"LOG_LEVEL", cfg.Log.Level,
		"GRPC_ADDRESS", cfg.GRPC.Address,
		"DATABASE_DSN", cfg.PG.DatabaseDSN,
		"MIGRATION_VERSION", cfg.PG.MigrationVersion,
	)

	// PostgreSQL initialization
	pgx, err := postgres.New(
		cfg.PG.DatabaseDSN,
		postgres.ConnectTimeout(cfg.PG.ConnectTimeout),
	)
	if err != nil {
		l.Log.Fatal(err)
	}
	defer pgx.Close()
	// migration
	if err = pgx.Migrate(cfg.PG.MigrationVersion); err != nil {
		l.Log.Fatal(err)
	}

	// Manager initialization
	jwtManager := jwtmanager.New(cfg.SecretKey, cfg.TokenExp) // for user endpoints
	cryptManager, err := crypt.New(cfg.SecretKey)             // cryptManager to encrypt data: key-value
	if err != nil {
		l.Log.Error(err)
	}
	// Service initialization with DB:
	userService := service.NewUserService(pg.NewUserStorage(pgx, cfg)) // pg.NewUserStorage(postgresDB) // inmemory.NewUserStorage()
	keyValueService := service.NewKeyValueService(
		pg.NewKeyValueStorage(pgx), // pg.NewKeyValueStorage(postgresDB) // inmemory.NewKeyValueStorage()
		cryptManager,
	)

	// Start gRPC server with services:
	// - userServer to store the users
	// - keyValueServer to store key-value data: login/pass
	l.Log.Infof("Starting gRPC server on %s", cfg.GRPC.Address)
	server := grpc.New(userService, keyValueService, jwtManager, cfg, l)

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	select {
	case s := <-interrupt:
		l.Log.Info("Got signal: " + s.String())

	case err = <-server.Notify():
		l.Log.Errorf("grpcServer.Notify: %v", err)
	}

	// Shutdown
	server.Shutdown()
}
