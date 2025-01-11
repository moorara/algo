package generic

import (
	"iter"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCollection1[T any] struct {
	items []T
}

func (c *testCollection1[T]) Add(vals ...T) {
	c.items = append(c.items, vals...)
}

func (c *testCollection1[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, v := range c.items {
			if !yield(v) {
				return
			}
		}
	}
}

type testCollection2[K comparable, V any] struct {
	keys []K
	vals []V
}

func (c *testCollection2[K, V]) Put(key K, val V) {
	c.keys = append(c.keys, key)
	c.vals = append(c.vals, val)
}

func (c *testCollection2[K, V]) All() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for i := range c.keys {
			if !yield(c.keys[i], c.vals[i]) {
				return
			}
		}
	}
}

func TestCollect1(t *testing.T) {
	c := new(testCollection1[string])
	c.Add("foo", "bar")

	expectedSlice := []string{"foo", "bar"}

	slice := Collect1(c.All())
	assert.Equal(t, expectedSlice, slice)
}

func TestCollect2(t *testing.T) {
	c := new(testCollection2[string, int])
	c.Put("foo", 1)
	c.Put("bar", 2)

	expectedSlice := []KeyValue[string, int]{
		{"foo", 1},
		{"bar", 2},
	}

	slice := Collect2(c.All())
	assert.Equal(t, expectedSlice, slice)
}
