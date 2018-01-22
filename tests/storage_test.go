package tests

import (
	"fmt"
	"testing"

	"github.com/AlexeyKremsa/CustomRedis/storage"
)

func BenchmarkSetStrAndGetStr(b *testing.B) {
	st := storage.Init()

	for i := 0; i < b.N; i++ {
		st.SetStr(fmt.Sprintf("key-%d", i), fmt.Sprintf("val-%d", i), 1)
	}

	for i := 0; i < b.N; i++ {
		st.GetStr(fmt.Sprintf("key-%d", i))
	}
}
