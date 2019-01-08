package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func (srv *server) setStr(w http.ResponseWriter, r *http.Request) {
	var req strRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeResponseMessage(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	if req.Key == "" {
		writeResponseMessage(w, r, http.StatusBadRequest, errEmptyKey)
		return
	}

	if req.Value == "" {
		writeResponseMessage(w, r, http.StatusBadRequest, errEmptyValue)
		return
	}

	srv.storage.SetStr(req.Key, req.Value, req.ExpirationSec)

	writeResponseEmpty(w, r, http.StatusCreated)
}

func (srv *server) setStrNX(w http.ResponseWriter, r *http.Request) {
	var req strRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeResponseMessage(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	if req.Key == "" {
		writeResponseMessage(w, r, http.StatusBadRequest, errEmptyKey)
		return
	}

	if req.Value == "" {
		writeResponseMessage(w, r, http.StatusBadRequest, errEmptyValue)
		return
	}

	err = srv.storage.SetStrNX(req.Key, req.Value, req.ExpirationSec)
	if err != nil {
		handleError(w, r, err)
		return
	}

	writeResponseEmpty(w, r, http.StatusCreated)
}

func (srv *server) getStr(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["key"]
	if key == "" {
		writeResponseMessage(w, r, http.StatusBadRequest, errEmptyKey)
		return
	}

	val, err := srv.storage.GetStr(key)
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
