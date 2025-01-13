package list

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/generic"
)

func TestSoftQueue(t *testing.T) {
	tests := []struct {
		name             string
		nodeSize         int
		enqueueValues    []string
		dequeuesCount    int
		expectedSize     int
		expectedIsEmpty  bool
		expectedPeek     string
		containsTests    []containsTest[string]
		expectedDequeues []string
		expectedValues   []string
	}{
		{
			name:             "Empty",
			nodeSize:         2,
			enqueueValues:    []string{},
			dequeuesCount:    0,
			expectedSize:     0,
			expectedIsEmpty:  true,
			expectedPeek:     "",
			containsTests:    []containsTest[string]{},
			expectedDequeues: []string{},
			expectedValues:   []string{},
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
			expectedDequeues: []string{"a", "b"},
			expectedValues:   []string{"a", "b"},
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
				{"a", true},
				{"b", true},
				{"c", true},
			},
			expectedDequeues: []string{"b", "c"},
			expectedValues:   []string{"a", "b", "c"},
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
				{"a", true},
				{"b", true},
				{"c", true},
				{"d", true},
				{"e", true},
				{"f", true},
				{"g", true},
			},
			expectedDequeues: []string{"c", "d", "e", "f", "g"},
			expectedValues:   []string{"a", "b", "c", "d", "e", "f", "g"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			equal := generic.NewEqualFunc[string]()
			squeue := NewSoftQueue[string](equal)

			t.Run("BeforeEnqueue", func(t *testing.T) {
				assert.Zero(t, squeue.Size())
				assert.True(t, squeue.IsEmpty())

				val, i := squeue.Dequeue()
				assert.Empty(t, val)
				assert.Equal(t, -1, i)

				val, i = squeue.Peek()
				assert.Empty(t, val)
				assert.Equal(t, -1, i)

				assert.Equal(t, -1, squeue.Contains(""))
			})

			t.Run("AfterEnqueue", func(t *testing.T) {
				for _, val := range tc.enqueueValues {
					squeue.Enqueue(val)
				}

				for i := 0; i < tc.dequeuesCount; i++ {
					squeue.Dequeue()
				}

				assert.Equal(t, tc.expectedSize, squeue.Size())
				assert.Equal(t, tc.expectedIsEmpty, squeue.IsEmpty())

				val, i := squeue.Peek()

				if tc.expectedSize == 0 {
					assert.Empty(t, val)
					assert.Equal(t, -1, i)
				} else {
					assert.Equal(t, tc.expectedPeek, val)
					assert.True(t, i >= 0)
				}

				for _, tc := range tc.containsTests {
					i := squeue.Contains(tc.val)
					assert.Equal(t, tc.expected, i >= 0)
				}

				for _, val := range tc.expectedDequeues {
					v, i := squeue.Dequeue()
					assert.Equal(t, val, v)
					assert.True(t, i >= 0)
				}

				vals := squeue.Values()
				assert.Equal(t, tc.expectedValues, vals)
			})

			t.Run("AfterDequeue", func(t *testing.T) {
				assert.Zero(t, squeue.Size())
				assert.True(t, squeue.IsEmpty())

				val, i := squeue.Dequeue()
				assert.Empty(t, val)
				assert.Equal(t, -1, i)

				val, i = squeue.Peek()
				assert.Empty(t, val)
				assert.Equal(t, -1, i)

				assert.Equal(t, -1, squeue.Contains(""))
			})
		})
	}
}

func BenchmarkSoftQueue(b *testing.B) {
	seed := time.Now().UTC().UnixNano()
	r := rand.New(rand.NewSource(seed))

	b.Run("Enqueue", func(b *testing.B) {
		squeue := NewSoftQueue[int](nil)

		b.ResetTimer()

		for n := 0; n < b.N; n++ {
			squeue.Enqueue(r.Int())
		}
	})

	b.Run("Dequeue", func(b *testing.B) {
		squeue := NewSoftQueue[int](nil)
		for n := 0; n < b.N; n++ {
			squeue.Enqueue(r.Int())
		}

		b.ResetTimer()

		for n := 0; n < b.N; n++ {
			squeue.Dequeue()
		}
	})
}
