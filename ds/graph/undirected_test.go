package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUndirected(t *testing.T) {
	type traverseVerticesTest struct {
		name                 string
		source               int
		strategy             TraversalStrategy
		order                TraversalOrder
		expectedVertexVisits []int
	}

	type traverseEdgesTest struct {
		name               string
		source             int
		strategy           TraversalStrategy
		expectedEdgeVisits [][2]int
	}

	tests := []struct {
		name                  string
		V                     int
		edges                 [][2]int
		expectedV             int
		expectedE             int
		expectedDegrees       []int
		expectedAdjacents     [][]int
		traverseVerticesTests []traverseVerticesTest
		traverseEdgesTests    []traverseEdgesTest
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
			traverseVerticesTests: []traverseVerticesTest{
				{
					name:                 "InvalidVertex",
					source:               -1,
					expectedVertexVisits: []int{},
				},
				{
					name:                 "InvalidStrategy",
					source:               0,
					strategy:             -1,
					expectedVertexVisits: []int{},
				},
				{
					name:                 "InvalidOrder",
					source:               0,
					strategy:             DFS,
					order:                -1,
					expectedVertexVisits: []int{},
				},
				{
					name:                 "PreOrderDFS",
					source:               0,
					strategy:             DFS,
					order:                PreOrder,
					expectedVertexVisits: []int{0, 1, 2, 5, 3, 4, 6},
				},
				{
					name:                 "PostOrderDFS",
					source:               0,
					strategy:             DFS,
					order:                PostOrder,
					expectedVertexVisits: []int{1, 2, 6, 4, 3, 5, 0},
				},
				{
					name:                 "PreOrderDFSi",
					source:               0,
					strategy:             DFSi,
					order:                PreOrder,
					expectedVertexVisits: []int{0, 1, 2, 5, 6, 4, 3},
				},
				{
					name:                 "PostOrderDFSi",
					source:               0,
					strategy:             DFSi,
					order:                PostOrder,
					expectedVertexVisits: []int{0, 6, 4, 3, 5, 2, 1},
				},
				{
					name:                 "PreOrderBFS",
					source:               0,
					strategy:             BFS,
					order:                PreOrder,
					expectedVertexVisits: []int{0, 1, 2, 5, 6, 3, 4},
				},
				{
					name:                 "PostOrderBFS",
					source:               0,
					strategy:             BFS,
					order:                PostOrder,
					expectedVertexVisits: []int{0, 1, 2, 5, 6, 3, 4},
				},
			},
			traverseEdgesTests: []traverseEdgesTest{
				{
					name:               "InvalidVertex",
					source:             -1,
					expectedEdgeVisits: [][2]int{},
				},
				{
					name:               "InvalidStrategy",
					source:             0,
					strategy:           -1,
					expectedEdgeVisits: [][2]int{},
				},
				{
					name:     "DFS",
					source:   0,
					strategy: DFS,
					expectedEdgeVisits: [][2]int{
						{0, 1},
						{0, 2},
						{0, 5},
						{5, 3},
						{3, 4},
						{4, 6},
					},
				},
				{
					name:     "DFSi",
					source:   0,
					strategy: DFSi,
					expectedEdgeVisits: [][2]int{
						{0, 1},
						{0, 2},
						{0, 5},
						{0, 6},
						{6, 4},
						{4, 3},
					},
				},
				{
					name:     "BFS",
					source:   0,
					strategy: BFS,
					expectedEdgeVisits: [][2]int{
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

			t.Run("TraverseVertices", func(t *testing.T) {
				for _, tc := range tc.traverseVerticesTests {
					t.Run(tc.name, func(t *testing.T) {
						visitor := newVisitor()
						g.TraverseVertices(tc.source, tc.strategy, tc.order, visitor)
						assert.Equal(t, tc.expectedVertexVisits, visitor.vertices)
					})
				}
			})

			t.Run("TraverseEdges", func(t *testing.T) {
				for _, tc := range tc.traverseEdgesTests {
					t.Run(tc.name, func(t *testing.T) {
						visitor := newVisitor()
						g.TraverseEdges(tc.source, tc.strategy, visitor)
						assert.Equal(t, tc.expectedEdgeVisits, visitor.edges)
					})
				}
			})

			assert.NotEmpty(t, g.Graphviz())
		})
	}
}
