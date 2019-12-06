package sort

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func merge(a, aux []interface{}, lo, mid, hi int, cmp CompareFunc) {
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

// MergeSort implements the iterative version of merge sort algorithm.
func MergeSort(a []interface{}, cmp CompareFunc) {
	n := len(a)
	aux := make([]interface{}, n)
	for sz := 1; sz < n; sz += sz {
		for lo := 0; lo < n-sz; lo += sz + sz {
			merge(a, aux, lo, lo+sz-1, min(lo+sz+sz-1, n-1), cmp)
		}
	}
}

func mergeSortRec(a, aux []interface{}, lo, hi int, cmp CompareFunc) {
	if hi <= lo {
		return
	}

	mid := lo + (hi-lo)/2
	mergeSortRec(a, aux, lo, mid, cmp)
	mergeSortRec(a, aux, mid+1, hi, cmp)
	if cmp(a[mid+1], a[mid]) >= 0 {
		return
	}
	merge(a, aux, lo, mid, hi, cmp)
}

// MergeSortRec implements the recursive version of merge sort algorithm.
func MergeSortRec(a []interface{}, cmp CompareFunc) {
	n := len(a)
	aux := make([]interface{}, n)

	mergeSortRec(a, aux, 0, n-1, cmp)
}
