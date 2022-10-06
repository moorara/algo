package heap

import (
	"math/rand"
	"testing"
	"time"

	"github.com/moorara/algo/generic"
)

const (
	minLen = 10
	maxLen = 100
	chars  = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func randIntSlice(size int) []int {
	vals := make([]int, size)
	for i := 0; i < len(vals); i++ {
		vals[i] = rand.Int()
	}

	return vals
}

func randStringKey(minLen, maxLen int) string {
	n := len(chars)
	l := minLen + rand.Intn(maxLen-minLen+1)
	b := make([]byte, l)

	for i := range b {
		b[i] = chars[rand.Intn(n)]
	}

	return string(b)
}

func randStringSlice(size int) []string {
	s := make([]string, size)
	for i := range s {
		s[i] = randStringKey(minLen, maxLen)
	}

	return s
}

func runHeapInsert(b *testing.B, heap Heap[int, string]) {
	keys := randIntSlice(b.N)
	vals := randStringSlice(b.N)

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		heap.Insert(keys[n], vals[n])
	}
}

func runHeapDelete(b *testing.B, heap Heap[int, string]) {
	keys := randIntSlice(b.N)
	vals := randStringSlice(b.N)

	for n := 0; n < b.N; n++ {
		heap.Insert(keys[n], vals[n])
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		heap.Delete()
	}
}

func runIndexedHeapInsert(b *testing.B, heap IndexedHeap[int, string]) {
	keys := randIntSlice(b.N)
	vals := randStringSlice(b.N)

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		heap.Insert(n, keys[n], vals[n])
	}
}

func runIndexedHeapDelete(b *testing.B, heap IndexedHeap[int, string]) {
	keys := randIntSlice(b.N)
	vals := randStringSlice(b.N)

	for n := 0; n < b.N; n++ {
		heap.Insert(n, keys[n], vals[n])
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		heap.Delete()
	}
}

func BenchmarkHeap_Insert(b *testing.B) {
	rand.Seed(time.Now().UTC().UnixNano())

	const size = 1024
	eqVal := generic.NewEqualFunc[string]()
	cmpMin := generic.NewCompareFunc[int]()
	cmpMax := generic.NewInvertedCompareFunc[int]()

	b.Run("BinaryMinHeap.Insert", func(b *testing.B) {
		heap := NewBinary(size, cmpMin, eqVal)
		runHeapInsert(b, heap)
	})

	b.Run("BinaryMaxHeap.Insert", func(b *testing.B) {
		heap := NewBinary(size, cmpMax, eqVal)
		runHeapInsert(b, heap)
	})
}

func BenchmarkHeap_Delete(b *testing.B) {
	rand.Seed(time.Now().UTC().UnixNano())

	const size = 1024
	eqVal := generic.NewEqualFunc[string]()
	cmpMin := generic.NewCompareFunc[int]()
	cmpMax := generic.NewInvertedCompareFunc[int]()

	b.Run("BinaryMinHeap.Delete", func(b *testing.B) {
		heap := NewBinary(size, cmpMin, eqVal)
		runHeapDelete(b, heap)
	})

	b.Run("BinaryMaxHeap.Delete", func(b *testing.B) {
		heap := NewBinary(size, cmpMax, eqVal)
		runHeapDelete(b, heap)
	})
}

func BenchmarkIndexedHeap_Insert(b *testing.B) {
	rand.Seed(time.Now().UTC().UnixNano())

	eqVal := generic.NewEqualFunc[string]()
	cmpMin := generic.NewCompareFunc[int]()
	cmpMax := generic.NewInvertedCompareFunc[int]()

	b.Run("BinaryMinIndexedHeap.Insert", func(b *testing.B) {
		heap := NewIndexedBinary[int, string](b.N, cmpMin, eqVal)
		runIndexedHeapInsert(b, heap)
	})

	b.Run("BinaryMaxIndexedHeap.Insert", func(b *testing.B) {
		heap := NewIndexedBinary[int, string](b.N, cmpMax, eqVal)
		runIndexedHeapInsert(b, heap)
	})
}

func BenchmarkIndexedHeap_Delete(b *testing.B) {
	rand.Seed(time.Now().UTC().UnixNano())

	eqVal := generic.NewEqualFunc[string]()
	cmpMin := generic.NewCompareFunc[int]()
	cmpMax := generic.NewInvertedCompareFunc[int]()

	b.Run("BinaryMinIndexedHeap.Delete", func(b *testing.B) {
		heap := NewIndexedBinary[int, string](b.N, cmpMin, eqVal)
		runIndexedHeapDelete(b, heap)
	})

	b.Run("BinaryMaxIndexedHeap.Delete", func(b *testing.B) {
		heap := NewIndexedBinary[int, string](b.N, cmpMax, eqVal)
		runIndexedHeapDelete(b, heap)
	})
}
