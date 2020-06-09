package sort

import (
	"math/rand"
	"sort"
	"testing"
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
	return compareInt(s[i], s[j]) < 0
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
			Heap(items, compareInt)
		}
	})

	b.Run("Insertion", func(b *testing.B) {
		rand.Seed(seed)
		items := randIntSlice(size, minInt, maxInt)

		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			Shuffle(items)
			Insertion(items, compareInt)
		}
	})

	b.Run("Merge", func(b *testing.B) {
		rand.Seed(seed)
		items := randIntSlice(size, minInt, maxInt)

		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			Shuffle(items)
			Merge(items, compareInt)
		}
	})

	b.Run("MergeRec", func(b *testing.B) {
		rand.Seed(seed)
		items := randIntSlice(size, minInt, maxInt)

		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			Shuffle(items)
			MergeRec(items, compareInt)
		}
	})

	b.Run("Quick", func(b *testing.B) {
		rand.Seed(seed)
		items := randIntSlice(size, minInt, maxInt)

		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			Shuffle(items)
			Quick(items, compareInt)
		}
	})

	b.Run("Quick3Way", func(b *testing.B) {
		rand.Seed(seed)
		items := randIntSlice(size, minInt, maxInt)

		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			Shuffle(items)
			Quick3Way(items, compareInt)
		}
	})

	b.Run("Shell", func(b *testing.B) {
		rand.Seed(seed)
		items := randIntSlice(size, minInt, maxInt)

		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			Shuffle(items)
			Shell(items, compareInt)
		}
	})
}
