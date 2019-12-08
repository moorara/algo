package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGraph(t *testing.T) {
	tests := []struct {
		name              string
		V                 int
		edges             [][2]int
		expectedV         int
		expectedE         int
		expectedDegrees   []int
		expectedAdjacents [][]int
	}{
		{
			name: "Disconnected",
			V:    13,
			edges: [][2]int{
				[2]int{0, 1},
				[2]int{0, 2},
				[2]int{0, 5},
				[2]int{0, 6},
				[2]int{3, 4},
				[2]int{3, 5},
				[2]int{4, 5},
				[2]int{4, 6},
				[2]int{7, 8},
				[2]int{9, 10},
				[2]int{9, 11},
				[2]int{9, 12},
				[2]int{11, 12},
			},
			expectedV:       13,
			expectedE:       13,
			expectedDegrees: []int{4, 1, 1, 2, 3, 3, 2, 1, 1, 3, 1, 2, 2},
			expectedAdjacents: [][]int{
				[]int{1, 2, 5, 6},
				[]int{0},
				[]int{0},
				[]int{4, 5},
				[]int{3, 5, 6},
				[]int{0, 3, 4},
				[]int{0, 4},
				[]int{8},
				[]int{7},
				[]int{10, 11, 12},
				[]int{9},
				[]int{9, 12},
				[]int{9, 11},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := NewGraph(tc.V, tc.edges...)

			assert.NotEmpty(t, g)
			assert.Equal(t, tc.expectedV, g.V())
			assert.Equal(t, tc.expectedE, g.E())

			assert.Equal(t, -1, g.Degree(-1))
			for v, expectedDegree := range tc.expectedDegrees {
				assert.Equal(t, expectedDegree, g.Degree(v))
			}

			assert.Nil(t, g.Adj(-1))
			for v, expectedAdj := range tc.expectedAdjacents {
				assert.Equal(t, expectedAdj, g.Adj(v))
			}

			assert.NotEmpty(t, g.Graphviz())
		})
	}
}
