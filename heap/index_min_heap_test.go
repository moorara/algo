package heap

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/generic"
)

func TestIndexMinHeap(t *testing.T) {
	type indexKeyValue struct {
		index int
		key   int
		val   string
	}

	tests := []struct {
		name                string
		cap                 int
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
			cap:                 10,
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
			name: "FewPairs",
			cap:  10,
			insertTests: []indexKeyValue{
				{0, 30, "thirty"},
				{1, 1, "ten"},
				{2, 200, "twenty"},
			},
			changeKeyTests: []indexKeyValue{
				{index: 1, key: 10},
				{index: 2, key: 20},
			},
			expectedSize:    3,
			expectedIsEmpty: false,
			expectedPeek:    indexKeyValue{1, 10, "ten"},
			expectedPeekIndex: []indexKeyValue{
				{1, 10, "ten"},
				{2, 20, "twenty"},
				{0, 30, "thirty"},
			},
			expectedContains: []indexKeyValue{
				{1, 10, "ten"},
				{2, 20, "twenty"},
				{0, 30, "thirty"},
			},
			expectedDelete: []indexKeyValue{
				{1, 10, "ten"},
			},
			expectedDeleteIndex: []indexKeyValue{
				{0, 30, "thirty"},
				{2, 20, "twenty"},
			},
		},
		{
			name: "SomePairs",
			cap:  10,
			insertTests: []indexKeyValue{
				{0, 50, "fifty"},
				{1, 30, "thirty"},
				{2, 4, "forty"},
				{3, 10, "ten"},
				{4, 200, "twenty"},
			},
			changeKeyTests: []indexKeyValue{
				{index: 2, key: 40},
				{index: 3, key: 10},
				{index: 4, key: 20},
			},
			expectedSize:    5,
			expectedIsEmpty: false,
			expectedPeek:    indexKeyValue{3, 10, "ten"},
			expectedPeekIndex: []indexKeyValue{
				{3, 10, "ten"},
				{4, 20, "twenty"},
				{1, 30, "thirty"},
				{2, 40, "forty"},
				{0, 50, "fifty"},
			},
			expectedContains: []indexKeyValue{
				{3, 10, "ten"},
				{4, 20, "twenty"},
				{1, 30, "thirty"},
				{2, 40, "forty"},
				{0, 50, "fifty"},
			},
			expectedDelete: []indexKeyValue{
				{3, 10, "ten"},
				{4, 20, "twenty"},
			},
			expectedDeleteIndex: []indexKeyValue{
				{0, 50, "fifty"},
				{2, 40, "forty"},
				{1, 30, "thirty"},
			},
		},
		{
			name: "MorePairs",
			cap:  10,
			insertTests: []indexKeyValue{
				{0, 90, "ninety"},
				{1, 80, "eighty"},
				{2, 70, "seventy"},
				{3, 40, "forty"},
				{4, 5, "fifty"},
				{5, 6, "sixty"},
				{6, 30, "thirty"},
				{7, 100, "ten"},
				{8, 200, "twenty"},
			},
			changeKeyTests: []indexKeyValue{
				{index: 4, key: 50},
				{index: 5, key: 60},
				{index: 6, key: 30},
				{index: 7, key: 10},
				{index: 8, key: 20},
			},
			expectedSize:    9,
			expectedIsEmpty: false,
			expectedPeek:    indexKeyValue{7, 10, "ten"},
			expectedPeekIndex: []indexKeyValue{
				{7, 10, "ten"},
				{8, 20, "twenty"},
				{6, 30, "thirty"},
				{3, 40, "forty"},
				{4, 50, "fifty"},
				{5, 60, "sixty"},
				{2, 70, "seventy"},
				{1, 80, "eighty"},
				{0, 90, "ninety"},
			},
			expectedContains: []indexKeyValue{
				{7, 10, "ten"},
				{8, 20, "twenty"},
				{6, 30, "thirty"},
				{3, 40, "forty"},
				{4, 50, "fifty"},
				{5, 60, "sixty"},
				{2, 70, "seventy"},
				{1, 80, "eighty"},
				{0, 90, "ninety"},
			},
			expectedDelete: []indexKeyValue{
				{7, 10, "ten"},
				{8, 20, "twenty"},
				{6, 30, "thirty"},
				{3, 40, "forty"},
			},
			expectedDeleteIndex: []indexKeyValue{
				{0, 90, "ninety"},
				{1, 80, "eighty"},
				{2, 70, "seventy"},
				{5, 60, "sixty"},
				{4, 50, "fifty"},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cmpKey := generic.NewCompareFunc[int]()
			eqVal := generic.NewEqualFunc[string]()
			heap := NewIndexMinHeap[int, string](tc.cap, cmpKey, eqVal)

			t.Run("BeforeInsert", func(t *testing.T) {
				assert.Zero(t, heap.Size())
				assert.True(t, heap.IsEmpty())
				assert.False(t, heap.ContainsIndex(0))
				assert.False(t, heap.ContainsKey(0))
				assert.False(t, heap.ContainsValue(""))

				peekIndex, peekKey, peekVal, peekOK := heap.Peek()
				assert.Equal(t, -1, peekIndex)
				assert.Zero(t, peekKey)
				assert.Empty(t, peekVal)
				assert.False(t, peekOK)

				peekKey, peekVal, peekOK = heap.PeekIndex(0)
				assert.Equal(t, -1, peekIndex)
				assert.Zero(t, peekKey)
				assert.Empty(t, peekVal)
				assert.False(t, peekOK)

				deleteIndex, deleteKey, deleteVal, deleteOK := heap.Delete()
				assert.Equal(t, -1, deleteIndex)
				assert.Zero(t, deleteKey)
				assert.Empty(t, deleteVal)
				assert.False(t, deleteOK)

				deleteKey, deleteVal, deleteOK = heap.DeleteIndex(0)
				assert.Zero(t, deleteKey)
				assert.Empty(t, deleteVal)
				assert.False(t, deleteOK)
			})

			t.Run("AfterInsert", func(t *testing.T) {
				for _, entry := range tc.insertTests {
					heap.Insert(entry.index, entry.key, entry.val)
				}

				for _, entry := range tc.changeKeyTests {
					heap.ChangeKey(entry.index, entry.key)
				}

				assert.Equal(t, tc.expectedSize, heap.Size())
				assert.Equal(t, tc.expectedIsEmpty, heap.IsEmpty())

				peekIndex, peekKey, peekVal, peekOK := heap.Peek()
				if tc.expectedSize == 0 {
					assert.Equal(t, -1, peekIndex)
					assert.Zero(t, peekKey)
					assert.Empty(t, peekVal)
					assert.False(t, peekOK)
				} else {
					assert.Equal(t, tc.expectedPeek.index, peekIndex)
					assert.Equal(t, tc.expectedPeek.key, peekKey)
					assert.Equal(t, tc.expectedPeek.val, peekVal)
					assert.True(t, peekOK)
				}

				for _, entry := range tc.expectedPeekIndex {
					peekKey, peekVal, peekOK = heap.PeekIndex(entry.index)
					assert.Equal(t, entry.key, peekKey)
					assert.Equal(t, entry.val, peekVal)
					assert.True(t, peekOK)
				}

				for _, entry := range tc.expectedContains {
					assert.True(t, heap.ContainsIndex(entry.index))
					assert.True(t, heap.ContainsKey(entry.key))
					assert.True(t, heap.ContainsValue(entry.val))
				}

				for _, entry := range tc.expectedDelete {
					deleteIndex, deleteKey, deleteVal, deleteOK := heap.Delete()
					assert.Equal(t, entry.index, deleteIndex)
					assert.Equal(t, entry.key, deleteKey)
					assert.Equal(t, entry.val, deleteVal)
					assert.True(t, deleteOK)
				}

				for _, entry := range tc.expectedDeleteIndex {
					deleteKey, deleteVal, deleteOK := heap.DeleteIndex(entry.index)
					assert.Equal(t, entry.key, deleteKey)
					assert.Equal(t, entry.val, deleteVal)
					assert.True(t, deleteOK)
				}
			})

			t.Run("AfterDelete", func(t *testing.T) {
				assert.Zero(t, heap.Size())
				assert.True(t, heap.IsEmpty())
				assert.False(t, heap.ContainsKey(0))
				assert.False(t, heap.ContainsValue(""))

				peekIndex, peekKey, peekVal, peekOK := heap.Peek()
				assert.Equal(t, -1, peekIndex)
				assert.Zero(t, peekKey)
				assert.Empty(t, peekVal)
				assert.False(t, peekOK)

				peekKey, peekVal, peekOK = heap.PeekIndex(0)
				assert.Equal(t, -1, peekIndex)
				assert.Zero(t, peekKey)
				assert.Empty(t, peekVal)
				assert.False(t, peekOK)

				deleteIndex, deleteKey, deleteVal, deleteOK := heap.Delete()
				assert.Equal(t, -1, deleteIndex)
				assert.Zero(t, deleteKey)
				assert.Empty(t, deleteVal)
				assert.False(t, deleteOK)

				deleteKey, deleteVal, deleteOK = heap.DeleteIndex(0)
				assert.Zero(t, deleteKey)
				assert.Empty(t, deleteVal)
				assert.False(t, deleteOK)
			})
		})
	}
}

func BenchmarkIndexMinHeap(b *testing.B) {
	rand.Seed(time.Now().UTC().UnixNano())

	b.Run("Insert", func(b *testing.B) {
		cmpKey := generic.NewCompareFunc[int]()
		heap := NewIndexMinHeap[int, string](b.N, cmpKey, nil)

		keys := randIntSlice(b.N)
		vals := randStringSlice(b.N)

		b.ResetTimer()

		for n := 0; n < b.N; n++ {
			heap.Insert(n, keys[n], vals[n])
		}
	})

	b.Run("Delete", func(b *testing.B) {
		cmpKey := generic.NewCompareFunc[int]()
		heap := NewIndexMinHeap[int, string](b.N, cmpKey, nil)

		keys := randIntSlice(b.N)
		vals := randStringSlice(b.N)

		for n := 0; n < b.N; n++ {
			heap.Insert(n, keys[n], vals[n])
		}

		b.ResetTimer()

		for n := 0; n < b.N; n++ {
			heap.Delete()
		}
	})
}
