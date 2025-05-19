package concurrency

import (
	"sync"
	"sync/atomic"
	"testing"
)

func BenchmarkAtomicSimple(b *testing.B) {
	var num int32
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		atomic.AddInt32(&num, 1)
	}
}

func BenchmarkMutexSimple(b *testing.B) {
	var num int32
	var mu sync.Mutex
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		mu.Lock()
		num++
		mu.Unlock()
	}
}

// Concurrent versions
func BenchmarkAtomicConcurrent(b *testing.B) {
	var num int32
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			atomic.AddInt32(&num, 1)
		}
	})
}

func BenchmarkMutexConcurrent(b *testing.B) {
	var num int32
	var mu sync.Mutex
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			mu.Lock()
			num++
			mu.Unlock()
		}
	})
}
