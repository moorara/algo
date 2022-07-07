package sort

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/moorara/algo/common"
)

const (
	seed   = 27
	size   = 1000
	minInt = 0
	maxInt = 1000000
)

func BenchmarkSort(b *testing.B) {
	cmp := common.NewCompareFunc[int]()

	b.Run("sort.Sort", func(b *testing.B) {
		rand.Seed(seed)
		items := sort.IntSlice(randIntSlice(size, minInt, maxInt))

		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			Shuffle[int](items)
			sort.Sort(items)
		}
	})

	b.Run("Heap", func(b *testing.B) {
		rand.Seed(seed)
		items := randIntSlice(size, minInt, maxInt)

		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			Shuffle[int](items)
			Heap[int](items, cmp)
		}
	})

	b.Run("Insertion", func(b *testing.B) {
		rand.Seed(seed)
		items := randIntSlice(size, minInt, maxInt)

		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			Shuffle[int](items)
			Insertion[int](items, cmp)
		}
	})

	b.Run("Merge", func(b *testing.B) {
		rand.Seed(seed)
		items := randIntSlice(size, minInt, maxInt)

		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			Shuffle[int](items)
			Merge[int](items, cmp)
		}
	})

	b.Run("MergeRec", func(b *testing.B) {
		rand.Seed(seed)
		items := randIntSlice(size, minInt, maxInt)

		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			Shuffle[int](items)
			MergeRec[int](items, cmp)
		}
	})

	b.Run("Quick", func(b *testing.B) {
		rand.Seed(seed)
		items := randIntSlice(size, minInt, maxInt)

		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			Shuffle[int](items)
			Quick[int](items, cmp)
		}
	})

	b.Run("Quick3Way", func(b *testing.B) {
		rand.Seed(seed)
		items := randIntSlice(size, minInt, maxInt)

		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			Shuffle[int](items)
			Quick3Way[int](items, cmp)
		}
	})

	b.Run("Shell", func(b *testing.B) {
		rand.Seed(seed)
		items := randIntSlice(size, minInt, maxInt)

		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			Shuffle[int](items)
			Shell[int](items, cmp)
		}
	})
}
