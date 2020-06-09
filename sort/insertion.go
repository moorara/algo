package sort

import "github.com/moorara/algo/compare"

// Insertion implements the insertion sort algorithm.
func Insertion(a []interface{}, cmp compare.Func) {
	n := len(a)
	for i := 0; i < n; i++ {
		for j := i; j > 0 && cmp(a[j], a[j-1]) < 0; j-- {
			a[j], a[j-1] = a[j-1], a[j]
		}
	}
}
