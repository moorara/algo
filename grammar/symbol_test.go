package grammar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
