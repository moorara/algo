package grammar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTerminalsAndEndmarker(t *testing.T) {
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
			f := newTerminalsAndEndmarker(tc.terms...)
			assert.NotNil(t, f)
			assert.True(t, f.Terminals.Contains(tc.terms...))
			assert.False(t, f.IncludesEndmarker)
		})
	}
}

func TestNewFollowTable(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		follow := newFollowTable()
		assert.NotNil(t, follow)
	})
}
