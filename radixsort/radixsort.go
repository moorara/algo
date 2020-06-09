// Package radixsort implements common radix sorting algorithms.
// Radix sorts are key-indexed counting for sorting keys with integer digits between 0 and R-1 (R is a small number).
package radixsort

import (
	"math/rand"
	"strings"
)

const r = 256 // uint8 (byte) size

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

func shuffleStringSlice(a []string) {
	n := len(a)
	for i := 0; i < n; i++ {
		r := i + rand.Intn(n-i)
		a[i], a[r] = a[r], a[i]
	}
}
