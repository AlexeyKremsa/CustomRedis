package main

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/AlexeyKremsa/CustomRedis/storage"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// Index returns a simple response to check if server is alive
func Index(w http.ResponseWriter, r *http.Request) {
	WriteResponseEmpty(w, r, http.StatusOK)
}

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

func (cr *CustomRedis) Delete(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["key"]
	if key == "" {
		WriteResponseMessage(w, r, http.StatusBadRequest, errEmptyKey)
		return
	}

	cr.Storage.RemoveItem(key)

	WriteResponseEmpty(w, r, http.StatusOK)
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
