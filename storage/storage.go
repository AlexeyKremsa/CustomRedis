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
	// pre-calculated value which is used to determine a shard
	shardCountDecremented uint64
	shards                []*shard
}

type shard struct {
	mutex sync.RWMutex
	// Also sync.Map could be considered for systems with heavy read operations and big amount of CPU cores
	keyValues map[string]Item
}

func Init(cleanupTimeoutSec, shardCount uint64) *Storage {
	strg := &Storage{shardCountDecremented: shardCount - 1}
	strg.shards = make([]*shard, shardCount)
	for i := 0; i < int(shardCount); i++ {
		strg.shards[i] = &shard{keyValues: make(map[string]Item)}
	}

	go strg.runCleanup(cleanupTimeoutSec)

	return strg
}

func (s *Storage) cleanup() {
	//TODO: pick up random shard to clean up

	// shard :=
	// 	log.Debugf("Cleanup started. Total items before cleanup: %d", len(s.keyValues))
	// s.getShard(key).mutex.Lock()
	// defer s.getShard(key).mutex.Unlock()

	// for key, val := range s.keyValues {
	// 	if isExpired(val.Expiration) {
	// 		delete(s.keyValues, key)
	// 	}
	// }
}

func (s *Storage) runCleanup(timeoutSec uint64) {
	ticker := time.NewTicker(time.Second * time.Duration(timeoutSec))

	for {
		select {
		case <-ticker.C:
			s.cleanup()
		}
	}
}

func (s *Storage) set(key string, value interface{}, expirationSec int64) {
	shard := s.getShard(key)

	var expTime int64
	if expirationSec > 0 {
		expTime = time.Now().Add(time.Second * time.Duration(expirationSec)).Unix()
	}
	item := Item{Value: value, Expiration: expTime}

	shard.mutex.Lock()
	defer shard.mutex.Unlock()
	shard.keyValues[key] = item
}

func (s *Storage) setNX(key string, value interface{}, expirationSec int64) error {
	shard := s.getShard(key)

	shard.mutex.Lock()
	defer shard.mutex.Unlock()

	if _, ok := shard.keyValues[key]; ok {
		return newErrCustom("Key already exists")
	}

	var expTime int64
	if expirationSec > 0 {
		expTime = time.Now().Add(time.Second * time.Duration(expirationSec)).Unix()
	}
	item := Item{Value: value, Expiration: expTime}

	shard.keyValues[key] = item

	return nil
}

func (s *Storage) get(key string) (interface{}, error) {
	shard := s.getShard(key)

	shard.mutex.RLock()
	defer shard.mutex.RUnlock()

	if item, ok := shard.keyValues[key]; ok {
		if isExpired(item.Expiration) {
			return nil, nil
		}
		return item.Value, nil
	}
	return nil, nil
}

func (s *Storage) RemoveItem(key string) {
	shard := s.getShard(key)

	shard.mutex.Lock()
	defer shard.mutex.Unlock()
	delete(shard.keyValues, key)
}

func isExpired(expiration int64) bool {
	if expiration > 0 && time.Now().Unix() > expiration {
		return true
	}
	return false
}
