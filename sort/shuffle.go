package sort

import "math/rand"

// Shuffle shuffles a slice in O(n) time.
func Shuffle[T any](a []T, r *rand.Rand) {
	n := len(a)
	for i := 0; i < n; i++ {
		r := i + r.Intn(n-i)
		a[i], a[r] = a[r], a[i]
	}
}
