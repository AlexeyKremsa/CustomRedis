package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func (srv *server) setMap(w http.ResponseWriter, r *http.Request) {
	var req mapRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeResponseMessage(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	if req.Key == "" {
		writeResponseMessage(w, r, http.StatusBadRequest, errEmptyKey)
		return
	}

	if len(req.Value) == 0 {
		writeResponseMessage(w, r, http.StatusBadRequest, errEmptyValue)
		return
	}

	srv.storage.SetMap(req.Key, req.Value, req.ExpirationSec)

	writeResponseEmpty(w, r, http.StatusCreated)
}

func (srv *server) getMap(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["key"]
	if key == "" {
		writeResponseMessage(w, r, http.StatusBadRequest, errEmptyKey)
		return
	}

	val, err := srv.storage.GetMap(key)
	if err != nil {
		handleError(w, r, err)
		return
	}

	if len(val) == 0 {
		writeResponseEmpty(w, r, http.StatusNoContent)
		return
	}

	writeResponseData(w, r, http.StatusOK, val)
}

func (srv *server) getMapItem(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["key"]
	if key == "" {
		writeResponseMessage(w, r, http.StatusBadRequest, errEmptyKey)
		return
	}

	itemKey := mux.Vars(r)["itemKey"]
	if itemKey == "" {
		writeResponseMessage(w, r, http.StatusBadRequest, errEmptyMapItemKey)
		return
	}

	val, err := srv.storage.GetMapItem(key, itemKey)
	if err != nil {
		handleError(w, r, err)
		return
	}

	if val == "" {
		writeResponseEmpty(w, r, http.StatusNoContent)
		return
	}

	writeResponseData(w, r, http.StatusOK, val)
}
