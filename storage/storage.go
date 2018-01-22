package storage

import (
	"errors"
	"sync"
	"time"
)

type TTL struct {
	Expiration int64
}

type Str struct {
	TTL
	Value string
}

type Storage struct {
	mutex sync.RWMutex
	// Also sync.Map could be considered for systems with heavy read operations and big amount of CPU cores
	keyValues map[string]interface{}
}

func Init() *Storage {
	strg := &Storage{}
	strg.keyValues = make(map[string]interface{})
	return strg
}

func (s *Storage) SetStr(key, value string, expirationSec int64) {
	var expTime int64
	if expirationSec > 0 {
		expTime = time.Now().Add(time.Second * time.Duration(expirationSec)).Unix()
	}
	item := Str{Value: value, TTL: TTL{Expiration: expTime}}

	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.keyValues[key] = item
}

func (s *Storage) SetStrNX(key, value string, expirationSec int64) error {
	if _, ok := s.keyValues[key]; ok {
		return newErrCustom("Key already exists")
	}

	var expTime int64
	if expirationSec > 0 {
		expTime = time.Now().Add(time.Second * time.Duration(expirationSec)).Unix()
	}
	item := Str{Value: value, TTL: TTL{Expiration: expTime}}

	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.keyValues[key] = item

	return nil
}

func (s *Storage) GetStr(key string) (interface{}, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	if item, ok := s.keyValues[key]; ok {
		if str, ok := item.(Str); ok {
			if isExpired(str.Expiration) {
				return nil, nil
			}

			return str.Value, nil
		}
		return nil, errors.New(errWrongType)
	}

	return nil, nil
}

func isExpired(expiration int64) bool {
	if expiration > 0 {
		if time.Now().Unix() > expiration {
			return true
		}
	}

	return false
}
