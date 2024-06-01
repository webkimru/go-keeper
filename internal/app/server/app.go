package server

import (
	"github.com/webkimru/go-keeper/internal/app/server/config"
	"github.com/webkimru/go-keeper/pkg/logger"
	"log"
)

func Run(cfg *config.Config) {
	l, err := logger.NewZap(cfg.Log.Level)
	if err != nil {
		log.Fatal(err)
	}

	l.Log.Infoln("Starting configuration:",
		"LOG_LEVEL", cfg.Log.Level,
	)
}
