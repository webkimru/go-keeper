package client

import (
	"context"
	"log"

	"github.com/webkimru/go-keeper/internal/app/client/cli/commands"
	"github.com/webkimru/go-keeper/internal/app/client/cli/grpc"
	"github.com/webkimru/go-keeper/internal/app/client/config"
	"github.com/webkimru/go-keeper/internal/app/client/models"
	sqliteStore "github.com/webkimru/go-keeper/internal/app/client/repository/store/sqlite"
	"github.com/webkimru/go-keeper/internal/app/client/service"
	"github.com/webkimru/go-keeper/pkg/crypt"
	"github.com/webkimru/go-keeper/pkg/jwtmanager"
	"github.com/webkimru/go-keeper/pkg/logger"
	"github.com/webkimru/go-keeper/pkg/sqlite"
)

// Run runs application.
func Run(cfg *config.Config) {
	l, err := logger.NewZap(
		logger.SetLevel(cfg.Log.Level),
		logger.SetOutput([]string{cfg.Log.LogSourcePath}),
	)
	if err != nil {
		log.Fatal(err)
	}

	l.Log.Infoln("Starting configuration:",
		"APP_TOKEN", cfg.Token,
		"APP_SECRET_KEY", cfg.App.SecretKey,
		"APP_TOKEN_EXP", cfg.App.TokenExp,
		"LOG_LEVEL", cfg.Log.Level,
		"GRPC_ADDRESS", cfg.GRPC.Address,
		"DATABASE_DSN", cfg.SQLite.DatabaseDSN,
		"MIGRATION_VERSION", cfg.SQLite.MigrationVersion,
	)

	db, err := sqlite.New(
		sqlite.PingInterval(cfg.SQLite.PingInterval),
		sqlite.DataSourcePath(cfg.SQLite.DatabaseDSN),
	)
	if err != nil {
		l.Log.Errorf("app - client - Run - sqlite.New()", err)
	}
	if err = db.Migrate(cfg.SQLite.MigrationVersion); err != nil {
		l.Log.Fatal(err)
	}

	l.Log.Infof("Initiating gRPC client on %s", cfg.GRPC.Address)
	client := grpc.NewClient(cfg, l)

	// jwtManager saving a token to the app config and reusing it in the commands
	jwtManager := jwtmanager.New(cfg.SecretKey, cfg.TokenExp)
	// userService business logic layer above the commands
	userService := service.NewUserService(
		sqliteStore.NewUserStorage(db, cfg),
		client,
		cfg,
		jwtManager,
		l,
	)

	// set context value after getting the token
	ctx := context.Background()
	if cfg.App.Token != "" {
		userID := jwtManager.GetUserID(cfg.App.Token)
		if userID != -1 {
			ctx = context.WithValue(ctx, models.ContextKey("userID"), userID)
		}
	}
	// cryptManager to encrypt local key-value data
	cryptManager, err := crypt.New(cfg.SecretKey)
	if err != nil {
		l.Log.Error(err)
	}
	keyValueService := service.NewKeyValueService(
		sqliteStore.NewKeyValueStorage(db, cfg),
		cryptManager,
		l,
	)

	commands.Execute(ctx, userService, keyValueService, cfg, l)

	err = client.Close()
	if err != nil {
		l.Log.Errorf("app - client - Run - client.Close()", err)
	}

}
