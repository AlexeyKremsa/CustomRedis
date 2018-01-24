package main

import (
	"encoding/json"
	"errors"
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

	cr.Storage.Set(req.Key, req.Value, req.ExpirationSec)

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

	err = cr.Storage.SetNX(req.Key, req.Value, req.ExpirationSec)
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

	val, err := cr.Storage.Get(key)
	if err != nil {
		handleError(w, r, err)
		return
	}

	if val == nil {
		WriteResponseEmpty(w, r, http.StatusNoContent)
		return
	}

	if str, ok := val.(string); ok {
		WriteResponseData(w, r, http.StatusOK, str)
		return
	}

	handleError(w, r, errors.New(errWrongType))
}
