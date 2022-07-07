package sort

import "github.com/moorara/algo/common"

// Selection implements the selection sort algorithm.
func Selection[T any](a []T, cmp common.CompareFunc[T]) {
	n := len(a)
	for i := 0; i < n; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if cmp(a[j], a[min]) < 0 {
				min = j
			}
		}
		a[i], a[min] = a[min], a[i]
	}
}
