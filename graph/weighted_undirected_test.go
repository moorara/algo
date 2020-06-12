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
			assert.Equal(t, -1, e.Other(99))
			assert.Equal(t, tc.weight, e.Weight())
		})
	}
}

func TestWeightedGraph(t *testing.T) {
	type traverseTest struct {
		name                          string
		source                        int
		strategy                      TraversalStrategy
		expectedVertexPreOrderVisits  []int
		expectedVertexPostOrderVisits []int
		expectedEdgePreOrderVisits    [][2]int
		expectedWeightPreOrderVisits  []float64
	}

	type pathsTest struct {
		name         string
		source       int
		strategy     TraversalStrategy
		vertex       int
		expectedPath []int
		expectedOK   bool
	}

	type ordersTest struct {
		name                     string
		strategy                 TraversalStrategy
		v                        int
		expectedPreRank          int
		expectedPostRank         int
		expectedPreOrder         []int
		expectedPostOrder        []int
		expectedReversePostOrder []int
	}

	type connectivityTest struct {
		name                string
		v                   int
		w                   int
		expectedID          int
		expectedIsConnected bool
	}

	tests := []struct {
		name                        string
		V                           int
		edges                       []UndirectedEdge
		expectedV                   int
		expectedE                   int
		expectedDegrees             []int
		expectedAdjacents           [][]UndirectedEdge
		expectedEdges               []UndirectedEdge
		traverseTests               []traverseTest
		pathsTests                  []pathsTest
		ordersTests                 []ordersTest
		expectedConnectedComponents [][]int
		connectivityTests           []connectivityTest
		expectedMSTWeight           float64
		expectedMSTEdges            []UndirectedEdge
	}{
		{
			name: "WeightedGraph",
			V:    8,
			edges: []UndirectedEdge{
				{0, 2, 0.26},
				{0, 4, 0.38},
				{0, 6, 0.58},
				{0, 7, 0.16},
				{1, 2, 0.36},
				{1, 3, 0.29},
				{1, 5, 0.32},
				{1, 7, 0.19},
				{2, 3, 0.17},
				{2, 6, 0.40},
				{2, 7, 0.34},
				{3, 6, 0.52},
				{4, 5, 0.35},
				{4, 6, 0.93},
				{4, 7, 0.37},
				{5, 7, 0.28},
			},
			expectedV:       8,
			expectedE:       16,
			expectedDegrees: []int{4, 4, 5, 3, 4, 3, 4},
			expectedAdjacents: [][]UndirectedEdge{
				{
					{0, 2, 0.26},
					{0, 4, 0.38},
					{0, 6, 0.58},
					{0, 7, 0.16},
				},
				{
					{1, 2, 0.36},
					{1, 3, 0.29},
					{1, 5, 0.32},
					{1, 7, 0.19},
				},
				{
					{0, 2, 0.26},
					{1, 2, 0.36},
					{2, 3, 0.17},
					{2, 6, 0.40},
					{2, 7, 0.34},
				},
				{
					{1, 3, 0.29},
					{2, 3, 0.17},
					{3, 6, 0.52},
				},
				{
					{0, 4, 0.38},
					{4, 5, 0.35},
					{4, 6, 0.93},
					{4, 7, 0.37},
				},
				{
					{1, 5, 0.32},
					{4, 5, 0.35},
					{5, 7, 0.28},
				},
				{
					{0, 6, 0.58},
					{2, 6, 0.40},
					{3, 6, 0.52},
					{4, 6, 0.93},
				},
				{
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
				{0, 6, 0.58},
				{0, 7, 0.16},
				{1, 2, 0.36},
				{1, 3, 0.29},
				{1, 5, 0.32},
				{1, 7, 0.19},
				{2, 3, 0.17},
				{2, 6, 0.40},
				{2, 7, 0.34},
				{3, 6, 0.52},
				{4, 5, 0.35},
				{4, 6, 0.93},
				{4, 7, 0.37},
				{5, 7, 0.28},
			},
			traverseTests: []traverseTest{
				{
					name:                          "InvalidVertex",
					source:                        -1,
					expectedVertexPreOrderVisits:  []int{},
					expectedVertexPostOrderVisits: []int{},
					expectedEdgePreOrderVisits:    [][2]int{},
					expectedWeightPreOrderVisits:  []float64{},
				},
				{
					name:                          "InvalidStrategy",
					source:                        0,
					strategy:                      -1,
					expectedVertexPreOrderVisits:  []int{},
					expectedVertexPostOrderVisits: []int{},
					expectedEdgePreOrderVisits:    [][2]int{},
					expectedWeightPreOrderVisits:  []float64{},
				},
				{
					name:                          "DFS",
					source:                        0,
					strategy:                      DFS,
					expectedVertexPreOrderVisits:  []int{0, 2, 1, 3, 6, 4, 5, 7},
					expectedVertexPostOrderVisits: []int{7, 5, 4, 6, 3, 1, 2, 0},
					expectedEdgePreOrderVisits: [][2]int{
						{0, 2},
						{2, 1},
						{1, 3},
						{3, 6},
						{6, 4},
						{4, 5},
						{5, 7},
					},
					expectedWeightPreOrderVisits: []float64{0.26, 0.36, 0.29, 0.52, 0.93, 0.35, 0.28},
				},
				{
					name:                          "DFSi",
					source:                        0,
					strategy:                      DFSi,
					expectedVertexPreOrderVisits:  []int{0, 2, 4, 6, 7, 1, 5, 3},
					expectedVertexPostOrderVisits: []int{0, 7, 5, 1, 3, 6, 4, 2},
					expectedEdgePreOrderVisits: [][2]int{
						{0, 2},
						{0, 4},
						{0, 6},
						{0, 7},
						{7, 1},
						{7, 5},
						{1, 3},
					},
					expectedWeightPreOrderVisits: []float64{0.26, 0.38, 0.58, 0.16, 0.19, 0.28, 0.29},
				},
				{
					name:                          "BFS",
					source:                        0,
					strategy:                      BFS,
					expectedVertexPreOrderVisits:  []int{0, 2, 4, 6, 7, 1, 3, 5},
					expectedVertexPostOrderVisits: []int{0, 2, 4, 6, 7, 1, 3, 5},
					expectedEdgePreOrderVisits: [][2]int{
						{0, 2},
						{0, 4},
						{0, 6},
						{0, 7},
						{2, 1},
						{2, 3},
						{4, 5},
					},
					expectedWeightPreOrderVisits: []float64{0.26, 0.38, 0.58, 0.16, 0.36, 0.17, 0.35},
				},
			},
			pathsTests: []pathsTest{
				{
					name:         "InvalidVertex",
					source:       -1,
					expectedPath: nil,
					expectedOK:   false,
				},
				{
					name:         "InvalidStrategy",
					source:       0,
					strategy:     -1,
					expectedPath: nil,
					expectedOK:   false,
				},
				{
					name:         "DFS",
					source:       0,
					strategy:     DFS,
					vertex:       5,
					expectedPath: []int{0, 2, 1, 3, 6, 4, 5},
					expectedOK:   true,
				},
				{
					name:         "DFSi",
					source:       0,
					strategy:     DFSi,
					vertex:       5,
					expectedPath: []int{0, 7, 5},
					expectedOK:   true,
				},
				{
					name:         "BFS",
					source:       0,
					strategy:     BFS,
					vertex:       5,
					expectedPath: []int{0, 4, 5},
					expectedOK:   true,
				},
			},
			ordersTests: []ordersTest{
				{
					name:                     "InvalidStrategy",
					strategy:                 -1,
					v:                        0,
					expectedPreRank:          0,
					expectedPostRank:         0,
					expectedPreOrder:         []int{},
					expectedPostOrder:        []int{},
					expectedReversePostOrder: []int{},
				},
				{
					name:                     "DFS",
					strategy:                 DFS,
					v:                        3,
					expectedPreRank:          3,
					expectedPostRank:         4,
					expectedPreOrder:         []int{0, 2, 1, 3, 6, 4, 5, 7},
					expectedPostOrder:        []int{7, 5, 4, 6, 3, 1, 2, 0},
					expectedReversePostOrder: []int{0, 2, 1, 3, 6, 4, 5, 7},
				},
				{
					name:                     "DFSi",
					strategy:                 DFSi,
					v:                        3,
					expectedPreRank:          7,
					expectedPostRank:         4,
					expectedPreOrder:         []int{0, 2, 4, 6, 7, 1, 5, 3},
					expectedPostOrder:        []int{0, 7, 5, 1, 3, 6, 4, 2},
					expectedReversePostOrder: []int{2, 4, 6, 3, 1, 5, 7, 0},
				},
				{
					name:                     "BFS",
					strategy:                 BFS,
					v:                        3,
					expectedPreRank:          6,
					expectedPostRank:         6,
					expectedPreOrder:         []int{0, 2, 4, 6, 7, 1, 3, 5},
					expectedPostOrder:        []int{0, 2, 4, 6, 7, 1, 3, 5},
					expectedReversePostOrder: []int{5, 3, 1, 7, 6, 4, 2, 0},
				},
			},
			expectedConnectedComponents: [][]int{
				{0, 1, 2, 3, 4, 5, 6, 7},
			},
			connectivityTests: []connectivityTest{
				{
					name:                "Connected#1",
					v:                   3,
					w:                   4,
					expectedID:          0,
					expectedIsConnected: true,
				},
				{
					name:                "Connected#2",
					v:                   5,
					w:                   6,
					expectedID:          0,
					expectedIsConnected: true,
				},
			},
			expectedMSTWeight: 1.81,
			expectedMSTEdges: []UndirectedEdge{
				{0, 7, 0.16},
				{0, 2, 0.26},
				{1, 7, 0.19},
				{2, 3, 0.17},
				{2, 6, 0.40},
				{4, 5, 0.35},
				{5, 7, 0.28},
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

			t.Run("Adjacency", func(t *testing.T) {
				assert.Nil(t, g.Adj(-1))
				for v, expectedAdj := range tc.expectedAdjacents {
					assert.Equal(t, expectedAdj, g.Adj(v))
				}
			})

			t.Run("Edges", func(t *testing.T) {
				assert.Equal(t, tc.expectedEdges, g.Edges())
			})

			t.Run("Traverse", func(t *testing.T) {
				for _, tc := range tc.traverseTests {
					t.Run(tc.name, func(t *testing.T) {
						tv := newTestVisitors()
						g.Traverse(tc.source, tc.strategy, tv.Visitors)
						assert.Equal(t, tc.expectedVertexPreOrderVisits, tv.preOrderVertices)
						assert.Equal(t, tc.expectedVertexPostOrderVisits, tv.postOrderVertices)
						assert.Equal(t, tc.expectedEdgePreOrderVisits, tv.preOrderEdges)
						assert.Equal(t, tc.expectedWeightPreOrderVisits, tv.preOrderWeights)
					})
				}
			})

			t.Run("Paths", func(t *testing.T) {
				for _, tc := range tc.pathsTests {
					t.Run(tc.name, func(t *testing.T) {
						path, ok := g.Paths(tc.source, tc.strategy).To(tc.vertex)
						assert.Equal(t, tc.expectedPath, path)
						assert.Equal(t, tc.expectedOK, ok)
					})
				}
			})

			t.Run("Orders", func(t *testing.T) {
				for _, tc := range tc.ordersTests {
					t.Run(tc.name, func(t *testing.T) {
						o := g.Orders(tc.strategy)
						assert.Equal(t, tc.expectedPreRank, o.PreRank(tc.v))
						assert.Equal(t, tc.expectedPostRank, o.PostRank(tc.v))
						assert.Equal(t, tc.expectedPreOrder, o.PreOrder())
						assert.Equal(t, tc.expectedPostOrder, o.PostOrder())
						assert.Equal(t, tc.expectedReversePostOrder, o.ReversePostOrder())
					})
				}
			})

			t.Run("ConnectedComponents", func(t *testing.T) {
				cc := g.ConnectedComponents()
				assert.Equal(t, tc.expectedConnectedComponents, cc.Components())
				for _, tc := range tc.connectivityTests {
					t.Run(tc.name, func(t *testing.T) {
						assert.Equal(t, tc.expectedID, cc.ID(tc.v))
						assert.Equal(t, tc.expectedIsConnected, cc.IsConnected(tc.v, tc.w))
					})
				}
			})

			t.Run("MinimumSpanningTree", func(t *testing.T) {
				mst := g.MinimumSpanningTree()
				edges := mst.Edges()
				assert.InEpsilon(t, tc.expectedMSTWeight, mst.Weight(), float64Epsilon)
				for _, expectedMSTEdge := range tc.expectedMSTEdges {
					assert.Contains(t, edges, expectedMSTEdge)
				}
			})

			assert.NotEmpty(t, g.Graphviz())
		})
	}
}
