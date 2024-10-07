package main

import (
	"log"

	app "github.com/webkimru/go-keeper/internal/app/client"
	"github.com/webkimru/go-keeper/internal/app/client/config"
)

const appName = "GophKeeper Client"

var (
	buildVersion string = "N/A"
	buildDate    string = "N/A"
	buildCommit  string = "N/A"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}
	cfg.App.BuildName = appName
	cfg.App.BuildVersion = buildVersion
	cfg.App.BuildCommit = buildCommit
	cfg.App.BuildDate = buildDate

	app.Run(cfg)
}
