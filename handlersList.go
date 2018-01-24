package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const errWrongType = "Operation against a key holding the wrong kind of value"

func (cr *CustomRedis) SetList(w http.ResponseWriter, r *http.Request) {
	var req ListRequest
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

	cr.Storage.Set(req.Key, req.Value, req.ExpirationSec)

	WriteResponseEmpty(w, r, http.StatusCreated)
}

func (cr *CustomRedis) PushList(w http.ResponseWriter, r *http.Request) {
	var req ListUpdateRequest
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

	err = cr.Storage.ListPush(req.Key, req.Value)
	if err != nil {
		handleError(w, r, err)
		return
	}

	WriteResponseEmpty(w, r, http.StatusCreated)
}

func (cr *CustomRedis) GetList(w http.ResponseWriter, r *http.Request) {
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

	if str, ok := val.([]string); ok {
		WriteResponseData(w, r, http.StatusOK, str)
		return
	}

	handleError(w, r, errors.New(errWrongType))
}

func (cr *CustomRedis) ListPop(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["key"]
	if key == "" {
		WriteResponseMessage(w, r, http.StatusBadRequest, errEmptyKey)
		return
	}

	val, err := cr.Storage.ListPop(key)
	if err != nil {
		handleError(w, r, err)
		return
	}

	WriteResponseData(w, r, http.StatusOK, val)
}
func (cr *CustomRedis) ListIndex(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["key"]
	if key == "" {
		WriteResponseMessage(w, r, http.StatusBadRequest, errEmptyKey)
		return
	}

	index := mux.Vars(r)["index"]
	if index == "" {
		WriteResponseMessage(w, r, http.StatusBadRequest, errEmptyIndex)
		return
	}

	i, err := strconv.Atoi(index)
	if err != nil {
		handleError(w, r, err)
		return
	}

	val, err := cr.Storage.ListIndex(key, i)
	if err != nil {
		handleError(w, r, err)
		return
	}

	WriteResponseData(w, r, http.StatusOK, val)
}
