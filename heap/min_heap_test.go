package heap

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/common"
)

func TestMinHeap(t *testing.T) {
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
			name: "SomePairs",
			size: 4,
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
			name: "MorePairs",
			size: 4,
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
			cmpKey := common.NewCompareFunc[int]()
			eqVal := common.NewEqualFunc[string]()
			heap := NewMinHeap[int, string](tc.size, cmpKey, eqVal)

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
				assert.False(t, deleteOK)
			})
		})
	}
}

func BenchmarkMinHeap(b *testing.B) {
	const heapSize = 1024

	rand.Seed(time.Now().UTC().UnixNano())

	b.Run("Insert", func(b *testing.B) {
		cmpKey := common.NewCompareFunc[int]()
		heap := NewMinHeap[int, string](heapSize, cmpKey, nil)

		keys := randIntSlice(b.N)
		vals := randStringSlice(b.N)

		b.ResetTimer()

		for n := 0; n < b.N; n++ {
			heap.Insert(keys[n], vals[n])
		}
	})

	b.Run("Delete", func(b *testing.B) {
		cmpKey := common.NewCompareFunc[int]()
		heap := NewMinHeap[int, string](heapSize, cmpKey, nil)

		keys := randIntSlice(b.N)
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
