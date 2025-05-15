package sort

import (
	"math/rand"
	"time"

	"github.com/moorara/algo/generic"
)

func partition[T any](a []T, lo, hi int, cmp generic.CompareFunc[T]) int {
	v := a[lo]
	i, j := lo, hi+1

	for {
		for i++; i < hi && cmp(a[i], v) < 0; i++ {
		}
		for j--; j > lo && cmp(a[j], v) > 0; j-- {
		}
		if i >= j {
			break
		}
		a[i], a[j] = a[j], a[i]
	}
	a[lo], a[j] = a[j], a[lo]

	return j
}

// Select finds the kth smallest item of an array in O(n) time on average.
func Select[T any](a []T, k int, cmp generic.CompareFunc[T]) T {
	seed := time.Now().UTC().UnixNano()
	r := rand.New(rand.NewSource(seed))
	Shuffle[T](a, r)

	lo, hi := 0, len(a)-1
	for lo < hi {
		j := partition[T](a, lo, hi, cmp)
		switch {
		case j < k:
			lo = j + 1
		case j > k:
			hi = j - 1
		default:
			return a[k]
		}
	}

	return a[k]
}

func quick[T any](a []T, lo, hi int, cmp generic.CompareFunc[T]) {
	if lo >= hi {
		return
	}

	j := partition[T](a, lo, hi, cmp)
	quick[T](a, lo, j-1, cmp)
	quick[T](a, j+1, hi, cmp)
}

// Quick implements the quick sort algorithm.
func Quick[T any](a []T, cmp generic.CompareFunc[T]) {
	seed := time.Now().UTC().UnixNano()
	r := rand.New(rand.NewSource(seed))
	Shuffle[T](a, r)

	quick[T](a, 0, len(a)-1, cmp)
}

func quick3Way[T any](a []T, lo, hi int, cmp generic.CompareFunc[T]) {
	if lo >= hi {
		return
	}

	v := a[lo]
	lt, i, gt := lo, lo+1, hi

	for i <= gt {
		c := cmp(a[i], v)
		switch {
		case c < 0:
			a[lt], a[i] = a[i], a[lt]
			lt++
			i++
		case c > 0:
			a[i], a[gt] = a[gt], a[i]
			gt--
		default:
			i++
		}
	}

	quick3Way[T](a, lo, lt-1, cmp)
	quick3Way[T](a, gt+1, hi, cmp)
}

// Quick3Way implements the 3-way version of quick sort algorithm.
func Quick3Way[T any](a []T, cmp generic.CompareFunc[T]) {
	quick3Way[T](a, 0, len(a)-1, cmp)
}
