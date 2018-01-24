package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func (cr *CustomRedis) SetMap(w http.ResponseWriter, r *http.Request) {
	var req MapRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		WriteResponseMessage(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	if req.Key == "" {
		WriteResponseMessage(w, r, http.StatusBadRequest, errEmptyKey)
		return
	}

	if len(req.Value) == 0 {
		WriteResponseMessage(w, r, http.StatusBadRequest, errEmptyValue)
		return
	}

	cr.Storage.SetMap(req.Key, req.Value, req.ExpirationSec)

	WriteResponseEmpty(w, r, http.StatusCreated)
}

func (cr *CustomRedis) GetMap(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["key"]
	if key == "" {
		WriteResponseMessage(w, r, http.StatusBadRequest, errEmptyKey)
		return
	}

	val, err := cr.Storage.GetMap(key)
	if err != nil {
		handleError(w, r, err)
		return
	}

	if len(val) == 0 {
		WriteResponseEmpty(w, r, http.StatusNoContent)
		return
	}

	WriteResponseData(w, r, http.StatusOK, val)
}

func (cr *CustomRedis) GetMapItem(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["key"]
	if key == "" {
		WriteResponseMessage(w, r, http.StatusBadRequest, errEmptyKey)
		return
	}

	itemKey := mux.Vars(r)["itemKey"]
	if itemKey == "" {
		WriteResponseMessage(w, r, http.StatusBadRequest, errEmptyMapItemKey)
		return
	}

	val, err := cr.Storage.GetMapItem(key, itemKey)
	if err != nil {
		handleError(w, r, err)
		return
	}

	if val == "" {
		WriteResponseEmpty(w, r, http.StatusNoContent)
		return
	}

	WriteResponseData(w, r, http.StatusOK, val)
}
