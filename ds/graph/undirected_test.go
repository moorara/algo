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
		edges                       [][2]int
		expectedV                   int
		expectedE                   int
		expectedDegrees             []int
		expectedAdjacents           [][]int
		traverseTests               []traverseTest
		pathsTests                  []pathsTest
		ordersTests                 []ordersTest
		expectedConnectedComponents [][]int
		connectivityTests           []connectivityTest
	}{
		{
			name: "Graph",
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
					vertex:       4,
					expectedPath: []int{0, 5, 3, 4},
					expectedOK:   true,
				},
				{
					name:         "DFSi",
					source:       0,
					strategy:     DFSi,
					vertex:       4,
					expectedPath: []int{0, 6, 4},
					expectedOK:   true,
				},
				{
					name:         "BFS",
					source:       0,
					strategy:     BFS,
					vertex:       4,
					expectedPath: []int{0, 5, 4},
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
					v:                        6,
					expectedPreRank:          6,
					expectedPostRank:         2,
					expectedPreOrder:         []int{0, 1, 2, 5, 3, 4, 6, 7, 8, 9, 10, 11, 12},
					expectedPostOrder:        []int{1, 2, 6, 4, 3, 5, 0, 8, 7, 10, 12, 11, 9},
					expectedReversePostOrder: []int{9, 11, 12, 10, 7, 8, 0, 5, 3, 4, 6, 2, 1},
				},
				{
					name:                     "DFSi",
					strategy:                 DFSi,
					v:                        6,
					expectedPreRank:          4,
					expectedPostRank:         1,
					expectedPreOrder:         []int{0, 1, 2, 5, 6, 4, 3, 7, 8, 9, 10, 11, 12},
					expectedPostOrder:        []int{0, 6, 4, 3, 5, 2, 1, 7, 8, 9, 12, 11, 10},
					expectedReversePostOrder: []int{10, 11, 12, 9, 8, 7, 1, 2, 5, 3, 4, 6, 0},
				},
				{
					name:                     "BFS",
					strategy:                 BFS,
					v:                        6,
					expectedPreRank:          4,
					expectedPostRank:         4,
					expectedPreOrder:         []int{0, 1, 2, 5, 6, 3, 4, 7, 8, 9, 10, 11, 12},
					expectedPostOrder:        []int{0, 1, 2, 5, 6, 3, 4, 7, 8, 9, 10, 11, 12},
					expectedReversePostOrder: []int{12, 11, 10, 9, 8, 7, 4, 3, 6, 5, 2, 1, 0},
				},
			},
			expectedConnectedComponents: [][]int{
				[]int{0, 1, 2, 3, 4, 5, 6},
				[]int{7, 8},
				[]int{9, 10, 11, 12},
			},
			connectivityTests: []connectivityTest{
				{
					name:                "Connected#1",
					v:                   0,
					w:                   4,
					expectedID:          0,
					expectedIsConnected: true,
				},
				{
					name:                "Connected#2",
					v:                   7,
					w:                   8,
					expectedID:          1,
					expectedIsConnected: true,
				},
				{
					name:                "Connected#3",
					v:                   10,
					w:                   11,
					expectedID:          2,
					expectedIsConnected: true,
				},
				{
					name:                "Disconnected#1",
					v:                   2,
					w:                   8,
					expectedID:          0,
					expectedIsConnected: false,
				},
				{
					name:                "Disconnected#2",
					v:                   2,
					w:                   10,
					expectedID:          0,
					expectedIsConnected: false,
				},
				{
					name:                "Disconnected#3",
					v:                   7,
					w:                   9,
					expectedID:          1,
					expectedIsConnected: false,
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

			assert.NotEmpty(t, g.Graphviz())
		})
	}
}
