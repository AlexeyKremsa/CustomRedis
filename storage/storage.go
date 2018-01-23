package storage

import (
	"sync"
	"time"
)

type Item struct {
	Value      interface{}
	Expiration int64
}

type Storage struct {
	mutex sync.RWMutex
	// Also sync.Map could be considered for systems with heavy read operations and big amount of CPU cores
	keyValues map[string]Item
}

func Init() *Storage {
	strg := &Storage{}
	strg.keyValues = make(map[string]Item)
	return strg
}

func (s *Storage) Set(key string, value interface{}, expirationSec int64) {
	var expTime int64
	if expirationSec > 0 {
		expTime = time.Now().Add(time.Second * time.Duration(expirationSec)).Unix()
	}
	item := Item{Value: value, Expiration: expTime}

	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.keyValues[key] = item
}

func (s *Storage) SetNX(key string, value interface{}, expirationSec int64) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if _, ok := s.keyValues[key]; ok {
		return newErrCustom("Key already exists")
	}

	var expTime int64
	if expirationSec > 0 {
		expTime = time.Now().Add(time.Second * time.Duration(expirationSec)).Unix()
	}
	item := Item{Value: value, Expiration: expTime}

	s.keyValues[key] = item

	return nil
}

func (s *Storage) Get(key string) (interface{}, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	if item, ok := s.keyValues[key]; ok {
		if isExpired(item.Expiration) {
			return nil, nil
		}
		return item.Value, nil
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
