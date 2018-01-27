package storage

import "testing"

func Test_RemoveItem(t *testing.T) {
	strg := Init(0, 1)

	key := "key1"
	valueToSet := "str1"
	strg.shards[0].keyValues[key] = Item{Value: valueToSet}

	strg.RemoveItem(key)

	if v, ok := strg.shards[0].keyValues[key]; ok {
		t.Fatalf("Value should have been removed, but got: %s", v.Value)
	}
}

// func Test_isExpired(t *testing.T) {
// 	strg := Init(0, 1)

// 	key := "key1"
// 	valueToSet := "str1"
// 	expiration := 1

// 	strg.SetStr(key, valueToSet, expiration)
// }
