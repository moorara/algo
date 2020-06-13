package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlowEdge(t *testing.T) {
	tests := []struct {
		name         string
		from         int
		to           int
		capacity     float64
		flow         float64
		v            int
		delta        float64
		expectedOK   bool
		expectedFlow float64
	}{
		{
			name:         "Forward",
			from:         1,
			to:           2,
			capacity:     0.50,
			flow:         0.25,
			v:            2,
			delta:        0.25,
			expectedOK:   true,
			expectedFlow: 0.50,
		},
		{
			name:         "Backward",
			from:         1,
			to:           2,
			capacity:     0.50,
			flow:         0.25,
			v:            1,
			delta:        0.25,
			expectedOK:   true,
			expectedFlow: 0.00,
		},
		{
			name:         "InvalidForward",
			from:         1,
			to:           2,
			capacity:     0.50,
			flow:         0.50,
			v:            2,
			delta:        0.25,
			expectedOK:   false,
			expectedFlow: 0.50,
		},
		{
			name:         "InvalidBackward",
			from:         1,
			to:           2,
			capacity:     0.50,
			flow:         0.00,
			v:            1,
			delta:        0.25,
			expectedOK:   false,
			expectedFlow: 0.00,
		},
		{
			name:         "InvalidVertex",
			from:         1,
			to:           2,
			capacity:     0.50,
			flow:         0.00,
			v:            0,
			delta:        0.25,
			expectedOK:   false,
			expectedFlow: 0.00,
		},
		{
			name:         "InvalidDelta",
			from:         1,
			to:           2,
			capacity:     0.50,
			flow:         0.00,
			v:            2,
			delta:        -0.25,
			expectedOK:   false,
			expectedFlow: 0.00,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			e := FlowEdge{tc.from, tc.to, tc.capacity, tc.flow}

			assert.NotEmpty(t, e)

			assert.Equal(t, tc.from, e.From())
			assert.Equal(t, tc.to, e.To())

			assert.Equal(t, tc.to, e.Other(tc.from))
			assert.Equal(t, tc.from, e.Other(tc.to))
			assert.Equal(t, -1, e.Other(99))

			assert.Equal(t, tc.capacity, e.Capacity())
			assert.Equal(t, tc.flow, e.Flow())

			assert.Equal(t, tc.flow, e.ResidualCapacityTo(e.from))
			assert.Equal(t, tc.capacity-tc.flow, e.ResidualCapacityTo(e.to))
			assert.Equal(t, float64(-1), e.ResidualCapacityTo(99))

			ok := e.AddResidualFlowTo(tc.v, tc.delta)
			assert.Equal(t, tc.expectedOK, ok)
			assert.Equal(t, tc.expectedFlow, e.flow, float64Epsilon)
		})
	}
}

func TestFlowNetwork(t *testing.T) {
	tests := []struct {
		name              string
		V                 int
		edges             []FlowEdge
		expectedV         int
		expectedE         int
		expectedAdjacents [][]FlowEdge
		expectedEdges     []FlowEdge
	}{
		{
			name: "Flow",
			V:    7,
			edges: []FlowEdge{
				{0, 1, 2.00, 2.00},
				{0, 2, 3.00, 1.00},
				{1, 3, 3.00, 2.00},
				{1, 4, 1.00, 0.00},
				{2, 3, 1.00, 0.00},
				{2, 4, 1.00, 1.00},
				{3, 5, 2.00, 2.00},
				{4, 5, 3.00, 1.00},
			},
			expectedV: 7,
			expectedE: 8,
			expectedAdjacents: [][]FlowEdge{
				{
					{0, 1, 2.00, 2.00},
					{0, 2, 3.00, 1.00},
				},
				{
					{0, 1, 2.00, 2.00},
					{1, 3, 3.00, 2.00},
					{1, 4, 1.00, 0.00},
				},
				{
					{0, 2, 3.00, 1.00},
					{2, 3, 1.00, 0.00},
					{2, 4, 1.00, 1.00},
				},
				{
					{1, 3, 3.00, 2.00},
					{2, 3, 1.00, 0.00},
					{3, 5, 2.00, 2.00},
				},
				{
					{1, 4, 1.00, 0.00},
					{2, 4, 1.00, 1.00},
					{4, 5, 3.00, 1.00},
				},
				{
					{3, 5, 2.00, 2.00},
					{4, 5, 3.00, 1.00},
				},
			},
			expectedEdges: []FlowEdge{
				{0, 1, 2.00, 2.00},
				{0, 2, 3.00, 1.00},
				{1, 3, 3.00, 2.00},
				{1, 4, 1.00, 0.00},
				{2, 3, 1.00, 0.00},
				{2, 4, 1.00, 1.00},
				{3, 5, 2.00, 2.00},
				{4, 5, 3.00, 1.00},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := NewFlowNetwork(tc.V, tc.edges...)

			assert.NotEmpty(t, g)
			assert.Equal(t, tc.expectedV, g.V())
			assert.Equal(t, tc.expectedE, g.E())

			t.Run("Adjacency", func(t *testing.T) {
				assert.Nil(t, g.Adj(-1))
				for v, expectedAdj := range tc.expectedAdjacents {
					assert.Equal(t, expectedAdj, g.Adj(v))
				}
			})

			t.Run("Edges", func(t *testing.T) {
				assert.Equal(t, tc.expectedEdges, g.Edges())
			})

			assert.NotEmpty(t, g.Graphviz())
		})
	}
}
