package heap

import (
	"math/rand"
	"testing"
	"time"

	"github.com/moorara/algo/compare"
	"github.com/stretchr/testify/assert"
)

func TestIndexMaxHeap(t *testing.T) {
	type indexKeyValue struct {
		index int
		key   int
		value string
	}

	tests := []struct {
		name                string
		capacity            int
		cmpKey              compare.Func
		cmpVal              compare.Func
		insertTests         []indexKeyValue
		changeKeyTests      []indexKeyValue
		expectedSize        int
		expectedIsEmpty     bool
		expectedPeek        indexKeyValue
		expectedPeekIndex   []indexKeyValue
		expectedContains    []indexKeyValue
		expectedDelete      []indexKeyValue
		expectedDeleteIndex []indexKeyValue
	}{
		{
			name:                "Empty",
			capacity:            10,
			cmpKey:              compare.Int,
			cmpVal:              compare.String,
			insertTests:         []indexKeyValue{},
			changeKeyTests:      []indexKeyValue{},
			expectedSize:        0,
			expectedIsEmpty:     true,
			expectedPeek:        indexKeyValue{},
			expectedPeekIndex:   []indexKeyValue{},
			expectedContains:    []indexKeyValue{},
			expectedDelete:      []indexKeyValue{},
			expectedDeleteIndex: []indexKeyValue{},
		},
		{
			name:     "FewPairs",
			capacity: 10,
			cmpKey:   compare.Int,
			cmpVal:   compare.String,
			insertTests: []indexKeyValue{
				{0, 10, "ten"},
				{1, 3, "thirty"},
				{2, 200, "twenty"},
			},
			changeKeyTests: []indexKeyValue{
				{index: 1, key: 30},
				{index: 2, key: 20},
			},
			expectedSize:    3,
			expectedIsEmpty: false,
			expectedPeek:    indexKeyValue{1, 30, "thirty"},
			expectedPeekIndex: []indexKeyValue{
				{1, 30, "thirty"},
				{2, 20, "twenty"},
				{0, 10, "ten"},
			},
			expectedContains: []indexKeyValue{
				{1, 30, "thirty"},
				{2, 20, "twenty"},
				{0, 10, "ten"},
			},
			expectedDelete: []indexKeyValue{
				{1, 30, "thirty"},
			},
			expectedDeleteIndex: []indexKeyValue{
				{0, 10, "ten"},
				{2, 20, "twenty"},
			},
		},
		{
			name:     "SomePairs",
			capacity: 10,
			cmpKey:   compare.Int,
			cmpVal:   compare.String,
			insertTests: []indexKeyValue{
				{0, 10, "ten"},
				{1, 30, "thirty"},
				{2, 2, "twenty"},
				{3, 50, "fifty"},
				{4, 400, "forty"},
			},
			changeKeyTests: []indexKeyValue{
				{index: 2, key: 20},
				{index: 3, key: 50},
				{index: 4, key: 40},
			},
			expectedSize:    5,
			expectedIsEmpty: false,
			expectedPeek:    indexKeyValue{3, 50, "fifty"},
			expectedPeekIndex: []indexKeyValue{
				{3, 50, "fifty"},
				{4, 40, "forty"},
				{1, 30, "thirty"},
				{2, 20, "twenty"},
				{0, 10, "ten"},
			},
			expectedContains: []indexKeyValue{
				{3, 50, "fifty"},
				{4, 40, "forty"},
				{1, 30, "thirty"},
				{2, 20, "twenty"},
				{0, 10, "ten"},
			},
			expectedDelete: []indexKeyValue{
				{3, 50, "fifty"},
				{4, 40, "forty"},
			},
			expectedDeleteIndex: []indexKeyValue{
				{0, 10, "ten"},
				{2, 20, "twenty"},
				{1, 30, "thirty"},
			},
		},
		{
			name:     "MorePairs",
			capacity: 10,
			cmpKey:   compare.Int,
			cmpVal:   compare.String,
			insertTests: []indexKeyValue{
				{0, 10, "ten"},
				{1, 30, "thirty"},
				{2, 20, "twenty"},
				{3, 50, "fifty"},
				{4, 4, "forty"},
				{5, 6, "sixty"},
				{6, 70, "seventy"},
				{7, 900, "ninety"},
				{8, 800, "eighty"},
			},
			changeKeyTests: []indexKeyValue{
				{index: 4, key: 40},
				{index: 5, key: 60},
				{index: 6, key: 70},
				{index: 7, key: 90},
				{index: 8, key: 80},
			},
			expectedSize:    9,
			expectedIsEmpty: false,
			expectedPeek:    indexKeyValue{7, 90, "ninety"},
			expectedPeekIndex: []indexKeyValue{
				{7, 90, "ninety"},
				{8, 80, "eighty"},
				{6, 70, "seventy"},
				{5, 60, "sixty"},
				{3, 50, "fifty"},
				{4, 40, "forty"},
				{1, 30, "thirty"},
				{2, 20, "twenty"},
				{0, 10, "ten"},
			},
			expectedContains: []indexKeyValue{
				{7, 90, "ninety"},
				{8, 80, "eighty"},
				{6, 70, "seventy"},
				{5, 60, "sixty"},
				{3, 50, "fifty"},
				{4, 40, "forty"},
				{1, 30, "thirty"},
				{2, 20, "twenty"},
				{0, 10, "ten"},
			},
			expectedDelete: []indexKeyValue{
				{7, 90, "ninety"},
				{8, 80, "eighty"},
				{6, 70, "seventy"},
				{5, 60, "sixty"},
			},
			expectedDeleteIndex: []indexKeyValue{
				{0, 10, "ten"},
				{2, 20, "twenty"},
				{1, 30, "thirty"},
				{4, 40, "forty"},
				{3, 50, "fifty"},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			heap := NewIndexMaxHeap(tc.capacity, tc.cmpKey, tc.cmpVal)

			// Heap initially should be empty
			assert.Zero(t, heap.Size())
			assert.True(t, heap.IsEmpty())
			assert.False(t, heap.ContainsIndex(0))
			assert.False(t, heap.ContainsKey(nil))
			assert.False(t, heap.ContainsValue(nil))

			peekIndex, peekKey, peekValue, peekOK := heap.Peek()
			assert.Equal(t, -1, peekIndex)
			assert.Nil(t, peekKey)
			assert.Nil(t, peekValue)
			assert.False(t, peekOK)

			peekKey, peekValue, peekOK = heap.PeekIndex(0)
			assert.Equal(t, -1, peekIndex)
			assert.Nil(t, peekKey)
			assert.Nil(t, peekValue)
			assert.False(t, peekOK)

			deleteIndex, deleteKey, deleteValue, deleteOK := heap.Delete()
			assert.Equal(t, -1, deleteIndex)
			assert.Nil(t, deleteKey)
			assert.Nil(t, deleteValue)
			assert.False(t, deleteOK)

			deleteKey, deleteValue, deleteOK = heap.DeleteIndex(0)
			assert.Nil(t, deleteKey)
			assert.Nil(t, deleteValue)
			assert.False(t, deleteOK)

			for _, entry := range tc.insertTests {
				heap.Insert(entry.index, entry.key, entry.value)
			}

			for _, entry := range tc.changeKeyTests {
				heap.ChangeKey(entry.index, entry.key)
			}

			assert.Equal(t, tc.expectedSize, heap.Size())
			assert.Equal(t, tc.expectedIsEmpty, heap.IsEmpty())

			peekIndex, peekKey, peekValue, peekOK = heap.Peek()
			if tc.expectedSize == 0 {
				assert.Equal(t, -1, peekIndex)
				assert.Nil(t, peekKey)
				assert.Nil(t, peekValue)
				assert.False(t, peekOK)
			} else {
				assert.Equal(t, tc.expectedPeek.index, peekIndex)
				assert.Equal(t, tc.expectedPeek.key, peekKey)
				assert.Equal(t, tc.expectedPeek.value, peekValue)
				assert.True(t, peekOK)
			}

			for _, entry := range tc.expectedPeekIndex {
				peekKey, peekValue, peekOK = heap.PeekIndex(entry.index)
				assert.Equal(t, entry.key, peekKey)
				assert.Equal(t, entry.value, peekValue)
				assert.True(t, peekOK)
			}

			for _, entry := range tc.expectedContains {
				assert.True(t, heap.ContainsIndex(entry.index))
				assert.True(t, heap.ContainsKey(entry.key))
				assert.True(t, heap.ContainsValue(entry.value))
			}

			for _, entry := range tc.expectedDelete {
				deleteIndex, deleteKey, deleteValue, deleteOK = heap.Delete()
				assert.Equal(t, entry.index, deleteIndex)
				assert.Equal(t, entry.key, deleteKey)
				assert.Equal(t, entry.value, deleteValue)
				assert.True(t, deleteOK)
			}

			for _, entry := range tc.expectedDeleteIndex {
				deleteKey, deleteValue, deleteOK = heap.DeleteIndex(entry.index)
				assert.Equal(t, entry.key, deleteKey)
				assert.Equal(t, entry.value, deleteValue)
				assert.True(t, deleteOK)
			}

			// Heap should be empty at the end
			assert.Zero(t, heap.Size())
			assert.True(t, heap.IsEmpty())
			assert.False(t, heap.ContainsKey(nil))
			assert.False(t, heap.ContainsValue(nil))

			peekIndex, peekKey, peekValue, peekOK = heap.Peek()
			assert.Equal(t, -1, peekIndex)
			assert.Nil(t, peekKey)
			assert.Nil(t, peekValue)
			assert.False(t, peekOK)

			peekKey, peekValue, peekOK = heap.PeekIndex(0)
			assert.Equal(t, -1, peekIndex)
			assert.Nil(t, peekKey)
			assert.Nil(t, peekValue)
			assert.False(t, peekOK)

			deleteIndex, deleteKey, deleteValue, deleteOK = heap.Delete()
			assert.Equal(t, -1, deleteIndex)
			assert.Nil(t, deleteKey)
			assert.Nil(t, deleteValue)
			assert.False(t, peekOK, deleteOK)

			deleteKey, deleteValue, deleteOK = heap.DeleteIndex(0)
			assert.Nil(t, deleteKey)
			assert.Nil(t, deleteValue)
			assert.False(t, deleteOK)
		})
	}
}

func BenchmarkIndexMaxHeap(b *testing.B) {
	minInt := 0
	maxInt := 1000000

	rand.Seed(time.Now().UTC().UnixNano())

	b.Run("Insert", func(b *testing.B) {
		heap := NewIndexMaxHeap(b.N, compare.Int, compare.String)
		keys := randIntSlice(b.N, minInt, maxInt)
		values := randStringSlice(b.N)

		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			heap.Insert(n, keys[n], values[n])
		}
	})

	b.Run("Delete", func(b *testing.B) {
		heap := NewIndexMaxHeap(b.N, compare.Int, compare.String)
		keys := randIntSlice(b.N, minInt, maxInt)
		values := randStringSlice(b.N)

		for n := 0; n < b.N; n++ {
			heap.Insert(n, keys[n], values[n])
		}

		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			heap.Delete()
		}
	})
}
