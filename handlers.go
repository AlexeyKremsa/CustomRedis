package main

import (
	"net/http"
)

func Heartbit(w http.ResponseWriter, r *http.Request) {
	WriteResponseEmpty(w,r, http.StatusOK)
}