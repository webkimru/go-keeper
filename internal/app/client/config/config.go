package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config is a general structure.
	Config struct {
		App    `json:"app"`
		Log    `json:"logger"`
		GRPC   `json:"grpc"`
		SQLite `json:"dsn"`
	}

	// App is an application configuration structure.
	App struct {
		BuildName    string `json:"build_name" env:"APP_BUILD_NAME"`
		BuildVersion string `json:"build_version" env:"APP_BUILD_VERSION"`
		BuildCommit  string `json:"build_commit" env:"APP_BUILD_COMMIT"`
		BuildDate    string `json:"build_date" env:"APP_BUILD_COMMIT"`
		SecretKey    string `json:"secret_key" env:"APP_SECRET_KEY" env-default:"secret"`
		TokenExp     int    `json:"token_exp" env:"APP_TOKEN_EXP" env-default:"120" env-description:"token expiration (minutes)"`
		Token        string `json:"token" env:"APP_TOKEN"`
	}

	// Log is a logging structure.
	Log struct {
		Level         string `json:"log_level" env:"LOG_LEVEL" env-default:"error"`
		LogSourcePath string `json:"log_source_path" env:"LOG_SOURCE_PATH" env-default:"client.log"` // IDE version - pkg/sqlite/data/data.db
	}

	// GRPC is a client structure working with the gRPC server.
	GRPC struct {
		Address      string `json:"grpc_address" env:"GRPC_ADDRESS" env-default:":3200"`
		QueryTimeout int    `json:"query_timeout" env:"GRPC_QUERY_TIMEOUT" env-default:"5"`
	}

	// SQLite is a SQLite structure.
	SQLite struct {
		DatabaseDSN      string `json:"database_dsn" env:"DATABASE_DSN" env-default:"client.db"`
		MigrationVersion int64  `json:"migration_version" env:"DB_MIGRATION_VERSION" env-default:"1"`
		PingInterval     int    `json:"ping_interval" env:"DB_PING_INTERVAL" env-default:"1"`
		QueryTimeout     int    `json:"query_timeout" env:"DB_QUERY_TIMEOUT" env-default:"5"`
	}
)

// New returns app config.
func New() (*Config, error) {
	var cfg Config

	// Read config from file
	_ = cleanenv.ReadConfig("config.json", &cfg) // silent loading config with token after auth

	// Read config from ENV
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
