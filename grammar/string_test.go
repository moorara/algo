package grammar

import (
	"bytes"
	"errors"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/generic"
)

func TestString(t *testing.T) {
	tests := []struct {
		name                   string
		s                      String[Symbol]
		expectedString         string
		containsSymbol         Symbol
		expectedContainsSymbol bool
		prefix                 String[Symbol]
		expectedHasPrefix      bool
		suffix                 String[Symbol]
		expectedHasSuffix      bool
		anyMatch               generic.Predicate1[Symbol]
		expectedAnyMatch       bool
		append                 []Symbol
		expectedAppend         String[Symbol]
		concat                 []String[Symbol]
		expectedConcat         String[Symbol]
		expectedTerminals      String[Terminal]
		expectedNonTerminals   String[NonTerminal]
	}{
		{
			name:                   "Empty",
			s:                      E,
			expectedString:         `ε`,
			containsSymbol:         Terminal(""),
			expectedContainsSymbol: false,
			prefix:                 String[Symbol]{},
			expectedHasPrefix:      true,
			suffix:                 String[Symbol]{},
			expectedHasSuffix:      true,
			anyMatch:               func(s Symbol) bool { return true },
			expectedAnyMatch:       false,
			append:                 []Symbol{},
			expectedAppend:         E,
			concat:                 []String[Symbol]{E},
			expectedConcat:         E,
			expectedTerminals:      String[Terminal]{},
			expectedNonTerminals:   String[NonTerminal]{},
		},
		{
			name:                   "AllTerminals",
			s:                      String[Symbol]{Terminal("a"), Terminal("b"), Terminal("c")},
			expectedString:         `"a" "b" "c"`,
			containsSymbol:         Terminal("b"),
			expectedContainsSymbol: true,
			prefix:                 String[Symbol]{Terminal("a"), Terminal("b")},
			expectedHasPrefix:      true,
			suffix:                 String[Symbol]{Terminal("a"), Terminal("c")},
			expectedHasSuffix:      false,
			anyMatch:               func(s Symbol) bool { return !s.IsTerminal() },
			expectedAnyMatch:       false,
			append:                 []Symbol{Terminal("d")},
			expectedAppend:         String[Symbol]{Terminal("a"), Terminal("b"), Terminal("c"), Terminal("d")},
			concat:                 []String[Symbol]{{Terminal("d"), Terminal("e"), Terminal("f")}},
			expectedConcat:         String[Symbol]{Terminal("a"), Terminal("b"), Terminal("c"), Terminal("d"), Terminal("e"), Terminal("f")},
			expectedTerminals:      String[Terminal]{"a", "b", "c"},
			expectedNonTerminals:   String[NonTerminal]{},
		},
		{
			name:                   "AllNonTerminals",
			s:                      String[Symbol]{NonTerminal("A"), NonTerminal("B"), NonTerminal("C")},
			expectedString:         `A B C`,
			containsSymbol:         NonTerminal("B"),
			expectedContainsSymbol: true,
			prefix:                 String[Symbol]{NonTerminal("A"), NonTerminal("C")},
			expectedHasPrefix:      false,
			suffix:                 String[Symbol]{NonTerminal("B"), NonTerminal("C")},
			expectedHasSuffix:      true,
			anyMatch:               func(s Symbol) bool { return s.IsTerminal() },
			expectedAnyMatch:       false,
			append:                 []Symbol{NonTerminal("D")},
			expectedAppend:         String[Symbol]{NonTerminal("A"), NonTerminal("B"), NonTerminal("C"), NonTerminal("D")},
			concat:                 []String[Symbol]{{NonTerminal("D"), NonTerminal("E"), NonTerminal("F")}},
			expectedConcat:         String[Symbol]{NonTerminal("A"), NonTerminal("B"), NonTerminal("C"), NonTerminal("D"), NonTerminal("E"), NonTerminal("F")},
			expectedTerminals:      String[Terminal]{},
			expectedNonTerminals:   String[NonTerminal]{"A", "B", "C"},
		},
		{
			name:                   "TerminalsAndNonTerminals",
			s:                      String[Symbol]{Terminal("a"), NonTerminal("A"), Terminal("b"), NonTerminal("B"), Terminal("c")},
			expectedString:         `"a" A "b" B "c"`,
			containsSymbol:         NonTerminal("C"),
			expectedContainsSymbol: false,
			prefix:                 String[Symbol]{Terminal("a"), NonTerminal("A"), Terminal("b"), NonTerminal("B"), Terminal("c")},
			expectedHasPrefix:      true,
			suffix:                 String[Symbol]{Terminal("a"), NonTerminal("A"), Terminal("b"), NonTerminal("B"), Terminal("c")},
			expectedHasSuffix:      true,
			anyMatch:               func(s Symbol) bool { return s.IsTerminal() },
			expectedAnyMatch:       true,
			append:                 []Symbol{NonTerminal("C"), Terminal("d"), NonTerminal("D")},
			expectedAppend:         String[Symbol]{Terminal("a"), NonTerminal("A"), Terminal("b"), NonTerminal("B"), Terminal("c"), NonTerminal("C"), Terminal("d"), NonTerminal("D")},
			concat:                 []String[Symbol]{{NonTerminal("C")}, {Terminal("d"), NonTerminal("D")}},
			expectedConcat:         String[Symbol]{Terminal("a"), NonTerminal("A"), Terminal("b"), NonTerminal("B"), Terminal("c"), NonTerminal("C"), Terminal("d"), NonTerminal("D")},
			expectedTerminals:      String[Terminal]{"a", "b", "c"},
			expectedNonTerminals:   String[NonTerminal]{"A", "B"},
		},
	}

	notEqual := String[Symbol]{Terminal("🙂"), NonTerminal("🙃")}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedString, tc.s.String())
			assert.True(t, tc.s.Equal(tc.s))
			assert.False(t, tc.s.Equal(notEqual))
			assert.Equal(t, tc.expectedContainsSymbol, tc.s.ContainsSymbol(tc.containsSymbol))
			assert.Equal(t, tc.expectedHasPrefix, tc.s.HasPrefix(tc.prefix))
			assert.Equal(t, tc.expectedHasSuffix, tc.s.HasSuffix(tc.suffix))
			assert.Equal(t, tc.expectedAnyMatch, tc.s.AnyMatch(tc.anyMatch))
			assert.Equal(t, tc.expectedAppend, tc.s.Append(tc.append...))
			assert.Equal(t, tc.expectedConcat, tc.s.Concat(tc.concat...))
			assert.Equal(t, tc.expectedTerminals, tc.s.Terminals())
			assert.Equal(t, tc.expectedNonTerminals, tc.s.NonTerminals())
		})
	}
}

func TestLongestCommonPrefixOf(t *testing.T) {
	tests := []struct {
		name                        string
		ss                          []String[Symbol]
		expectedLongestCommonPrefix String[Symbol]
	}{
		{
			name:                        "Empty",
			ss:                          []String[Symbol]{},
			expectedLongestCommonPrefix: E,
		},
		{
			name: "NoCommonPrefix",
			ss: []String[Symbol]{
				{NonTerminal("expr"), Terminal("?"), NonTerminal("stmt"), Terminal(":"), NonTerminal("stmt")},
				{Terminal("if"), NonTerminal("expr"), Terminal("then"), NonTerminal("stmt"), Terminal("else"), NonTerminal("stmt")},
			},
			expectedLongestCommonPrefix: E,
		},
		{
			name: "CommonPrefix",
			ss: []String[Symbol]{
				{Terminal("if"), NonTerminal("expr"), Terminal("then"), NonTerminal("stmt"), Terminal("else"), NonTerminal("stmt")},
				{Terminal("if"), NonTerminal("expr"), Terminal("then"), NonTerminal("stmt")},
			},
			expectedLongestCommonPrefix: String[Symbol]{Terminal("if"), NonTerminal("expr"), Terminal("then"), NonTerminal("stmt")},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			lcp := LongestCommonPrefixOf(tc.ss...)
			assert.Equal(t, tc.expectedLongestCommonPrefix, lcp)
		})
	}
}

func TestWriteString(t *testing.T) {
	tests := []struct {
		name          string
		w             io.Writer
		s             String[Symbol]
		expectedN     int
		expectedError error
	}{
		{
			name:          "OK",
			w:             new(bytes.Buffer),
			s:             String[Symbol]{Terminal("if"), NonTerminal("expr"), Terminal("then"), NonTerminal("stmt")},
			expectedN:     18,
			expectedError: nil,
		},
		{
			name: "Error",
			w: &MockWriter{
				WriteMocks: []WriteMock{
					{OutN: 0, OutError: errors.New("error on write")},
				},
			},
			s:             String[Symbol]{Terminal("if"), NonTerminal("expr"), Terminal("then"), NonTerminal("stmt")},
			expectedN:     0,
			expectedError: errors.New("error on write"),
		},
	}

	for _, tc := range tests {
		n, err := WriteString(tc.w, tc.s)
		assert.Equal(t, tc.expectedN, n)
		assert.Equal(t, tc.expectedError, err)
	}
}

func TestHashFuncForString(t *testing.T) {
	tests := []struct {
		s            String[Symbol]
		expectedHash uint64
	}{
		{
			s:            String[Symbol]{Terminal("if"), NonTerminal("expr"), Terminal("then"), NonTerminal("stmt")},
			expectedHash: 0xb0616925421a7df6,
		},
		{
			s:            String[Symbol]{Terminal("if"), NonTerminal("expr"), Terminal("then"), NonTerminal("stmt"), Terminal("else"), NonTerminal("stmt")},
			expectedHash: 0xdf211ff9239df1ed,
		},
	}

	for _, tc := range tests {
		hash := hashFuncForString()(tc.s)
		assert.Equal(t, tc.expectedHash, hash)
	}
}
