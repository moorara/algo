package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEdge(t *testing.T) {
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
			e := NewEdge(tc.v, tc.w, tc.weight)

			assert.NotEmpty(t, e)
			assert.Equal(t, tc.v, e.Either())
			assert.Equal(t, tc.v, e.Other(tc.w))
			assert.Equal(t, tc.w, e.Other(tc.v))
			assert.Equal(t, tc.weight, e.Weight())
		})
	}
}

func TestWeightedGraph(t *testing.T) {
	tests := []struct {
		name              string
		V                 int
		edges             []Edge
		expectedV         int
		expectedE         int
		expectedDegrees   []int
		expectedAdjacents [][]Edge
		expectedEdges     []Edge
	}{
		{
			name: "Connected",
			V:    8,
			edges: []Edge{
				&edge{0, 2, 0.26},
				&edge{0, 4, 0.38},
				&edge{0, 7, 0.16},
				&edge{1, 2, 0.36},
				&edge{1, 3, 0.29},
				&edge{1, 5, 0.32},
				&edge{1, 7, 0.19},
				&edge{2, 3, 0.17},
				&edge{2, 7, 0.34},
				&edge{3, 6, 0.52},
				&edge{4, 5, 0.35},
				&edge{4, 6, 0.93},
				&edge{4, 7, 0.37},
				&edge{5, 7, 0.28},
				&edge{6, 0, 0.58},
				&edge{6, 2, 0.40},
			},
			expectedV:       8,
			expectedE:       16,
			expectedDegrees: []int{4, 4, 5, 3, 4, 3, 4},
			expectedAdjacents: [][]Edge{
				[]Edge{
					&edge{0, 2, 0.26},
					&edge{0, 4, 0.38},
					&edge{0, 7, 0.16},
					&edge{6, 0, 0.58},
				},
				[]Edge{
					&edge{1, 2, 0.36},
					&edge{1, 3, 0.29},
					&edge{1, 5, 0.32},
					&edge{1, 7, 0.19},
				},
				[]Edge{
					&edge{0, 2, 0.26},
					&edge{1, 2, 0.36},
					&edge{2, 3, 0.17},
					&edge{2, 7, 0.34},
					&edge{6, 2, 0.40},
				},
				[]Edge{
					&edge{1, 3, 0.29},
					&edge{2, 3, 0.17},
					&edge{3, 6, 0.52},
				},
				[]Edge{
					&edge{0, 4, 0.38},
					&edge{4, 5, 0.35},
					&edge{4, 6, 0.93},
					&edge{4, 7, 0.37},
				},
				[]Edge{
					&edge{1, 5, 0.32},
					&edge{4, 5, 0.35},
					&edge{5, 7, 0.28},
				},
				[]Edge{
					&edge{3, 6, 0.52},
					&edge{4, 6, 0.93},
					&edge{6, 0, 0.58},
					&edge{6, 2, 0.40},
				},
				[]Edge{
					&edge{0, 7, 0.16},
					&edge{1, 7, 0.19},
					&edge{2, 7, 0.34},
					&edge{4, 7, 0.37},
					&edge{5, 7, 0.28},
				},
			},
			expectedEdges: []Edge{
				&edge{0, 2, 0.26},
				&edge{0, 4, 0.38},
				&edge{0, 7, 0.16},
				&edge{6, 0, 0.58},
				&edge{1, 2, 0.36},
				&edge{1, 3, 0.29},
				&edge{1, 5, 0.32},
				&edge{1, 7, 0.19},
				&edge{2, 3, 0.17},
				&edge{2, 7, 0.34},
				&edge{6, 2, 0.40},
				&edge{3, 6, 0.52},
				&edge{4, 5, 0.35},
				&edge{4, 6, 0.93},
				&edge{4, 7, 0.37},
				&edge{5, 7, 0.28},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := NewWeightedGraph(tc.V, tc.edges...)

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
			assert.NotEmpty(t, g.Graphviz())
		})
	}
}
