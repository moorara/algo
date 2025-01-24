package automata

import (
	"testing"

	"github.com/moorara/algo/symboltable"
	"github.com/stretchr/testify/assert"
)

func TestGroup_Equal(t *testing.T) {
	tests := []struct {
		name          string
		g             group
		rhs           group
		expectedEqual bool
	}{
		{
			name:          "Equal",
			g:             group{rep: 0, states: States{0, 1, 2}},
			rhs:           group{rep: 1, states: States{0, 1, 2}},
			expectedEqual: true,
		},
		{
			name:          "NotEqual",
			g:             group{rep: 0, states: States{0, 1, 2}},
			rhs:           group{rep: 1, states: States{0, 1, 2, 3}},
			expectedEqual: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedEqual, tc.g.Equal(tc.rhs))
		})
	}
}

func TestGroups_Contains(t *testing.T) {
	tests := []struct {
		name             string
		g                groups
		h                group
		expectedContains bool
	}{
		{
			name: "Yes",
			g: groups{
				group{rep: 0, states: States{0}},
				group{rep: 1, states: States{1, 2}},
				group{rep: 2, states: States{3, 4}},
			},
			h:                group{rep: 1, states: States{1, 2}},
			expectedContains: true,
		},
		{
			name: "No",
			g: groups{
				group{rep: 0, states: States{0}},
				group{rep: 1, states: States{1, 2}},
				group{rep: 2, states: States{3, 4}},
			},
			h:                group{rep: 1, states: States{2, 4}},
			expectedContains: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedContains, tc.g.Contains(tc.h))
		})
	}
}

func TestGroups_Equal(t *testing.T) {
	tests := []struct {
		name          string
		g             groups
		rhs           groups
		expectedEqual bool
	}{
		{
			name: "Equal",
			g: groups{
				group{rep: 0, states: States{0}},
				group{rep: 1, states: States{1, 2}},
				group{rep: 2, states: States{3, 4}},
			},
			rhs: groups{
				group{rep: 0, states: States{0}},
				group{rep: 1, states: States{1, 2}},
				group{rep: 2, states: States{3, 4}},
			},
			expectedEqual: true,
		},
		{
			name: "NotEqual",
			g: groups{
				group{rep: 0, states: States{0}},
				group{rep: 1, states: States{1, 2}},
				group{rep: 2, states: States{3, 4}},
			},
			rhs: groups{
				group{rep: 0, states: States{0}},
				group{rep: 1, states: States{1, 2}},
				group{rep: 2, states: States{3}},
				group{rep: 3, states: States{4}},
			},
			expectedEqual: false,
		},
		{
			name: "NotEqual",
			g: groups{
				group{rep: 0, states: States{0}},
				group{rep: 1, states: States{1, 2}},
				group{rep: 2, states: States{3, 4}},
			},
			rhs: groups{
				group{rep: 0, states: States{0}},
				group{rep: 1, states: States{1, 2}},
				group{rep: 2, states: States{3, 4}},
				group{rep: 3, states: States{5, 6}},
			},
			expectedEqual: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedEqual, tc.g.Equal(tc.rhs))
		})
	}
}

func TestNewPartition(t *testing.T) {
	p := newPartition()
	assert.NotNil(t, p)
}

func TestPartition_Equal(t *testing.T) {
	tests := []struct {
		name          string
		p             *partition
		rhs           *partition
		expectedEqual bool
	}{
		{
			name: "Equal",
			p: &partition{
				groups: groups{
					group{rep: 0, states: States{0}},
					group{rep: 1, states: States{1, 2}},
				},
				nextRep: 2,
			},
			rhs: &partition{
				groups: groups{
					group{rep: 0, states: States{0}},
					group{rep: 1, states: States{1, 2}},
				},
				nextRep: 2,
			},
			expectedEqual: true,
		},
		{
			name: "NotEqual",
			p: &partition{
				groups: groups{
					group{rep: 0, states: States{0}},
					group{rep: 1, states: States{1, 2}},
				},
				nextRep: 2,
			},
			rhs: &partition{
				groups: groups{
					group{rep: 0, states: States{0}},
					group{rep: 1, states: States{1}},
					group{rep: 2, states: States{2}},
				},
				nextRep: 3,
			},
			expectedEqual: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedEqual, tc.p.Equal(tc.rhs))
		})
	}
}

func TestPartition_Add(t *testing.T) {
	tests := []struct {
		name              string
		p                 *partition
		groups            []States
		expectedPartition *partition
	}{
		{
			name: "OK",
			p: &partition{
				groups: groups{
					group{rep: 0, states: States{0}},
				},
				nextRep: 1,
			},
			groups: []States{
				{1, 2},
				{3, 4},
			},
			expectedPartition: &partition{
				groups: groups{
					group{rep: 0, states: States{0}},
					group{rep: 1, states: States{1, 2}},
					group{rep: 2, states: States{3, 4}},
				},
				nextRep: 4,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.p.Add(tc.groups...)
			assert.True(t, tc.p.Equal(tc.expectedPartition))
		})
	}
}

func TestPartition_Rep(t *testing.T) {
	tests := []struct {
		name        string
		p           *partition
		s           State
		expectedOK  bool
		expectedRep State
	}{
		{
			name: "OK",
			p: &partition{
				groups: groups{
					group{rep: 0, states: States{0}},
					group{rep: 1, states: States{1, 2}},
					group{rep: 2, states: States{3, 4}},
				},
				nextRep: 4,
			},
			s:           State(4),
			expectedOK:  true,
			expectedRep: State(2),
		},
		{
			name: "NotOK",
			p: &partition{
				groups: groups{
					group{rep: 0, states: States{0}},
					group{rep: 1, states: States{1, 2}},
					group{rep: 2, states: States{3, 4}},
				},
				nextRep: 4,
			},
			s:           State(8),
			expectedOK:  false,
			expectedRep: -1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			rep, ok := tc.p.Rep(tc.s)
			assert.Equal(t, tc.expectedOK, ok)
			assert.Equal(t, tc.expectedRep, rep)
		})
	}
}

func TestPartition_BuildGroupTrans(t *testing.T) {
	dfa := getTestDFAs()[2]

	trans := symboltable.NewRedBlack(cmpState, eqSymbolState)

	s2Trans := symboltable.NewRedBlack(cmpSymbol, eqState)
	s2Trans.Put('b', 1)
	trans.Put(2, s2Trans)

	s4Trans := symboltable.NewRedBlack(cmpSymbol, eqState)
	s4Trans.Put('a', 1)
	trans.Put(4, s4Trans)

	tests := []struct {
		name          string
		p             *partition
		dfa           *DFA
		G             group
		expectedTrans doubleKeyMap[State, Symbol, State]
	}{
		{
			name: "OK",
			p: &partition{
				groups: groups{
					group{rep: 0, states: States{0, 1, 3}},
					group{rep: 1, states: States{2, 4}},
				},
				nextRep: 2,
			},
			dfa:           dfa,
			G:             group{rep: 1, states: States{2, 4}},
			expectedTrans: trans,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			trans := tc.p.BuildGroupTrans(tc.dfa, tc.G)
			assert.True(t, trans.Equal(tc.expectedTrans))
		})
	}
}

func TestPartition_PartitionAndAddGroups(t *testing.T) {
	trans := symboltable.NewRedBlack(cmpState, eqSymbolState)

	s2Trans := symboltable.NewRedBlack(cmpSymbol, eqState)
	s2Trans.Put('b', 1)
	trans.Put(2, s2Trans)

	s4Trans := symboltable.NewRedBlack(cmpSymbol, eqState)
	s4Trans.Put('a', 1)
	trans.Put(4, s4Trans)

	tests := []struct {
		name              string
		p                 *partition
		Gtrans            doubleKeyMap[State, Symbol, State]
		expectedPartition *partition
	}{
		{
			name: "OK",
			p: &partition{
				groups:  groups{},
				nextRep: 0,
			},
			Gtrans: trans,
			expectedPartition: &partition{
				groups: groups{
					group{rep: 0, states: States{2}},
					group{rep: 1, states: States{4}},
				},
				nextRep: 2,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.p.PartitionAndAddGroups(tc.Gtrans)
			assert.True(t, tc.p.Equal(tc.expectedPartition))
		})
	}
}
