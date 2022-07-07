package sort

import "github.com/moorara/algo/common"

// Insertion implements the insertion sort algorithm.
func Insertion[T any](a []T, cmp common.CompareFunc[T]) {
	n := len(a)
	for i := 0; i < n; i++ {
		for j := i; j > 0 && cmp(a[j], a[j-1]) < 0; j-- {
			a[j], a[j-1] = a[j-1], a[j]
		}
	}
}
