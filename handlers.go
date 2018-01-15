package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Index returns a simple response to check if server is alive
func Index(w http.ResponseWriter, r *http.Request) {
	WriteResponseEmpty(w, r, http.StatusOK)
}

func SetKey(w http.ResponseWriter, r *http.Request) {
	var kv KeyValue
	err := json.NewDecoder(r.Body).Decode(&kv)
	if err != nil {
		WriteResponseMessage(w, r, http.StatusBadRequest, err.Error())
	}

	fmt.Println(kv.Key + " " + kv.Value)

}
