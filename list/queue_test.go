package list

import (
	"testing"

	"github.com/moorara/algo/compare"
	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	tests := []struct {
		name                 string
		cmp                  compare.Func
		nodeSize             int
		enqueueItems         []string
		expectedSize         int
		expectedIsEmpty      bool
		expectedPeek         string
		expectedContains     []string
		expectedDequeueItems []string
	}{
		{
			"Empty",
			compare.String,
			2,
			[]string{},
			0, true,
			"",
			[]string{},
			[]string{},
		},
		{
			"OneNode",
			compare.String,
			2,
			[]string{"a", "b"},
			2, false,
			"a",
			[]string{"a", "b"},
			[]string{"a", "b"},
		},
		{
			"TwoNodes",
			compare.String,
			2,
			[]string{"a", "b", "c"},
			3, false,
			"a",
			[]string{"a", "b", "c"},
			[]string{"a", "b", "c"},
		},
		{
			"MoreNodes",
			compare.String,
			2,
			[]string{"a", "b", "c", "d", "e", "f", "g"},
			7, false,
			"a",
			[]string{"a", "b", "c", "d", "e", "f", "g"},
			[]string{"a", "b", "c", "d", "e", "f", "g"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			queue := NewQueue(tc.nodeSize)

			// Queue initially should be empty
			assert.Zero(t, queue.Size())
			assert.True(t, queue.IsEmpty())
			assert.Nil(t, queue.Peek())
			queue.Contains(nil, tc.cmp)
			assert.Nil(t, queue.Dequeue())

			for _, item := range tc.enqueueItems {
				queue.Enqueue(item)
			}

			assert.Equal(t, tc.expectedSize, queue.Size())
			assert.Equal(t, tc.expectedIsEmpty, queue.IsEmpty())

			if tc.expectedSize == 0 {
				assert.Nil(t, queue.Peek())
			} else {
				assert.Equal(t, tc.expectedPeek, queue.Peek())
			}

			for _, item := range tc.expectedContains {
				assert.True(t, queue.Contains(item, tc.cmp))
			}

			for _, item := range tc.expectedDequeueItems {
				assert.Equal(t, item, queue.Dequeue())
			}

			// Queue should be empty at the end
			assert.Zero(t, queue.Size())
			assert.True(t, queue.IsEmpty())
			assert.Nil(t, queue.Peek())
			queue.Contains(nil, tc.cmp)
			assert.Nil(t, queue.Dequeue())
		})
	}
}

func BenchmarkQueue(b *testing.B) {
	nodeSize := 1024
	item := 27

	b.Run("Enqueue", func(b *testing.B) {
		queue := NewQueue(nodeSize)
		b.ResetTimer()

		for n := 0; n < b.N; n++ {
			queue.Enqueue(item)
		}
	})

	b.Run("Dequeue", func(b *testing.B) {
		queue := NewQueue(nodeSize)
		for n := 0; n < b.N; n++ {
			queue.Enqueue(item)
		}
		b.ResetTimer()

		for n := 0; n < b.N; n++ {
			queue.Dequeue()
		}
	})
}
