package sort

import "github.com/moorara/algo/compare"

// Selection implements the selection sort algorithm.
func Selection(a []interface{}, cmp compare.Func) {
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
