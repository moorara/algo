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

	type strongConnectivityTest struct {
		name                        string
		v                           int
		w                           int
		expectedID                  int
		expectedIsStronglyConnected bool
	}

	tests := []struct {
		name                                string
		V                                   int
		edges                               []DirectedEdge
		expectedV                           int
		expectedE                           int
		expectedInDegrees                   []int
		expectedOutDegrees                  []int
		expectedAdjacents                   [][]DirectedEdge
		expectedEdges                       []DirectedEdge
		expectedReverse                     *WeightedDirected
		traverseTests                       []traverseTest
		pathsTests                          []pathsTest
		ordersTests                         []ordersTest
		expectedStronglyConnectedComponents [][]int
		strongConnectivityTests             []strongConnectivityTest
	}{
		{
			name: "WeightedDigraph",
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
				{
					{0, 2, 0.26},
					{0, 4, 0.38},
				},
				{
					{1, 3, 0.29},
				},
				{
					{2, 7, 0.34},
				},
				{
					{3, 6, 0.52},
				},
				{
					{4, 5, 0.35},
					{4, 7, 0.37},
				},
				{
					{5, 1, 0.32},
					{5, 4, 0.35},
					{5, 7, 0.28},
				},
				{
					{6, 0, 0.58},
					{6, 2, 0.40},
					{6, 4, 0.93},
				},
				{
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
					expectedVertexPreOrderVisits:  []int{0, 2, 7, 3, 6, 4, 5, 1},
					expectedVertexPostOrderVisits: []int{1, 5, 4, 6, 3, 7, 2, 0},
					expectedEdgePreOrderVisits: [][2]int{
						{0, 2},
						{2, 7},
						{7, 3},
						{3, 6},
						{6, 4},
						{4, 5},
						{5, 1},
					},
					expectedWeightPreOrderVisits: []float64{0.26, 0.34, 0.39, 0.52, 0.93, 0.35, 0.32},
				},
				{
					name:                          "DFSi",
					source:                        0,
					strategy:                      DFSi,
					expectedVertexPreOrderVisits:  []int{0, 2, 4, 5, 7, 3, 6, 1},
					expectedVertexPostOrderVisits: []int{0, 4, 7, 3, 6, 5, 1, 2},
					expectedEdgePreOrderVisits: [][2]int{
						{0, 2},
						{0, 4},
						{4, 5},
						{4, 7},
						{7, 3},
						{3, 6},
						{5, 1},
					},
					expectedWeightPreOrderVisits: []float64{0.26, 0.38, 0.35, 0.37, 0.39, 0.52, 0.32},
				},
				{
					name:                          "BFS",
					source:                        0,
					strategy:                      BFS,
					expectedVertexPreOrderVisits:  []int{0, 2, 4, 7, 5, 3, 1, 6},
					expectedVertexPostOrderVisits: []int{0, 2, 4, 7, 5, 3, 1, 6},
					expectedEdgePreOrderVisits: [][2]int{
						{0, 2},
						{0, 4},
						{2, 7},
						{4, 5},
						{7, 3},
						{5, 1},
						{3, 6},
					},
					expectedWeightPreOrderVisits: []float64{0.26, 0.38, 0.34, 0.35, 0.39, 0.32, 0.52},
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
					vertex:       6,
					expectedPath: []int{0, 2, 7, 3, 6},
					expectedOK:   true,
				},
				{
					name:         "DFSi",
					source:       0,
					strategy:     DFSi,
					vertex:       6,
					expectedPath: []int{0, 4, 7, 3, 6},
					expectedOK:   true,
				},
				{
					name:         "BFS",
					source:       0,
					strategy:     BFS,
					vertex:       6,
					expectedPath: []int{0, 2, 7, 3, 6},
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
					v:                        5,
					expectedPreRank:          6,
					expectedPostRank:         1,
					expectedPreOrder:         []int{0, 2, 7, 3, 6, 4, 5, 1},
					expectedPostOrder:        []int{1, 5, 4, 6, 3, 7, 2, 0},
					expectedReversePostOrder: []int{0, 2, 7, 3, 6, 4, 5, 1},
				},
				{
					name:                     "DFSi",
					strategy:                 DFSi,
					v:                        5,
					expectedPreRank:          3,
					expectedPostRank:         5,
					expectedPreOrder:         []int{0, 2, 4, 5, 7, 3, 6, 1},
					expectedPostOrder:        []int{0, 4, 7, 3, 6, 5, 1, 2},
					expectedReversePostOrder: []int{2, 1, 5, 6, 3, 7, 4, 0},
				},
				{
					name:                     "BFS",
					strategy:                 BFS,
					v:                        5,
					expectedPreRank:          4,
					expectedPostRank:         4,
					expectedPreOrder:         []int{0, 2, 4, 7, 5, 3, 1, 6},
					expectedPostOrder:        []int{0, 2, 4, 7, 5, 3, 1, 6},
					expectedReversePostOrder: []int{6, 1, 3, 5, 7, 4, 2, 0},
				},
			},
			expectedStronglyConnectedComponents: [][]int{
				{0, 1, 2, 3, 4, 5, 6, 7},
			},
			strongConnectivityTests: []strongConnectivityTest{
				{
					name:                        "Connected#1",
					v:                           0,
					w:                           4,
					expectedID:                  0,
					expectedIsStronglyConnected: true,
				},
				{
					name:                        "Connected#2",
					v:                           3,
					w:                           5,
					expectedID:                  0,
					expectedIsStronglyConnected: true,
				},
				{
					name:                        "Connected#3",
					v:                           1,
					w:                           6,
					expectedID:                  0,
					expectedIsStronglyConnected: true,
				},
				{
					name:                        "Connected#4",
					v:                           2,
					w:                           7,
					expectedID:                  0,
					expectedIsStronglyConnected: true,
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

			t.Run("Adjacency", func(t *testing.T) {
				assert.Nil(t, g.Adj(-1))
				for v, expectedAdj := range tc.expectedAdjacents {
					assert.Equal(t, expectedAdj, g.Adj(v))
				}
			})

			t.Run("Edges", func(t *testing.T) {
				assert.Equal(t, tc.expectedEdges, g.Edges())
			})

			t.Run("Reverse", func(t *testing.T) {
				assert.Equal(t, tc.expectedReverse, g.Reverse())
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

			t.Run("StronglyConnectedComponents", func(t *testing.T) {
				scc := g.StronglyConnectedComponents()
				assert.Equal(t, tc.expectedStronglyConnectedComponents, scc.Components())
				for _, tc := range tc.strongConnectivityTests {
					t.Run(tc.name, func(t *testing.T) {
						assert.Equal(t, tc.expectedID, scc.ID(tc.v))
						assert.Equal(t, tc.expectedIsStronglyConnected, scc.IsStronglyConnected(tc.v, tc.w))
					})
				}
			})

			assert.NotEmpty(t, g.Graphviz())
		})
	}
}
