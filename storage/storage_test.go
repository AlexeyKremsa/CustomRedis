package storage

import (
	"os"
	"testing"
	"time"
)

var strg *Storage

func TestMain(m *testing.M) {
	strg = Init(11111, 1)
	exitCode := m.Run()
	os.Exit(exitCode)
}

func Test_RemoveItem(t *testing.T) {
	key := "key1"
	valueToSet := "str1"
	strg.shards[0].keyValues[key] = item{value: valueToSet}

	strg.RemoveItem(key)

	if v, ok := strg.shards[0].keyValues[key]; ok {
		t.Fatalf("Value should have been removed, but got: %s", v.value)
	}
}

func Test_isExpired(t *testing.T) {
	testTime := time.Now().AddDate(0, 0, -1).Unix()

	res := isExpired(uint64(testTime))

	if !res {
		t.Fatal("Expected time to be expired")
	}
}

func Test_cleanup(t *testing.T) {
	strg.SetStr("k1", "v1", 1)
	strg.SetStr("k2", "v2", 1)
	strg.SetStr("k3", "v3", 0)
	strg.SetStr("k4", "v4", 0)

	time.Sleep(2 * time.Second)
	strg.cleanup()

	if len(strg.shards[0].keyValues) != 2 {
		t.Fatalf("Expected to delete 2 elements from 4")
	}
}

func Test_GetAllKeys_1shard(t *testing.T) {
	s1 := Init(1111, 1)

	s1.shards[0].keyValues["k1"] = item{value: "t1"}
	s1.shards[0].keyValues["k2"] = item{value: "t1"}
	s1.shards[0].keyValues["k3"] = item{value: "t1"}

	res := s1.GetAllKeys()

	// keys are unordered!
	for _, val := range res {
		if val == "k1" || val == "k2" || val == "k3" {
			continue
		}
		t.Fatalf("Unexpected key: %s", val)
	}
}

func Test_GetAllKeys_2shards(t *testing.T) {
	s2 := Init(1111, 2)

	s2.shards[0].keyValues["k1"] = item{value: "t1"}
	s2.shards[0].keyValues["k2"] = item{value: "t1"}
	s2.shards[1].keyValues["k3"] = item{value: "t1"}

	res := s2.GetAllKeys()

	// keys are unordered!
	for _, val := range res {
		if val == "k1" || val == "k2" || val == "k3" {
			continue
		}
		t.Fatalf("Unexpected key: %s", val)
	}
}
