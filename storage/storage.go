package storage

import (
	"math/rand"
	"sync"
	"time"

	"github.com/labstack/gommon/log"
)

type item struct {
	value      interface{}
	expiration uint64
}

// Storage describes storage settings and fields
type Storage struct {
	// pre-calculated value which is used to determine a shard
	shardCountDecremented uint64
	shards                []*shard
}

type shard struct {
	mutex sync.RWMutex
	// Also sync.Map could be considered for systems with heavy read operations and big amount of CPU cores
	keyValues map[string]item
}

// Init creates Storage object
func Init(cleanupTimeoutSec, shardCount uint64) *Storage {
	strg := &Storage{shardCountDecremented: shardCount - 1}
	strg.shards = make([]*shard, shardCount)
	for i := 0; i < int(shardCount); i++ {
		strg.shards[i] = &shard{keyValues: make(map[string]item)}
	}

	if cleanupTimeoutSec > 0 {
		go strg.runCleanup(cleanupTimeoutSec)
	} else {
		log.Warn("cleanupTimeoutSec is set to 0! No cleanup will be performed!")
	}

	return strg
}

func isExpired(expiration uint64) bool {
	if expiration > 0 && time.Now().Unix() > int64(expiration) {
		return true
	}
	return false
}

func (s *Storage) cleanup() {
	var shardIndex int
	if s.shardCountDecremented == 0 {
		shardIndex = 0
	} else {
		seed := rand.NewSource(time.Now().UnixNano())
		rnd := rand.New(seed)
		shardIndex = rnd.Intn(int(s.shardCountDecremented))
	}

	shard := s.shards[shardIndex]
	log.Debugf("Cleanup started")

	shard.mutex.Lock()
	defer shard.mutex.Unlock()

	for key, val := range shard.keyValues {
		if isExpired(val.expiration) {
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

func (s *Storage) set(key string, value interface{}, expirationSec uint64) {
	shard := s.getShard(key)

	var expTime int64
	if expirationSec > 0 {
		expTime = time.Now().Add(time.Second * time.Duration(expirationSec)).Unix()
	}
	itm := item{value: value, expiration: uint64(expTime)}

	shard.mutex.Lock()
	defer shard.mutex.Unlock()
	shard.keyValues[key] = itm
}

func (s *Storage) setNX(key string, value interface{}, expirationSec uint64) error {
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
	itm := item{value: value, expiration: uint64(expTime)}

	shard.keyValues[key] = itm

	return nil
}

func (s *Storage) get(key string) *item {
	shard := s.getShard(key)

	shard.mutex.RLock()
	defer shard.mutex.RUnlock()

	if itm, ok := shard.keyValues[key]; ok {
		if isExpired(itm.expiration) {
			return nil
		}
		return &itm
	}
	return nil
}

func (s *Storage) RemoveItem(key string) {
	shard := s.getShard(key)

	shard.mutex.Lock()
	defer shard.mutex.Unlock()
	delete(shard.keyValues, key)
}

// GetAllKeys returns all keys
func (s *Storage) GetAllKeys() []string {
	var wg sync.WaitGroup
	resCh := make(chan []string, len(s.shards))
	allKeys := make([]string, 0)

	for _, v := range s.shards {
		wg.Add(1)
		go collectKeys(v, resCh, &wg)
	}

	wg.Wait()
	close(resCh)

	for keys := range resCh {
		allKeys = append(allKeys, keys...)
	}

	return allKeys
}

func collectKeys(shard *shard, resCh chan []string, wg *sync.WaitGroup) {
	res := make([]string, 0)

	shard.mutex.RLock()
	defer shard.mutex.RUnlock()

	for key, val := range shard.keyValues {
		if !isExpired(val.expiration) {
			res = append(res, key)
		}
	}

	resCh <- res
	wg.Done()
}
