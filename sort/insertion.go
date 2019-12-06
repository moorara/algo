package sort

// InsertionSort implements the insertion sort algorithm.
func InsertionSort(a []interface{}, cmp CompareFunc) {
	n := len(a)
	for i := 0; i < n; i++ {
		for j := i; j > 0 && cmp(a[j], a[j-1]) < 0; j-- {
			a[j], a[j-1] = a[j-1], a[j]
		}
	}
}
