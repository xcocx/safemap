package safemap

import "testing"

const key string = "key"

var m = New()

func BenchmarkSafeMap_GetBool(b *testing.B) {
	m.Set(key, true)
	for i := 0; i < b.N; i++ {
		if m.GetBool(key) != true {
			b.Fatalf("GetBool() is error")
		}
	}
}

func BenchmarkSafeMap_GetString(b *testing.B) {
	m.Set(key, "this is string")
	for i := 0; i < b.N; i++ {
		if m.GetString(key) != "this is string" {
			b.Fatalf("GetString() is error")
		}
	}
}
