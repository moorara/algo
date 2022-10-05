package trie

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBitPattern(t *testing.T) {
	tests := []struct {
		name    string
		pattern string
	}{
		{
			name:    "OK",
			pattern: "a**d",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			b := newBitPattern(tc.pattern)
			assert.NotNil(t, b)
		})
	}
}

func TestBitPattern_Bit(t *testing.T) {
	tests := []struct {
		name         string
		b            *bitPattern
		expectedBits []byte
	}{
		{
			name: "OK",
			b:    newBitPattern("G*a*i*y"),
			expectedBits: []byte{
				'0', '1', '0', '0', '0', '1', '1', '1',
				'*', '*', '*', '*', '*', '*', '*', '*',
				'0', '1', '1', '0', '0', '0', '0', '1',
				'*', '*', '*', '*', '*', '*', '*', '*',
				'0', '1', '1', '0', '1', '0', '0', '1',
				'*', '*', '*', '*', '*', '*', '*', '*',
				'0', '1', '1', '1', '1', '0', '0', '1',
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			for i, expectedBit := range tc.expectedBits {
				assert.Equal(t, expectedBit, tc.b.Bit(i+1))
			}
		})
	}
}
