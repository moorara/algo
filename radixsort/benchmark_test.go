package radixsort

import (
	"math/rand"
	"testing"
	"time"
)

const (
	slen = 128
	size = 1000000

	chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func randString(l int) string {
	n := len(chars)
	b := make([]byte, l)

	for i := range b {
		b[i] = chars[rand.Intn(n)]
	}

	return string(b)
}

func BenchmarkString(b *testing.B) {
	rand.Seed(time.Now().UnixNano())

	// generate a sequence of random strings
	a := make([]string, size)
	for i := range a {
		a[i] = randString(slen)
	}

	shuffle[string](a)

	b.Run("LSDString", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			LSDString(a, slen)
		}
	})

	b.Run("MSDString", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			MSDString(a)
		}
	})

	b.Run("Quick3WayString", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			Quick3WayString(a)
		}
	})
}

func BenchmarkInt(b *testing.B) {
	rand.Seed(time.Now().UnixNano())

	// generate a sequence of random integers (signed).
	a := make([]int, size)
	for i := range a {
		a[i] = rand.Int()
	}

	shuffle[int](a)

	b.Run("LSDInt", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			LSDInt(a)
		}
	})

	b.Run("MSDInt", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			MSDInt(a)
		}
	})
}

func BenchmarkUint(b *testing.B) {
	rand.Seed(time.Now().UnixNano())

	// generate a sequence of random integers (unsigned).
	a := make([]uint, size)
	for i := range a {
		a[i] = (uint(rand.Uint32()) << 32) + uint(rand.Uint32())
	}

	shuffle[uint](a)

	b.Run("LSDUint", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			LSDUint(a)
		}
	})

	b.Run("MSDUint", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			MSDUint(a)
		}
	})
}
