package main

import (
	"os"

	"github.com/AlexeyKremsa/CustomRedis/api"
	"github.com/AlexeyKremsa/CustomRedis/config"
	"github.com/AlexeyKremsa/CustomRedis/storage"

	log "github.com/sirupsen/logrus"
)

// CustomRedis describes whole API service
type CustomRedis struct {
	Storage *storage.Storage
}

func main() {
	cfg := config.LoadConfig("../../config/config.toml")
	setupLogger(cfg)
	st := storage.Init(uint64(cfg.CleanupTimeoutSec), uint64(cfg.ShardCount))

	api.ServerStart(st, cfg)
}

func setupLogger(config *config.Config) {
	log.SetOutput(os.Stdout)
	lvl, err := log.ParseLevel(config.LogLevel)
	if err != nil {
		log.Fatalf("Failed to parse log level. %v", err)
	}
	log.SetLevel(lvl)
}
