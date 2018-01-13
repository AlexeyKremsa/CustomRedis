package main

import (
	"fmt"
	"net/http"
	"github.com/AlexeyKremsa/CustomRedis/config"

	log "github.com/sirupsen/logrus"
)

func main() {
	cfg := config.LoadConfig()
	setupLogger(cfg)

	router := NewRouter()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), router))
}