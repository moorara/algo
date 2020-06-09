package heap

import (
	"math/rand"
	"testing"
	"time"

	"github.com/moorara/algo/compare"
	"github.com/stretchr/testify/assert"
)

func TestMinHeap(t *testing.T) {
	type keyValue struct {
		key   int
		value string
	}

	tests := []struct {
		name             string
		initialCapacity  int
		cmpKey           compare.Func
		cmpVal           compare.Func
		insertTests      []keyValue
		expectedSize     int
		expectedIsEmpty  bool
		expectedPeek     keyValue
		expectedContains []keyValue
		expectedDelete   []keyValue
	}{
		{
			name:             "Empty",
			initialCapacity:  2,
			cmpKey:           compare.Int,
			cmpVal:           compare.String,
			insertTests:      []keyValue{},
			expectedSize:     0,
			expectedIsEmpty:  true,
			expectedPeek:     keyValue{0, ""},
			expectedContains: []keyValue{},
			expectedDelete:   []keyValue{},
		},
		{
			name:            "FewPairs",
			initialCapacity: 2,
			cmpKey:          compare.Int,
			cmpVal:          compare.String,
			insertTests: []keyValue{
				{30, "thirty"},
				{10, "ten"},
				{20, "twenty"},
			},
			expectedSize:    3,
			expectedIsEmpty: false,
			expectedPeek:    keyValue{10, "ten"},
			expectedContains: []keyValue{
				{10, "ten"},
				{20, "twenty"},
				{30, "thirty"},
			},
			expectedDelete: []keyValue{
				{10, "ten"},
				{20, "twenty"},
				{30, "thirty"},
			},
		},
		{
			name:            "SomePairs",
			initialCapacity: 4,
			cmpKey:          compare.Int,
			cmpVal:          compare.String,
			insertTests: []keyValue{
				{50, "fifty"},
				{30, "thirty"},
				{40, "forty"},
				{10, "ten"},
				{20, "twenty"},
			},
			expectedSize:    5,
			expectedIsEmpty: false,
			expectedPeek:    keyValue{10, "ten"},
			expectedContains: []keyValue{
				{10, "ten"},
				{20, "twenty"},
				{30, "thirty"},
				{40, "forty"},
				{50, "fifty"},
			},
			expectedDelete: []keyValue{
				{10, "ten"},
				{20, "twenty"},
				{30, "thirty"},
				{40, "forty"},
				{50, "fifty"},
			},
		},
		{
			name:            "MorePairs",
			initialCapacity: 4,
			cmpKey:          compare.Int,
			cmpVal:          compare.String,
			insertTests: []keyValue{
				{90, "ninety"},
				{80, "eighty"},
				{70, "seventy"},
				{40, "forty"},
				{50, "fifty"},
				{60, "sixty"},
				{30, "thirty"},
				{10, "ten"},
				{20, "twenty"},
			},
			expectedSize:    9,
			expectedIsEmpty: false,
			expectedPeek:    keyValue{10, "ten"},
			expectedContains: []keyValue{
				{10, "ten"},
				{20, "twenty"},
				{30, "thirty"},
				{40, "forty"},
				{50, "fifty"},
				{60, "sixty"},
				{70, "seventy"},
				{80, "eighty"},
				{90, "ninety"},
			},
			expectedDelete: []keyValue{
				{10, "ten"},
				{20, "twenty"},
				{30, "thirty"},
				{40, "forty"},
				{50, "fifty"},
				{60, "sixty"},
				{70, "seventy"},
				{80, "eighty"},
				{90, "ninety"},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			heap := NewMinHeap(tc.initialCapacity, tc.cmpKey, tc.cmpVal)

			// Heap initially should be empty
			assert.Zero(t, heap.Size())
			assert.True(t, heap.IsEmpty())
			assert.False(t, heap.ContainsKey(nil))
			assert.False(t, heap.ContainsValue(nil))

			peekKey, peekValue, peekOK := heap.Peek()
			assert.Nil(t, peekKey)
			assert.Nil(t, peekValue)
			assert.False(t, peekOK)

			deleteKey, deleteValue, deleteOK := heap.Delete()
			assert.Nil(t, deleteKey)
			assert.Nil(t, deleteValue)
			assert.False(t, deleteOK)

			for _, entry := range tc.insertTests {
				heap.Insert(entry.key, entry.value)
			}

			assert.Equal(t, tc.expectedSize, heap.Size())
			assert.Equal(t, tc.expectedIsEmpty, heap.IsEmpty())

			peekKey, peekValue, peekOK = heap.Peek()
			if tc.expectedSize == 0 {
				assert.Nil(t, peekKey)
				assert.Nil(t, peekValue)
				assert.False(t, peekOK)
			} else {
				assert.Equal(t, tc.expectedPeek.key, peekKey)
				assert.Equal(t, tc.expectedPeek.value, peekValue)
				assert.True(t, peekOK)
			}

			for _, entry := range tc.expectedContains {
				assert.True(t, heap.ContainsKey(entry.key))
				assert.True(t, heap.ContainsValue(entry.value))
			}

			for _, entry := range tc.expectedDelete {
				deleteKey, deleteValue, deleteOK = heap.Delete()
				assert.Equal(t, entry.key, deleteKey)
				assert.Equal(t, entry.value, deleteValue)
				assert.True(t, deleteOK)
			}

			// Heap should be empty at the end
			assert.Zero(t, heap.Size())
			assert.True(t, heap.IsEmpty())
			assert.False(t, heap.ContainsKey(nil))
			assert.False(t, heap.ContainsValue(nil))

			peekKey, peekValue, peekOK = heap.Peek()
			assert.Nil(t, peekKey)
			assert.Nil(t, peekValue)
			assert.False(t, peekOK)

			deleteKey, deleteValue, deleteOK = heap.Delete()
			assert.Nil(t, deleteKey)
			assert.Nil(t, deleteValue)
			assert.False(t, deleteOK)
		})
	}
}

func BenchmarkMinHeap(b *testing.B) {
	heapSize := 1024
	minInt := 0
	maxInt := 1000000

	rand.Seed(time.Now().UTC().UnixNano())

	b.Run("Insert", func(b *testing.B) {
		heap := NewMinHeap(heapSize, compare.Int, compare.String)
		keys := randIntSlice(b.N, minInt, maxInt)
		values := randStringSlice(b.N)

		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			heap.Insert(keys[n], values[n])
		}
	})

	b.Run("Delete", func(b *testing.B) {
		heap := NewMinHeap(heapSize, compare.Int, compare.String)
		keys := randIntSlice(b.N, minInt, maxInt)
		values := randStringSlice(b.N)

		for n := 0; n < b.N; n++ {
			heap.Insert(keys[n], values[n])
		}

		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			heap.Delete()
		}
	})
}
