package heap

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMaxHeap(t *testing.T) {
	type keyValue struct {
		key   int
		value string
	}

	tests := []struct {
		name             string
		initialCapacity  int
		cmpKey           CompareFunc
		cmpVal           CompareFunc
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
			cmpKey:           compareInt,
			cmpVal:           compareString,
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
			cmpKey:          compareInt,
			cmpVal:          compareString,
			insertTests: []keyValue{
				{10, "ten"},
				{30, "thirty"},
				{20, "twenty"},
			},
			expectedSize:    3,
			expectedIsEmpty: false,
			expectedPeek:    keyValue{30, "thirty"},
			expectedContains: []keyValue{
				{30, "thirty"},
				{20, "twenty"},
				{10, "ten"},
			},
			expectedDelete: []keyValue{
				{30, "thirty"},
				{20, "twenty"},
				{10, "ten"},
			},
		},
		{
			name:            "SomePairs",
			initialCapacity: 4,
			cmpKey:          compareInt,
			cmpVal:          compareString,
			insertTests: []keyValue{
				{10, "ten"},
				{30, "thirty"},
				{20, "twenty"},
				{50, "fifty"},
				{40, "forty"},
			},
			expectedSize:    5,
			expectedIsEmpty: false,
			expectedPeek:    keyValue{50, "fifty"},
			expectedContains: []keyValue{
				{50, "fifty"},
				{40, "forty"},
				{30, "thirty"},
				{20, "twenty"},
				{10, "ten"},
			},
			expectedDelete: []keyValue{
				{50, "fifty"},
				{40, "forty"},
				{30, "thirty"},
				{20, "twenty"},
				{10, "ten"},
			},
		},
		{
			name:            "MorePairs",
			initialCapacity: 4,
			cmpKey:          compareInt,
			cmpVal:          compareString,
			insertTests: []keyValue{
				{10, "ten"},
				{30, "thirty"},
				{20, "twenty"},
				{50, "fifty"},
				{40, "forty"},
				{60, "sixty"},
				{70, "seventy"},
				{90, "ninety"},
				{80, "eighty"},
			},
			expectedSize:    9,
			expectedIsEmpty: false,
			expectedPeek:    keyValue{90, "ninety"},
			expectedContains: []keyValue{
				{90, "ninety"},
				{80, "eighty"},
				{70, "seventy"},
				{60, "sixty"},
				{50, "fifty"},
				{40, "forty"},
				{30, "thirty"},
				{20, "twenty"},
				{10, "ten"},
			},
			expectedDelete: []keyValue{
				{90, "ninety"},
				{80, "eighty"},
				{70, "seventy"},
				{60, "sixty"},
				{50, "fifty"},
				{40, "forty"},
				{30, "thirty"},
				{20, "twenty"},
				{10, "ten"},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			heap := NewMaxHeap(tc.initialCapacity, tc.cmpKey, tc.cmpVal)

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
			assert.False(t, peekOK, deleteOK)
		})
	}
}

func BenchmarkMaxHeap(b *testing.B) {
	heapSize := 1024
	minInt := 0
	maxInt := 1000000

	rand.Seed(time.Now().UTC().UnixNano())

	b.Run("Insert", func(b *testing.B) {
		heap := NewMaxHeap(heapSize, compareInt, compareString)
		keys := randIntSlice(b.N, minInt, maxInt)
		values := randStringSlice(b.N)

		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			heap.Insert(keys[n], values[n])
		}
	})

	b.Run("Delete", func(b *testing.B) {
		heap := NewMaxHeap(heapSize, compareInt, compareString)
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
