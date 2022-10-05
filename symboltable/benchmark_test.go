package symboltable

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

func runPutBenchmark(b *testing.B, ost OrderedSymbolTable[string, int]) {
	keys := randStringSlice(b.N)
	vals := randIntSlice(b.N)

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		ost.Put(keys[n], vals[n])
	}
}

func runGetBenchmark(b *testing.B, ost OrderedSymbolTable[string, int]) {
	keys := randStringSlice(b.N)
	vals := randIntSlice(b.N)

	for n := 0; n < b.N; n++ {
		ost.Put(keys[n], vals[n])
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		ost.Get(keys[n])
	}
}

func runDeleteBenchmark(b *testing.B, ost OrderedSymbolTable[string, int]) {
	keys := randStringSlice(b.N)
	vals := randIntSlice(b.N)

	for n := 0; n < b.N; n++ {
		ost.Put(keys[n], vals[n])
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		ost.Delete(keys[n])
	}
}

func BenchmarkOrderedSymbolTable_Put(b *testing.B) {
	rand.Seed(time.Now().UTC().UnixNano())

	cmpKey := generic.NewCompareFunc[string]()

	b.Run("BST.Put", func(b *testing.B) {
		ost := NewBST[string, int](cmpKey)
		runPutBenchmark(b, ost)
	})

	b.Run("AVL.Put", func(b *testing.B) {
		ost := NewAVL[string, int](cmpKey)
		runPutBenchmark(b, ost)
	})

	b.Run("RedBlack.Put", func(b *testing.B) {
		ost := NewRedBlack[string, int](cmpKey)
		runPutBenchmark(b, ost)
	})
}

func BenchmarkOrderedSymbolTable_Get(b *testing.B) {
	rand.Seed(time.Now().UTC().UnixNano())

	cmpKey := generic.NewCompareFunc[string]()

	b.Run("BST.Get", func(b *testing.B) {
		ost := NewBST[string, int](cmpKey)
		runGetBenchmark(b, ost)
	})

	b.Run("AVL.Get", func(b *testing.B) {
		ost := NewAVL[string, int](cmpKey)
		runGetBenchmark(b, ost)
	})

	b.Run("RedBlack.Get", func(b *testing.B) {
		ost := NewRedBlack[string, int](cmpKey)
		runGetBenchmark(b, ost)
	})
}

func BenchmarkOrderedSymbolTable_Delete(b *testing.B) {
	rand.Seed(time.Now().UTC().UnixNano())

	cmpKey := generic.NewCompareFunc[string]()

	b.Run("BST.Delete", func(b *testing.B) {
		ost := NewBST[string, int](cmpKey)
		runDeleteBenchmark(b, ost)
	})

	b.Run("AVL.Delete", func(b *testing.B) {
		ost := NewAVL[string, int](cmpKey)
		runDeleteBenchmark(b, ost)
	})

	b.Run("RedBlack.Delete", func(b *testing.B) {
		ost := NewRedBlack[string, int](cmpKey)
		runDeleteBenchmark(b, ost)
	})
}
