package storage

import (
	"math/rand"
	"sync"
	"time"

	"github.com/labstack/gommon/log"
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

	if cleanupTimeoutSec > 0 {
		go strg.runCleanup(cleanupTimeoutSec)
	} else {
		log.Warn("cleanupTimeoutSec is set to 0! No cleanup will be performed!")
	}

	return strg
}

func (s *Storage) cleanup() {
	seed := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(seed)
	shardIndex := rnd.Intn(int(s.shardCountDecremented))
	shard := s.shards[shardIndex]
	log.Debugf("Cleanup started")

	shard.mutex.Lock()
	defer shard.mutex.Unlock()

	for key, val := range shard.keyValues {
		if isExpired(val.Expiration) {
			delete(shard.keyValues, key)
		}
	}
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
		return newErrCustom(errKeyExists)
	}

	var expTime int64
	if expirationSec > 0 {
		expTime = time.Now().Add(time.Second * time.Duration(expirationSec)).Unix()
	}
	item := Item{Value: value, Expiration: expTime}

	shard.keyValues[key] = item

	return nil
}

func (s *Storage) get(key string) interface{} {
	shard := s.getShard(key)

	shard.mutex.RLock()
	defer shard.mutex.RUnlock()

	if item, ok := shard.keyValues[key]; ok {
		if isExpired(item.Expiration) {
			return nil
		}
		return item.Value
	}
	return nil
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
