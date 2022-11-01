package list

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/generic"
)

func TestStack(t *testing.T) {
	tests := []struct {
		name              string
		nodeSize          int
		pushValues        []string
		popsCount         int
		expectedSize      int
		expectedIsEmpty   bool
		expectedPeek      string
		containsTests     []containsTest[string]
		expectedPopValues []string
	}{
		{
			name:              "Empty",
			nodeSize:          2,
			pushValues:        []string{},
			popsCount:         0,
			expectedSize:      0,
			expectedIsEmpty:   true,
			expectedPeek:      "",
			containsTests:     []containsTest[string]{},
			expectedPopValues: []string{},
		},
		{
			name:            "OneNode",
			nodeSize:        2,
			pushValues:      []string{"a", "b"},
			popsCount:       0,
			expectedSize:    2,
			expectedIsEmpty: false,
			expectedPeek:    "b",
			containsTests: []containsTest[string]{
				{"a", true},
				{"b", true},
				{"c", false},
			},
			expectedPopValues: []string{"b", "a"},
		},
		{
			name:            "TwoNodes",
			nodeSize:        2,
			pushValues:      []string{"a", "b", "c"},
			popsCount:       1,
			expectedSize:    2,
			expectedIsEmpty: false,
			expectedPeek:    "b",
			containsTests: []containsTest[string]{
				{"a", true},
				{"b", true},
				{"c", false},
			},
			expectedPopValues: []string{"b", "a"},
		},
		{
			name:            "MoreNodes",
			nodeSize:        2,
			pushValues:      []string{"a", "b", "c", "d", "e", "f", "g"},
			popsCount:       2,
			expectedSize:    5,
			expectedIsEmpty: false,
			expectedPeek:    "e",
			containsTests: []containsTest[string]{
				{"a", true},
				{"b", true},
				{"c", true},
				{"d", true},
				{"e", true},
				{"f", false},
				{"g", false},
			},
			expectedPopValues: []string{"e", "d", "c", "b", "a"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			equal := generic.NewEqualFunc[string]()
			stack := NewStack[string](tc.nodeSize, equal)

			t.Run("BeforePush", func(t *testing.T) {
				assert.Zero(t, stack.Size())
				assert.True(t, stack.IsEmpty())

				val, ok := stack.Pop()
				assert.False(t, ok)
				assert.Empty(t, val)

				val, ok = stack.Peek()
				assert.False(t, ok)
				assert.Empty(t, val)

				assert.False(t, stack.Contains(""))
			})

			t.Run("AfterPush", func(t *testing.T) {
				for _, val := range tc.pushValues {
					stack.Push(val)
				}

				for i := 0; i < tc.popsCount; i++ {
					stack.Pop()
				}

				assert.Equal(t, tc.expectedSize, stack.Size())
				assert.Equal(t, tc.expectedIsEmpty, stack.IsEmpty())

				val, ok := stack.Peek()

				if tc.expectedSize == 0 {
					assert.False(t, ok)
					assert.Empty(t, val)
				} else {
					assert.True(t, ok)
					assert.Equal(t, tc.expectedPeek, val)
				}

				for _, tc := range tc.containsTests {
					assert.Equal(t, tc.expected, stack.Contains(tc.val))
				}

				for _, val := range tc.expectedPopValues {
					v, ok := stack.Pop()
					assert.True(t, ok)
					assert.Equal(t, val, v)
				}
			})

			t.Run("AfterPop", func(t *testing.T) {
				assert.Zero(t, stack.Size())
				assert.True(t, stack.IsEmpty())

				val, ok := stack.Pop()
				assert.False(t, ok)
				assert.Empty(t, val)

				val, ok = stack.Peek()
				assert.False(t, ok)
				assert.Empty(t, val)

				assert.False(t, stack.Contains(""))
			})
		})
	}
}

func BenchmarkStack(b *testing.B) {
	const nodeSize = 1024

	rand.Seed(time.Now().UTC().UnixNano())

	b.Run("Push", func(b *testing.B) {
		stack := NewStack[int](nodeSize, nil)

		b.ResetTimer()

		for n := 0; n < b.N; n++ {
			stack.Push(rand.Int())
		}
	})

	b.Run("Pop", func(b *testing.B) {
		stack := NewStack[int](nodeSize, nil)
		for n := 0; n < b.N; n++ {
			stack.Push(rand.Int())
		}

		b.ResetTimer()

		for n := 0; n < b.N; n++ {
			stack.Pop()
		}
	})
}
