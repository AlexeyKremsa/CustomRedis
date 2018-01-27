package main

// StrRequest holds fields for string key value storage
type StrRequest struct {
	Key           string `json:"key"`
	Value         string `json:"strValue"`
	ExpirationSec uint64 `json:"expirationSec,omitempty"`
}

// ListRequest holds fields for list key value storage
type ListRequest struct {
	Key           string   `json:"key"`
	Value         []string `json:"listValue"`
	ExpirationSec uint64   `json:"expirationSec,omitempty"`
}

// ListUpdateRequest holds fields used to update list
type ListUpdateRequest struct {
	Key   string   `json:"key"`
	Value []string `json:"listValue"`
}

// MapRequest holds fields for map key value storage
type MapRequest struct {
	Key           string            `json:"key"`
	Value         map[string]string `json:"mapValue"`
	ExpirationSec uint64            `json:"expirationSec,omitempty"`
}
