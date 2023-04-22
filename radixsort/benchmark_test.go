package radixsort

import (
	"math/rand"
	"testing"
	"time"
)

const (
	slen  = 128
	size  = 100000
	chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

var r *rand.Rand

func randString(l int) string {
	n := len(chars)
	b := make([]byte, l)

	for i := range b {
		b[i] = chars[r.Intn(n)]
	}

	return string(b)
}

func BenchmarkString(b *testing.B) {
	seed := time.Now().UTC().UnixNano()
	r = rand.New(rand.NewSource(seed))

	// generate a sequence of random strings
	vals := make([]string, size)
	for i := range vals {
		vals[i] = randString(slen)
	}

	b.Run("LSDString", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			a := make([]string, len(vals))
			copy(a, vals)
			shuffle[string](a)
			LSDString(a, slen)
		}
	})

	b.Run("MSDString", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			a := make([]string, len(vals))
			copy(a, vals)
			shuffle[string](a)
			MSDString(a)
		}
	})

	b.Run("Quick3WayString", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			a := make([]string, len(vals))
			copy(a, vals)
			shuffle[string](a)
			Quick3WayString(a)
		}
	})
}

func BenchmarkInt(b *testing.B) {
	seed := time.Now().UTC().UnixNano()
	r = rand.New(rand.NewSource(seed))

	// generate a sequence of random integers (signed).
	nums := make([]int, size)
	for i := range nums {
		nums[i] = r.Int()
	}

	b.Run("LSDInt", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			a := make([]int, len(nums))
			copy(a, nums)
			shuffle[int](a)
			LSDInt(a)
		}
	})

	b.Run("MSDInt", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			a := make([]int, len(nums))
			copy(a, nums)
			shuffle[int](a)
			MSDInt(a)
		}
	})
}

func BenchmarkUint(b *testing.B) {
	seed := time.Now().UTC().UnixNano()
	r = rand.New(rand.NewSource(seed))

	// generate a sequence of random integers (unsigned).
	nums := make([]uint, size)
	for i := range nums {
		nums[i] = (uint(r.Uint32()) << 32) + uint(r.Uint32())
	}

	b.Run("LSDUint", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			a := make([]uint, len(nums))
			copy(a, nums)
			shuffle[uint](a)
			LSDUint(a)
		}
	})

	b.Run("MSDUint", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			a := make([]uint, len(nums))
			copy(a, nums)
			shuffle[uint](a)
			MSDUint(a)
		}
	})
}
