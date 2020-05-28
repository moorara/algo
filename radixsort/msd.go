package radixsort

import "strings"

const cutoff = 15

func charAt(s string, d int) int {
	if d < len(s) {
		return int(s[d])
	}
	return -1
}

func insertionString(a []string, lo, hi, d int) {
	for i := lo; i <= hi; i++ {
		for j := i; j > 0 && strings.Compare(a[j], a[j-1]) < 0; j-- {
			a[j], a[j-1] = a[j-1], a[j]
		}
	}
}

func msdString(a, aux []string, lo, hi, d int) {
	if hi <= lo+cutoff {
		insertionString(a, lo, hi, d)
		return
	}

	count := make([]int, r+2)

	// compute frequency counts
	for i := lo; i <= hi; i++ {
		c := charAt(a[i], d)
		count[c+2]++
	}

	// transform counts to indicies
	for i := 0; i < r+1; i++ {
		count[i+1] += count[i]
	}

	// distribute keys to aux
	for i := lo; i <= hi; i++ {
		c := charAt(a[i], d)
		aux[count[c+1]] = a[i]
		count[c+1]++
	}

	// copy back aux to a
	for i := lo; i <= hi; i++ {
		a[i] = aux[i-lo]
	}

	// recursively sort for each character (excludes sentinel -1)
	for i := 0; i < r; i++ {
		msdString(a, aux, lo+count[i], lo+count[i+1]-1, d+1)
	}
}

// MSDString is the MSD (most significant digit) sorting algorithm for string keys with variable length.
func MSDString(a []string) {
	n := len(a)
	aux := make([]string, n)
	msdString(a, aux, 0, n-1, 0)
}

func msdInt(a, aux []int, lo, hi, d int) {
	// TODO:
	// Ref: https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/MSD.java.html
}

// MSDInt is the MSD (most significant digit) sorting algorithm for integer numbers.
func MSDInt(a []int) {
	// TODO:
	// Ref: https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/MSD.java.html
}
