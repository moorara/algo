package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDirected(t *testing.T) {
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
		expectedInDegrees     []int
		expectedOutDegrees    []int
		expectedAdjacents     [][]int
		expectedReverse       *Directed
		traverseVerticesTests []traverseVerticesTest
		traverseEdgesTests    []traverseEdgesTest
	}{
		{
			name: "Disconnected",
			V:    13,
			edges: [][2]int{
				[2]int{0, 1},
				[2]int{0, 5},
				[2]int{2, 0},
				[2]int{2, 3},
				[2]int{3, 2},
				[2]int{3, 5},
				[2]int{4, 2},
				[2]int{4, 3},
				[2]int{5, 4},
				[2]int{6, 0},
				[2]int{6, 4},
				[2]int{6, 9},
				[2]int{7, 6},
				[2]int{7, 8},
				[2]int{8, 7},
				[2]int{8, 9},
				[2]int{9, 10},
				[2]int{9, 11},
				[2]int{10, 12},
				[2]int{11, 4},
				[2]int{11, 12},
				[2]int{12, 9},
			},
			expectedV:          13,
			expectedE:          22,
			expectedInDegrees:  []int{2, 1, 2, 2, 3, 2, 1, 1, 1, 3, 1, 1, 2},
			expectedOutDegrees: []int{2, 0, 2, 2, 2, 1, 3, 2, 2, 2, 1, 2, 1},
			expectedAdjacents: [][]int{
				[]int{1, 5},
				[]int{},
				[]int{0, 3},
				[]int{2, 5},
				[]int{2, 3},
				[]int{4},
				[]int{0, 4, 9},
				[]int{6, 8},
				[]int{7, 9},
				[]int{10, 11},
				[]int{12},
				[]int{4, 12},
				[]int{9},
			},
			expectedReverse: &Directed{
				v:   13,
				e:   22,
				ins: []int{2, 0, 2, 2, 2, 1, 3, 2, 2, 2, 1, 2, 1},
				adj: [][]int{
					[]int{2, 6},
					[]int{0},
					[]int{3, 4},
					[]int{2, 4},
					[]int{5, 6, 11},
					[]int{0, 3},
					[]int{7},
					[]int{8},
					[]int{7},
					[]int{6, 8, 12},
					[]int{9},
					[]int{9},
					[]int{10, 11},
				},
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
					name:                 "InvalidOrderDFS",
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
					expectedVertexVisits: []int{0, 1, 5, 4, 2, 3},
				},
				{
					name:                 "PostOrderDFS",
					source:               0,
					strategy:             DFS,
					order:                PostOrder,
					expectedVertexVisits: []int{1, 3, 2, 4, 5, 0},
				},
				{
					name:                 "PreOrderDFSi",
					source:               0,
					strategy:             DFSi,
					order:                PreOrder,
					expectedVertexVisits: []int{0, 1, 5, 4, 2, 3},
				},
				{
					name:                 "PostOrderDFSi",
					source:               0,
					strategy:             DFSi,
					order:                PostOrder,
					expectedVertexVisits: []int{0, 5, 4, 3, 2, 1},
				},
				{
					name:                 "PreOrderBFS",
					source:               0,
					strategy:             BFS,
					order:                PreOrder,
					expectedVertexVisits: []int{0, 1, 5, 4, 2, 3},
				},
				{
					name:                 "PostOrderBFS",
					source:               0,
					strategy:             BFS,
					order:                PostOrder,
					expectedVertexVisits: []int{0, 1, 5, 4, 2, 3},
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
						{0, 5},
						{5, 4},
						{4, 2},
						{2, 3},
					},
				},
				{
					name:     "DFSi",
					source:   0,
					strategy: DFSi,
					expectedEdgeVisits: [][2]int{
						{0, 1},
						{0, 5},
						{5, 4},
						{4, 2},
						{4, 3},
					},
				},
				{
					name:     "BFS",
					source:   0,
					strategy: BFS,
					expectedEdgeVisits: [][2]int{
						{0, 1},
						{0, 5},
						{5, 4},
						{4, 2},
						{4, 3},
					},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := NewDirected(tc.V, tc.edges...)

			assert.NotEmpty(t, g)
			assert.Equal(t, tc.expectedV, g.V())
			assert.Equal(t, tc.expectedE, g.E())

			assert.Equal(t, -1, g.InDegree(-1))
			for v, expectedInDegree := range tc.expectedInDegrees {
				assert.Equal(t, expectedInDegree, g.InDegree(v))
			}

			assert.Equal(t, -1, g.OutDegree(-1))
			for v, expectedOutDegree := range tc.expectedOutDegrees {
				assert.Equal(t, expectedOutDegree, g.OutDegree(v))
			}

			assert.Nil(t, g.Adj(-1))
			for v, expectedAdj := range tc.expectedAdjacents {
				assert.Equal(t, expectedAdj, g.Adj(v))
			}

			assert.Equal(t, tc.expectedReverse, g.Reverse())

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
