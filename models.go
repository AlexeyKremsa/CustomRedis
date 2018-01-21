package main

// KeyValue holds fields for string key value storage
type KeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value,omitempty"`
}
