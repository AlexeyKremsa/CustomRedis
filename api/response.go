package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AlexeyKremsa/CustomRedis/storage"
	"github.com/labstack/gommon/log"
)

type responseJSON struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func handleError(w http.ResponseWriter, r *http.Request, err error) {
	switch err.(type) {
	// business case error happened, return status 200 because request was handled properly
	case storage.ErrBusiness:
		log.Debug(err.Error())
		writeResponseMessage(w, r, http.StatusOK, err.Error())

	// unexpected error happened
	default:
		log.Error(err.Error())
		writeResponseMessage(w, r, http.StatusInternalServerError, err.Error())
	}
}

// writeResponseData sends a response with status and data
func writeResponseData(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	writeResponseMessageWithData(w, r, status, "", data)
}

// writeResponseMessage sends a response with status and message
func writeResponseMessage(w http.ResponseWriter, r *http.Request, status int, message string) {
	writeResponseMessageWithData(w, r, status, message, nil)
}

// writeResponseMessageWithData sends a response with status, message and data
func writeResponseMessageWithData(w http.ResponseWriter, r *http.Request, status int, message string, data interface{}) {
	response := responseJSON{
		Message: message,
		Data:    data,
	}
	writeResponseJSON(w, r, status, response)
}

// writeResponseJSON sends a response with status and data
func writeResponseJSON(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
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

// writeResponseEmpty sends a response only with status
func writeResponseEmpty(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
}
