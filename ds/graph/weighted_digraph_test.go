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
			e := NewDirectedEdge(tc.from, tc.to, tc.weight)

			assert.NotEmpty(t, e)
			assert.Equal(t, tc.from, e.From())
			assert.Equal(t, tc.to, e.To())
			assert.Equal(t, tc.weight, e.Weight())
		})
	}
}

func TestWeightedDigraph(t *testing.T) {
	tests := []struct {
		name               string
		V                  int
		edges              []DirectedEdge
		expectedV          int
		expectedE          int
		expectedInDegrees  []int
		expectedOutDegrees []int
		expectedAdjacents  [][]DirectedEdge
		expectedEdges      []DirectedEdge
		expectedReverse    WeightedDigraph
	}{
		{
			name: "",
			V:    8,
			edges: []DirectedEdge{
				&directedEdge{0, 2, 0.26},
				&directedEdge{0, 4, 0.38},
				&directedEdge{1, 3, 0.29},
				&directedEdge{2, 7, 0.34},
				&directedEdge{3, 6, 0.52},
				&directedEdge{4, 5, 0.35},
				&directedEdge{4, 7, 0.37},
				&directedEdge{5, 1, 0.32},
				&directedEdge{5, 4, 0.35},
				&directedEdge{5, 7, 0.28},
				&directedEdge{6, 0, 0.58},
				&directedEdge{6, 2, 0.40},
				&directedEdge{6, 4, 0.93},
				&directedEdge{7, 3, 0.39},
				&directedEdge{7, 5, 0.28},
			},
			expectedV:          8,
			expectedE:          15,
			expectedInDegrees:  []int{1, 1, 2, 2, 3, 2, 1, 3},
			expectedOutDegrees: []int{2, 1, 1, 1, 2, 3, 3, 2},
			expectedAdjacents: [][]DirectedEdge{
				[]DirectedEdge{
					&directedEdge{0, 2, 0.26},
					&directedEdge{0, 4, 0.38},
				},
				[]DirectedEdge{
					&directedEdge{1, 3, 0.29},
				},
				[]DirectedEdge{
					&directedEdge{2, 7, 0.34},
				},
				[]DirectedEdge{
					&directedEdge{3, 6, 0.52},
				},
				[]DirectedEdge{
					&directedEdge{4, 5, 0.35},
					&directedEdge{4, 7, 0.37},
				},
				[]DirectedEdge{
					&directedEdge{5, 1, 0.32},
					&directedEdge{5, 4, 0.35},
					&directedEdge{5, 7, 0.28},
				},
				[]DirectedEdge{
					&directedEdge{6, 0, 0.58},
					&directedEdge{6, 2, 0.40},
					&directedEdge{6, 4, 0.93},
				},
				[]DirectedEdge{
					&directedEdge{7, 3, 0.39},
					&directedEdge{7, 5, 0.28},
				},
			},
			expectedEdges: []DirectedEdge{
				&directedEdge{0, 2, 0.26},
				&directedEdge{0, 4, 0.38},
				&directedEdge{1, 3, 0.29},
				&directedEdge{2, 7, 0.34},
				&directedEdge{3, 6, 0.52},
				&directedEdge{4, 5, 0.35},
				&directedEdge{4, 7, 0.37},
				&directedEdge{5, 1, 0.32},
				&directedEdge{5, 4, 0.35},
				&directedEdge{5, 7, 0.28},
				&directedEdge{6, 0, 0.58},
				&directedEdge{6, 2, 0.40},
				&directedEdge{6, 4, 0.93},
				&directedEdge{7, 3, 0.39},
				&directedEdge{7, 5, 0.28},
			},
			expectedReverse: &weightedDigraph{
				v:   8,
				e:   15,
				ins: []int{2, 1, 1, 1, 2, 3, 3, 2},
				adj: [][]DirectedEdge{
					[]DirectedEdge{
						&directedEdge{0, 6, 0.58},
					},
					[]DirectedEdge{
						&directedEdge{1, 5, 0.32},
					},
					[]DirectedEdge{
						&directedEdge{2, 0, 0.26},
						&directedEdge{2, 6, 0.40},
					},
					[]DirectedEdge{
						&directedEdge{3, 1, 0.29},
						&directedEdge{3, 7, 0.39},
					},
					[]DirectedEdge{
						&directedEdge{4, 0, 0.38},
						&directedEdge{4, 5, 0.35},
						&directedEdge{4, 6, 0.93},
					},
					[]DirectedEdge{
						&directedEdge{5, 4, 0.35},
						&directedEdge{5, 7, 0.28},
					},
					[]DirectedEdge{
						&directedEdge{6, 3, 0.52},
					},
					[]DirectedEdge{
						&directedEdge{7, 2, 0.34},
						&directedEdge{7, 4, 0.37},
						&directedEdge{7, 5, 0.28},
					},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := NewWeightedDigraph(tc.V, tc.edges...)

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
			assert.NotEmpty(t, g.Graphviz())
		})
	}
}
