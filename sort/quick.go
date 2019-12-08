package sort

func partition(a []interface{}, lo, hi int, compare CompareFunc) int {
	v := a[lo]
	var i, j int = lo, hi + 1

	for {
		for i++; i < hi && compare(a[i], v) < 0; i++ {
		}
		for j--; j > lo && compare(a[j], v) > 0; j-- {
		}
		if i >= j {
			break
		}
		a[i], a[j] = a[j], a[i]
	}
	a[lo], a[j] = a[j], a[lo]

	return j
}

// Select finds the kth smallest item of an array in O(n) time on average.
func Select(a []interface{}, k int, compare CompareFunc) interface{} {
	Shuffle(a)
	var lo, hi int = 0, len(a) - 1
	for lo < hi {
		j := partition(a, lo, hi, compare)
		switch {
		case j < k:
			lo = j + 1
		case j > k:
			hi = j - 1
		default:
			return a[k]
		}
	}

	return a[k]
}

func quickSort(a []interface{}, lo, hi int, compare CompareFunc) {
	if lo >= hi {
		return
	}

	j := partition(a, lo, hi, compare)
	quickSort(a, lo, j-1, compare)
	quickSort(a, j+1, hi, compare)
}

// QuickSort implements the quick sort algorithm.
func QuickSort(a []interface{}, compare CompareFunc) {
	Shuffle(a)
	quickSort(a, 0, len(a)-1, compare)
}

func quickSort3Way(a []interface{}, lo, hi int, compare CompareFunc) {
	if lo >= hi {
		return
	}

	v := a[lo]
	var lt, i, gt int = lo, lo + 1, hi

	for i <= gt {
		c := compare(a[i], v)
		switch {
		case c < 0:
			a[lt], a[i] = a[i], a[lt]
			lt++
			i++
		case c > 0:
			a[i], a[gt] = a[gt], a[i]
			gt--
		default:
			i++
		}
	}

	quickSort3Way(a, lo, lt-1, compare)
	quickSort3Way(a, gt+1, hi, compare)
}

// QuickSort3Way implements the 3-way version of quick sort algorithm.
func QuickSort3Way(a []interface{}, compare CompareFunc) {
	quickSort3Way(a, 0, len(a)-1, compare)
}
