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

	tests := []struct {
		name          string
		seq           iter.Seq[string]
		expectedSlice []string
	}{
		{
			name:          "Nil",
			seq:           nil,
			expectedSlice: nil,
		},
		{
			name:          "OK",
			seq:           c.All(),
			expectedSlice: []string{"foo", "bar"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			slice := Collect1(tc.seq)
			assert.Equal(t, tc.expectedSlice, slice)
		})
	}
}

func TestCollect2(t *testing.T) {
	c := new(testCollection2[string, int])
	c.Put("foo", 1)
	c.Put("bar", 2)

	tests := []struct {
		name          string
		seq2          iter.Seq2[string, int]
		expectedSlice []KeyValue[string, int]
	}{
		{
			name:          "Nil",
			seq2:          nil,
			expectedSlice: nil,
		},
		{
			name: "OK",
			seq2: c.All(),
			expectedSlice: []KeyValue[string, int]{
				{"foo", 1},
				{"bar", 2},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			slice := Collect2(tc.seq2)
			assert.Equal(t, tc.expectedSlice, slice)
		})
	}

}

func TestAnyMatch(t *testing.T) {
	tests := []struct {
		name             string
		s                []string
		p                Predicate1[string]
		expectedAnyMatch bool
	}{
		{
			name:             "True",
			s:                []string{"Rose", "Lily", "Jasmine", "Camellia"},
			p:                func(s string) bool { return s == "Rose" },
			expectedAnyMatch: true,
		},
		{
			name:             "False",
			s:                []string{"Rose", "Lily", "Jasmine", "Camellia"},
			p:                func(s string) bool { return s == "Tulip" },
			expectedAnyMatch: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedAnyMatch, AnyMatch(tc.s, tc.p))
		})
	}
}

func TestAllMatch(t *testing.T) {
	tests := []struct {
		name             string
		s                []string
		p                Predicate1[string]
		expectedAllMatch bool
	}{
		{
			name:             "True",
			s:                []string{"Apple", "Apricot", "Avocado", "Acerola"},
			p:                func(s string) bool { return s[0] == 'A' },
			expectedAllMatch: true,
		},
		{
			name:             "False",
			s:                []string{"Apple", "Apricot", "Avocado", "Acerola"},
			p:                func(s string) bool { return s[1] == 'p' },
			expectedAllMatch: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedAllMatch, AllMatch(tc.s, tc.p))
		})
	}
}

func TestSelectMatch(t *testing.T) {
	tests := []struct {
		name                string
		s                   []string
		p                   Predicate1[string]
		expectedSelectMatch []string
	}{
		{
			name:                "OK",
			s:                   []string{"Eagle", "Sparrow", "Owl", "Hummingbird", "Falcon", "Parrot", "Swan", "Seagull"},
			p:                   func(s string) bool { return s[0] == 'S' },
			expectedSelectMatch: []string{"Sparrow", "Swan", "Seagull"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedSelectMatch, SelectMatch(tc.s, tc.p))
		})
	}
}
