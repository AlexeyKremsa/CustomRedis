package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// WriteResponseMessage sends a response with status and message
func WriteResponseMessage(w http.ResponseWriter, r *http.Request, status int, message string) {
	WriteResponseMessageWithData(w, r, status, message, nil)
}

// WriteResponseMessageWithData sends a response with status, message and data
func WriteResponseMessageWithData(w http.ResponseWriter, r *http.Request, status int, message string, data interface{}) {
	type responseJSON struct {
		Message string      `json:"message,omitempty"`
		Data    interface{} `json:"data,omitempty"`
	}
	response := responseJSON{
		Message: message,
		Data:    data,
	}
	WriteResponseJSON(w, r, status, response)
}

// WriteResponseJSON sends a response with status and data
func WriteResponseJSON(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Failed build response")
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		fmt.Fprint(w, string(jsonData))
	}
}

// WriteResponseEmpty sends a response only with status
func WriteResponseEmpty(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
}
