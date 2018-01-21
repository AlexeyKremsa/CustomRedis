package main

import (
	"fmt"
	"net/http"

	"github.com/AlexeyKremsa/CustomRedis/config"
	"github.com/AlexeyKremsa/CustomRedis/storage"

	log "github.com/sirupsen/logrus"
)

// CustomRedis describes whole API service
type CustomRedis struct {
	Storage *storage.Storage
}

func main() {
	cfg := config.LoadConfig()
	setupLogger(cfg)
	st := storage.Init()

	cr := &CustomRedis{
		Storage: st,
	}
	router := NewRouter(cr)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), router))
}
