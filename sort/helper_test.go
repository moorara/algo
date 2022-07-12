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

func randIntSlice(size int) []int {
	a := make([]int, size)
	for i := range a {
		a[i] = rand.Int()
	}

	return a
}
