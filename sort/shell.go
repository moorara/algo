package sort

import "github.com/moorara/algo/common"

// Shell implements the shell sort algorithm.
func Shell[T any](a []T, cmp common.CompareFunc[T]) {
	n := len(a)
	h := 1
	for h < n/3 {
		h = 3*h + 1 // 1, 4, 13, 40, 121, 364, ...
	}

	for ; h >= 1; h /= 3 {
		for i := h; i < n; i++ { // h-sort the array
			for j := i; j >= h && cmp(a[j], a[j-h]) < 0; j -= h {
				a[j], a[j-h] = a[j-h], a[j]
			}
		}
	}
}
