package sort

import (
	"math/rand"
	"sort"
	"testing"
	"time"

	. "github.com/moorara/algo/generic"
)

const size = 100000

func BenchmarkSort(b *testing.B) {
	seed := time.Now().UTC().UnixNano()
	r := rand.New(rand.NewSource(seed))

	nums := randIntSlice(size)
	cmp := NewCompareFunc[int]()

	b.Run("sort.Sort", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			a := make([]int, len(nums))
			copy(a, nums)
			Shuffle[int](a, r)
			sort.Sort(sort.IntSlice(a)) // nolint directives: SA1032
		}
	})

	b.Run("Heap", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			a := make([]int, len(nums))
			copy(a, nums)
			Shuffle[int](a, r)
			Heap[int](a, cmp)
		}
	})

	b.Run("Insertion", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			a := make([]int, len(nums))
			copy(a, nums)
			Shuffle[int](a, r)
			Insertion[int](a, cmp)
		}
	})

	b.Run("Merge", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			a := make([]int, len(nums))
			copy(a, nums)
			Shuffle[int](a, r)
			Merge[int](a, cmp)
		}
	})

	b.Run("MergeRec", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			a := make([]int, len(nums))
			copy(a, nums)
			Shuffle[int](a, r)
			MergeRec[int](a, cmp)
		}
	})

	b.Run("Quick", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			a := make([]int, len(nums))
			copy(a, nums)
			Shuffle[int](a, r)
			Quick[int](a, cmp)
		}
	})

	b.Run("Quick3Way", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			a := make([]int, len(nums))
			copy(a, nums)
			Shuffle[int](a, r)
			Quick3Way[int](a, cmp)
		}
	})

	b.Run("Shell", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			a := make([]int, len(nums))
			copy(a, nums)
			Shuffle[int](a, r)
			Shell[int](a, cmp)
		}
	})
}
