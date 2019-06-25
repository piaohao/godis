package godis

import "testing"

func BenchmarkSet(b *testing.B) {
	b.ResetTimer()
	pool := NewPool(nil, option)
	for i := 0; i < b.N; i++ {
		redis, _ := pool.Get()
		redis.Set("godis", "good")
		redis.Close()
	}
}

func BenchmarkGet(b *testing.B) {
	b.ResetTimer()
	pool := NewPool(nil, option)
	for i := 0; i < b.N; i++ {
		redis, _ := pool.Get()
		redis.Get("godis")
		redis.Close()
	}
}

func BenchmarkIncr(b *testing.B) {
	flushAll()
	b.ResetTimer()
	pool := NewPool(nil, option)
	i := 0
	for ; i < b.N; i++ {
		redis, _ := pool.Get()
		redis.Incr("godis")
		redis.Close()
	}
}
