package set

import (
	"testing"

	"github.com/moorara/algo/generic"
	"github.com/moorara/algo/hash"
)

const (
	minLen = 10
	maxLen = 100
	chars  = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func randString(minLen, maxLen int) string {
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
		s[i] = randString(minLen, maxLen)
	}

	return s
}

func runAddBenchmark(b *testing.B, set Set[string]) {
	vals := randStringSlice(b.N)

	b.ResetTimer()

	set.Add(vals...)
}

func runRemoveBenchmark(b *testing.B, set Set[string]) {
	vals := randStringSlice(b.N)
	set.Add(vals...)

	b.ResetTimer()

	set.Remove(vals...)
}

func runContainsBenchmark(b *testing.B, set Set[string]) {
	vals := randStringSlice(b.N)
	set.Add(vals...)

	b.ResetTimer()

	set.Contains(vals...)
}

func BenchmarkSet_Add(b *testing.B) {
	hashString := hash.HashFuncForString[string](nil)
	eqString := generic.NewEqualFunc[string]()
	cmpString := generic.NewCompareFunc[string]()
	opts := HashSetOpts{}

	b.Run("Set.Add", func(b *testing.B) {
		ht := New(eqString)
		runAddBenchmark(b, ht)
	})

	b.Run("StableSet.Add", func(b *testing.B) {
		ht := NewStableSet(eqString)
		runAddBenchmark(b, ht)
	})

	b.Run("SortedSet.Add", func(b *testing.B) {
		ht := NewSortedSet(cmpString)
		runAddBenchmark(b, ht)
	})

	b.Run("HashSet.Add", func(b *testing.B) {
		ht := NewHashSet(hashString, eqString, opts)
		runAddBenchmark(b, ht)
	})
}

func BenchmarkSet_Remove(b *testing.B) {
	hashString := hash.HashFuncForString[string](nil)
	eqString := generic.NewEqualFunc[string]()
	cmpString := generic.NewCompareFunc[string]()
	opts := HashSetOpts{}

	b.Run("Set.Remove", func(b *testing.B) {
		ht := New(eqString)
		runRemoveBenchmark(b, ht)
	})

	b.Run("StableSet.Remove", func(b *testing.B) {
		ht := NewStableSet(eqString)
		runRemoveBenchmark(b, ht)
	})

	b.Run("SortedSet.Remove", func(b *testing.B) {
		ht := NewSortedSet(cmpString)
		runRemoveBenchmark(b, ht)
	})

	b.Run("HashSet.Remove", func(b *testing.B) {
		ht := NewHashSet(hashString, eqString, opts)
		runRemoveBenchmark(b, ht)
	})
}

func BenchmarkSet_Contains(b *testing.B) {
	hashString := hash.HashFuncForString[string](nil)
	eqString := generic.NewEqualFunc[string]()
	cmpString := generic.NewCompareFunc[string]()
	opts := HashSetOpts{}

	b.Run("Set.Contains", func(b *testing.B) {
		ht := New(eqString)
		runContainsBenchmark(b, ht)
	})

	b.Run("StableSet.Contains", func(b *testing.B) {
		ht := NewStableSet(eqString)
		runContainsBenchmark(b, ht)
	})

	b.Run("SortedSet.Contains", func(b *testing.B) {
		ht := NewSortedSet(cmpString)
		runContainsBenchmark(b, ht)
	})

	b.Run("HashSet.Contains", func(b *testing.B) {
		ht := NewHashSet(hashString, eqString, opts)
		runContainsBenchmark(b, ht)
	})
}
