package sort

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/moorara/algo/compare"
)

const (
	seed   = 27
	size   = 1000
	minInt = 0
	maxInt = 1000000
)

// sliceInterface implements sort.Interface
type sliceInterface []interface{}

func (s sliceInterface) Len() int {
	return len(s)
}

func (s sliceInterface) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sliceInterface) Less(i, j int) bool {
	return compare.Int(s[i], s[j]) < 0
}

func BenchmarkSort(b *testing.B) {
	b.Run("sort.Sort", func(b *testing.B) {
		rand.Seed(seed)
		items := sliceInterface(randIntSlice(size, minInt, maxInt))

		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			Shuffle(items)
			sort.Sort(items)
		}
	})

	b.Run("Heap", func(b *testing.B) {
		rand.Seed(seed)
		items := randIntSlice(size, minInt, maxInt)

		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			Shuffle(items)
			Heap(items, compare.Int)
		}
	})

	b.Run("Insertion", func(b *testing.B) {
		rand.Seed(seed)
		items := randIntSlice(size, minInt, maxInt)

		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			Shuffle(items)
			Insertion(items, compare.Int)
		}
	})

	b.Run("Merge", func(b *testing.B) {
		rand.Seed(seed)
		items := randIntSlice(size, minInt, maxInt)

		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			Shuffle(items)
			Merge(items, compare.Int)
		}
	})

	b.Run("MergeRec", func(b *testing.B) {
		rand.Seed(seed)
		items := randIntSlice(size, minInt, maxInt)

		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			Shuffle(items)
			MergeRec(items, compare.Int)
		}
	})

	b.Run("Quick", func(b *testing.B) {
		rand.Seed(seed)
		items := randIntSlice(size, minInt, maxInt)

		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			Shuffle(items)
			Quick(items, compare.Int)
		}
	})

	b.Run("Quick3Way", func(b *testing.B) {
		rand.Seed(seed)
		items := randIntSlice(size, minInt, maxInt)

		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			Shuffle(items)
			Quick3Way(items, compare.Int)
		}
	})

	b.Run("Shell", func(b *testing.B) {
		rand.Seed(seed)
		items := randIntSlice(size, minInt, maxInt)

		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			Shuffle(items)
			Shell(items, compare.Int)
		}
	})
}
