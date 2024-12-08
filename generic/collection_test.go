package generic

import (
	"iter"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCollection[K comparable, V any] struct {
	keys []K
	vals []V
}

func (c *testCollection[K, V]) Add(key K, val V) {
	c.keys = append(c.keys, key)
	c.vals = append(c.vals, val)
}

func (c *testCollection[K, V]) All() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for i := range c.keys {
			if !yield(c.keys[i], c.vals[i]) {
				return
			}
		}
	}
}

func TestCollect(t *testing.T) {
	c := new(testCollection[string, int])
	c.Add("foo", 1)
	c.Add("bar", 2)

	expectedSlice := []KeyValue[string, int]{
		{"foo", 1},
		{"bar", 2},
	}

	slice := Collect(c.All())
	assert.Equal(t, expectedSlice, slice)
}
