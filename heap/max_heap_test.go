package heap

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/common"
)

func TestMaxHeap(t *testing.T) {
	type keyValue struct {
		key int
		val string
	}

	tests := []struct {
		name             string
		size             int
		insertTests      []keyValue
		expectedSize     int
		expectedIsEmpty  bool
		expectedPeek     keyValue
		expectedContains []keyValue
		expectedDelete   []keyValue
	}{
		{
			name:             "Empty",
			size:             2,
			insertTests:      []keyValue{},
			expectedSize:     0,
			expectedIsEmpty:  true,
			expectedPeek:     keyValue{0, ""},
			expectedContains: []keyValue{},
			expectedDelete:   []keyValue{},
		},
		{
			name: "FewPairs",
			size: 2,
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
			name: "SomePairs",
			size: 4,
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
			name: "MorePairs",
			size: 4,
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
			cmpKey := common.NewCompareFunc[int]()
			eqVal := common.NewEqualFunc[string]()
			heap := NewMaxHeap[int, string](tc.size, cmpKey, eqVal)

			t.Run("BeforeInsert", func(t *testing.T) {
				assert.Zero(t, heap.Size())
				assert.True(t, heap.IsEmpty())
				assert.False(t, heap.ContainsKey(0))
				assert.False(t, heap.ContainsValue(""))

				peekKey, peekVal, peekOK := heap.Peek()
				assert.Zero(t, peekKey)
				assert.Empty(t, peekVal)
				assert.False(t, peekOK)

				deleteKey, deleteVal, deleteOK := heap.Delete()
				assert.Zero(t, deleteKey)
				assert.Empty(t, deleteVal)
				assert.False(t, deleteOK)
			})

			t.Run("AfterInsert", func(t *testing.T) {
				for _, entry := range tc.insertTests {
					heap.Insert(entry.key, entry.val)
				}

				assert.Equal(t, tc.expectedSize, heap.Size())
				assert.Equal(t, tc.expectedIsEmpty, heap.IsEmpty())

				peekKey, peekVal, peekOK := heap.Peek()
				if tc.expectedSize == 0 {
					assert.Zero(t, peekKey)
					assert.Empty(t, peekVal)
					assert.False(t, peekOK)
				} else {
					assert.Equal(t, tc.expectedPeek.key, peekKey)
					assert.Equal(t, tc.expectedPeek.val, peekVal)
					assert.True(t, peekOK)
				}

				for _, entry := range tc.expectedContains {
					assert.True(t, heap.ContainsKey(entry.key))
					assert.True(t, heap.ContainsValue(entry.val))
				}

				for _, entry := range tc.expectedDelete {
					deleteKey, deleteVal, deleteOK := heap.Delete()
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

				peekKey, peekVal, peekOK := heap.Peek()
				assert.Zero(t, peekKey)
				assert.Empty(t, peekVal)
				assert.False(t, peekOK)

				deleteKey, deleteVal, deleteOK := heap.Delete()
				assert.Zero(t, deleteKey)
				assert.Empty(t, deleteVal)
				assert.False(t, peekOK, deleteOK)
			})
		})
	}
}

func BenchmarkMaxHeap(b *testing.B) {
	heapSize := 1024
	minInt := 0
	maxInt := 1000000

	rand.Seed(time.Now().UTC().UnixNano())

	b.Run("Insert", func(b *testing.B) {
		cmpKey := common.NewCompareFunc[int]()
		heap := NewMaxHeap[int, string](heapSize, cmpKey, nil)

		keys := randIntSlice(b.N, minInt, maxInt)
		vals := randStringSlice(b.N)

		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			heap.Insert(keys[n], vals[n])
		}
	})

	b.Run("Delete", func(b *testing.B) {
		cmpKey := common.NewCompareFunc[int]()
		heap := NewMaxHeap[int, string](heapSize, cmpKey, nil)

		keys := randIntSlice(b.N, minInt, maxInt)
		vals := randStringSlice(b.N)

		for n := 0; n < b.N; n++ {
			heap.Insert(keys[n], vals[n])
		}

		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			heap.Delete()
		}
	})
}
