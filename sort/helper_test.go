package sort

import (
	"math/rand"

	"github.com/moorara/algo/common"
)

func isSorted[T any](items []T, cmp common.CompareFunc[T]) bool {
	for i := 0; i < len(items)-1; i++ {
		if cmp(items[i], items[i+1]) > 0 {
			return false
		}
	}

	return true
}

func randIntSlice(size, min, max int) []int {
	items := make([]int, size)
	for i := 0; i < len(items); i++ {
		items[i] = min + rand.Intn(max-min+1)
	}

	return items
}
