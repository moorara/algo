package sort

func sink(a []interface{}, k, n int, cmp CompareFunc) {
	for 2*k <= n {
		j := 2 * k
		if j < n && cmp(a[j], a[j+1]) < 0 {
			j++
		}
		if cmp(a[k], a[j]) >= 0 {
			break
		}
		a[k], a[j] = a[j], a[k]
		k = j
	}
}

func heapSort(a []interface{}, cmp CompareFunc) {
	n := len(a) - 1

	// build max-heap bottom-up
	for k := n / 2; k >= 1; k-- {
		sink(a, k, n, cmp)
	}

	// remove the maximum, one at a time
	for n > 1 {
		a[1], a[n] = a[n], a[1]
		n--
		sink(a, 1, n, cmp)
	}
}

// HeapSort implements the heap sort algorithm.
func HeapSort(a []interface{}, cmp CompareFunc) {
	// Heap elements need to start from position 1
	aux := append([]interface{}{nil}, a...)
	heapSort(aux, cmp)
	copy(a, aux[1:])
}
