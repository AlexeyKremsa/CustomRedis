package main

// StrRequest holds fields for string key value storage
type StrRequest struct {
	Key           string `json:"key"`
	Value         string `json:"value"`
	ExpirationSec int64  `json:"expirationSec"`
}

// ListRequest holds fields for list key value storage
type ListRequest struct {
	Key           string   `json:"key"`
	Value         []string `json:"value"`
	ExpirationSec int64    `json:"expirationSec"`
}
