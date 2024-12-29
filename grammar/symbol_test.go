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
	tests := []struct {
		value          string
		expectedString string
	}{
		{value: "a", expectedString: `"a"`},
		{value: "b", expectedString: `"b"`},
		{value: "c", expectedString: `"c"`},
		{value: "0", expectedString: `"0"`},
		{value: "1", expectedString: `"1"`},
		{value: "2", expectedString: `"2"`},
		{value: "+", expectedString: `"+"`},
		{value: "*", expectedString: `"*"`},
		{value: "(", expectedString: `"("`},
		{value: ")", expectedString: `")"`},
		{value: "id", expectedString: `"id"`},
		{value: "if", expectedString: `"if"`},
	}

	notEqual := Terminal("🙂")

	for _, tc := range tests {
		t.Run(tc.value, func(t *testing.T) {
			tr := Terminal(tc.value)
			assert.Equal(t, tc.expectedString, tr.String())
			assert.Equal(t, tc.value, tr.Name())
			assert.True(t, tr.Equals(Terminal(tc.value)))
			assert.False(t, tr.Equals(NonTerminal(tc.value)))
			assert.False(t, tr.Equals(notEqual))
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
			assert.Equal(t, tc.value, n.Name())
			assert.True(t, n.Equals(NonTerminal(tc.value)))
			assert.False(t, n.Equals(Terminal(tc.value)))
			assert.False(t, n.Equals(notEqual))
			assert.False(t, n.IsTerminal())
		})
	}
}
