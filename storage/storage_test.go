package storage

import (
	"os"
	"testing"
	"time"
)

var strg *Storage

func TestMain(m *testing.M) {
	strg = Init(0, 1)
	exitCode := m.Run()
	os.Exit(exitCode)
}

func Test_RemoveItem(t *testing.T) {
	key := "key1"
	valueToSet := "str1"
	strg.shards[0].keyValues[key] = Item{Value: valueToSet}

	strg.RemoveItem(key)

	if v, ok := strg.shards[0].keyValues[key]; ok {
		t.Fatalf("Value should have been removed, but got: %s", v.Value)
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
