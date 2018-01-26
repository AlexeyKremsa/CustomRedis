package storage

import (
	"reflect"
	"testing"
)

func Test_SetList(t *testing.T) {
	strg := Init(0, 1)

	key := "key1"
	valueToSet := []string{"str1"}

	strg.SetList(key, valueToSet, 0)

	if v, ok := strg.shards[0].keyValues[key]; ok {
		if !reflect.DeepEqual(valueToSet, v.Value) {
			t.Fatalf("Expected: %v, actual: %v", valueToSet, v.Value)
		}
	} else {
		t.Fatalf("Expected to find value `%v` by key `%s`", valueToSet, key)
	}
}

func Test_SetList_GetErrWrongType(t *testing.T) {
	strg := Init(0, 1)

	key := "key1"
	valueToSet := 2
	strg.shards[0].keyValues[key] = Item{Value: valueToSet}

	val, err := strg.GetList(key)
	if err == nil {
		t.Fatalf("Expected error: `%s` but got nil", errWrongType)
	}

	if len(val) != 0 {
		t.Fatalf("Expected `val` to be empty, but got: %v", val)
	}

	if _, ok := err.(ErrBusiness); !ok {
		t.Fatal("Unexpected error type")
	}
}

func Test_GetList(t *testing.T) {
	strg := Init(0, 1)

	key := "key1"
	valueToSet := []string{"str1", "str2"}
	strg.shards[0].keyValues[key] = Item{Value: valueToSet}

	res, err := strg.GetList(key)
	if err != nil {
		t.Fatalf("Unexpected error: %s", err.Error())
	}

	if !reflect.DeepEqual(valueToSet, res) {
		t.Fatalf("Expected:`%v`, actual: `%v`", valueToSet, res)
	}
}

func Test_ListPush(t *testing.T) {
	strg := Init(0, 1)

	key := "key1"
	valueToAdd := []string{"str2"}
	initialArr := []string{"str1"}
	expected := []string{"str1", "str2"}

	strg.shards[0].keyValues[key] = Item{Value: initialArr}

	err := strg.ListPush(key, valueToAdd)
	if err != nil {
		t.Fatalf("Unexpected error: %s", err.Error())
	}

	if v, ok := strg.shards[0].keyValues[key]; ok {
		if !reflect.DeepEqual(expected, v.Value) {
			t.Fatalf("Expected: %v, actual: %v", expected, v.Value)
		}
	} else {
		t.Fatalf("Expected to find value `%v` by key `%s`", expected, key)
	}
}

func Test_ListPush_GetErrWrongType(t *testing.T) {
	strg := Init(0, 1)

	key := "key1"
	valueToAdd := []string{"str2"}
	initialArr := 2

	strg.shards[0].keyValues[key] = Item{Value: initialArr}

	err := strg.ListPush(key, valueToAdd)
	if err == nil {
		t.Fatalf("Expected error: `%s` but got nil", errWrongType)
	}

	if _, ok := err.(ErrBusiness); !ok {
		t.Fatal("Unexpected error type")
	}

	if err.Error() != errWrongType {
		t.Fatalf("Expected error: `%s`, actual: `%s`", errWrongType, err.Error())
	}
}

func Test_ListPush_GetErrNotExist(t *testing.T) {
	strg := Init(0, 1)

	valueToAdd := []string{"str2"}

	err := strg.ListPush("newKey", valueToAdd)
	if err == nil {
		t.Fatalf("Expected error: `%s` but got nil", errNotExist)
	}

	if _, ok := err.(ErrBusiness); !ok {
		t.Fatalf("Unexpected error type")
	}

	if err.Error() != errNotExist {
		t.Fatalf("Expected error: `%s`, actual: `%s`", errNotExist, err.Error())
	}
}

func Test_ListPop(t *testing.T) {
	strg := Init(0, 1)

	key := "key1"
	valueToGet := "str2"
	initialArr := []string{"str1", valueToGet}
	expectedArr := []string{"str1"}

	strg.shards[0].keyValues[key] = Item{Value: initialArr}

	res, err := strg.ListPop(key)
	if err != nil {
		t.Fatalf("Unexpected error: %s", err.Error())
	}

	if res != valueToGet {
		t.Fatalf("Expecetd: %s, actual: %s", valueToGet, res)
	}

	modifiedArr := strg.shards[0].keyValues[key].Value
	if !reflect.DeepEqual(modifiedArr, expectedArr) {
		t.Fatalf("Expected: %v, actual: %v", modifiedArr, expectedArr)
	}
}

func Test_ListPop_GetErrWrongType(t *testing.T) {
	strg := Init(0, 1)

	key := "key1"
	initialArr := 2

	strg.shards[0].keyValues[key] = Item{Value: initialArr}

	val, err := strg.ListPop(key)
	if err == nil {
		t.Fatalf("Expected error: `%s` but got nil", errWrongType)
	}

	if val != "" {
		t.Fatalf("Expected to have an empty value, but got: %s", val)
	}

	if _, ok := err.(ErrBusiness); !ok {
		t.Fatal("Unexpected error type")
	}

	if err.Error() != errWrongType {
		t.Fatalf("Expected error: `%s`, actual: `%s`", errWrongType, err.Error())
	}
}

func Test_ListPop_GetErrNotExist(t *testing.T) {
	strg := Init(0, 1)

	key := "key1"
	initialArr := []string{"str1"}

	strg.shards[0].keyValues[key] = Item{Value: initialArr}

	val, err := strg.ListPop("newKey")
	if err == nil {
		t.Fatalf("Expected error: `%s` but got nil", errWrongType)
	}

	if val != "" {
		t.Fatalf("Expected to have an empty value, but got: %s", val)
	}

	if _, ok := err.(ErrBusiness); !ok {
		t.Fatalf("Unexpected error type")
	}

	if err.Error() != errNotExist {
		t.Fatalf("Expected error: `%s`, actual: `%s`", errNotExist, err.Error())
	}
}

func Test_ListIndex(t *testing.T) {
	strg := Init(0, 1)

	key := "key1"
	expected := "str2"
	valueToSet := []string{"str1", expected}
	strg.shards[0].keyValues[key] = Item{Value: valueToSet}

	res, err := strg.ListIndex(key, 1)
	if err != nil {
		t.Fatalf("Unexpected error: %s", err.Error())
	}

	if res != expected {
		t.Fatalf("Expected: %s, actual: %s", expected, res)
	}
}

func Test_ListIndex_GetErrOutOfRange(t *testing.T) {
	strg := Init(0, 1)

	key := "key1"
	valueToSet := []string{"str1", "str2"}
	strg.shards[0].keyValues[key] = Item{Value: valueToSet}

	val, err := strg.ListIndex(key, 15)
	if err == nil {
		t.Fatalf("Expected error: `%s` but got nil", errIndexOutOfRange)
	}

	if val != "" {
		t.Fatalf("Expected to have an empty value, but got: %s", val)
	}

	if _, ok := err.(ErrBusiness); !ok {
		t.Fatalf("Unexpected error type")
	}

	if err.Error() != errIndexOutOfRange {
		t.Fatalf("Expected error: `%s`, actual: `%s`", errNotExist, err.Error())
	}
}

func Test_ListIndex_GetErrWrongType(t *testing.T) {
	strg := Init(0, 1)

	key := "key1"
	initialArr := 2

	strg.shards[0].keyValues[key] = Item{Value: initialArr}

	val, err := strg.ListIndex(key, 0)
	if err == nil {
		t.Fatalf("Expected error: `%s` but got nil", errWrongType)
	}

	if val != "" {
		t.Fatalf("Expected to have an empty value, but got: %s", val)
	}

	if _, ok := err.(ErrBusiness); !ok {
		t.Fatal("Unexpected error type")
	}

	if err.Error() != errWrongType {
		t.Fatalf("Expected error: `%s`, actual: `%s`", errWrongType, err.Error())
	}
}

func Test_ListIndex_GetErrNotExist(t *testing.T) {
	strg := Init(0, 1)

	key := "key1"
	initialArr := []string{"str1"}

	strg.shards[0].keyValues[key] = Item{Value: initialArr}

	val, err := strg.ListIndex("newKey", 2)
	if err == nil {
		t.Fatalf("Expected error: `%s` but got nil", errWrongType)
	}

	if val != "" {
		t.Fatalf("Expected to have an empty value, but got: %s", val)
	}

	if _, ok := err.(ErrBusiness); !ok {
		t.Fatalf("Unexpected error type")
	}

	if err.Error() != errNotExist {
		t.Fatalf("Expected error: `%s`, actual: `%s`", errNotExist, err.Error())
	}
}
