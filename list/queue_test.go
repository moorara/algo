package list

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/common"
)

func TestQueue(t *testing.T) {
	tests := []struct {
		name                  string
		nodeSize              int
		enqueueValues         []string
		expectedSize          int
		expectedIsEmpty       bool
		expectedPeek          string
		expectedContains      []string
		expectedDequeueValues []string
	}{
		{
			"Empty",
			2,
			[]string{},
			0, true,
			"",
			[]string{},
			[]string{},
		},
		{
			"OneNode",
			2,
			[]string{"a", "b"},
			2, false,
			"a",
			[]string{"a", "b"},
			[]string{"a", "b"},
		},
		{
			"TwoNodes",
			2,
			[]string{"a", "b", "c"},
			3, false,
			"a",
			[]string{"a", "b", "c"},
			[]string{"a", "b", "c"},
		},
		{
			"MoreNodes",
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
			equal := common.NewEqualFunc[string]()
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

				for _, val := range tc.expectedContains {
					assert.True(t, queue.Contains(val))
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
