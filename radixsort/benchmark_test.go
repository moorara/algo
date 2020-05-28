package radixsort

import (
	"math/rand"
	"testing"
)

const (
	seed   = 27
	size   = 1000
	keyLen = 16

	chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func randStringKey(l int) string {
	b := make([]byte, l)
	for i := range b {
		n := len(chars)
		b[i] = chars[rand.Intn(n)]
	}
	return string(b)
}

func randStringSlice(size, keyLen int) []string {
	// make sure benchmarks are deterministic
	rand.Seed(seed)

	// generate
	a := make([]string, size)
	for i := range a {
		a[i] = randStringKey(keyLen)
	}

	// shuffle
	n := len(a)
	for i := 0; i < n; i++ {
		r := i + rand.Intn(n-i)
		a[i], a[r] = a[r], a[i]
	}

	return a
}

func BenchmarkString(b *testing.B) {
	b.Run("LSDString", func(b *testing.B) {
		items := randStringSlice(size, keyLen)
		b.ResetTimer()

		for n := 0; n < b.N; n++ {
			LSDString(items, keyLen)
		}
	})
}
