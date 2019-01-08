package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (srv *server) setList(w http.ResponseWriter, r *http.Request) {
	var req listRequest
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

	srv.storage.SetList(req.Key, req.Value, req.ExpirationSec)

	writeResponseEmpty(w, r, http.StatusCreated)
}

func (srv *server) listInsert(w http.ResponseWriter, r *http.Request) {
	var req listRequest
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

	count, err := srv.storage.ListInsert(req.Key, req.Value)
	if err != nil {
		handleError(w, r, err)
		return
	}

	writeResponseData(w, r, http.StatusCreated, count)
}

func (srv *server) getList(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["key"]
	if key == "" {
		writeResponseMessage(w, r, http.StatusBadRequest, errEmptyKey)
		return
	}

	val, err := srv.storage.GetList(key)
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

func (srv *server) listPop(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["key"]
	if key == "" {
		writeResponseMessage(w, r, http.StatusBadRequest, errEmptyKey)
		return
	}

	val, err := srv.storage.ListPop(key)
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
func (srv *server) listIndex(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["key"]
	if key == "" {
		writeResponseMessage(w, r, http.StatusBadRequest, errEmptyKey)
		return
	}

	index := mux.Vars(r)["index"]
	if index == "" {
		writeResponseMessage(w, r, http.StatusBadRequest, errEmptyIndex)
		return
	}

	i, err := strconv.Atoi(index)
	if err != nil {
		handleError(w, r, err)
		return
	}

	val, err := srv.storage.ListIndex(key, i)
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
