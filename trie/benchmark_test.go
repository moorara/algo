package trie

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

func BenchmarkTrie_Put(b *testing.B) {
	rand.Seed(time.Now().UTC().UnixNano())

	eqVal := generic.NewEqualFunc[int]()

	b.Run("BinaryTrie.Put", func(b *testing.B) {
		trie := NewBinary[int](eqVal)
		runPutBenchmark(b, trie)
	})

	b.Run("Patricia.Put", func(b *testing.B) {
		trie := NewPatricia[int](eqVal)
		runPutBenchmark(b, trie)
	})
}

func BenchmarkTrie_Get(b *testing.B) {
	rand.Seed(time.Now().UTC().UnixNano())

	eqVal := generic.NewEqualFunc[int]()

	b.Run("BinaryTrie.Get", func(b *testing.B) {
		trie := NewBinary[int](eqVal)
		runGetBenchmark(b, trie)
	})

	b.Run("Patricia.Get", func(b *testing.B) {
		trie := NewPatricia[int](eqVal)
		runGetBenchmark(b, trie)
	})
}

func BenchmarkTrie_Delete(b *testing.B) {
	rand.Seed(time.Now().UTC().UnixNano())

	eqVal := generic.NewEqualFunc[int]()

	b.Run("BinaryTrie.Delete", func(b *testing.B) {
		trie := NewBinary[int](eqVal)
		runDeleteBenchmark(b, trie)
	})

	b.Run("Patricia.Delete", func(b *testing.B) {
		trie := NewPatricia[int](eqVal)
		runDeleteBenchmark(b, trie)
	})
}
