package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUndirected(t *testing.T) {
	type traverseTest struct {
		name                          string
		source                        int
		strategy                      TraversalStrategy
		expectedVertexPreOrderVisits  []int
		expectedVertexPostOrderVisits []int
		expectedEdgePreOrderVisits    [][2]int
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
					name:                          "InvalidVertex",
					source:                        -1,
					expectedVertexPreOrderVisits:  []int{},
					expectedVertexPostOrderVisits: []int{},
					expectedEdgePreOrderVisits:    [][2]int{},
				},
				{
					name:                          "InvalidStrategy",
					source:                        0,
					strategy:                      -1,
					expectedVertexPreOrderVisits:  []int{},
					expectedVertexPostOrderVisits: []int{},
					expectedEdgePreOrderVisits:    [][2]int{},
				},
				{
					name:                          "DFS",
					source:                        0,
					strategy:                      DFS,
					expectedVertexPreOrderVisits:  []int{0, 1, 2, 5, 3, 4, 6},
					expectedVertexPostOrderVisits: []int{1, 2, 6, 4, 3, 5, 0},
					expectedEdgePreOrderVisits: [][2]int{
						{0, 1},
						{0, 2},
						{0, 5},
						{5, 3},
						{3, 4},
						{4, 6},
					},
				},
				{
					name:                          "DFSi",
					source:                        0,
					strategy:                      DFSi,
					expectedVertexPreOrderVisits:  []int{0, 1, 2, 5, 6, 4, 3},
					expectedVertexPostOrderVisits: []int{0, 6, 4, 3, 5, 2, 1},
					expectedEdgePreOrderVisits: [][2]int{
						{0, 1},
						{0, 2},
						{0, 5},
						{0, 6},
						{6, 4},
						{4, 3},
					},
				},
				{
					name:                          "BFS",
					source:                        0,
					strategy:                      BFS,
					expectedVertexPreOrderVisits:  []int{0, 1, 2, 5, 6, 3, 4},
					expectedVertexPostOrderVisits: []int{0, 1, 2, 5, 6, 3, 4},
					expectedEdgePreOrderVisits: [][2]int{
						{0, 1},
						{0, 2},
						{0, 5},
						{0, 6},
						{5, 3},
						{5, 4},
					},
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

			t.Run("Traverse", func(t *testing.T) {
				for _, tc := range tc.traverseTests {
					t.Run(tc.name, func(t *testing.T) {
						tv := newTestVisitors()
						g.Traverse(tc.source, tc.strategy, tv.Visitors)
						assert.Equal(t, tc.expectedVertexPreOrderVisits, tv.preOrderVertices)
						assert.Equal(t, tc.expectedVertexPostOrderVisits, tv.postOrderVertices)
						assert.Equal(t, tc.expectedEdgePreOrderVisits, tv.preOrderEdges)
					})
				}
			})

			assert.NotEmpty(t, g.Graphviz())
		})
	}
}
