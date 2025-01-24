package grammar

import (
	"bytes"
	"errors"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteSymbol(t *testing.T) {
	tests := []struct {
		name          string
		w             io.Writer
		s             Symbol
		expectedN     int
		expectedError error
	}{
		{
			name:          "OK",
			w:             new(bytes.Buffer),
			s:             NonTerminal("expr"),
			expectedN:     4,
			expectedError: nil,
		},
		{
			name: "Error",
			w: &MockWriter{
				WriteMocks: []WriteMock{
					{OutN: 0, OutError: errors.New("error on write")},
				},
			},
			s:             NonTerminal("expr"),
			expectedN:     0,
			expectedError: errors.New("error on write"),
		},
	}

	for _, tc := range tests {
		n, err := WriteSymbol(tc.w, tc.s)
		assert.Equal(t, tc.expectedN, n)
		assert.Equal(t, tc.expectedError, err)
	}
}

func TestCompareFuncForSymbol(t *testing.T) {
	tests := []struct {
		name            string
		lhs             Symbol
		rhs             Symbol
		expectedCompare int
	}{
		{
			name:            "Terminal_NonTerminal",
			lhs:             Terminal("a"),
			rhs:             NonTerminal("A"),
			expectedCompare: -1,
		},
		{
			name:            "NonTerminal_Terminal",
			lhs:             NonTerminal("A"),
			rhs:             Terminal("a"),
			expectedCompare: 1,
		},
		{
			name:            "Terminal_Terminal",
			lhs:             Terminal("a"),
			rhs:             Terminal("b"),
			expectedCompare: -1,
		},
		{
			name:            "NonTerminal_NonTerminal",
			lhs:             NonTerminal("B"),
			rhs:             NonTerminal("A"),
			expectedCompare: 1,
		},
		{
			name:            "Terminal_Terminal_Equal",
			lhs:             Terminal("a"),
			rhs:             Terminal("a"),
			expectedCompare: 0,
		},
		{
			name:            "NonTerminal_NonTerminal_Equal",
			lhs:             NonTerminal("A"),
			rhs:             NonTerminal("A"),
			expectedCompare: 0,
		},
	}

	for _, tc := range tests {
		cmp := compareFuncForSymbol()(tc.lhs, tc.rhs)
		assert.Equal(t, tc.expectedCompare, cmp)
	}
}

func TestHashFuncForSymbol(t *testing.T) {
	tests := []struct {
		s            Symbol
		expectedHash uint64
	}{
		{
			s:            Terminal("map"),
			expectedHash: 0x7d481fd1761e83a5,
		},
		{
			s:            NonTerminal("map"),
			expectedHash: 0xd8b3c5186b8ca065,
		},
	}

	for _, tc := range tests {
		hash := hashFuncForSymbol()(tc.s)
		assert.Equal(t, tc.expectedHash, hash)
	}
}

func TestTerminal(t *testing.T) {
	t.Run("Endmarker", func(t *testing.T) {
		assert.Equal(t, "$", Endmarker.String())
		assert.Equal(t, "$", Endmarker.Name())
		assert.True(t, Endmarker.Equal(Terminal("\uEEEE")))
		assert.False(t, Endmarker.Equal(NonTerminal("\uEEEE")))
		assert.False(t, Endmarker.Equal(Terminal("\uEEEF")))
		assert.True(t, Endmarker.IsTerminal())
	})

	tests := []struct {
		value          string
		expectedString string
	}{
		{value: "a", expectedString: `"a"`},
		{value: "b", expectedString: `"b"`},
		{value: "0", expectedString: `"0"`},
		{value: "1", expectedString: `"1"`},
		{value: "2", expectedString: `"2"`},
		{value: "+", expectedString: `"+"`},
		{value: "*", expectedString: `"*"`},
		{value: "(", expectedString: `"("`},
		{value: ")", expectedString: `")"`},
		{value: "id", expectedString: `"id"`},
		{value: "if", expectedString: `"if"`},
		{value: "if", expectedString: `"if"`},
	}

	notEqual := Terminal("🙂")

	for _, tc := range tests {
		t.Run(tc.value, func(t *testing.T) {
			tr := Terminal(tc.value)
			assert.Equal(t, tc.expectedString, tr.String())
			assert.True(t, tr.Equal(Terminal(tc.value)))
			assert.False(t, tr.Equal(NonTerminal(tc.value)))
			assert.False(t, tr.Equal(notEqual))
			assert.Equal(t, tc.value, tr.Name())
			assert.True(t, tr.IsTerminal())
		})
	}
}

func TestNonTerminal(t *testing.T) {
	tests := []struct {
		value string
	}{
		{value: "A"},
		{value: "B"},
		{value: "C"},
		{value: "S"},
		{value: "expr"},
		{value: "stmt"},
	}

	notEqual := NonTerminal("🙂")

	for _, tc := range tests {
		t.Run(tc.value, func(t *testing.T) {
			n := NonTerminal(tc.value)
			assert.Equal(t, tc.value, n.String())
			assert.True(t, n.Equal(NonTerminal(tc.value)))
			assert.False(t, n.Equal(Terminal(tc.value)))
			assert.False(t, n.Equal(notEqual))
			assert.Equal(t, tc.value, n.Name())
			assert.False(t, n.IsTerminal())
		})
	}
}
