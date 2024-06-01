package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

type (
	// Config is an application configuration structure.
	Config struct {
		Log `json:"logger"`
	}

	Log struct {
		Level string `json:"log_level" env:"LOG_LEVEL" env-default:"info"`
	}
)

// Args command-line parameters.
type Args struct {
	ConfigPath string
}

// New returns app config.
func New() (*Config, error) {
	cfg := &Config{}

	// Read config from flags
	args, err := FlagArgs(cfg)
	if err != nil {
		return nil, err
	}

	// Read config from file
	if err = cleanenv.ReadConfig(args.ConfigPath, cfg); args.ConfigPath != "" && err != nil {
		return nil, err
	}

	// Read config from ENV
	if err = cleanenv.ReadEnv(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

// FlagArgs processes and handles CLI arguments.
func FlagArgs(cfg *Config) (Args, error) {
	var a Args
	f := flag.NewFlagSet("GophKeeper Server", flag.ContinueOnError)
	f.StringVar(&a.ConfigPath, "c", "", "path to config file")

	f.Usage = cleanenv.FUsage(f.Output(), &cfg, nil, f.Usage)

	err := f.Parse(os.Args[1:])
	if err != nil {
		return a, err
	}

	return a, nil
}
