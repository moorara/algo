package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type containsTest[T any] struct {
	val      T
	expected bool
}

func TestNewArrayNode(t *testing.T) {
	tests := []struct {
		size int
		next *arrayNode[string]
	}{
		{64, nil},
		{256, nil},
		{1024, &arrayNode[string]{}},
		{4096, &arrayNode[string]{}},
	}

	for _, tc := range tests {
		n := newArrayNode(tc.size, tc.next)

		assert.Equal(t, tc.next, n.next)
		assert.Equal(t, tc.size, len(n.block))
	}
}
