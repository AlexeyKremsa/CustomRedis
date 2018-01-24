package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func (cr *CustomRedis) SetStr(w http.ResponseWriter, r *http.Request) {
	var req StrRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		WriteResponseMessage(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	if req.Key == "" {
		WriteResponseMessage(w, r, http.StatusBadRequest, errEmptyKey)
		return
	}

	if req.Value == "" {
		WriteResponseMessage(w, r, http.StatusBadRequest, errEmptyValue)
		return
	}

	cr.Storage.SetStr(req.Key, req.Value, req.ExpirationSec)

	WriteResponseEmpty(w, r, http.StatusCreated)
}

func (cr *CustomRedis) SetStrNX(w http.ResponseWriter, r *http.Request) {
	var req StrRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		WriteResponseMessage(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	if req.Key == "" {
		WriteResponseMessage(w, r, http.StatusBadRequest, errEmptyKey)
		return
	}

	if req.Value == "" {
		WriteResponseMessage(w, r, http.StatusBadRequest, errEmptyValue)
		return
	}

	err = cr.Storage.SetStrNX(req.Key, req.Value, req.ExpirationSec)
	if err != nil {
		handleError(w, r, err)
		return
	}

	WriteResponseEmpty(w, r, http.StatusCreated)
}

func (cr *CustomRedis) GetStr(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["key"]
	if key == "" {
		WriteResponseMessage(w, r, http.StatusBadRequest, errEmptyKey)
		return
	}

	val, err := cr.Storage.GetStr(key)
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
