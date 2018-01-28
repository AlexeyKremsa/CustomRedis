package storage

import (
	"fmt"
	"testing"
	"time"
)

func Benchmark_SetStr_1shard(b *testing.B) {
	s1 := Init(11111, 1)
	for i := 0; i < b.N; i++ {
		now := time.Now().UnixNano()
		go s1.SetStr(fmt.Sprintf("key-%d", now), fmt.Sprintf("val-%d", now), 1111)
	}
}

func Benchmark_SetStr_25shards(b *testing.B) {
	s25 := Init(11111, 25)
	for i := 0; i < b.N; i++ {
		now := time.Now().UnixNano()
		go s25.SetStr(fmt.Sprintf("key-%d", now), fmt.Sprintf("val-%d", now), 1111)
	}
}

func Benchmark_GetAllKeys_1shard(b *testing.B) {
	s1 := Init(11111, 1)

	allKeysRes := make([]string, 0)
	for i := 0; i < 10000; i++ {
		now := time.Now().UnixNano()
		go s1.SetStr(fmt.Sprintf("key-%d", now), fmt.Sprintf("val-%d", now), 1111)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		allKeysRes = append(allKeysRes, s1.GetAllKeys()...)
	}
}

func Benchmark_GetAllKeys_25shards(b *testing.B) {
	s25 := Init(11111, 25)

	allKeysRes := make([]string, 0)
	for i := 0; i < 10000; i++ {
		now := time.Now().UnixNano()
		go s25.SetStr(fmt.Sprintf("key-%d", now), fmt.Sprintf("val-%d", now), 1111)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		allKeysRes = append(allKeysRes, s25.GetAllKeys()...)
	}
}
