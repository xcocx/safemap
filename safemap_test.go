package safemap

import "testing"

const key string = "key"

var m = New()

func BenchmarkSafeMap_Set(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m.Set(key, "this is content")
	}
}

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

func BenchmarkSafeMap_GetInt64(b *testing.B) {
	m.Set(key, int64(100))
	for i := 0; i < b.N; i++ {
		if m.GetInt64(key) != int64(100) {
			b.Fatalf("GetInt64() is error")
		}
	}
}

func BenchmarkSafeMap_GetFloat64(b *testing.B) {
	m.Set(key, 100.01)
	for i := 0; i < b.N; i++ {
		if m.GetFloat64(key) != 100.01 {
			b.Fatalf("GetFloat64() is error")
		}
	}
}
