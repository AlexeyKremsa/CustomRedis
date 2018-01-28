package storage

import (
	"testing"
)

func Test_SetStr(t *testing.T) {
	key := "key1"
	valueToSet := "str1"

	strg.SetStr(key, valueToSet, 0)

	if v, ok := strg.shards[0].keyValues[key]; ok {
		if valueToSet != v.value {
			t.Fatalf("Expected: %s, actual: %s", valueToSet, v.value)
		}
	} else {
		t.Fatalf("Expected to find value `%s` by key `%s`", valueToSet, key)
	}
}

func Test_SetStrNX(t *testing.T) {
	key := "keyNX"
	valueToSet := "str1"

	strg.SetStrNX(key, valueToSet, 0)

	if v, ok := strg.shards[0].keyValues[key]; ok {
		if valueToSet != v.value {
			t.Fatalf("Expected: %s, actual: %s", valueToSet, v.value)
		}
	} else {
		t.Fatalf("Expected to find value `%s` by key `%s`", valueToSet, key)
	}
}

func Test_SetStrNX_GetKeyExistsError(t *testing.T) {
	key := "key1"
	valueToSet := "str1"
	strg.shards[0].keyValues[key] = item{value: valueToSet}

	err := strg.SetStrNX(key, valueToSet, 0)
	if errCustom, ok := err.(ErrBusiness); ok {
		if errCustom.Error() != errKeyExists {
			t.Fatalf("Expected error: `%s`, actual: `%s`", errKeyExists, errCustom.Error())
		}
	} else {
		t.Fatal("Unexpected error type. Expected ErrBusiness")
	}
}

func Test_GetStr(t *testing.T) {
	key := "key1"
	valueToSet := "str1"
	strg.shards[0].keyValues[key] = item{value: valueToSet}

	val, err := strg.GetStr(key)
	if err != nil {
		t.Fatalf("Unexpected error: %s", err.Error())
	}

	if val != valueToSet {
		t.Fatalf("Expected: %s, actual: %s", valueToSet, val)
	}
}

func Test_GetStr_NotExists_ReturnEmptyString(t *testing.T) {
	val, err := strg.GetStr("any")
	if err != nil {
		t.Fatalf("Unexpected error: %s", err.Error())
	}

	if val != "" {
		t.Fatalf("Result expected to be empty, but got: %s", val)
	}
}

func Test_GetStr_GetErrWronType(t *testing.T) {
	key := "key1"
	strg.shards[0].keyValues[key] = item{value: []string{"q", "w"}}

	val, err := strg.GetStr(key)
	if err == nil {
		t.Fatal("Expected errWronType but got nil")
	}

	if val != "" {
		t.Fatalf("Expected `val` to be empty, but got: %s", val)
	}

	if _, ok := err.(ErrBusiness); !ok {
		t.Fatal("Unexpected error type")
	}
}
