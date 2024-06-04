package main

import (
	app "github.com/webkimru/go-keeper/internal/app/server"
	"github.com/webkimru/go-keeper/internal/app/server/config"
	"log"
	"runtime"
)

const appName = "GophKeeper Server"

var (
	buildVersion string = "N/A"
	buildDate    string = "N/A"
	buildCommit  string = "N/A"
)

func main() {
	log.Printf("******************************************")
	log.Printf("** %s%s%s built in %s", "\033[31m", appName, "\033[0m", runtime.Version())
	log.Printf("**----------------------------------------")
	log.Printf("** Running with %d Processors", runtime.NumCPU())
	log.Printf("** Running on %s", runtime.GOOS)
	log.Printf("** Build version: %s", buildVersion)
	log.Printf("** Build date: %s", buildDate)
	log.Printf("** Build commit: %s", buildCommit)
	log.Printf("******************************************")

	// Configuration
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
