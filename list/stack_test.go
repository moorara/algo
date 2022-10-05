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
		expectedSize      int
		expectedIsEmpty   bool
		expectedPeek      string
		expectedContains  []string
		expectedPopValues []string
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
			"b",
			[]string{"a", "b"},
			[]string{"b", "a"},
		},
		{
			"TwoNodes",
			2,
			[]string{"a", "b", "c"},
			3, false,
			"c",
			[]string{"a", "b", "c"},
			[]string{"c", "b", "a"},
		},
		{
			"MoreNodes",
			2,
			[]string{"a", "b", "c", "d", "e", "f", "g"},
			7, false,
			"g",
			[]string{"a", "b", "c", "d", "e", "f", "g"},
			[]string{"g", "f", "e", "d", "c", "b", "a"},
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

				for _, val := range tc.expectedContains {
					assert.True(t, stack.Contains(val))
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
