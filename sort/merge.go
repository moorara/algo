package sort

import "github.com/moorara/algo/compare"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func merge(a, aux []interface{}, lo, mid, hi int, cmp compare.Func) {
	var i, j int = lo, mid + 1
	copy(aux[lo:hi+1], a[lo:hi+1])
	for k := lo; k <= hi; k++ {
		switch {
		case i > mid:
			a[k] = aux[j]
			j++
		case j > hi:
			a[k] = aux[i]
			i++
		case cmp(aux[j], aux[i]) < 0:
			a[k] = aux[j]
			j++
		default:
			a[k] = aux[i]
			i++
		}
	}
}

// Merge implements the iterative version of merge sort algorithm.
func Merge(a []interface{}, cmp compare.Func) {
	n := len(a)
	aux := make([]interface{}, n)
	for sz := 1; sz < n; sz += sz {
		for lo := 0; lo < n-sz; lo += sz + sz {
			merge(a, aux, lo, lo+sz-1, min(lo+sz+sz-1, n-1), cmp)
		}
	}
}

func mergeRec(a, aux []interface{}, lo, hi int, cmp compare.Func) {
	if hi <= lo {
		return
	}

	mid := (lo + hi) / 2
	mergeRec(a, aux, lo, mid, cmp)
	mergeRec(a, aux, mid+1, hi, cmp)
	if cmp(a[mid+1], a[mid]) >= 0 {
		return
	}
	merge(a, aux, lo, mid, hi, cmp)
}

// MergeRec implements the recursive version of merge sort algorithm.
func MergeRec(a []interface{}, cmp compare.Func) {
	n := len(a)
	aux := make([]interface{}, n)

	mergeRec(a, aux, 0, n-1, cmp)
}
