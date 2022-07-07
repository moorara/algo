package radixsort

import "golang.org/x/exp/constraints"

func isSorted[T constraints.Ordered](a []T) bool {
	for i := 0; i < len(a)-1; i++ {
		if a[i] > a[i+1] {
			return false
		}
	}
	return true
}
