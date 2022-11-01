package list

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/generic"
)

func TestQueue(t *testing.T) {
	tests := []struct {
		name                  string
		nodeSize              int
		enqueueValues         []string
		dequeuesCount         int
		expectedSize          int
		expectedIsEmpty       bool
		expectedPeek          string
		containsTests         []containsTest[string]
		expectedDequeueValues []string
	}{
		{
			name:                  "Empty",
			nodeSize:              2,
			enqueueValues:         []string{},
			dequeuesCount:         0,
			expectedSize:          0,
			expectedIsEmpty:       true,
			expectedPeek:          "",
			containsTests:         []containsTest[string]{},
			expectedDequeueValues: []string{},
		},
		{
			name:            "OneNode",
			nodeSize:        2,
			enqueueValues:   []string{"a", "b"},
			dequeuesCount:   0,
			expectedSize:    2,
			expectedIsEmpty: false,
			expectedPeek:    "a",
			containsTests: []containsTest[string]{
				{"a", true},
				{"b", true},
				{"c", false},
			},
			expectedDequeueValues: []string{"a", "b"},
		},
		{
			name:            "TwoNodes",
			nodeSize:        2,
			enqueueValues:   []string{"a", "b", "c"},
			dequeuesCount:   1,
			expectedSize:    2,
			expectedIsEmpty: false,
			expectedPeek:    "b",
			containsTests: []containsTest[string]{
				{"a", false},
				{"b", true},
				{"c", true},
			},
			expectedDequeueValues: []string{"b", "c"},
		},
		{
			name:            "MoreNodes",
			nodeSize:        2,
			enqueueValues:   []string{"a", "b", "c", "d", "e", "f", "g"},
			dequeuesCount:   2,
			expectedSize:    5,
			expectedIsEmpty: false,
			expectedPeek:    "c",
			containsTests: []containsTest[string]{
				{"a", false},
				{"b", false},
				{"c", true},
				{"d", true},
				{"e", true},
				{"f", true},
				{"g", true},
			},
			expectedDequeueValues: []string{"c", "d", "e", "f", "g"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			equal := generic.NewEqualFunc[string]()
			queue := NewQueue[string](tc.nodeSize, equal)

			t.Run("BeforeEnqueue", func(t *testing.T) {
				assert.Zero(t, queue.Size())
				assert.True(t, queue.IsEmpty())

				val, ok := queue.Dequeue()
				assert.False(t, ok)
				assert.Empty(t, val)

				val, ok = queue.Peek()
				assert.False(t, ok)
				assert.Empty(t, val)

				assert.False(t, queue.Contains(""))
			})

			t.Run("AfterEnqueue", func(t *testing.T) {
				for _, val := range tc.enqueueValues {
					queue.Enqueue(val)
				}

				for i := 0; i < tc.dequeuesCount; i++ {
					queue.Dequeue()
				}

				assert.Equal(t, tc.expectedSize, queue.Size())
				assert.Equal(t, tc.expectedIsEmpty, queue.IsEmpty())

				val, ok := queue.Peek()

				if tc.expectedSize == 0 {
					assert.False(t, ok)
					assert.Empty(t, val)
				} else {
					assert.True(t, ok)
					assert.Equal(t, tc.expectedPeek, val)
				}

				for _, tc := range tc.containsTests {
					assert.Equal(t, tc.expected, queue.Contains(tc.val))
				}

				for _, val := range tc.expectedDequeueValues {
					v, ok := queue.Dequeue()
					assert.True(t, ok)
					assert.Equal(t, val, v)
				}
			})

			t.Run("AfterDequeue", func(t *testing.T) {
				assert.Zero(t, queue.Size())
				assert.True(t, queue.IsEmpty())

				val, ok := queue.Dequeue()
				assert.False(t, ok)
				assert.Empty(t, val)

				val, ok = queue.Peek()
				assert.False(t, ok)
				assert.Empty(t, val)

				assert.False(t, queue.Contains(""))
			})
		})
	}
}

func BenchmarkQueue(b *testing.B) {
	const nodeSize = 1024

	rand.Seed(time.Now().UTC().UnixNano())

	b.Run("Enqueue", func(b *testing.B) {
		queue := NewQueue[int](nodeSize, nil)

		b.ResetTimer()

		for n := 0; n < b.N; n++ {
			queue.Enqueue(rand.Int())
		}
	})

	b.Run("Dequeue", func(b *testing.B) {
		queue := NewQueue[int](nodeSize, nil)
		for n := 0; n < b.N; n++ {
			queue.Enqueue(rand.Int())
		}

		b.ResetTimer()

		for n := 0; n < b.N; n++ {
			queue.Dequeue()
		}
	})
}
