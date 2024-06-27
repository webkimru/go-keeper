package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config is a general structure.
	Config struct {
		App  `json:"app"`
		Log  `json:"logger"`
		GRPC `json:"grpc"`
		PG   `json:"dsn"`
	}

	// App is an application configuration structure.
	App struct {
		SecretKey string `json:"secret_key" env:"APP_SECRET_KEY" env-default:"secret"`
		TokenExp  int    `json:"token_exp" env:"APP_TOKEN_EXP" env-default:"120" env-description:"token expiration (minutes)"`
	}

	// Log is a logging structure.
	Log struct {
		Level string `json:"log_level" env:"LOG_LEVEL" env-default:"info"`
	}

	// GRPC is a gRPC server structure.
	GRPC struct {
		Address string `json:"grpc_address" env:"GRPC_ADDRESS" env-default:":3200"`
	}

	// PG is a PostgreSQL structure.
	PG struct {
		DatabaseDSN      string `json:"database_dsn" env:"DATABASE_DSN"`
		MigrationVersion int64  `json:"migration_version" env:"DB_MIGRATION_VERSION" env-default:"1"`
		ConnectTimeout   int    `json:"connect_timeout" env:"DB_CONNECT_TIMEOUT" env-default:"60"`
		QueryTimeout     int    `json:"query_timeout" env:"DB_QUERY_TIMEOUT" env-default:"10"`
	}
)

// Args command-line parameters.
type Args struct {
	ConfigPath string
}

// New returns app config.
func New() (*Config, error) {
	var cfg Config

	// Read config from flags
	args, err := FlagArgs(&cfg)
	if err != nil {
		return nil, fmt.Errorf("config - New - FlagArgs(): %w", err)
	}

	// Read config from file
	if err = cleanenv.ReadConfig(args.ConfigPath, &cfg); args.ConfigPath != "" && err != nil {
		return nil, fmt.Errorf("config - New - cleanenv.ReadConfig(): %w", err)
	}

	// Read config from ENV
	if err = cleanenv.ReadEnv(&cfg); err != nil {
		return nil, fmt.Errorf("config - New - cleanenv.ReadEnv(): %w", err)
	}

	return &cfg, nil
}

// FlagArgs processes and handles CLI arguments.
func FlagArgs(cfg *Config) (Args, error) {
	var a Args
	f := flag.NewFlagSet("GophKeeper Server", 1) // debug - flag.ContinueOnError
	f.StringVar(&a.ConfigPath, "c", "", "path to config file")

	f.Usage = cleanenv.FUsage(f.Output(), cfg, nil, f.Usage)

	err := f.Parse(os.Args[1:])
	if err != nil {
		return a, fmt.Errorf("config - FlagArgs - f.Parse(): %w", err)
	}

	return a, nil
}
