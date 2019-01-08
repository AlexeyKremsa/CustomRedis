package api

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
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

func Index(w http.ResponseWriter, r *http.Request) {
	writeResponseEmpty(w, r, http.StatusOK)
}

func (srv *server) Delete(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["key"]
	if key == "" {
		writeResponseMessage(w, r, http.StatusBadRequest, errEmptyKey)
		return
	}

	srv.storage.RemoveItem(key)

	writeResponseEmpty(w, r, http.StatusOK)
}

func (srv *server) Keys(w http.ResponseWriter, r *http.Request) {
	keys := srv.storage.GetAllKeys()

	if len(keys) == 0 {
		writeResponseEmpty(w, r, http.StatusNoContent)
		return
	}

	writeResponseData(w, r, http.StatusOK, keys)
}
