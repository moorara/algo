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

	type directedCycleTest struct {
		expectedCycle []int
		expectedOK    bool
	}

	type topologicalRankTest struct {
		v            int
		expectedRank int
		expectedOK   bool
	}

	type topologicalTest struct {
		expectedOrder        []int
		expectedOK           bool
		topologicalRankTests []topologicalRankTest
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
		directedCycleTest                   directedCycleTest
		topologicalTest                     topologicalTest
	}{
		{
			name: "Digraph",
			V:    13,
			edges: [][2]int{
				{0, 1},
				{0, 5},
				{2, 0},
				{2, 3},
				{3, 2},
				{3, 5},
				{4, 2},
				{4, 3},
				{5, 4},
				{6, 0},
				{6, 4},
				{6, 9},
				{7, 6},
				{7, 8},
				{8, 7},
				{8, 9},
				{9, 10},
				{9, 11},
				{10, 12},
				{11, 4},
				{11, 12},
				{12, 9},
			},
			expectedV:          13,
			expectedE:          22,
			expectedInDegrees:  []int{2, 1, 2, 2, 3, 2, 1, 1, 1, 3, 1, 1, 2},
			expectedOutDegrees: []int{2, 0, 2, 2, 2, 1, 3, 2, 2, 2, 1, 2, 1},
			expectedAdjacents: [][]int{
				{1, 5},
				{},
				{0, 3},
				{2, 5},
				{2, 3},
				{4},
				{0, 4, 9},
				{6, 8},
				{7, 9},
				{10, 11},
				{12},
				{4, 12},
				{9},
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
				{1},
				{0, 2, 3, 4, 5},
				{9, 10, 11, 12},
				{6},
				{7, 8},
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
			directedCycleTest: directedCycleTest{
				expectedCycle: []int{2, 0, 5, 4, 2},
				expectedOK:    true,
			},
			topologicalTest: topologicalTest{
				expectedOrder: nil,
				expectedOK:    false,
				topologicalRankTests: []topologicalRankTest{
					{
						v:            0,
						expectedRank: -1,
						expectedOK:   false,
					},
				},
			},
		},
		{
			name: "DAG",
			V:    13,
			edges: [][2]int{
				{0, 1},
				{0, 5},
				{0, 6},
				{2, 0},
				{2, 3},
				{3, 5},
				{5, 4},
				{6, 4},
				{6, 9},
				{7, 6},
				{8, 7},
				{9, 10},
				{9, 11},
				{9, 12},
				{11, 12},
			},
			expectedV:          13,
			expectedE:          15,
			expectedInDegrees:  []int{1, 1, 0, 1, 2, 2, 2, 1, 0, 1, 1, 1, 2},
			expectedOutDegrees: []int{3, 0, 2, 1, 0, 1, 2, 1, 1, 3, 0, 1, 0},
			expectedAdjacents: [][]int{
				{1, 5, 6},
				{},
				{0, 3},
				{5},
				{},
				{4},
				{4, 9},
				{6},
				{7},
				{10, 11, 12},
				{},
				{12},
				{},
			},
			expectedReverse: &Directed{
				v:   13,
				e:   15,
				ins: []int{3, 0, 2, 1, 0, 1, 2, 1, 1, 3, 0, 1, 0},
				adj: [][]int{
					[]int{2},
					[]int{0},
					[]int{},
					[]int{2},
					[]int{5, 6},
					[]int{0, 3},
					[]int{0, 7},
					[]int{8},
					[]int{},
					[]int{6},
					[]int{9},
					[]int{9},
					[]int{9, 11},
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
					expectedVertexPreOrderVisits:  []int{0, 1, 5, 4, 6, 9, 10, 11, 12},
					expectedVertexPostOrderVisits: []int{1, 4, 5, 10, 12, 11, 9, 6, 0},
					expectedEdgePreOrderVisits: [][2]int{
						{0, 1},
						{0, 5},
						{5, 4},
						{0, 6},
						{6, 9},
						{9, 10},
						{9, 11},
						{11, 12},
					},
				},
				{
					name:                          "DFSi",
					source:                        0,
					strategy:                      DFSi,
					expectedVertexPreOrderVisits:  []int{0, 1, 5, 6, 4, 9, 10, 11, 12},
					expectedVertexPostOrderVisits: []int{0, 6, 9, 12, 11, 10, 4, 5, 1},
					expectedEdgePreOrderVisits: [][2]int{
						{0, 1},
						{0, 5},
						{0, 6},
						{6, 4},
						{6, 9},
						{9, 10},
						{9, 11},
						{9, 12},
					},
				},
				{
					name:                          "BFS",
					source:                        0,
					strategy:                      BFS,
					expectedVertexPreOrderVisits:  []int{0, 1, 5, 6, 4, 9, 10, 11, 12},
					expectedVertexPostOrderVisits: []int{0, 1, 5, 6, 4, 9, 10, 11, 12},
					expectedEdgePreOrderVisits: [][2]int{
						{0, 1},
						{0, 5},
						{0, 6},
						{5, 4},
						{6, 9},
						{9, 10},
						{9, 11},
						{9, 12},
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
					vertex:       12,
					expectedPath: []int{0, 6, 9, 11, 12},
					expectedOK:   true,
				},
				{
					name:         "DFSi",
					source:       0,
					strategy:     DFSi,
					vertex:       12,
					expectedPath: []int{0, 6, 9, 12},
					expectedOK:   true,
				},
				{
					name:         "BFS",
					source:       0,
					strategy:     BFS,
					vertex:       12,
					expectedPath: []int{0, 6, 9, 12},
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
					expectedPreRank:          6,
					expectedPostRank:         3,
					expectedPreOrder:         []int{0, 1, 5, 4, 6, 9, 10, 11, 12, 2, 3, 7, 8},
					expectedPostOrder:        []int{1, 4, 5, 10, 12, 11, 9, 6, 0, 3, 2, 7, 8},
					expectedReversePostOrder: []int{8, 7, 2, 3, 0, 6, 9, 11, 12, 10, 5, 4, 1},
				},
				{
					name:                     "DFSi",
					strategy:                 DFSi,
					v:                        10,
					expectedPreRank:          6,
					expectedPostRank:         5,
					expectedPreOrder:         []int{0, 1, 5, 6, 4, 9, 10, 11, 12, 2, 3, 7, 8},
					expectedPostOrder:        []int{0, 6, 9, 12, 11, 10, 4, 5, 1, 2, 3, 7, 8},
					expectedReversePostOrder: []int{8, 7, 3, 2, 1, 5, 4, 10, 11, 12, 9, 6, 0},
				},
				{
					name:                     "BFS",
					strategy:                 BFS,
					v:                        10,
					expectedPreRank:          6,
					expectedPostRank:         6,
					expectedPreOrder:         []int{0, 1, 5, 6, 4, 9, 10, 11, 12, 2, 3, 7, 8},
					expectedPostOrder:        []int{0, 1, 5, 6, 4, 9, 10, 11, 12, 2, 3, 7, 8},
					expectedReversePostOrder: []int{8, 7, 3, 2, 12, 11, 10, 9, 4, 6, 5, 1, 0},
				},
			},
			expectedStronglyConnectedComponents: [][]int{
				{12},
				{11},
				{10},
				{9},
				{4},
				{6},
				{7},
				{8},
				{5},
				{3},
				{1},
				{0},
				{2},
			},
			strongConnectivityTests: []strongConnectivityTest{
				{
					name:                        "Disconnected#1",
					v:                           0,
					w:                           10,
					expectedID:                  11,
					expectedIsStronglyConnected: false,
				},
				{
					name:                        "Disconnected#2",
					v:                           1,
					w:                           11,
					expectedID:                  10,
					expectedIsStronglyConnected: false,
				},
			},
			directedCycleTest: directedCycleTest{
				expectedCycle: nil,
				expectedOK:    false,
			},
			topologicalTest: topologicalTest{
				expectedOrder: []int{8, 7, 2, 3, 0, 6, 9, 11, 12, 10, 5, 4, 1},
				expectedOK:    true,
				topologicalRankTests: []topologicalRankTest{
					{
						v:            0,
						expectedRank: 4,
						expectedOK:   true,
					},
					{
						v:            6,
						expectedRank: 5,
						expectedOK:   true,
					},
					{
						v:            12,
						expectedRank: 8,
						expectedOK:   true,
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

			t.Run("Adjacency", func(t *testing.T) {
				assert.Nil(t, g.Adj(-1))
				for v, expectedAdj := range tc.expectedAdjacents {
					assert.Equal(t, expectedAdj, g.Adj(v))
				}
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

			t.Run("DirectedCycle", func(t *testing.T) {
				cycle, ok := g.DirectedCycle().Cycle()
				assert.Equal(t, tc.directedCycleTest.expectedCycle, cycle)
				assert.Equal(t, tc.directedCycleTest.expectedOK, ok)
			})

			t.Run("Topological", func(t *testing.T) {
				tp := g.Topological()
				order, ok := tp.Order()
				assert.Equal(t, tc.topologicalTest.expectedOrder, order)
				assert.Equal(t, tc.topologicalTest.expectedOK, ok)
				for _, tc := range tc.topologicalTest.topologicalRankTests {
					rank, ok := tp.Rank(tc.v)
					assert.Equal(t, tc.expectedRank, rank)
					assert.Equal(t, tc.expectedOK, ok)
				}
			})

			assert.NotEmpty(t, g.Graphviz())
		})
	}
}
