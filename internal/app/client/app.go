package client

import (
	"log"

	"github.com/webkimru/go-keeper/internal/app/client/cli/commands"
	"github.com/webkimru/go-keeper/internal/app/client/cli/grpc"
	"github.com/webkimru/go-keeper/internal/app/client/config"
	sqliteStore "github.com/webkimru/go-keeper/internal/app/client/repository/store/sqlite"
	"github.com/webkimru/go-keeper/internal/app/client/service"
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
		"APP_SECRET_KEY", cfg.SecretKey,
		"APP_TOKEN_EXP", cfg.TokenExp,
		"LOG_LEVEL", cfg.Log.Level,
		"GRPC_ADDRESS", cfg.GRPC.Address,
		"DATABASE_DSN", cfg.SQLite.DatabaseDSN,
		"MIGRATION_VERSION", cfg.SQLite.MigrationVersion,
	)

	db, err := sqlite.New(
		sqlite.PingInterval(cfg.SQLite.PingInterval),
		sqlite.DataSourcePath(cfg.SQLite.DataSourcePath),
	)
	if err != nil {
		l.Log.Errorf("app - client - Run - sqlite.New()", err)
	}
	if err = db.Migrate(cfg.SQLite.MigrationVersion); err != nil {
		l.Log.Fatal(err)
	}

	l.Log.Infof("Initiating gRPC client on %s", cfg.GRPC.Address)
	client := grpc.NewClient(cfg, l)

	userService := service.NewUserService(
		sqliteStore.NewUserStorage(db, cfg),
		client,
		cfg,
		l,
	)

	commands.Execute(userService, cfg, l)

	err = client.Close()
	if err != nil {
		l.Log.Errorf("app - client - Run - client.Close()", err)
	}

}
