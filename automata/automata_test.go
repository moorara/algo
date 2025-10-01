package automata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStates(t *testing.T) {
	tests := []struct {
		name string
		s    []State
	}{
		{
			name: "OK",
			s:    []State{0, 1, 2, 3},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			set := NewStates(tc.s...)

			assert.NotNil(t, set)
			assert.True(t, set.Contains(tc.s...))
		})
	}
}

func TestNewSymbols(t *testing.T) {
	tests := []struct {
		name string
		a    []Symbol
	}{
		{
			name: "OK",
			a:    []Symbol{'a', 'b', 'c', 'd'},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			set := NewSymbols(tc.a...)

			assert.NotNil(t, set)
			assert.True(t, set.Contains(tc.a...))
		})
	}
}
