package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDirected(t *testing.T) {
	type traverseTest struct {
		name                          string
		source                        int
		strategy                      TraversalStrategy
		expectedVertexPreOrderVisits  []int
		expectedVertexPostOrderVisits []int
		expectedEdgePreOrderVisits    [][2]int
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
		edges                               [][2]int
		expectedV                           int
		expectedE                           int
		expectedInDegrees                   []int
		expectedOutDegrees                  []int
		expectedAdjacents                   [][]int
		expectedReverse                     *Directed
		traverseTests                       []traverseTest
		pathsTests                          []pathsTest
		ordersTests                         []ordersTest
		expectedStronglyConnectedComponents [][]int
		strongConnectivityTests             []strongConnectivityTest
	}{
		{
			name: "Digraph",
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
					expectedVertexPreOrderVisits:  []int{0, 1, 5, 4, 2, 3},
					expectedVertexPostOrderVisits: []int{1, 3, 2, 4, 5, 0},
					expectedEdgePreOrderVisits: [][2]int{
						{0, 1},
						{0, 5},
						{5, 4},
						{4, 2},
						{2, 3},
					},
				},
				{
					name:                          "DFSi",
					source:                        0,
					strategy:                      DFSi,
					expectedVertexPreOrderVisits:  []int{0, 1, 5, 4, 2, 3},
					expectedVertexPostOrderVisits: []int{0, 5, 4, 3, 2, 1},
					expectedEdgePreOrderVisits: [][2]int{
						{0, 1},
						{0, 5},
						{5, 4},
						{4, 2},
						{4, 3},
					},
				},
				{
					name:                          "BFS",
					source:                        0,
					strategy:                      BFS,
					expectedVertexPreOrderVisits:  []int{0, 1, 5, 4, 2, 3},
					expectedVertexPostOrderVisits: []int{0, 1, 5, 4, 2, 3},
					expectedEdgePreOrderVisits: [][2]int{
						{0, 1},
						{0, 5},
						{5, 4},
						{4, 2},
						{4, 3},
					},
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
					vertex:       3,
					expectedPath: []int{0, 5, 4, 2, 3},
					expectedOK:   true,
				},
				{
					name:         "DFSi",
					source:       0,
					strategy:     DFSi,
					vertex:       3,
					expectedPath: []int{0, 5, 4, 3},
					expectedOK:   true,
				},
				{
					name:         "BFS",
					source:       0,
					strategy:     BFS,
					vertex:       3,
					expectedPath: []int{0, 5, 4, 3},
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
					v:                        10,
					expectedPreRank:          8,
					expectedPostRank:         7,
					expectedPreOrder:         []int{0, 1, 5, 4, 2, 3, 6, 9, 10, 12, 11, 7, 8},
					expectedPostOrder:        []int{1, 3, 2, 4, 5, 0, 12, 10, 11, 9, 6, 8, 7},
					expectedReversePostOrder: []int{7, 8, 6, 9, 11, 10, 12, 0, 5, 4, 2, 3, 1},
				},
				{
					name:                     "DFSi",
					strategy:                 DFSi,
					v:                        10,
					expectedPreRank:          8,
					expectedPostRank:         10,
					expectedPreOrder:         []int{0, 1, 5, 4, 2, 3, 6, 9, 10, 11, 12, 7, 8},
					expectedPostOrder:        []int{0, 5, 4, 3, 2, 1, 6, 9, 11, 12, 10, 7, 8},
					expectedReversePostOrder: []int{8, 7, 10, 12, 11, 9, 6, 1, 2, 3, 4, 5, 0},
				},
				{
					name:                     "BFS",
					strategy:                 BFS,
					v:                        10,
					expectedPreRank:          8,
					expectedPostRank:         8,
					expectedPreOrder:         []int{0, 1, 5, 4, 2, 3, 6, 9, 10, 11, 12, 7, 8},
					expectedPostOrder:        []int{0, 1, 5, 4, 2, 3, 6, 9, 10, 11, 12, 7, 8},
					expectedReversePostOrder: []int{8, 7, 12, 11, 10, 9, 6, 3, 2, 4, 5, 1, 0},
				},
			},
			expectedStronglyConnectedComponents: [][]int{
				[]int{1},
				[]int{0, 2, 3, 4, 5},
				[]int{9, 10, 11, 12},
				[]int{6},
				[]int{7, 8},
			},
			strongConnectivityTests: []strongConnectivityTest{
				{
					name:                        "Connected#1",
					v:                           0,
					w:                           2,
					expectedID:                  1,
					expectedIsStronglyConnected: true,
				},
				{
					name:                        "Connected#2",
					v:                           3,
					w:                           5,
					expectedID:                  1,
					expectedIsStronglyConnected: true,
				},
				{
					name:                        "Connected#3",
					v:                           7,
					w:                           8,
					expectedID:                  4,
					expectedIsStronglyConnected: true,
				},
				{
					name:                        "Connected#4",
					v:                           10,
					w:                           12,
					expectedID:                  2,
					expectedIsStronglyConnected: true,
				},
				{
					name:                        "Disconnected#1",
					v:                           0,
					w:                           1,
					expectedID:                  1,
					expectedIsStronglyConnected: false,
				},
				{
					name:                        "Disconnected#2",
					v:                           4,
					w:                           6,
					expectedID:                  1,
					expectedIsStronglyConnected: false,
				},
				{
					name:                        "Disconnected#3",
					v:                           6,
					w:                           8,
					expectedID:                  3,
					expectedIsStronglyConnected: false,
				},
				{
					name:                        "Disconnected#4",
					v:                           8,
					w:                           12,
					expectedID:                  4,
					expectedIsStronglyConnected: false,
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
