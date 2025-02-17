package generic

import (
	"iter"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testC1[T any] struct {
	items []T
}

func (c *testC1[T]) Add(vals ...T) {
	c.items = append(c.items, vals...)
}

func (c *testC1[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, v := range c.items {
			if !yield(v) {
				return
			}
		}
	}
}

type testC2[K comparable, V any] struct {
	keys []K
	vals []V
}

func (c *testC2[K, V]) Put(key K, val V) {
	c.keys = append(c.keys, key)
	c.vals = append(c.vals, val)
}

func (c *testC2[K, V]) All() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for i := range c.keys {
			if !yield(c.keys[i], c.vals[i]) {
				return
			}
		}
	}
}

func TestCollect1(t *testing.T) {
	c := new(testC1[string])
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
	c := new(testC2[string, int])
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

func TestFind(t *testing.T) {
	tests := []struct {
		name          string
		s             []string
		eq            EqualFunc[string]
		val           string
		expectedIndex int
	}{
		{
			name:          "Found",
			s:             []string{"Football", "Basketball", "Volleyball", "Handball"},
			eq:            NewEqualFunc[string](),
			val:           "Football",
			expectedIndex: 0,
		},
		{
			name:          "NotFound",
			s:             []string{"Football", "Basketball", "Volleyball", "Handball"},
			eq:            NewEqualFunc[string](),
			val:           "Water Polo",
			expectedIndex: -1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedIndex, Find(tc.s, tc.eq, tc.val))
		})
	}
}

func TestContains(t *testing.T) {
	tests := []struct {
		name             string
		s                []string
		eq               EqualFunc[string]
		vals             []string
		expectedContains bool
	}{
		{
			name:             "True",
			s:                []string{"Shih Tzu", "Pomeranian", "Chihuahua", "Maltese"},
			eq:               NewEqualFunc[string](),
			vals:             []string{"Shih Tzu", "Maltese"},
			expectedContains: true,
		},
		{
			name:             "False",
			s:                []string{"Shih Tzu", "Pomeranian", "Chihuahua", "Maltese"},
			eq:               NewEqualFunc[string](),
			vals:             []string{"Shih Tzu", "Bichon Frise"},
			expectedContains: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedContains, Contains(tc.s, tc.eq, tc.vals...))
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

func TestFirstMatch(t *testing.T) {
	tests := []struct {
		name          string
		s             []string
		p             Predicate1[string]
		expectedValue string
		expectedOK    bool
	}{
		{
			name:          "Empty",
			s:             []string{},
			p:             func(s string) bool { return s[0] == 'F' },
			expectedValue: "",
			expectedOK:    false,
		},
		{
			name:          "NoMatch",
			s:             []string{"Eagle", "Sparrow", "Owl", "Hummingbird", "Falcon", "Parrot", "Swan", "Seagull"},
			p:             func(s string) bool { return s[0] == 'M' },
			expectedValue: "",
			expectedOK:    false,
		},
		{
			name:          "OK",
			s:             []string{"Eagle", "Sparrow", "Owl", "Hummingbird", "Falcon", "Parrot", "Swan", "Seagull"},
			p:             func(s string) bool { return s[0] == 'S' },
			expectedValue: "Sparrow",
			expectedOK:    true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			val, ok := FirstMatch(tc.s, tc.p)
			assert.Equal(t, tc.expectedValue, val)
			assert.Equal(t, tc.expectedOK, ok)
		})
	}
}

func TestSelectMatch(t *testing.T) {
	tests := []struct {
		name             string
		s                []string
		p                Predicate1[string]
		expectedSelected []string
	}{
		{
			name:             "OK",
			s:                []string{"Eagle", "Sparrow", "Owl", "Hummingbird", "Falcon", "Parrot", "Swan", "Seagull"},
			p:                func(s string) bool { return s[0] == 'S' },
			expectedSelected: []string{"Sparrow", "Swan", "Seagull"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedSelected, SelectMatch(tc.s, tc.p))
		})
	}
}

func TestPartitionMatch(t *testing.T) {
	tests := []struct {
		name              string
		s                 []string
		p                 Predicate1[string]
		expectedMatched   []string
		expectedUnmatched []string
	}{
		{
			name:              "OK",
			s:                 []string{"Eagle", "Sparrow", "Owl", "Hummingbird", "Falcon", "Parrot", "Swan", "Seagull"},
			p:                 func(s string) bool { return s[0] == 'S' },
			expectedMatched:   []string{"Sparrow", "Swan", "Seagull"},
			expectedUnmatched: []string{"Eagle", "Owl", "Hummingbird", "Falcon", "Parrot"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			matched, unmatched := PartitionMatch(tc.s, tc.p)

			assert.Equal(t, tc.expectedMatched, matched)
			assert.Equal(t, tc.expectedUnmatched, unmatched)
		})
	}
}

func TestTransform(t *testing.T) {
	tests := []struct {
		name           string
		s              []string
		f              func(string) int
		expectedResult []int
	}{
		{
			name: "OK",
			s:    []string{"27", "69"},
			f: func(s string) int {
				i, _ := strconv.Atoi(s)
				return i
			},
			expectedResult: []int{27, 69},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedResult, Transform(tc.s, tc.f))
		})
	}
}
