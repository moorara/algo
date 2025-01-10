package heap

import (
	"math/rand"
	"testing"
	"time"
)

const (
	size   = 1024
	minLen = 10
	maxLen = 100
	chars  = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

var r *rand.Rand

func randIntSlice(size int) []int {
	vals := make([]int, size)
	for i := 0; i < len(vals); i++ {
		vals[i] = r.Int()
	}

	return vals
}

func randStringKey(minLen, maxLen int) string {
	n := len(chars)
	l := minLen + r.Intn(maxLen-minLen+1)
	b := make([]byte, l)

	for i := range b {
		b[i] = chars[r.Intn(n)]
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

func runIndexedHeapChangeKey(b *testing.B, heap IndexedHeap[int, string]) {
	keys := randIntSlice(b.N)
	vals := randStringSlice(b.N)
	newKeys := randIntSlice(b.N)

	for n := 0; n < b.N; n++ {
		heap.Insert(n, keys[n], vals[n])
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		heap.ChangeKey(n, newKeys[n])
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

func runMergeableHeapMerge(b *testing.B, h1, h2 MergeableHeap[int, string]) {
	keys := randIntSlice(b.N)
	vals := randStringSlice(b.N)

	for n := 0; n < b.N; n++ {
		h1.Insert(keys[n], vals[n])
	}

	keys = randIntSlice(b.N)
	vals = randStringSlice(b.N)

	for n := 0; n < b.N; n++ {
		h2.Insert(keys[n], vals[n])
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		h1.Merge(h2)
	}
}

func BenchmarkHeap_Insert(b *testing.B) {
	seed := time.Now().UTC().UnixNano()
	r = rand.New(rand.NewSource(seed))

	b.Run("BinaryMinHeap.Insert", func(b *testing.B) {
		heap := NewBinary(size, cmpMin, eqVal)
		runHeapInsert(b, heap)
	})

	b.Run("BinaryMaxHeap.Insert", func(b *testing.B) {
		heap := NewBinary(size, cmpMax, eqVal)
		runHeapInsert(b, heap)
	})

	b.Run("BinomialMinHeap.Insert", func(b *testing.B) {
		heap := NewBinomial(cmpMin, eqVal)
		runHeapInsert(b, heap)
	})

	b.Run("BinomialMaxHeap.Insert", func(b *testing.B) {
		heap := NewBinomial(cmpMax, eqVal)
		runHeapInsert(b, heap)
	})

	b.Run("FibonacciMinHeap.Insert", func(b *testing.B) {
		heap := NewFibonacci(cmpMin, eqVal)
		runHeapInsert(b, heap)
	})

	b.Run("FibonacciMaxHeap.Insert", func(b *testing.B) {
		heap := NewFibonacci(cmpMax, eqVal)
		runHeapInsert(b, heap)
	})
}

func BenchmarkHeap_Delete(b *testing.B) {
	seed := time.Now().UTC().UnixNano()
	r = rand.New(rand.NewSource(seed))

	b.Run("BinaryMinHeap.Delete", func(b *testing.B) {
		heap := NewBinary(size, cmpMin, eqVal)
		runHeapDelete(b, heap)
	})

	b.Run("BinaryMaxHeap.Delete", func(b *testing.B) {
		heap := NewBinary(size, cmpMax, eqVal)
		runHeapDelete(b, heap)
	})

	b.Run("BinomialMinHeap.Delete", func(b *testing.B) {
		heap := NewBinomial(cmpMin, eqVal)
		runHeapDelete(b, heap)
	})

	b.Run("BinomialMaxHeap.Delete", func(b *testing.B) {
		heap := NewBinomial(cmpMax, eqVal)
		runHeapDelete(b, heap)
	})

	b.Run("FibonacciMinHeap.Delete", func(b *testing.B) {
		heap := NewFibonacci(cmpMin, eqVal)
		runHeapDelete(b, heap)
	})

	b.Run("FibonacciMaxHeap.Delete", func(b *testing.B) {
		heap := NewFibonacci(cmpMax, eqVal)
		runHeapDelete(b, heap)
	})
}

func BenchmarkIndexedHeap_Insert(b *testing.B) {
	seed := time.Now().UTC().UnixNano()
	r = rand.New(rand.NewSource(seed))

	b.Run("BinaryMinIndexedHeap.Insert", func(b *testing.B) {
		heap := NewIndexedBinary(b.N, cmpMin, eqVal)
		runIndexedHeapInsert(b, heap)
	})

	b.Run("BinaryMaxIndexedHeap.Insert", func(b *testing.B) {
		heap := NewIndexedBinary(b.N, cmpMax, eqVal)
		runIndexedHeapInsert(b, heap)
	})

	b.Run("BinomialMinIndexedHeap.Insert", func(b *testing.B) {
		heap := NewIndexedBinomial(b.N, cmpMin, eqVal)
		runIndexedHeapInsert(b, heap)
	})

	b.Run("BinomialMaxIndexedHeap.Insert", func(b *testing.B) {
		heap := NewIndexedBinomial(b.N, cmpMax, eqVal)
		runIndexedHeapInsert(b, heap)
	})

	b.Run("FibonacciMinIndexedHeap.Insert", func(b *testing.B) {
		heap := NewIndexedFibonacci(b.N, cmpMin, eqVal)
		runIndexedHeapInsert(b, heap)
	})

	b.Run("FibonacciMaxIndexedHeap.Insert", func(b *testing.B) {
		heap := NewIndexedFibonacci(b.N, cmpMax, eqVal)
		runIndexedHeapInsert(b, heap)
	})
}

func BenchmarkIndexedHeap_ChangeKey(b *testing.B) {
	seed := time.Now().UTC().UnixNano()
	r = rand.New(rand.NewSource(seed))

	b.Run("BinaryMinIndexedHeap.ChangeKey", func(b *testing.B) {
		heap := NewIndexedBinary(b.N, cmpMin, eqVal)
		runIndexedHeapChangeKey(b, heap)
	})

	b.Run("BinaryMaxIndexedHeap.ChangeKey", func(b *testing.B) {
		heap := NewIndexedBinary(b.N, cmpMax, eqVal)
		runIndexedHeapChangeKey(b, heap)
	})

	b.Run("BinomialMinIndexedHeap.ChangeKey", func(b *testing.B) {
		heap := NewIndexedBinomial(b.N, cmpMin, eqVal)
		runIndexedHeapChangeKey(b, heap)
	})

	b.Run("BinomialMaxIndexedHeap.ChangeKey", func(b *testing.B) {
		heap := NewIndexedBinomial(b.N, cmpMax, eqVal)
		runIndexedHeapChangeKey(b, heap)
	})

	b.Run("FibonacciMinIndexedHeap.ChangeKey", func(b *testing.B) {
		heap := NewIndexedFibonacci(b.N, cmpMin, eqVal)
		runIndexedHeapChangeKey(b, heap)
	})

	b.Run("FibonacciMaxIndexedHeap.ChangeKey", func(b *testing.B) {
		heap := NewIndexedFibonacci(b.N, cmpMax, eqVal)
		runIndexedHeapChangeKey(b, heap)
	})
}

func BenchmarkIndexedHeap_Delete(b *testing.B) {
	seed := time.Now().UTC().UnixNano()
	r = rand.New(rand.NewSource(seed))

	b.Run("BinaryMinIndexedHeap.Delete", func(b *testing.B) {
		heap := NewIndexedBinary(b.N, cmpMin, eqVal)
		runIndexedHeapDelete(b, heap)
	})

	b.Run("BinaryMaxIndexedHeap.Delete", func(b *testing.B) {
		heap := NewIndexedBinary(b.N, cmpMax, eqVal)
		runIndexedHeapDelete(b, heap)
	})

	b.Run("BinomialMinIndexedHeap.Delete", func(b *testing.B) {
		heap := NewIndexedBinomial(b.N, cmpMin, eqVal)
		runIndexedHeapDelete(b, heap)
	})

	b.Run("BinomialMaxIndexedHeap.Delete", func(b *testing.B) {
		heap := NewIndexedBinomial(b.N, cmpMax, eqVal)
		runIndexedHeapDelete(b, heap)
	})

	b.Run("FibonacciMinIndexedHeap.Delete", func(b *testing.B) {
		heap := NewIndexedFibonacci(b.N, cmpMin, eqVal)
		runIndexedHeapDelete(b, heap)
	})

	b.Run("FibonacciMaxIndexedHeap.Delete", func(b *testing.B) {
		heap := NewIndexedFibonacci(b.N, cmpMax, eqVal)
		runIndexedHeapDelete(b, heap)
	})
}

func BenchmarkHeap_Merge(b *testing.B) {
	seed := time.Now().UTC().UnixNano()
	r = rand.New(rand.NewSource(seed))

	b.Run("BinomialMinHeap.Merge", func(b *testing.B) {
		h1 := NewBinomial(cmpMin, eqVal)
		h2 := NewBinomial(cmpMin, eqVal)
		runMergeableHeapMerge(b, h1, h2)
	})

	b.Run("BinomialMaxHeap.Merge", func(b *testing.B) {
		h1 := NewBinomial(cmpMax, eqVal)
		h2 := NewBinomial(cmpMin, eqVal)
		runMergeableHeapMerge(b, h1, h2)
	})

	b.Run("FibonacciMinHeap.Merge", func(b *testing.B) {
		h1 := NewFibonacci(cmpMin, eqVal)
		h2 := NewFibonacci(cmpMin, eqVal)
		runMergeableHeapMerge(b, h1, h2)
	})

	b.Run("FibonacciMaxHeap.Merge", func(b *testing.B) {
		h1 := NewFibonacci(cmpMax, eqVal)
		h2 := NewFibonacci(cmpMin, eqVal)
		runMergeableHeapMerge(b, h1, h2)
	})
}
