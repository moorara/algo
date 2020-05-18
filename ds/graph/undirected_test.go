package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUndirected(t *testing.T) {
	type traverseTest struct {
		name           string
		start          int
		strategy       TraverseStrategy
		order          TraverseOrder
		expectedVisits []int
	}

	tests := []struct {
		name              string
		V                 int
		edges             [][2]int
		expectedV         int
		expectedE         int
		expectedDegrees   []int
		expectedAdjacents [][]int
		traverseTests     []traverseTest
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
			traverseTests: []traverseTest{
				{
					name:           "InvalidVertex",
					start:          -1,
					expectedVisits: []int{},
				},
				{
					name:           "InvalidStrategy",
					start:          0,
					strategy:       -1,
					expectedVisits: []int{},
				},
				{
					name:           "InvalidOrder",
					start:          0,
					strategy:       DFS,
					order:          -1,
					expectedVisits: []int{},
				},
				{
					name:           "PreOrderDFS",
					start:          0,
					strategy:       DFS,
					order:          PreOrder,
					expectedVisits: []int{0, 1, 2, 5, 3, 4, 6},
				},
				{
					name:           "PostOrderDFS",
					start:          0,
					strategy:       DFS,
					order:          PostOrder,
					expectedVisits: []int{1, 2, 6, 4, 3, 5, 0},
				},
				{
					name:           "PreOrderDFSIterative",
					start:          0,
					strategy:       DFSIterative,
					order:          PreOrder,
					expectedVisits: []int{0, 1, 2, 5, 6, 4, 3},
				},
				{
					name:           "PostOrderDFSIterative",
					start:          0,
					strategy:       DFSIterative,
					order:          PostOrder,
					expectedVisits: []int{0, 6, 4, 3, 5, 2, 1},
				},
				{
					name:           "PreOrderBFS",
					start:          0,
					strategy:       BFS,
					order:          PreOrder,
					expectedVisits: []int{0, 1, 2, 5, 6, 3, 4},
				},
				{
					name:           "PostOrderBFS",
					start:          0,
					strategy:       BFS,
					order:          PostOrder,
					expectedVisits: []int{0, 1, 2, 5, 6, 3, 4},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := NewUndirected(tc.V, tc.edges...)

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

			for _, traverse := range tc.traverseTests {
				t.Run(traverse.name, func(t *testing.T) {
					visited := make([]int, 0)
					g.Traverse(traverse.start, traverse.strategy, traverse.order, &Visitor{
						VisitVertex: func(v int) {
							visited = append(visited, v)
						},
					})
					assert.Equal(t, traverse.expectedVisits, visited)
				})
			}

			assert.NotEmpty(t, g.Graphviz())
		})
	}
}
