package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	errEmptyKey   = "Key can not be empty"
	errEmptyValue = "Value can not be empty"
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
	}

	if kv.Key == "" {
		WriteResponseMessage(w, r, http.StatusBadRequest, errEmptyKey)
	}

	if kv.Value == "" {
		WriteResponseMessage(w, r, http.StatusBadRequest, errEmptyValue)
	}

	cr.Storage.SetStr(kv.Key, kv.Value)

	WriteResponseEmpty(w, r, http.StatusCreated)
}

func (cr *CustomRedis) GetStr(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["key"]
	if key == "" {
		WriteResponseMessage(w, r, http.StatusBadRequest, errEmptyKey)
	}

	val, err := cr.Storage.GetStr(key)
	if err != nil {
		WriteResponseMessage(w, r, http.StatusInternalServerError, err.Error())
	}

	if val == "" {
		WriteResponseEmpty(w, r, http.StatusNoContent)
	}

	WriteResponseJSON(w, r, http.StatusOK, val)
}
