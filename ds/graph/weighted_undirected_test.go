package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUndirectedEdge(t *testing.T) {
	tests := []struct {
		name   string
		v, w   int
		weight float64
	}{
		{"Edge1", 1, 2, 0.27},
		{"Edge2", 3, 4, 0.69},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			e := UndirectedEdge{tc.v, tc.w, tc.weight}

			assert.NotEmpty(t, e)
			assert.Equal(t, tc.v, e.Either())
			assert.Equal(t, tc.v, e.Other(tc.w))
			assert.Equal(t, tc.w, e.Other(tc.v))
			assert.Equal(t, tc.weight, e.Weight())
		})
	}
}

func TestWeightedGraph(t *testing.T) {
	type traverseVerticesTest struct {
		name                 string
		source               int
		strategy             TraversalStrategy
		order                TraversalOrder
		expectedVertexVisits []int
	}

	type traverseEdgesTest struct {
		name                 string
		source               int
		strategy             TraversalStrategy
		expectedEdgeVisits   [][2]int
		expectedWeightVisits []float64
	}

	tests := []struct {
		name                  string
		V                     int
		edges                 []UndirectedEdge
		expectedV             int
		expectedE             int
		expectedDegrees       []int
		expectedAdjacents     [][]UndirectedEdge
		expectedEdges         []UndirectedEdge
		traverseVerticesTests []traverseVerticesTest
		traverseEdgesTests    []traverseEdgesTest
	}{
		{
			name: "Connected",
			V:    8,
			edges: []UndirectedEdge{
				{0, 2, 0.26},
				{0, 4, 0.38},
				{0, 7, 0.16},
				{1, 2, 0.36},
				{1, 3, 0.29},
				{1, 5, 0.32},
				{1, 7, 0.19},
				{2, 3, 0.17},
				{2, 7, 0.34},
				{3, 6, 0.52},
				{4, 5, 0.35},
				{4, 6, 0.93},
				{4, 7, 0.37},
				{5, 7, 0.28},
				{6, 0, 0.58},
				{6, 2, 0.40},
			},
			expectedV:       8,
			expectedE:       16,
			expectedDegrees: []int{4, 4, 5, 3, 4, 3, 4},
			expectedAdjacents: [][]UndirectedEdge{
				[]UndirectedEdge{
					{0, 2, 0.26},
					{0, 4, 0.38},
					{0, 7, 0.16},
					{6, 0, 0.58},
				},
				[]UndirectedEdge{
					{1, 2, 0.36},
					{1, 3, 0.29},
					{1, 5, 0.32},
					{1, 7, 0.19},
				},
				[]UndirectedEdge{
					{0, 2, 0.26},
					{1, 2, 0.36},
					{2, 3, 0.17},
					{2, 7, 0.34},
					{6, 2, 0.40},
				},
				[]UndirectedEdge{
					{1, 3, 0.29},
					{2, 3, 0.17},
					{3, 6, 0.52},
				},
				[]UndirectedEdge{
					{0, 4, 0.38},
					{4, 5, 0.35},
					{4, 6, 0.93},
					{4, 7, 0.37},
				},
				[]UndirectedEdge{
					{1, 5, 0.32},
					{4, 5, 0.35},
					{5, 7, 0.28},
				},
				[]UndirectedEdge{
					{3, 6, 0.52},
					{4, 6, 0.93},
					{6, 0, 0.58},
					{6, 2, 0.40},
				},
				[]UndirectedEdge{
					{0, 7, 0.16},
					{1, 7, 0.19},
					{2, 7, 0.34},
					{4, 7, 0.37},
					{5, 7, 0.28},
				},
			},
			expectedEdges: []UndirectedEdge{
				{0, 2, 0.26},
				{0, 4, 0.38},
				{0, 7, 0.16},
				{6, 0, 0.58},
				{1, 2, 0.36},
				{1, 3, 0.29},
				{1, 5, 0.32},
				{1, 7, 0.19},
				{2, 3, 0.17},
				{2, 7, 0.34},
				{6, 2, 0.40},
				{3, 6, 0.52},
				{4, 5, 0.35},
				{4, 6, 0.93},
				{4, 7, 0.37},
				{5, 7, 0.28},
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
					expectedVertexVisits: []int{0, 2, 1, 3, 6, 4, 5, 7},
				},
				{
					name:                 "PostOrderDFS",
					source:               0,
					strategy:             DFS,
					order:                PostOrder,
					expectedVertexVisits: []int{7, 5, 4, 6, 3, 1, 2, 0},
				},
				{
					name:                 "PreOrderDFSi",
					source:               0,
					strategy:             DFSi,
					order:                PreOrder,
					expectedVertexVisits: []int{0, 2, 4, 7, 6, 3, 1, 5},
				},
				{
					name:                 "PostOrderDFSi",
					source:               0,
					strategy:             DFSi,
					order:                PostOrder,
					expectedVertexVisits: []int{0, 6, 3, 1, 5, 7, 4, 2},
				},
				{
					name:                 "PreOrderBFS",
					source:               0,
					strategy:             BFS,
					order:                PreOrder,
					expectedVertexVisits: []int{0, 2, 4, 7, 6, 1, 3, 5},
				},
				{
					name:                 "PostOrderBFS",
					source:               0,
					strategy:             BFS,
					order:                PostOrder,
					expectedVertexVisits: []int{0, 2, 4, 7, 6, 1, 3, 5},
				},
			},
			traverseEdgesTests: []traverseEdgesTest{
				{
					name:                 "InvalidVertex",
					source:               -1,
					expectedEdgeVisits:   [][2]int{},
					expectedWeightVisits: []float64{},
				},
				{
					name:                 "InvalidStrategy",
					source:               0,
					strategy:             -1,
					expectedEdgeVisits:   [][2]int{},
					expectedWeightVisits: []float64{},
				},
				{
					name:     "DFS",
					source:   0,
					strategy: DFS,
					expectedEdgeVisits: [][2]int{
						{0, 2},
						{2, 1},
						{1, 3},
						{3, 6},
						{6, 4},
						{4, 5},
						{5, 7},
					},
					expectedWeightVisits: []float64{0.26, 0.36, 0.29, 0.52, 0.93, 0.35, 0.28},
				},
				{
					name:     "DFSi",
					source:   0,
					strategy: DFSi,
					expectedEdgeVisits: [][2]int{
						{0, 2},
						{0, 4},
						{0, 7},
						{0, 6},
						{6, 3},
						{3, 1},
						{1, 5},
					},
					expectedWeightVisits: []float64{0.26, 0.38, 0.16, 0.58, 0.52, 0.29, 0.32},
				},
				{
					name:     "BFS",
					source:   0,
					strategy: BFS,
					expectedEdgeVisits: [][2]int{
						{0, 2},
						{0, 4},
						{0, 7},
						{0, 6},
						{2, 1},
						{2, 3},
						{4, 5},
					},
					expectedWeightVisits: []float64{0.26, 0.38, 0.16, 0.58, 0.36, 0.17, 0.35},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := NewWeightedUndirected(tc.V, tc.edges...)

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

			assert.Equal(t, tc.expectedEdges, g.Edges())

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
						assert.Equal(t, tc.expectedWeightVisits, visitor.weights)
					})
				}
			})

			assert.NotEmpty(t, g.Graphviz())
		})
	}
}
