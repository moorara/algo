package radixsort

func quick3WayString(a []string, lo, hi, d int) {
	// cutoff to insertion sort for small subarrays
	if hi <= lo+cutoff {
		insertionString(a, lo, hi, d)
		return
	}

	// 3-way partitioning on dth char
	v := charAt(a[lo], d)
	lt, i, gt := lo, lo+1, hi
	for i <= gt {
		c := charAt(a[i], d)
		switch {
		case c < v:
			a[lt], a[i] = a[i], a[lt]
			lt++
			i++
		case c > v:
			a[i], a[gt] = a[gt], a[i]
			gt--
		default:
			i++
		}

		// a[lo..lt-1] < v = a[lt..gt] < a[gt+1..hi]
		quick3WayString(a, lo, lt-1, d)
		if v >= 0 {
			quick3WayString(a, lt, gt, d+1)
		}
		quick3WayString(a, gt+1, hi, d)
	}
}

// Quick3WayString is the 3-Way Radix Quick sorting algorithm for string keys with variable length.
func Quick3WayString(a []string) {
	shuffleString(a)
	quick3WayString(a, 0, len(a)-1, 0)
}
