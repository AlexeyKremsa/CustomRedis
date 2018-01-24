package main

import (
	"net/http"
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/AlexeyKremsa/CustomRedis/config"
	"github.com/AlexeyKremsa/CustomRedis/storage"
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

func handleError(w http.ResponseWriter, r *http.Request, err error) {
	switch err.(type) {
	// business case error happened, return status 200 because request was handled properly
	case storage.ErrBusiness:
		log.Debug(err.Error())
		WriteResponseMessage(w, r, http.StatusOK, err.Error())

	// unexpected error happened
	default:
		log.Error(err.Error())
		WriteResponseMessage(w, r, http.StatusInternalServerError, err.Error())
	}
}
