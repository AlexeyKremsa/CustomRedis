package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	WriteResponseEmpty(w, r, http.StatusOK)
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

func (cr *CustomRedis) Keys(w http.ResponseWriter, r *http.Request) {
	keys := cr.Storage.GetAllKeys()

	if len(keys) == 0 {
		WriteResponseEmpty(w, r, http.StatusNoContent)
	}

	WriteResponseData(w, r, http.StatusOK, keys)
}
