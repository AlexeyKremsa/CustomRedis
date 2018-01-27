package storage

import (
	"reflect"
	"testing"
)

func Test_SetMap(t *testing.T) {
	strg := Init(0, 1)

	key := "key1"
	valueToSet := make(map[string]string)
	valueToSet["k"] = "v"

	strg.SetMap(key, valueToSet, 0)

	if v, ok := strg.shards[0].keyValues[key]; ok {
		if !reflect.DeepEqual(valueToSet, v.Value) {
			t.Fatalf("Expected: %v, actual: %v", valueToSet, v.Value)
		}
	} else {
		t.Fatalf("Expected to find value `%v` by key `%s`", valueToSet, key)
	}
}

func Test_GetMap(t *testing.T) {
	strg := Init(0, 1)

	key := "key1"
	valueToSet := make(map[string]string)
	valueToSet["k"] = "v"
	strg.shards[0].keyValues[key] = Item{Value: valueToSet}

	res, err := strg.GetMap(key)
	if err != nil {
		t.Fatalf("Unexpected error: %s", err.Error())
	}

	if !reflect.DeepEqual(valueToSet, res) {
		t.Fatalf("Expected:`%v`, actual: `%v`", valueToSet, res)
	}
}

func Test_GetMap_MapNotExists_ReturnNil(t *testing.T) {
	strg := Init(0, 1)

	res, err := strg.GetMap("any")
	if err != nil {
		t.Fatalf("Unexpected error: %s", err.Error())
	}

	if res != nil {
		t.Fatalf("Expected result to nil, but got: %v", res)
	}
}

func Test_GetMap_GetErrWrongType(t *testing.T) {
	strg := Init(0, 1)

	key := "key1"
	valueToSet := 2
	strg.shards[0].keyValues[key] = Item{Value: valueToSet}

	res, err := strg.GetMap(key)
	if err == nil {
		t.Fatalf("Expected error: `%s` but got nil", errWrongType)
	}

	if res != nil {
		t.Fatalf("Expected result to be nil, but got: %v", res)
	}

	if _, ok := err.(ErrBusiness); !ok {
		t.Fatal("Unexpected error type")
	}

	if err.Error() != errWrongType {
		t.Fatalf("Expected error: `%s`, actual: `%s`", errWrongType, err.Error())
	}
}

func Test_GetMapItem(t *testing.T) {
	strg := Init(0, 1)

	key := "key1"
	valueToSet := make(map[string]string)
	itemKey := "k"
	itemValue := "v"
	valueToSet[itemKey] = itemValue
	strg.shards[0].keyValues[key] = Item{Value: valueToSet}

	res, err := strg.GetMapItem(key, itemKey)
	if err != nil {
		t.Fatalf("Unexpected error: %s", err.Error())
	}

	if itemValue != res {
		t.Fatalf("Expected: %s, actual: %s", itemValue, res)
	}
}

func Test_GetMapItem_MapNotExists_ReturnNil(t *testing.T) {
	strg := Init(0, 1)

	res, err := strg.GetMapItem("any", "any2")
	if err != nil {
		t.Fatalf("Unexpected error: %s", err.Error())
	}

	if res != "" {
		t.Fatalf("Expected result to empty, but got: %v", res)
	}
}

func Test_GetMapItem_GetErrWrongType(t *testing.T) {
	strg := Init(0, 1)

	key := "key1"
	valueToSet := 2
	strg.shards[0].keyValues[key] = Item{Value: valueToSet}

	res, err := strg.GetMapItem(key, "any")
	if err == nil {
		t.Fatalf("Expected error: `%s` but got nil", errWrongType)
	}

	if res != "" {
		t.Fatalf("Expected result to be empty, but got: %v", res)
	}

	if _, ok := err.(ErrBusiness); !ok {
		t.Fatal("Unexpected error type")
	}

	if err.Error() != errWrongType {
		t.Fatalf("Expected error: `%s`, actual: `%s`", errWrongType, err.Error())
	}
}

func Test_GetMapItem_KeyNotExists_ReturnEmptyString(t *testing.T) {
	strg := Init(0, 1)

	key := "key1"
	valueToSet := make(map[string]string)
	itemKey := "k"
	itemValue := "v"
	valueToSet[itemKey] = itemValue
	strg.shards[0].keyValues[key] = Item{Value: valueToSet}

	res, err := strg.GetMapItem(key, "any2")
	if err != nil {
		t.Fatalf("Unexpected error: `%s`", err.Error())
	}

	if res != "" {
		t.Fatalf("Expected result to be empty, but got: %v", res)
	}
}
