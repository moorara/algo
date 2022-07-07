// Package radixsort implements common radix sorting algorithms.
// Radix sorts are key-indexed counting for sorting keys with integer digits between 0 and R-1 (R is a small number).
// Radix sorting algorithms are efficient for a large number of keys with small constant width.
package radixsort

import (
	"math/rand"

	"golang.org/x/exp/constraints"
)

func charAt(s string, d int) int {
	if d < len(s) {
		return int(s[d])
	}
	return -1
}

func insertion[T constraints.Ordered](a []T, lo, hi int) {
	for i := lo; i <= hi; i++ {
		for j := i; j > lo && a[j] < a[j-1]; j-- {
			a[j], a[j-1] = a[j-1], a[j]
		}
	}
}

func shuffle[T any](a []T) {
	n := len(a)
	for i := 0; i < n; i++ {
		r := i + rand.Intn(n-i)
		a[i], a[r] = a[r], a[i]
	}
}
