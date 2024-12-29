package grammar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTerminalsAndEmpty(t *testing.T) {
	tests := []struct {
		name  string
		terms []Terminal
	}{
		{
			name:  "OK",
			terms: []Terminal{"a", "b", "c", "d", "e", "f"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			f := newTerminalsAndEmpty(tc.terms...)
			assert.NotNil(t, f)
			assert.True(t, f.Terminals.Contains(tc.terms...))
			assert.False(t, f.IncludesEmpty)
		})
	}
}

func TestNewFirstBySymbolTable(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		firstBySymbol := newFirstBySymbolTable()
		assert.NotNil(t, firstBySymbol)
	})
}

func TestNewFirstByStringTable(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		firstByString := newFirstByStringTable()
		assert.NotNil(t, firstByString)
	})
}