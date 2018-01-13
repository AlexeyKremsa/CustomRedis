package main

import (
	"net/http"
	"time"
	log "github.com/sirupsen/logrus"
	"os"

	"github.com/AlexeyKremsa/CustomRedis/config"
)

// loggerHandler Log all HTTP requests to output in a proper format.
func loggerHandler(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		inner.ServeHTTP(w, r)
		log.Debugf("%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}

func setupLogger(config *config.CRConfig) {
	log.SetOutput(os.Stdout)
	lvl, err := log.ParseLevel(config.LogLevel)
	if err != nil {
		log.Fatalf("Failed to parse log level. %v", err)
	}
	log.SetLevel(lvl)
}
