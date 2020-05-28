package radixsort

// LSDString is the LSD (least significant digit) sorting algorithm for string keys with fixed length w.
func LSDString(a []string, w int) {
	n := len(a)
	aux := make([]string, n)

	// sort by key-indexed counting on dth char (stable)
	for d := w - 1; d >= 0; d-- {
		count := make([]int, r+1)

		// compute frequency counts
		for _, s := range a {
			count[s[d]+1]++
		}

		// compute cumulative counts
		for i := 0; i < r; i++ {
			count[i+1] += count[i]
		}

		// distribute keys to aux
		for _, s := range a {
			aux[count[s[d]]] = s
			count[s[d]]++
		}

		// copy back aux to a
		for i := 0; i < n; i++ {
			a[i] = aux[i]
		}
	}
}

// LSDInt is the LSD (least significant digit) sorting algorithm for integer numbers.
func LSDInt(a []int) {
	// TODO:
	// Ref: https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/LSD.java.html
}
