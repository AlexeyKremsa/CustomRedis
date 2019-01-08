package api

// strRequest holds fields for string key value storage
type strRequest struct {
	Key           string `json:"key"`
	Value         string `json:"value"`
	ExpirationSec uint64 `json:"expirationSec,omitempty"`
}

// listRequest holds fields for list key value storage
type listRequest struct {
	Key           string   `json:"key"`
	Value         []string `json:"value"`
	ExpirationSec uint64   `json:"expirationSec,omitempty"`
}

// mapRequest holds fields for map key value storage
type mapRequest struct {
	Key           string            `json:"key"`
	Value         map[string]string `json:"value"`
	ExpirationSec uint64            `json:"expirationSec,omitempty"`
}
