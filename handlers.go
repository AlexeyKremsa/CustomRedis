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
	var kv KeyValue
	err := json.NewDecoder(r.Body).Decode(&kv)
	if err != nil {
		WriteResponseMessage(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	if kv.Key == "" {
		WriteResponseMessage(w, r, http.StatusBadRequest, errEmptyKey)
		return
	}

	if kv.Value == "" {
		WriteResponseMessage(w, r, http.StatusBadRequest, errEmptyValue)
		return
	}

	cr.Storage.Set(kv.Key, kv.Value, kv.ExpirationSec)

	WriteResponseEmpty(w, r, http.StatusCreated)
	return
}

func (cr *CustomRedis) SetStrNX(w http.ResponseWriter, r *http.Request) {
	var kv KeyValue
	err := json.NewDecoder(r.Body).Decode(&kv)
	if err != nil {
		WriteResponseMessage(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	if kv.Key == "" {
		WriteResponseMessage(w, r, http.StatusBadRequest, errEmptyKey)
		return
	}

	if kv.Value == "" {
		WriteResponseMessage(w, r, http.StatusBadRequest, errEmptyValue)
		return
	}

	err = cr.Storage.SetNX(kv.Key, kv.Value, kv.ExpirationSec)
	if err != nil {
		HandleError(w, r, err)
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
		HandleError(w, r, err)
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

	HandleError(w, r, errors.New(errWrongType))
}

func HandleError(w http.ResponseWriter, r *http.Request, err error) {
	switch err.(type) {
	// business case error happened, return status 200 because request was handled properly
	case storage.ErrBusiness:
		WriteResponseMessage(w, r, http.StatusOK, err.Error())

	// unexpected error happened
	default:
		log.Error(err.Error())
		WriteResponseMessage(w, r, http.StatusInternalServerError, err.Error())
	}
}
