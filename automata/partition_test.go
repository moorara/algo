package automata

import (
	"testing"

	"github.com/moorara/algo/symboltable"
	"github.com/stretchr/testify/assert"
)

func TestNewGroups(t *testing.T) {
	tests := []struct {
		name string
		g    []group
	}{
		{
			name: "OK",
			g: []group{
				{States: NewStates(0), rep: 0},
				{States: NewStates(1, 2), rep: 1},
				{States: NewStates(3, 4), rep: 2},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			set := newGroups(tc.g...)

			assert.NotNil(t, set)
			assert.True(t, set.Contains(tc.g...))
		})
	}
}

func TestNewPartition(t *testing.T) {
	P := newPartition()
	assert.NotNil(t, P)
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
				groups: newGroups(
					group{States: NewStates(0), rep: 0},
				),
				nextRep: 1,
			},
			groups: []States{
				NewStates(1, 2),
				NewStates(3, 4),
			},
			expectedPartition: &partition{
				groups: newGroups(
					group{States: NewStates(0), rep: 0},
					group{States: NewStates(1, 2), rep: 1},
					group{States: NewStates(3, 4), rep: 2},
				),
				nextRep: 3,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.p.Add(tc.groups...)

			assert.True(t, tc.p.groups.Equal(tc.expectedPartition.groups))
			assert.Equal(t, tc.expectedPartition.nextRep, tc.p.nextRep)
		})
	}
}

func TestPartition_Rep(t *testing.T) {
	tests := []struct {
		name        string
		p           *partition
		s           State
		expectedRep State
	}{
		{
			name: "OK",
			p: &partition{
				groups: newGroups(
					group{States: NewStates(0), rep: 0},
					group{States: NewStates(1, 2), rep: 1},
					group{States: NewStates(3, 4), rep: 2},
				),
				nextRep: 4,
			},
			s:           State(4),
			expectedRep: State(2),
		},
		{
			name: "NotOK",
			p: &partition{
				groups: newGroups(
					group{States: NewStates(0), rep: 0},
					group{States: NewStates(1, 2), rep: 1},
					group{States: NewStates(3, 4), rep: 2},
				),
				nextRep: 4,
			},
			s:           State(8),
			expectedRep: State(-1),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			rep := tc.p.Rep(tc.s)

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
				groups: newGroups(
					group{States: NewStates(0, 1, 3), rep: 0},
					group{States: NewStates(2, 4), rep: 1},
				),
				nextRep: 2,
			},
			dfa:           dfa,
			G:             group{States: NewStates(2, 4), rep: 1},
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
				groups:  newGroups(),
				nextRep: 0,
			},
			Gtrans: trans,
			expectedPartition: &partition{
				groups: newGroups(
					group{States: NewStates(2), rep: 0},
					group{States: NewStates(4), rep: 1},
				),
				nextRep: 2,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.p.PartitionAndAddGroups(tc.Gtrans)

			assert.True(t, tc.p.groups.Equal(tc.expectedPartition.groups))
			assert.Equal(t, tc.expectedPartition.nextRep, tc.p.nextRep)
		})
	}
}
