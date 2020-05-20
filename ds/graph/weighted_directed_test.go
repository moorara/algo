package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDirectedEdge(t *testing.T) {
	tests := []struct {
		name     string
		from, to int
		weight   float64
	}{
		{"Edge1", 1, 2, 0.27},
		{"Edge2", 3, 4, 0.69},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			e := DirectedEdge{tc.from, tc.to, tc.weight}

			assert.NotEmpty(t, e)
			assert.Equal(t, tc.from, e.From())
			assert.Equal(t, tc.to, e.To())
			assert.Equal(t, tc.weight, e.Weight())
		})
	}
}

func TestWeightedDirected(t *testing.T) {
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
		edges                 []DirectedEdge
		expectedV             int
		expectedE             int
		expectedInDegrees     []int
		expectedOutDegrees    []int
		expectedAdjacents     [][]DirectedEdge
		expectedEdges         []DirectedEdge
		expectedReverse       *WeightedDirected
		traverseVerticesTests []traverseVerticesTest
		traverseEdgesTests    []traverseEdgesTest
	}{
		{
			name: "",
			V:    8,
			edges: []DirectedEdge{
				{0, 2, 0.26},
				{0, 4, 0.38},
				{1, 3, 0.29},
				{2, 7, 0.34},
				{3, 6, 0.52},
				{4, 5, 0.35},
				{4, 7, 0.37},
				{5, 1, 0.32},
				{5, 4, 0.35},
				{5, 7, 0.28},
				{6, 0, 0.58},
				{6, 2, 0.40},
				{6, 4, 0.93},
				{7, 3, 0.39},
				{7, 5, 0.28},
			},
			expectedV:          8,
			expectedE:          15,
			expectedInDegrees:  []int{1, 1, 2, 2, 3, 2, 1, 3},
			expectedOutDegrees: []int{2, 1, 1, 1, 2, 3, 3, 2},
			expectedAdjacents: [][]DirectedEdge{
				[]DirectedEdge{
					{0, 2, 0.26},
					{0, 4, 0.38},
				},
				[]DirectedEdge{
					{1, 3, 0.29},
				},
				[]DirectedEdge{
					{2, 7, 0.34},
				},
				[]DirectedEdge{
					{3, 6, 0.52},
				},
				[]DirectedEdge{
					{4, 5, 0.35},
					{4, 7, 0.37},
				},
				[]DirectedEdge{
					{5, 1, 0.32},
					{5, 4, 0.35},
					{5, 7, 0.28},
				},
				[]DirectedEdge{
					{6, 0, 0.58},
					{6, 2, 0.40},
					{6, 4, 0.93},
				},
				[]DirectedEdge{
					{7, 3, 0.39},
					{7, 5, 0.28},
				},
			},
			expectedEdges: []DirectedEdge{
				{0, 2, 0.26},
				{0, 4, 0.38},
				{1, 3, 0.29},
				{2, 7, 0.34},
				{3, 6, 0.52},
				{4, 5, 0.35},
				{4, 7, 0.37},
				{5, 1, 0.32},
				{5, 4, 0.35},
				{5, 7, 0.28},
				{6, 0, 0.58},
				{6, 2, 0.40},
				{6, 4, 0.93},
				{7, 3, 0.39},
				{7, 5, 0.28},
			},
			expectedReverse: &WeightedDirected{
				v:   8,
				e:   15,
				ins: []int{2, 1, 1, 1, 2, 3, 3, 2},
				adj: [][]DirectedEdge{
					[]DirectedEdge{
						{0, 6, 0.58},
					},
					[]DirectedEdge{
						{1, 5, 0.32},
					},
					[]DirectedEdge{
						{2, 0, 0.26},
						{2, 6, 0.40},
					},
					[]DirectedEdge{
						{3, 1, 0.29},
						{3, 7, 0.39},
					},
					[]DirectedEdge{
						{4, 0, 0.38},
						{4, 5, 0.35},
						{4, 6, 0.93},
					},
					[]DirectedEdge{
						{5, 4, 0.35},
						{5, 7, 0.28},
					},
					[]DirectedEdge{
						{6, 3, 0.52},
					},
					[]DirectedEdge{
						{7, 2, 0.34},
						{7, 4, 0.37},
						{7, 5, 0.28},
					},
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
					expectedVertexVisits: []int{0, 2, 7, 3, 6, 4, 5, 1},
				},
				{
					name:                 "PostOrderDFS",
					source:               0,
					strategy:             DFS,
					order:                PostOrder,
					expectedVertexVisits: []int{1, 5, 4, 6, 3, 7, 2, 0},
				},
				{
					name:                 "PreOrderDFSi",
					source:               0,
					strategy:             DFSi,
					order:                PreOrder,
					expectedVertexVisits: []int{0, 2, 4, 5, 7, 3, 6, 1},
				},
				{
					name:                 "PostOrderDFSi",
					source:               0,
					strategy:             DFSi,
					order:                PostOrder,
					expectedVertexVisits: []int{0, 4, 7, 3, 6, 5, 1, 2},
				},
				{
					name:                 "PreOrderBFS",
					source:               0,
					strategy:             BFS,
					order:                PreOrder,
					expectedVertexVisits: []int{0, 2, 4, 7, 5, 3, 1, 6},
				},
				{
					name:                 "PostOrderBFS",
					source:               0,
					strategy:             BFS,
					order:                PostOrder,
					expectedVertexVisits: []int{0, 2, 4, 7, 5, 3, 1, 6},
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
						{2, 7},
						{7, 3},
						{3, 6},
						{6, 4},
						{4, 5},
						{5, 1},
					},
					expectedWeightVisits: []float64{0.26, 0.34, 0.39, 0.52, 0.93, 0.35, 0.32},
				},
				{
					name:     "DFSi",
					source:   0,
					strategy: DFSi,
					expectedEdgeVisits: [][2]int{
						{0, 2},
						{0, 4},
						{4, 5},
						{4, 7},
						{7, 3},
						{3, 6},
						{5, 1},
					},
					expectedWeightVisits: []float64{0.26, 0.38, 0.35, 0.37, 0.39, 0.52, 0.32},
				},
				{
					name:     "BFS",
					source:   0,
					strategy: BFS,
					expectedEdgeVisits: [][2]int{
						{0, 2},
						{0, 4},
						{2, 7},
						{4, 5},
						{7, 3},
						{5, 1},
						{3, 6},
					},
					expectedWeightVisits: []float64{0.26, 0.38, 0.34, 0.35, 0.39, 0.32, 0.52},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := NewWeightedDirected(tc.V, tc.edges...)

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

			assert.Equal(t, tc.expectedEdges, g.Edges())
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
						assert.Equal(t, tc.expectedWeightVisits, visitor.weights)
					})
				}
			})

			assert.NotEmpty(t, g.Graphviz())
		})
	}
}
