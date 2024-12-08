package sort

import . "github.com/moorara/algo/generic"

func sink[T any](a []T, k, n int, cmp CompareFunc[T]) {
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

func heap[T any](a []T, cmp CompareFunc[T]) {
	n := len(a) - 1

	// build max-heap bottom-up
	for k := n / 2; k >= 1; k-- {
		sink[T](a, k, n, cmp)
	}

	// remove the maximum, one at a time
	for n > 1 {
		a[1], a[n] = a[n], a[1]
		n--
		sink[T](a, 1, n, cmp)
	}
}

// Heap implements the heap sort algorithm.
func Heap[T any](a []T, cmp CompareFunc[T]) {
	// Heap elements need to start from position 1
	var zero T
	aux := append([]T{zero}, a...)
	heap[T](aux, cmp)
	copy(a, aux[1:])
}
