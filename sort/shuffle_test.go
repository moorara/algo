package sort

import (
	"math/rand"
	"testing"
	"time"
)

func TestShuffle_int(t *testing.T) {
	tests := []struct {
		items []int
	}{
		{[]int{10, 20, 30, 40, 50, 60, 70, 80, 90}},
	}

	rand.Seed(time.Now().UTC().UnixNano())

	for _, tc := range tests {
		Shuffle[int](tc.items)
	}
}

func TestShuffle_string(t *testing.T) {
	tests := []struct {
		items []string
	}{
		{[]string{"Alice", "Bob", "Dan", "Edgar", "Helen", "Karen", "Milad", "Peter", "Sam", "Wesley"}},
	}

	rand.Seed(time.Now().UTC().UnixNano())

	for _, tc := range tests {
		Shuffle[string](tc.items)
	}
}
