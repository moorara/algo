package trie

import (
	"math/rand"
	"testing"
	"time"
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

func runPutBenchmark(b *testing.B, trie Trie[int]) {
	keys := randStringSlice(b.N)
	vals := randIntSlice(b.N)

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		trie.Put(keys[n], vals[n])
	}
}

func runGetBenchmark(b *testing.B, trie Trie[int]) {
	keys := randStringSlice(b.N)
	vals := randIntSlice(b.N)

	for n := 0; n < b.N; n++ {
		trie.Put(keys[n], vals[n])
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		trie.Get(keys[n])
	}
}

func runDeleteBenchmark(b *testing.B, trie Trie[int]) {
	keys := randStringSlice(b.N)
	vals := randIntSlice(b.N)

	for n := 0; n < b.N; n++ {
		trie.Put(keys[n], vals[n])
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		trie.Delete(keys[n])
	}
}

func BenchmarkOrderedSymbolTable_Put(b *testing.B) {
	rand.Seed(time.Now().UTC().UnixNano())

	b.Run("BinaryTrie.Put", func(b *testing.B) {
		trie := NewBinary[int]()
		runPutBenchmark(b, trie)
	})

	// TODO:

	b.Run("Patricia.Put", func(b *testing.B) {
		trie := NewPatricia[int]()
		runPutBenchmark(b, trie)
	})
}

func BenchmarkOrderedSymbolTable_Get(b *testing.B) {
	rand.Seed(time.Now().UTC().UnixNano())

	b.Run("BinaryTrie.Get", func(b *testing.B) {
		trie := NewBinary[int]()
		runGetBenchmark(b, trie)
	})

	// TODO:

	b.Run("Patricia.Get", func(b *testing.B) {
		trie := NewPatricia[int]()
		runGetBenchmark(b, trie)
	})
}

func BenchmarkOrderedSymbolTable_Delete(b *testing.B) {
	rand.Seed(time.Now().UTC().UnixNano())

	b.Run("BinaryTrie.Delete", func(b *testing.B) {
		trie := NewBinary[int]()
		runDeleteBenchmark(b, trie)
	})

	// TODO:

	b.Run("Patricia.Delete", func(b *testing.B) {
		trie := NewPatricia[int]()
		runDeleteBenchmark(b, trie)
	})
}
