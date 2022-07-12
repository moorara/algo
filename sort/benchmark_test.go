package sort

import (
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/moorara/algo/common"
)

const size = 100000

func BenchmarkSort(b *testing.B) {
	rand.Seed(time.Now().UTC().UnixNano())

	nums := randIntSlice(size)
	cmp := common.NewCompareFunc[int]()

	b.Run("sort.Sort", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			a := make([]int, len(nums))
			copy(a, nums)
			Shuffle[int](a)
			sort.Sort(sort.IntSlice(a))
		}
	})

	b.Run("Heap", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			a := make([]int, len(nums))
			copy(a, nums)
			Shuffle[int](a)
			Heap[int](a, cmp)
		}
	})

	b.Run("Insertion", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			a := make([]int, len(nums))
			copy(a, nums)
			Shuffle[int](a)
			Insertion[int](a, cmp)
		}
	})

	b.Run("Merge", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			a := make([]int, len(nums))
			copy(a, nums)
			Shuffle[int](a)
			Merge[int](a, cmp)
		}
	})

	b.Run("MergeRec", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			a := make([]int, len(nums))
			copy(a, nums)
			Shuffle[int](a)
			MergeRec[int](a, cmp)
		}
	})

	b.Run("Quick", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			a := make([]int, len(nums))
			copy(a, nums)
			Shuffle[int](a)
			Quick[int](a, cmp)
		}
	})

	b.Run("Quick3Way", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			a := make([]int, len(nums))
			copy(a, nums)
			Shuffle[int](a)
			Quick3Way[int](a, cmp)
		}
	})

	b.Run("Shell", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			a := make([]int, len(nums))
			copy(a, nums)
			Shuffle[int](a)
			Shell[int](a, cmp)
		}
	})
}
