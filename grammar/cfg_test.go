package grammar

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/set"
)

var CFGrammars = []CFG{
	NewCFG(
		[]Terminal{"0", "1"},
		[]NonTerminal{"S", "X", "Y"},
		[]Production{
			{"S", String[Symbol]{NonTerminal("X"), NonTerminal("Y"), NonTerminal("X")}}, // S → XYX
			{"X", String[Symbol]{Terminal("0"), NonTerminal("X")}},                      // X → 0X
			{"X", ε}, // X → ε
			{"Y", String[Symbol]{Terminal("1"), NonTerminal("Y")}}, // Y → 1Y
			{"Y", ε}, // Y → ε
		},
		"S",
	),
	NewCFG(
		[]Terminal{"a", "b"},
		[]NonTerminal{"S"},
		[]Production{
			{"S", String[Symbol]{Terminal("a"), NonTerminal("S"), Terminal("b"), NonTerminal("S")}}, // S → aSbS
			{"S", String[Symbol]{Terminal("b"), NonTerminal("S"), Terminal("a"), NonTerminal("S")}}, // S → bSaS
			{"S", ε}, // S → ε
		},
		"S",
	),
	NewCFG(
		[]Terminal{"a", "b"},
		[]NonTerminal{"S", "A", "B"},
		[]Production{
			{"S", String[Symbol]{Terminal("a"), NonTerminal("B"), Terminal("a")}}, // S → aBa
			{"S", String[Symbol]{NonTerminal("A"), Terminal("b")}},                // S → Ab
			{"S", String[Symbol]{Terminal("a")}},                                  // S → a
			{"A", String[Symbol]{Terminal("b")}},                                  // A → b
			{"A", ε},                                                              // A → ε
			{"B", String[Symbol]{NonTerminal("A")}},                               // B → A
			{"B", String[Symbol]{Terminal("b")}},                                  // B → b
		},
		"S",
	),
	NewCFG(
		[]Terminal{"b", "c", "d", "s"},
		[]NonTerminal{"S", "A", "B", "C", "D"},
		[]Production{
			{"S", String[Symbol]{NonTerminal("A")}}, // S → A
			{"S", String[Symbol]{Terminal("s")}},    // S → s
			{"A", String[Symbol]{NonTerminal("B")}}, // A → B
			{"B", String[Symbol]{NonTerminal("C")}}, // B → C
			{"B", String[Symbol]{Terminal("b")}},    // B → b
			{"C", String[Symbol]{NonTerminal("D")}}, // C → D
			{"D", String[Symbol]{Terminal("d")}},    // D → d
		},
		"S",
	),
	NewCFG(
		[]Terminal{"a", "b", "c", "d"},
		[]NonTerminal{"S", "A", "B", "C", "D"},
		[]Production{
			{"S", String[Symbol]{NonTerminal("A"), NonTerminal("B")}}, // S → AB
			{"A", String[Symbol]{Terminal("a"), NonTerminal("A")}},    // A → aA
			{"A", String[Symbol]{Terminal("a")}},                      // A → a
			{"B", String[Symbol]{Terminal("b"), NonTerminal("B")}},    // B → bB
			{"B", String[Symbol]{Terminal("b")}},                      // B → b
			{"C", String[Symbol]{Terminal("c"), NonTerminal("C")}},    // C → cC
			{"C", String[Symbol]{Terminal("c")}},                      // C → c
			{"D", String[Symbol]{Terminal("d")}},                      // D → d
		},
		"S",
	),
	NewCFG(
		[]Terminal{"+", "-", "*", "/", "(", ")", "id"},
		[]NonTerminal{"S", "E"},
		[]Production{
			{"S", String[Symbol]{NonTerminal("E")}},                                  // S → E
			{"E", String[Symbol]{NonTerminal("E"), Terminal("+"), NonTerminal("E")}}, // E → E + E
			{"E", String[Symbol]{NonTerminal("E"), Terminal("-"), NonTerminal("E")}}, // E → E - E
			{"E", String[Symbol]{NonTerminal("E"), Terminal("*"), NonTerminal("E")}}, // E → E * E
			{"E", String[Symbol]{NonTerminal("E"), Terminal("/"), NonTerminal("E")}}, // E → E / E
			{"E", String[Symbol]{Terminal("("), NonTerminal("E"), Terminal(")")}},    // E → ( E )
			{"E", String[Symbol]{Terminal("-"), NonTerminal("E")}},                   // E → - E
			{"E", String[Symbol]{Terminal("id")}},                                    // E → id
		},
		"S",
	),
	NewCFG(
		[]Terminal{"+", "-", "*", "/", "(", ")", "id"},
		[]NonTerminal{"S", "E", "T", "F"},
		[]Production{
			{"S", String[Symbol]{NonTerminal("E")}},                                  // S → E
			{"E", String[Symbol]{NonTerminal("E"), Terminal("+"), NonTerminal("T")}}, // E → E + T
			{"E", String[Symbol]{NonTerminal("E"), Terminal("-"), NonTerminal("T")}}, // E → E - T
			{"E", String[Symbol]{NonTerminal("T")}},                                  // E → T
			{"T", String[Symbol]{NonTerminal("T"), Terminal("*"), NonTerminal("F")}}, // T → T * F
			{"T", String[Symbol]{NonTerminal("T"), Terminal("/"), NonTerminal("F")}}, // T → T / F
			{"T", String[Symbol]{NonTerminal("F")}},                                  // T → F
			{"F", String[Symbol]{Terminal("("), NonTerminal("E"), Terminal(")")}},    // F → ( E )
			{"F", String[Symbol]{Terminal("id")}},                                    // F → id
		},
		"S",
	),
	NewCFG(
		[]Terminal{"=", "|", "(", ")", "[", "]", "{", "}", "{{", "}}", "GRAMMAR", "IDENT", "TOKEN", "STRING", "REGEX"},
		[]NonTerminal{"grammar", "name", "decls", "decl", "token", "rule", "lhs", "rhs", "nonterm", "term"},
		[]Production{
			{"grammar", String[Symbol]{NonTerminal("name"), NonTerminal("decls")}}, // grammar → name decls
			{"name", String[Symbol]{Terminal("GRAMMAR"), Terminal("IDENT")}},       // name → GRAMMAR IDENT
			{"decls", String[Symbol]{NonTerminal("decls"), NonTerminal("decl")}},   // decls → decls decl
			{"decls", ε}, // decls → ε
			{"decl", String[Symbol]{NonTerminal("token")}},                                  // decl → token
			{"decl", String[Symbol]{NonTerminal("rule")}},                                   // decl → rule
			{"token", String[Symbol]{Terminal("TOKEN"), Terminal("="), Terminal("STRING")}}, // token → TOKEN "=" STRING
			{"token", String[Symbol]{Terminal("TOKEN"), Terminal("="), Terminal("REGEX")}},  // token → TOKEN "=" REGEX
			{"rule", String[Symbol]{NonTerminal("lhs"), Terminal("="), NonTerminal("rhs")}}, // rule → lhs "=" rhs
			{"rule", String[Symbol]{NonTerminal("lhs"), Terminal("=")}},                     // rule → lhs "="
			{"lhs", String[Symbol]{NonTerminal("nonterm")}},                                 // lhs → nonterm
			{"rhs", String[Symbol]{NonTerminal("rhs"), NonTerminal("rhs")}},                 // rhs → rhs rhs
			{"rhs", String[Symbol]{NonTerminal("rhs"), Terminal("|"), NonTerminal("rhs")}},  // rhs → rhs "|" rhs
			{"rhs", String[Symbol]{Terminal("("), NonTerminal("rhs"), Terminal(")")}},       // rhs → "(" rhs ")"
			{"rhs", String[Symbol]{Terminal("["), NonTerminal("rhs"), Terminal("]")}},       // rhs → "[" rhs "]"
			{"rhs", String[Symbol]{Terminal("{"), NonTerminal("rhs"), Terminal("}")}},       // rhs → "{" rhs "}"
			{"rhs", String[Symbol]{Terminal("{{"), NonTerminal("rhs"), Terminal("}}")}},     // rhs → "{{" rhs "}}"
			{"rhs", String[Symbol]{NonTerminal("nonterm")}},                                 // rhs → nonterm
			{"rhs", String[Symbol]{NonTerminal("term")}},                                    // rhs → term
			{"nonterm", String[Symbol]{Terminal("IDENT")}},                                  // nonterm → IDENT
			{"term", String[Symbol]{Terminal("TOKEN")}},                                     // term → TOKEN
			{"term", String[Symbol]{Terminal("STRING")}},                                    // term → STRING
		},
		"grammar",
	),
}

func TestNewCFG(t *testing.T) {
	tests := []struct {
		name     string
		terms    []Terminal
		nonTerms []NonTerminal
		prods    []Production
		start    NonTerminal
	}{
		{
			name:     "MatchingPairs",
			terms:    []Terminal{"a", "b"},
			nonTerms: []NonTerminal{"S"},
			prods: []Production{
				{"S", String[Symbol]{Terminal("a"), NonTerminal("S"), Terminal("b")}}, //  S → aSb
				{"S", ε}, //  S → ε
			},
			start: "S",
		},
		{
			name:     "WellformedParantheses",
			terms:    []Terminal{"(", ")"},
			nonTerms: []NonTerminal{"S"},
			prods: []Production{
				{"S", String[Symbol]{NonTerminal("S"), NonTerminal("S")}},             //  S → SS
				{"S", String[Symbol]{Terminal("("), NonTerminal("S"), Terminal(")")}}, //  S → (S)
				{"S", String[Symbol]{Terminal("("), Terminal(")")}},                   //  S → ()
			},
			start: "S",
		},
		{
			name:     "WellformedParanthesesAndBrackets",
			terms:    []Terminal{"(", ")", "[", "]"},
			nonTerms: []NonTerminal{"S"},
			prods: []Production{
				{"S", String[Symbol]{NonTerminal("S"), NonTerminal("S")}},             //  S → SS
				{"S", String[Symbol]{Terminal("("), NonTerminal("S"), Terminal(")")}}, //  S → (S)
				{"S", String[Symbol]{Terminal("["), NonTerminal("S"), Terminal("]")}}, //  S → [S]
				{"S", String[Symbol]{Terminal("("), Terminal(")")}},                   //  S → ()
				{"S", String[Symbol]{Terminal("["), Terminal("]")}},                   //  S → []
			},
			start: "S",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := NewCFG(tc.terms, tc.nonTerms, tc.prods, tc.start)
			assert.NotEmpty(t, g)
			assert.NoError(t, g.Verify())
		})
	}
}

func TestCFG_Verify(t *testing.T) {
	tests := []struct {
		name          string
		g             CFG
		expectedError string
	}{
		{
			name: "StartSymbolNotDeclared",
			g: NewCFG(
				[]Terminal{},
				[]NonTerminal{},
				[]Production{},
				"S",
			),
			expectedError: "start symbol S not in the set of non-terminal symbols\nno production rule for start symbol S",
		},
		{
			name: "StartSymbolHasNoProduction",
			g: NewCFG(
				[]Terminal{},
				[]NonTerminal{"S"},
				[]Production{},
				"S",
			),
			expectedError: "no production rule for start symbol S\nno production rule for non-terminal symbol S",
		},
		{
			name: "NonTerminalHasNoProduction",
			g: NewCFG(
				[]Terminal{},
				[]NonTerminal{"A", "S"},
				[]Production{
					{"S", ε}, // S → ε
				},
				"S",
			),
			expectedError: "no production rule for non-terminal symbol A",
		},
		{
			name: "ProductionHeadNotDeclared",
			g: NewCFG(
				[]Terminal{},
				[]NonTerminal{"A", "S"},
				[]Production{
					{"S", String[Symbol]{NonTerminal("A")}}, // S → A
					{"A", ε},                                // A → ε
					{"B", ε},                                // B → ε
				},
				"S",
			),
			expectedError: "production head B not in the set of non-terminal symbols",
		},
		{
			name: "TerminalNotDeclared",
			g: NewCFG(
				[]Terminal{},
				[]NonTerminal{"A", "B", "S"},
				[]Production{
					{"S", String[Symbol]{NonTerminal("A")}}, // S → A
					{"A", String[Symbol]{Terminal("a")}},    // A → a
					{"B", ε},                                // B → ε
				},
				"S",
			),
			expectedError: "terminal symbol \"a\" not in the set of terminal symbols",
		},
		{
			name: "NonTerminalNotDeclared",
			g: NewCFG(
				[]Terminal{"a"},
				[]NonTerminal{"A", "B", "S"},
				[]Production{
					{"S", String[Symbol]{NonTerminal("A")}}, // S → A
					{"A", String[Symbol]{Terminal("a")}},    // A → a
					{"B", String[Symbol]{NonTerminal("C")}}, // B → C
				},
				"S",
			),
			expectedError: "non-terminal symbol C not in the set of non-terminal symbols",
		},
		{
			name: "Valid",
			g: NewCFG(
				[]Terminal{"a", "b"},
				[]NonTerminal{"A", "B", "S"},
				[]Production{
					{"S", String[Symbol]{NonTerminal("A")}}, // S → A
					{"S", String[Symbol]{NonTerminal("B")}}, // S → B
					{"A", String[Symbol]{Terminal("a")}},    // A → a
					{"B", String[Symbol]{Terminal("b")}},    // B → b
				},
				"S",
			),
			expectedError: "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.g.Verify()

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}
		})
	}
}

func TestCFG_String(t *testing.T) {
	tests := []struct {
		name           string
		g              CFG
		expectedString string
	}{
		{
			name:           "1st",
			g:              CFGrammars[0],
			expectedString: "Terminal Symbols: \"0\" \"1\"\nNon-Terminal Symbols: S X Y\nStart Symbol: S\nProduction Rules:\n  S → X Y X\n  X → \"0\" X | ε\n  Y → \"1\" Y | ε\n",
		},
		{
			name:           "2nd",
			g:              CFGrammars[1],
			expectedString: "Terminal Symbols: \"a\" \"b\"\nNon-Terminal Symbols: S\nStart Symbol: S\nProduction Rules:\n  S → \"a\" S \"b\" S | \"b\" S \"a\" S | ε\n",
		},
		{
			name:           "3rd",
			g:              CFGrammars[2],
			expectedString: "Terminal Symbols: \"a\" \"b\"\nNon-Terminal Symbols: S B A\nStart Symbol: S\nProduction Rules:\n  S → \"a\" B \"a\" | A \"b\" | \"a\"\n  B → A | \"b\"\n  A → \"b\" | ε\n",
		},
		{
			name:           "4th",
			g:              CFGrammars[3],
			expectedString: "Terminal Symbols: \"b\" \"c\" \"d\" \"s\"\nNon-Terminal Symbols: S A B C D\nStart Symbol: S\nProduction Rules:\n  S → A | \"s\"\n  A → B\n  B → C | \"b\"\n  C → D\n  D → \"d\"\n",
		},
		{
			name:           "5th",
			g:              CFGrammars[4],
			expectedString: "Terminal Symbols: \"a\" \"b\" \"c\" \"d\"\nNon-Terminal Symbols: S A B C D\nStart Symbol: S\nProduction Rules:\n  S → A B\n  A → \"a\" A | \"a\"\n  B → \"b\" B | \"b\"\n  C → \"c\" C | \"c\"\n  D → \"d\"\n",
		},
		{
			name:           "6th",
			g:              CFGrammars[5],
			expectedString: "Terminal Symbols: \"(\" \")\" \"*\" \"+\" \"-\" \"/\" \"id\"\nNon-Terminal Symbols: S E\nStart Symbol: S\nProduction Rules:\n  S → E\n  E → E \"*\" E | E \"+\" E | E \"-\" E | E \"/\" E | \"(\" E \")\" | \"-\" E | \"id\"\n",
		},
		{
			name:           "7th",
			g:              CFGrammars[6],
			expectedString: "Terminal Symbols: \"(\" \")\" \"*\" \"+\" \"-\" \"/\" \"id\"\nNon-Terminal Symbols: S E T F\nStart Symbol: S\nProduction Rules:\n  S → E\n  E → E \"+\" T | E \"-\" T | T\n  T → T \"*\" F | T \"/\" F | F\n  F → \"(\" E \")\" | \"id\"\n",
		},
		{
			name:           "8th",
			g:              CFGrammars[7],
			expectedString: "Terminal Symbols: \"(\" \")\" \"=\" \"GRAMMAR\" \"IDENT\" \"REGEX\" \"STRING\" \"TOKEN\" \"[\" \"]\" \"{\" \"{{\" \"|\" \"}\" \"}}\"\nNon-Terminal Symbols: grammar name decls decl rule token lhs rhs nonterm term\nStart Symbol: grammar\nProduction Rules:\n  grammar → name decls\n  name → \"GRAMMAR\" \"IDENT\"\n  decls → decls decl | ε\n  decl → rule | token\n  rule → lhs \"=\" rhs | lhs \"=\"\n  token → \"TOKEN\" \"=\" \"REGEX\" | \"TOKEN\" \"=\" \"STRING\"\n  lhs → nonterm\n  rhs → rhs \"|\" rhs | rhs rhs | \"(\" rhs \")\" | \"[\" rhs \"]\" | \"{\" rhs \"}\" | \"{{\" rhs \"}}\" | nonterm | term\n  nonterm → \"IDENT\"\n  term → \"STRING\" | \"TOKEN\"\n",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NoError(t, tc.g.Verify())
			assert.Equal(t, tc.expectedString, tc.g.String())
		})
	}
}

func TestCFG_Clone(t *testing.T) {
	tests := []struct {
		name string
		g    CFG
	}{
		{
			name: "OK",
			g:    CFGrammars[1],
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NoError(t, tc.g.Verify())
			newG := tc.g.Clone()
			assert.NoError(t, newG.Verify())
			assert.False(t, newG == tc.g)
			assert.True(t, newG.Equals(tc.g))
		})
	}
}

func TestCFG_Equals(t *testing.T) {
	tests := []struct {
		name           string
		lhs            CFG
		rhs            CFG
		expectedEquals bool
	}{
		{
			name: "TerminalsNotEqual",
			lhs: NewCFG(
				[]Terminal{"a", "b"},
				[]NonTerminal{"A", "B", "S"},
				[]Production{},
				"S",
			),
			rhs: NewCFG(
				[]Terminal{"a", "b", "c"},
				[]NonTerminal{"A", "B", "S"},
				[]Production{},
				"S",
			),
			expectedEquals: false,
		},
		{
			name: "NonTerminalsNotEqual",
			lhs: NewCFG(
				[]Terminal{"a", "b"},
				[]NonTerminal{"A", "B", "C", "S"},
				[]Production{},
				"S",
			),
			rhs: NewCFG(
				[]Terminal{"a", "b"},
				[]NonTerminal{"A", "B", "S"},
				[]Production{},
				"S",
			),
			expectedEquals: false,
		},
		{
			name: "ProductionsNotEqual",
			lhs: NewCFG(
				[]Terminal{"a", "b"},
				[]NonTerminal{"A", "B", "S"},
				[]Production{
					{"S", String[Symbol]{Terminal("a"), NonTerminal("A")}}, // S → aA
					{"S", String[Symbol]{Terminal("b"), NonTerminal("B")}}, // S → bB
					{"A", String[Symbol]{Terminal("a"), NonTerminal("S")}}, // A → aS
					{"A", String[Symbol]{Terminal("b"), NonTerminal("A")}}, // A → bA
					{"A", ε}, // A → ε
					{"B", String[Symbol]{Terminal("b"), NonTerminal("S")}}, // B → bS
					{"B", String[Symbol]{Terminal("a"), NonTerminal("B")}}, // B → aB
					{"B", ε}, // B → ε
				},
				"S",
			),
			rhs: NewCFG(
				[]Terminal{"a", "b"},
				[]NonTerminal{"A", "B", "S"},
				[]Production{
					{"S", String[Symbol]{Terminal("a"), NonTerminal("A")}}, // S → aA
					{"S", String[Symbol]{Terminal("b"), NonTerminal("B")}}, // S → bB
					{"A", String[Symbol]{Terminal("a"), NonTerminal("S")}}, // A → aS
					{"A", String[Symbol]{Terminal("b"), NonTerminal("A")}}, // A → bA
					{"B", String[Symbol]{Terminal("b"), NonTerminal("S")}}, // B → bS
					{"B", String[Symbol]{Terminal("a"), NonTerminal("B")}}, // B → aB
					{"B", ε}, // B → ε
				},
				"S",
			),
			expectedEquals: false,
		},
		{
			name: "StartSymbolsNotEqual",
			lhs: NewCFG(
				[]Terminal{"a", "b"},
				[]NonTerminal{"A", "B", "S"},
				[]Production{
					{"S", String[Symbol]{Terminal("a"), NonTerminal("A")}}, // S → aA
					{"S", String[Symbol]{Terminal("b"), NonTerminal("B")}}, // S → bB
					{"A", String[Symbol]{Terminal("a"), NonTerminal("S")}}, // A → aS
					{"A", String[Symbol]{Terminal("b"), NonTerminal("A")}}, // A → bA
					{"A", ε}, // A → ε
					{"B", String[Symbol]{Terminal("b"), NonTerminal("S")}}, // B → bS
					{"B", String[Symbol]{Terminal("a"), NonTerminal("B")}}, // B → aB
					{"B", ε}, // B → ε
				},
				"S",
			),
			rhs: NewCFG(
				[]Terminal{"a", "b"},
				[]NonTerminal{"A", "B", "S"},
				[]Production{
					{"S", String[Symbol]{Terminal("a"), NonTerminal("A")}}, // S → aA
					{"S", String[Symbol]{Terminal("b"), NonTerminal("B")}}, // S → bB
					{"A", String[Symbol]{Terminal("a"), NonTerminal("S")}}, // A → aS
					{"A", String[Symbol]{Terminal("b"), NonTerminal("A")}}, // A → bA
					{"A", ε}, // A → ε
					{"B", String[Symbol]{Terminal("b"), NonTerminal("S")}}, // B → bS
					{"B", String[Symbol]{Terminal("a"), NonTerminal("B")}}, // B → aB
					{"B", ε}, // B → ε
				},
				"A",
			),
			expectedEquals: false,
		},
		{
			name: "Equal",
			lhs: NewCFG(
				[]Terminal{"+", "-", "*", "/", "(", ")", "id"},
				[]NonTerminal{"S", "E", "T", "F"},
				[]Production{
					{"S", String[Symbol]{NonTerminal("E")}},                                  // S → E
					{"E", String[Symbol]{NonTerminal("E"), Terminal("+"), NonTerminal("T")}}, // E → E + T
					{"E", String[Symbol]{NonTerminal("E"), Terminal("-"), NonTerminal("T")}}, // E → E - T
					{"E", String[Symbol]{NonTerminal("T")}},                                  // E → T
					{"T", String[Symbol]{NonTerminal("T"), Terminal("*"), NonTerminal("F")}}, // T → T * F
					{"T", String[Symbol]{NonTerminal("T"), Terminal("/"), NonTerminal("F")}}, // T → T / F
					{"T", String[Symbol]{NonTerminal("F")}},                                  // T → F
					{"F", String[Symbol]{Terminal("("), NonTerminal("E"), Terminal(")")}},    // F → ( E )
					{"F", String[Symbol]{Terminal("id")}},                                    // F → id
				},
				"S",
			),
			rhs: NewCFG(
				[]Terminal{"id", "(", ")", "+", "-", "*", "/"},
				[]NonTerminal{"F", "T", "E", "S"},
				[]Production{
					{"F", String[Symbol]{Terminal("id")}},                                    // F → id
					{"F", String[Symbol]{Terminal("("), NonTerminal("E"), Terminal(")")}},    // F → ( E )
					{"T", String[Symbol]{NonTerminal("F")}},                                  // T → F
					{"T", String[Symbol]{NonTerminal("T"), Terminal("*"), NonTerminal("F")}}, // T → T * F
					{"T", String[Symbol]{NonTerminal("T"), Terminal("/"), NonTerminal("F")}}, // T → T / F
					{"E", String[Symbol]{NonTerminal("T")}},                                  // E → T
					{"E", String[Symbol]{NonTerminal("E"), Terminal("+"), NonTerminal("T")}}, // E → E + T
					{"E", String[Symbol]{NonTerminal("E"), Terminal("-"), NonTerminal("T")}}, // E → E - T
					{"S", String[Symbol]{NonTerminal("E")}},                                  // S → E
				},
				"S",
			),
			expectedEquals: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedEquals, tc.lhs.Equals(tc.rhs))
		})
	}
}

func TestCFG_expectedIsCNF(t *testing.T) {
	tests := []struct {
		name          string
		g             CFG
		expectedIsCNF bool
	}{
		{
			name: "NotCNF",
			g: NewCFG(
				[]Terminal{"a", "b"},
				[]NonTerminal{"S", "A", "B"},
				[]Production{
					{"S", String[Symbol]{NonTerminal("A"), NonTerminal("B")}}, // S → AB
					{"S", ε}, // S → ε
					{"A", String[Symbol]{Terminal("a"), NonTerminal("A")}}, // A → aA
					{"A", String[Symbol]{Terminal("a")}},                   // A → a
					{"B", String[Symbol]{Terminal("b"), NonTerminal("B")}}, // B → bB
					{"B", String[Symbol]{Terminal("b")}},                   // B → b
				},
				"S",
			),
			expectedIsCNF: false,
		},
		{
			name: "CNF",
			g: NewCFG(
				[]Terminal{"a", "b"},
				[]NonTerminal{"S", "A", "A₁", "B", "B₁"},
				[]Production{
					{"S", String[Symbol]{NonTerminal("A"), NonTerminal("B")}}, // S → AB
					{"S", ε}, // S → ε
					{"A", String[Symbol]{NonTerminal("A₁"), NonTerminal("A")}}, // A → A₁A
					{"A", String[Symbol]{Terminal("a")}},                       // A → a
					{"A₁", String[Symbol]{Terminal("a")}},                      // A₁ → a
					{"B", String[Symbol]{NonTerminal("B₁"), NonTerminal("B")}}, // B → B₁B
					{"B", String[Symbol]{Terminal("b")}},                       // B → b
					{"B₁", String[Symbol]{Terminal("b")}},                      // B₁ → b
				},
				"S",
			),
			expectedIsCNF: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NoError(t, tc.g.Verify())
			assert.Equal(t, tc.expectedIsCNF, tc.g.IsCNF())
		})
	}
}

func TestCFG_NullableNonTerminals(t *testing.T) {
	tests := []struct {
		name              string
		g                 CFG
		expectedNullables []NonTerminal
	}{
		{
			name:              "1st",
			g:                 CFGrammars[0],
			expectedNullables: []NonTerminal{"S", "X", "Y"},
		},
		{
			name:              "2nd",
			g:                 CFGrammars[1],
			expectedNullables: []NonTerminal{"S"},
		},
		{
			name:              "3rd",
			g:                 CFGrammars[2],
			expectedNullables: []NonTerminal{"A", "B"},
		},
		{
			name:              "4th",
			g:                 CFGrammars[3],
			expectedNullables: []NonTerminal{},
		},
		{
			name:              "5th",
			g:                 CFGrammars[4],
			expectedNullables: []NonTerminal{},
		},
		{
			name:              "6th",
			g:                 CFGrammars[5],
			expectedNullables: []NonTerminal{},
		},
		{
			name:              "7th",
			g:                 CFGrammars[6],
			expectedNullables: []NonTerminal{},
		},
		{
			name:              "8th",
			g:                 CFGrammars[7],
			expectedNullables: []NonTerminal{"decls"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NoError(t, tc.g.Verify())
			nullables := tc.g.NullableNonTerminals()

			for nullable := range nullables.All() {
				assert.Contains(t, tc.expectedNullables, nullable)
			}

			for _, expectedNullable := range tc.expectedNullables {
				assert.True(t, nullables.Contains(expectedNullable))
			}
		})
	}
}

func TestCFG_EliminateEmptyProductions(t *testing.T) {
	tests := []struct {
		name            string
		g               CFG
		expectedGrammar CFG
	}{
		{
			name: "1st",
			g:    CFGrammars[0],
			expectedGrammar: NewCFG(
				[]Terminal{"0", "1"},
				[]NonTerminal{"S′", "S", "X", "Y"},
				[]Production{
					{"S′", String[Symbol]{NonTerminal("S")}}, // S′ → S
					{"S′", ε}, // S′ → ε
					{"S", String[Symbol]{NonTerminal("X"), NonTerminal("Y"), NonTerminal("X")}}, // S → XYX
					{"S", String[Symbol]{NonTerminal("X"), NonTerminal("X")}},                   // S → XX
					{"S", String[Symbol]{NonTerminal("X"), NonTerminal("Y")}},                   // S → XY
					{"S", String[Symbol]{NonTerminal("Y"), NonTerminal("X")}},                   // S → YX
					{"S", String[Symbol]{NonTerminal("X")}},                                     // S → X
					{"S", String[Symbol]{NonTerminal("Y")}},                                     // S → Y
					{"X", String[Symbol]{Terminal("0"), NonTerminal("X")}},                      // X → 0X
					{"X", String[Symbol]{Terminal("0")}},                                        // X → 0
					{"Y", String[Symbol]{Terminal("1"), NonTerminal("Y")}},                      // Y → 1Y
					{"Y", String[Symbol]{Terminal("1")}},                                        // Y → 1
				},
				"S′",
			),
		},
		{
			name: "2nd",
			g:    CFGrammars[1],
			expectedGrammar: NewCFG(
				[]Terminal{"a", "b"},
				[]NonTerminal{"S′", "S"},
				[]Production{
					{"S′", String[Symbol]{NonTerminal("S")}}, // S′ → S
					{"S′", ε}, // S′ → ε
					{"S", String[Symbol]{Terminal("a"), NonTerminal("S"), Terminal("b"), NonTerminal("S")}}, // S → aSbS
					{"S", String[Symbol]{Terminal("b"), NonTerminal("S"), Terminal("a"), NonTerminal("S")}}, // S → bSaS
					{"S", String[Symbol]{Terminal("a"), NonTerminal("S"), Terminal("b")}},                   // S → aSb
					{"S", String[Symbol]{Terminal("a"), Terminal("b"), NonTerminal("S")}},                   // S → abS
					{"S", String[Symbol]{Terminal("b"), NonTerminal("S"), Terminal("a")}},                   // S → bSa
					{"S", String[Symbol]{Terminal("b"), Terminal("a"), NonTerminal("S")}},                   // S → baS
					{"S", String[Symbol]{Terminal("a"), Terminal("b")}},                                     // S → ab
					{"S", String[Symbol]{Terminal("b"), Terminal("a")}},                                     // S → ba
				},
				"S′",
			),
		},
		{
			name: "3rd",
			g:    CFGrammars[2],
			expectedGrammar: NewCFG(
				[]Terminal{"a", "b"},
				[]NonTerminal{"S", "A", "B"},
				[]Production{
					{"S", String[Symbol]{Terminal("a"), NonTerminal("B"), Terminal("a")}}, // S → aBa
					{"S", String[Symbol]{NonTerminal("A"), Terminal("b")}},                // S → Ab
					{"S", String[Symbol]{Terminal("a"), Terminal("a")}},                   // S → aa
					{"S", String[Symbol]{Terminal("a")}},                                  // S → a
					{"S", String[Symbol]{Terminal("b")}},                                  // S → b
					{"A", String[Symbol]{Terminal("b")}},                                  // A → b
					{"B", String[Symbol]{NonTerminal("A")}},                               // B → A
					{"B", String[Symbol]{Terminal("b")}},                                  // B → b
				},
				"S",
			),
		},
		{
			name:            "4th",
			g:               CFGrammars[3],
			expectedGrammar: CFGrammars[3],
		},
		{
			name:            "5th",
			g:               CFGrammars[4],
			expectedGrammar: CFGrammars[4],
		},
		{
			name:            "6th",
			g:               CFGrammars[5],
			expectedGrammar: CFGrammars[5],
		},
		{
			name:            "7th",
			g:               CFGrammars[6],
			expectedGrammar: CFGrammars[6],
		},
		{
			name: "8th",
			g:    CFGrammars[7],
			expectedGrammar: NewCFG(
				[]Terminal{"=", "|", "(", ")", "[", "]", "{", "}", "{{", "}}", "GRAMMAR", "IDENT", "TOKEN", "STRING", "REGEX"},
				[]NonTerminal{"grammar", "name", "decls", "decl", "token", "rule", "lhs", "rhs", "nonterm", "term"},
				[]Production{
					{"grammar", String[Symbol]{NonTerminal("name")}},                                // grammar → name
					{"grammar", String[Symbol]{NonTerminal("name"), NonTerminal("decls")}},          // grammar → name decls
					{"name", String[Symbol]{Terminal("GRAMMAR"), Terminal("IDENT")}},                // name → GRAMMAR IDENT
					{"decls", String[Symbol]{NonTerminal("decls"), NonTerminal("decl")}},            // decls → decls decl
					{"decls", String[Symbol]{NonTerminal("decl")}},                                  // decls → decl
					{"decl", String[Symbol]{NonTerminal("token")}},                                  // decl → token
					{"decl", String[Symbol]{NonTerminal("rule")}},                                   // decl → rule
					{"token", String[Symbol]{Terminal("TOKEN"), Terminal("="), Terminal("STRING")}}, // token → TOKEN "=" STRING
					{"token", String[Symbol]{Terminal("TOKEN"), Terminal("="), Terminal("REGEX")}},  // token → TOKEN "=" REGEX
					{"rule", String[Symbol]{NonTerminal("lhs"), Terminal("="), NonTerminal("rhs")}}, // rule → lhs "=" rhs
					{"rule", String[Symbol]{NonTerminal("lhs"), Terminal("=")}},                     // rule → lhs "="
					{"lhs", String[Symbol]{NonTerminal("nonterm")}},                                 // lhs → nonterm
					{"rhs", String[Symbol]{NonTerminal("rhs"), NonTerminal("rhs")}},                 // rhs → rhs rhs
					{"rhs", String[Symbol]{NonTerminal("rhs"), Terminal("|"), NonTerminal("rhs")}},  // rhs → rhs "|" rhs
					{"rhs", String[Symbol]{Terminal("("), NonTerminal("rhs"), Terminal(")")}},       // rhs → "(" rhs ")"
					{"rhs", String[Symbol]{Terminal("["), NonTerminal("rhs"), Terminal("]")}},       // rhs → "[" rhs "]"
					{"rhs", String[Symbol]{Terminal("{"), NonTerminal("rhs"), Terminal("}")}},       // rhs → "{" rhs "}"
					{"rhs", String[Symbol]{Terminal("{{"), NonTerminal("rhs"), Terminal("}}")}},     // rhs → "{{" rhs "}}"
					{"rhs", String[Symbol]{NonTerminal("nonterm")}},                                 // rhs → nonterm
					{"rhs", String[Symbol]{NonTerminal("term")}},                                    // rhs → term
					{"nonterm", String[Symbol]{Terminal("IDENT")}},                                  // nonterm → IDENT
					{"term", String[Symbol]{Terminal("TOKEN")}},                                     // term → TOKEN
					{"term", String[Symbol]{Terminal("STRING")}},                                    // term → STRING
				},
				"grammar",
			),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NoError(t, tc.g.Verify())
			g := tc.g.EliminateEmptyProductions()
			assert.NoError(t, g.Verify())
			assert.True(t, g.Equals(tc.expectedGrammar))
		})
	}
}

func TestCFG_EliminateSingleProductions(t *testing.T) {
	tests := []struct {
		name            string
		g               CFG
		expectedGrammar CFG
	}{
		{
			name:            "1st",
			g:               CFGrammars[0],
			expectedGrammar: CFGrammars[0],
		},
		{
			name:            "2nd",
			g:               CFGrammars[1],
			expectedGrammar: CFGrammars[1],
		},
		{
			name: "3rd",
			g:    CFGrammars[2],
			expectedGrammar: NewCFG(
				[]Terminal{"a", "b"},
				[]NonTerminal{"S", "A", "B"},
				[]Production{
					{"S", String[Symbol]{Terminal("a"), NonTerminal("B"), Terminal("a")}}, // S → aBa
					{"S", String[Symbol]{NonTerminal("A"), Terminal("b")}},                // S → Ab
					{"S", String[Symbol]{Terminal("a")}},                                  // S → a
					{"A", String[Symbol]{Terminal("b")}},                                  // A → b
					{"A", ε},                                                              // A → ε
					{"B", String[Symbol]{Terminal("b")}},                                  // B → b
					{"B", ε},                                                              // B → ε
				},
				"S",
			),
		},
		{
			name: "4th",
			g:    CFGrammars[3],
			expectedGrammar: NewCFG(
				[]Terminal{"b", "c", "d", "s"},
				[]NonTerminal{"S", "A", "B", "C", "D"},
				[]Production{
					{"S", String[Symbol]{Terminal("b")}}, // S → b
					{"S", String[Symbol]{Terminal("d")}}, // S → d
					{"S", String[Symbol]{Terminal("s")}}, // S → s
					{"A", String[Symbol]{Terminal("b")}}, // A → b
					{"A", String[Symbol]{Terminal("d")}}, // A → d
					{"B", String[Symbol]{Terminal("b")}}, // B → b
					{"B", String[Symbol]{Terminal("d")}}, // B → d
					{"C", String[Symbol]{Terminal("d")}}, // C → d
					{"D", String[Symbol]{Terminal("d")}}, // D → d
				},
				"S",
			),
		},
		{
			name:            "5th",
			g:               CFGrammars[4],
			expectedGrammar: CFGrammars[4],
		},
		{
			name: "6th",
			g:    CFGrammars[5],
			expectedGrammar: NewCFG(
				[]Terminal{"+", "-", "*", "/", "(", ")", "id"},
				[]NonTerminal{"S", "E"},
				[]Production{
					{"S", String[Symbol]{NonTerminal("E"), Terminal("+"), NonTerminal("E")}}, // S → E + E
					{"S", String[Symbol]{NonTerminal("E"), Terminal("-"), NonTerminal("E")}}, // S → E - E
					{"S", String[Symbol]{NonTerminal("E"), Terminal("*"), NonTerminal("E")}}, // S → E * E
					{"S", String[Symbol]{NonTerminal("E"), Terminal("/"), NonTerminal("E")}}, // S → E / E
					{"S", String[Symbol]{Terminal("("), NonTerminal("E"), Terminal(")")}},    // S → ( E )
					{"S", String[Symbol]{Terminal("-"), NonTerminal("E")}},                   // S → - E
					{"S", String[Symbol]{Terminal("id")}},                                    // S → id
					{"E", String[Symbol]{NonTerminal("E"), Terminal("+"), NonTerminal("E")}}, // E → E + E
					{"E", String[Symbol]{NonTerminal("E"), Terminal("-"), NonTerminal("E")}}, // E → E - E
					{"E", String[Symbol]{NonTerminal("E"), Terminal("*"), NonTerminal("E")}}, // E → E * E
					{"E", String[Symbol]{NonTerminal("E"), Terminal("/"), NonTerminal("E")}}, // E → E / E
					{"E", String[Symbol]{Terminal("("), NonTerminal("E"), Terminal(")")}},    // E → ( E )
					{"E", String[Symbol]{Terminal("-"), NonTerminal("E")}},                   // E → - E
					{"E", String[Symbol]{Terminal("id")}},                                    // E → id
				},
				"S",
			),
		},
		{
			name: "7th",
			g:    CFGrammars[6],
			expectedGrammar: NewCFG(
				[]Terminal{"+", "-", "*", "/", "(", ")", "id"},
				[]NonTerminal{"S", "E", "T", "F"},
				[]Production{
					{"S", String[Symbol]{NonTerminal("E"), Terminal("+"), NonTerminal("T")}}, // S → E + T
					{"S", String[Symbol]{NonTerminal("E"), Terminal("-"), NonTerminal("T")}}, // S → E - T
					{"S", String[Symbol]{NonTerminal("T"), Terminal("*"), NonTerminal("F")}}, // S → T * F
					{"S", String[Symbol]{NonTerminal("T"), Terminal("/"), NonTerminal("F")}}, // S → T / F
					{"S", String[Symbol]{Terminal("("), NonTerminal("E"), Terminal(")")}},    // S → ( E )
					{"S", String[Symbol]{Terminal("id")}},                                    // S → id
					{"E", String[Symbol]{NonTerminal("E"), Terminal("+"), NonTerminal("T")}}, // E → E + T
					{"E", String[Symbol]{NonTerminal("E"), Terminal("-"), NonTerminal("T")}}, // E → E - T
					{"E", String[Symbol]{NonTerminal("T"), Terminal("*"), NonTerminal("F")}}, // E → T * F
					{"E", String[Symbol]{NonTerminal("T"), Terminal("/"), NonTerminal("F")}}, // E → T / F
					{"E", String[Symbol]{Terminal("("), NonTerminal("E"), Terminal(")")}},    // E → ( E )
					{"E", String[Symbol]{Terminal("id")}},                                    // E → id
					{"T", String[Symbol]{NonTerminal("T"), Terminal("*"), NonTerminal("F")}}, // T → T * F
					{"T", String[Symbol]{NonTerminal("T"), Terminal("/"), NonTerminal("F")}}, // T → T / F
					{"T", String[Symbol]{Terminal("("), NonTerminal("E"), Terminal(")")}},    // T → ( E )
					{"T", String[Symbol]{Terminal("id")}},                                    // T → id
					{"F", String[Symbol]{Terminal("("), NonTerminal("E"), Terminal(")")}},    // F → ( E )
					{"F", String[Symbol]{Terminal("id")}},                                    // F → id
				},
				"S",
			),
		},
		{
			name: "8th",
			g:    CFGrammars[7],
			expectedGrammar: NewCFG(
				[]Terminal{"=", "|", "(", ")", "[", "]", "{", "}", "{{", "}}", "GRAMMAR", "IDENT", "TOKEN", "STRING", "REGEX"},
				[]NonTerminal{"grammar", "name", "decls", "decl", "token", "rule", "lhs", "rhs", "nonterm", "term"},
				[]Production{
					{"grammar", String[Symbol]{NonTerminal("name"), NonTerminal("decls")}}, // grammar → name decls
					{"name", String[Symbol]{Terminal("GRAMMAR"), Terminal("IDENT")}},       // name → GRAMMAR IDENT
					{"decls", String[Symbol]{NonTerminal("decls"), NonTerminal("decl")}},   // decls → decls decl
					{"decls", ε}, // decls → ε
					{"decl", String[Symbol]{NonTerminal("lhs"), Terminal("="), NonTerminal("rhs")}}, // decl → lhs "=" rhs
					{"decl", String[Symbol]{NonTerminal("lhs"), Terminal("=")}},                     // decl → lhs "="
					{"decl", String[Symbol]{Terminal("TOKEN"), Terminal("="), Terminal("STRING")}},  // decl → TOKEN "=" STRING
					{"decl", String[Symbol]{Terminal("TOKEN"), Terminal("="), Terminal("REGEX")}},   // decl → TOKEN "=" REGEX
					{"token", String[Symbol]{Terminal("TOKEN"), Terminal("="), Terminal("STRING")}}, // token → TOKEN "=" STRING
					{"token", String[Symbol]{Terminal("TOKEN"), Terminal("="), Terminal("REGEX")}},  // token → TOKEN "=" REGEX
					{"rule", String[Symbol]{NonTerminal("lhs"), Terminal("="), NonTerminal("rhs")}}, // rule → lhs "=" rhs
					{"rule", String[Symbol]{NonTerminal("lhs"), Terminal("=")}},                     // rule → lhs "="
					{"lhs", String[Symbol]{Terminal("IDENT")}},                                      // lhs → IDENT
					{"rhs", String[Symbol]{NonTerminal("rhs"), NonTerminal("rhs")}},                 // rhs → rhs rhs
					{"rhs", String[Symbol]{NonTerminal("rhs"), Terminal("|"), NonTerminal("rhs")}},  // rhs → rhs "|" rhs
					{"rhs", String[Symbol]{Terminal("("), NonTerminal("rhs"), Terminal(")")}},       // rhs → "(" rhs ")"
					{"rhs", String[Symbol]{Terminal("["), NonTerminal("rhs"), Terminal("]")}},       // rhs → "[" rhs "]"
					{"rhs", String[Symbol]{Terminal("{"), NonTerminal("rhs"), Terminal("}")}},       // rhs → "{" rhs "}"
					{"rhs", String[Symbol]{Terminal("{{"), NonTerminal("rhs"), Terminal("}}")}},     // rhs → "{{" rhs "}}"
					{"rhs", String[Symbol]{Terminal("IDENT")}},                                      // rhs → IDENT
					{"rhs", String[Symbol]{Terminal("TOKEN")}},                                      // rhs → TOKEN
					{"rhs", String[Symbol]{Terminal("STRING")}},                                     // rhs → STRING
					{"nonterm", String[Symbol]{Terminal("IDENT")}},                                  // nonterm → IDENT
					{"term", String[Symbol]{Terminal("TOKEN")}},                                     // term → TOKEN
					{"term", String[Symbol]{Terminal("STRING")}},                                    // term → STRING
				},
				"grammar",
			),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NoError(t, tc.g.Verify())
			g := tc.g.EliminateSingleProductions()
			assert.NoError(t, g.Verify())
			assert.True(t, g.Equals(tc.expectedGrammar))
		})
	}
}

func TestCFG_EliminateUnreachableProductions(t *testing.T) {
	tests := []struct {
		name            string
		g               CFG
		expectedGrammar CFG
	}{
		{
			name:            "1st",
			g:               CFGrammars[0],
			expectedGrammar: CFGrammars[0],
		},
		{
			name:            "2nd",
			g:               CFGrammars[1],
			expectedGrammar: CFGrammars[1],
		},
		{
			name:            "3rd",
			g:               CFGrammars[2],
			expectedGrammar: CFGrammars[2],
		},
		{
			name: "4th",
			g:    CFGrammars[3],
			expectedGrammar: NewCFG(
				[]Terminal{"b", "d", "s"},
				[]NonTerminal{"S", "A", "B", "C", "D"},
				[]Production{
					{"S", String[Symbol]{NonTerminal("A")}}, // S → A
					{"S", String[Symbol]{Terminal("s")}},    // S → s
					{"A", String[Symbol]{NonTerminal("B")}}, // A → B
					{"B", String[Symbol]{NonTerminal("C")}}, // B → C
					{"B", String[Symbol]{Terminal("b")}},    // B → b
					{"C", String[Symbol]{NonTerminal("D")}}, // C → D
					{"D", String[Symbol]{Terminal("d")}},    // D → d
				},
				"S",
			),
		},
		{
			name: "5th",
			g:    CFGrammars[4],
			expectedGrammar: NewCFG(
				[]Terminal{"a", "b"},
				[]NonTerminal{"S", "A", "B"},
				[]Production{
					{"S", String[Symbol]{NonTerminal("A"), NonTerminal("B")}}, // S → AB
					{"A", String[Symbol]{Terminal("a"), NonTerminal("A")}},    // A → aA
					{"A", String[Symbol]{Terminal("a")}},                      // A → a
					{"B", String[Symbol]{Terminal("b"), NonTerminal("B")}},    // B → bB
					{"B", String[Symbol]{Terminal("b")}},                      // B → b
				},
				"S",
			),
		},
		{
			name:            "6th",
			g:               CFGrammars[5],
			expectedGrammar: CFGrammars[5],
		},
		{
			name:            "7th",
			g:               CFGrammars[6],
			expectedGrammar: CFGrammars[6],
		},
		{
			name:            "8th",
			g:               CFGrammars[7],
			expectedGrammar: CFGrammars[7],
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NoError(t, tc.g.Verify())
			g := tc.g.EliminateUnreachableProductions()
			assert.NoError(t, g.Verify())
			assert.True(t, g.Equals(tc.expectedGrammar))
		})
	}
}

func TestCFG_EliminateCycles(t *testing.T) {
	tests := []struct {
		name            string
		g               CFG
		expectedGrammar CFG
	}{
		{
			name: "1st",
			g:    CFGrammars[0],
			expectedGrammar: NewCFG(
				[]Terminal{"0", "1"},
				[]NonTerminal{"S′", "X", "Y"},
				[]Production{
					{"S′", String[Symbol]{NonTerminal("X"), NonTerminal("Y"), NonTerminal("X")}}, // S′ → XYX
					{"S′", String[Symbol]{NonTerminal("X"), NonTerminal("X")}},                   // S′ → XX
					{"S′", String[Symbol]{NonTerminal("X"), NonTerminal("Y")}},                   // S′ → XY
					{"S′", String[Symbol]{NonTerminal("Y"), NonTerminal("X")}},                   // S′ → YX
					{"S′", String[Symbol]{Terminal("0"), NonTerminal("X")}},                      // S′ → 0X
					{"S′", String[Symbol]{Terminal("1"), NonTerminal("Y")}},                      // S′ → 1Y
					{"S′", String[Symbol]{Terminal("0")}},                                        // S′ → 0
					{"S′", String[Symbol]{Terminal("1")}},                                        // S′ → 1
					{"S′", ε},                                                                    // S′ → ε
					{"X", String[Symbol]{Terminal("0"), NonTerminal("X")}},                       // X → 0X
					{"X", String[Symbol]{Terminal("0")}},                                         // X → 0
					{"Y", String[Symbol]{Terminal("1"), NonTerminal("Y")}},                       // Y → 1Y
					{"Y", String[Symbol]{Terminal("1")}},                                         // Y → 1
				},
				"S′",
			),
		},
		{
			name: "2nd",
			g:    CFGrammars[1],
			expectedGrammar: NewCFG(
				[]Terminal{"a", "b"},
				[]NonTerminal{"S′", "S"},
				[]Production{
					{"S′", String[Symbol]{Terminal("a"), NonTerminal("S"), Terminal("b"), NonTerminal("S")}}, // S′ → aSbS
					{"S′", String[Symbol]{Terminal("b"), NonTerminal("S"), Terminal("a"), NonTerminal("S")}}, // S′ → bSaS
					{"S′", String[Symbol]{Terminal("a"), NonTerminal("S"), Terminal("b")}},                   // S′ → aSb
					{"S′", String[Symbol]{Terminal("a"), Terminal("b"), NonTerminal("S")}},                   // S′ → abS
					{"S′", String[Symbol]{Terminal("b"), NonTerminal("S"), Terminal("a")}},                   // S′ → bSa
					{"S′", String[Symbol]{Terminal("b"), Terminal("a"), NonTerminal("S")}},                   // S′ → baS
					{"S′", String[Symbol]{Terminal("a"), Terminal("b")}},                                     // S′ → ab
					{"S′", String[Symbol]{Terminal("b"), Terminal("a")}},                                     // S′ → ba
					{"S′", ε}, // S′ → ε
					{"S", String[Symbol]{Terminal("a"), NonTerminal("S"), Terminal("b"), NonTerminal("S")}}, // S → aSbS
					{"S", String[Symbol]{Terminal("b"), NonTerminal("S"), Terminal("a"), NonTerminal("S")}}, // S → bSaS
					{"S", String[Symbol]{Terminal("a"), NonTerminal("S"), Terminal("b")}},                   // S → aSb
					{"S", String[Symbol]{Terminal("a"), Terminal("b"), NonTerminal("S")}},                   // S → abS
					{"S", String[Symbol]{Terminal("b"), NonTerminal("S"), Terminal("a")}},                   // S → bSa
					{"S", String[Symbol]{Terminal("b"), Terminal("a"), NonTerminal("S")}},                   // S → baS
					{"S", String[Symbol]{Terminal("a"), Terminal("b")}},                                     // S → ab
					{"S", String[Symbol]{Terminal("b"), Terminal("a")}},                                     // S → ba
				},
				"S′",
			),
		},
		{
			name: "3rd",
			g:    CFGrammars[2],
			expectedGrammar: NewCFG(
				[]Terminal{"a", "b"},
				[]NonTerminal{"S", "A", "B"},
				[]Production{
					{"S", String[Symbol]{Terminal("a"), NonTerminal("B"), Terminal("a")}}, // S → aBa
					{"S", String[Symbol]{NonTerminal("A"), Terminal("b")}},                // S → Ab
					{"S", String[Symbol]{Terminal("a"), Terminal("a")}},                   // S → aa
					{"S", String[Symbol]{Terminal("a")}},                                  // S → a
					{"S", String[Symbol]{Terminal("b")}},                                  // S → b
					{"A", String[Symbol]{Terminal("b")}},                                  // A → b
					{"B", String[Symbol]{Terminal("b")}},                                  // B → b
				},
				"S",
			),
		},
		{
			name: "4th",
			g:    CFGrammars[3],
			expectedGrammar: NewCFG(
				[]Terminal{"b", "d", "s"},
				[]NonTerminal{"S"},
				[]Production{
					{"S", String[Symbol]{Terminal("b")}}, // S → b
					{"S", String[Symbol]{Terminal("d")}}, // S → d
					{"S", String[Symbol]{Terminal("s")}}, // S → s
				},
				"S",
			),
		},
		{
			name: "5th",
			g:    CFGrammars[4],
			expectedGrammar: NewCFG(
				[]Terminal{"a", "b"},
				[]NonTerminal{"S", "A", "B"},
				[]Production{
					{"S", String[Symbol]{NonTerminal("A"), NonTerminal("B")}}, // S → AB
					{"A", String[Symbol]{Terminal("a"), NonTerminal("A")}},    // A → aA
					{"A", String[Symbol]{Terminal("a")}},                      // A → a
					{"B", String[Symbol]{Terminal("b"), NonTerminal("B")}},    // B → bB
					{"B", String[Symbol]{Terminal("b")}},                      // B → b
				},
				"S",
			),
		},
		{
			name: "6th",
			g:    CFGrammars[5],
			expectedGrammar: NewCFG(
				[]Terminal{"+", "-", "*", "/", "(", ")", "id"},
				[]NonTerminal{"S", "E"},
				[]Production{
					{"S", String[Symbol]{NonTerminal("E"), Terminal("+"), NonTerminal("E")}}, // S → E + E
					{"S", String[Symbol]{NonTerminal("E"), Terminal("-"), NonTerminal("E")}}, // S → E - E
					{"S", String[Symbol]{NonTerminal("E"), Terminal("*"), NonTerminal("E")}}, // S → E * E
					{"S", String[Symbol]{NonTerminal("E"), Terminal("/"), NonTerminal("E")}}, // S → E / E
					{"S", String[Symbol]{Terminal("("), NonTerminal("E"), Terminal(")")}},    // S → ( E )
					{"S", String[Symbol]{Terminal("-"), NonTerminal("E")}},                   // S → - E
					{"S", String[Symbol]{Terminal("id")}},                                    // S → id
					{"E", String[Symbol]{NonTerminal("E"), Terminal("+"), NonTerminal("E")}}, // E → E + E
					{"E", String[Symbol]{NonTerminal("E"), Terminal("-"), NonTerminal("E")}}, // E → E - E
					{"E", String[Symbol]{NonTerminal("E"), Terminal("*"), NonTerminal("E")}}, // E → E * E
					{"E", String[Symbol]{NonTerminal("E"), Terminal("/"), NonTerminal("E")}}, // E → E / E
					{"E", String[Symbol]{Terminal("("), NonTerminal("E"), Terminal(")")}},    // E → ( E )
					{"E", String[Symbol]{Terminal("-"), NonTerminal("E")}},                   // E → - E
					{"E", String[Symbol]{Terminal("id")}},                                    // E → id
				},
				"S",
			),
		},
		{
			name: "7th",
			g:    CFGrammars[6],
			expectedGrammar: NewCFG(
				[]Terminal{"+", "-", "*", "/", "(", ")", "id"},
				[]NonTerminal{"S", "E", "T", "F"},
				[]Production{
					{"S", String[Symbol]{NonTerminal("E"), Terminal("+"), NonTerminal("T")}}, // S → E + T
					{"S", String[Symbol]{NonTerminal("E"), Terminal("-"), NonTerminal("T")}}, // S → E - T
					{"S", String[Symbol]{NonTerminal("T"), Terminal("*"), NonTerminal("F")}}, // S → T * F
					{"S", String[Symbol]{NonTerminal("T"), Terminal("/"), NonTerminal("F")}}, // S → T / F
					{"S", String[Symbol]{Terminal("("), NonTerminal("E"), Terminal(")")}},    // S → ( E )
					{"S", String[Symbol]{Terminal("id")}},                                    // S → id
					{"E", String[Symbol]{NonTerminal("E"), Terminal("+"), NonTerminal("T")}}, // E → E + T
					{"E", String[Symbol]{NonTerminal("E"), Terminal("-"), NonTerminal("T")}}, // E → E - T
					{"E", String[Symbol]{NonTerminal("T"), Terminal("*"), NonTerminal("F")}}, // E → T * F
					{"E", String[Symbol]{NonTerminal("T"), Terminal("/"), NonTerminal("F")}}, // E → T / F
					{"E", String[Symbol]{Terminal("("), NonTerminal("E"), Terminal(")")}},    // E → ( E )
					{"E", String[Symbol]{Terminal("id")}},                                    // E → id
					{"T", String[Symbol]{NonTerminal("T"), Terminal("*"), NonTerminal("F")}}, // T → T * F
					{"T", String[Symbol]{NonTerminal("T"), Terminal("/"), NonTerminal("F")}}, // T → T / F
					{"T", String[Symbol]{Terminal("("), NonTerminal("E"), Terminal(")")}},    // T → ( E )
					{"T", String[Symbol]{Terminal("id")}},                                    // T → id
					{"F", String[Symbol]{Terminal("("), NonTerminal("E"), Terminal(")")}},    // F → ( E )
					{"F", String[Symbol]{Terminal("id")}},                                    // F → id
				},
				"S",
			),
		},
		{
			name: "8th",
			g:    CFGrammars[7],
			expectedGrammar: NewCFG(
				[]Terminal{"=", "|", "(", ")", "[", "]", "{", "}", "{{", "}}", "GRAMMAR", "IDENT", "TOKEN", "STRING", "REGEX"},
				[]NonTerminal{"grammar", "name", "decls", "decl", "lhs", "rhs"},
				[]Production{
					{"grammar", String[Symbol]{NonTerminal("name"), NonTerminal("decls")}},           // grammar → name decls
					{"grammar", String[Symbol]{Terminal("GRAMMAR"), Terminal("IDENT")}},              // grammar → GRAMMAR IDENT
					{"name", String[Symbol]{Terminal("GRAMMAR"), Terminal("IDENT")}},                 // name → GRAMMAR IDENT
					{"decls", String[Symbol]{NonTerminal("decls"), NonTerminal("decl")}},             // decls → decls decl
					{"decls", String[Symbol]{NonTerminal("lhs"), Terminal("="), NonTerminal("rhs")}}, // decls → lhs "=" rhs
					{"decls", String[Symbol]{NonTerminal("lhs"), Terminal("=")}},                     // decls → lhs "="
					{"decls", String[Symbol]{Terminal("TOKEN"), Terminal("="), Terminal("STRING")}},  // decls → TOKEN "=" STRING
					{"decls", String[Symbol]{Terminal("TOKEN"), Terminal("="), Terminal("REGEX")}},   // decls → TOKEN "=" REGEX
					{"decl", String[Symbol]{NonTerminal("lhs"), Terminal("="), NonTerminal("rhs")}},  // decl → lhs "=" rhs
					{"decl", String[Symbol]{NonTerminal("lhs"), Terminal("=")}},                      // decl → lhs "="
					{"decl", String[Symbol]{Terminal("TOKEN"), Terminal("="), Terminal("STRING")}},   // decl → TOKEN "=" STRING
					{"decl", String[Symbol]{Terminal("TOKEN"), Terminal("="), Terminal("REGEX")}},    // decl → TOKEN "=" REGEX
					{"lhs", String[Symbol]{Terminal("IDENT")}},                                       // lhs → IDENT
					{"rhs", String[Symbol]{NonTerminal("rhs"), NonTerminal("rhs")}},                  // rhs → rhs rhs
					{"rhs", String[Symbol]{NonTerminal("rhs"), Terminal("|"), NonTerminal("rhs")}},   // rhs → rhs "|" rhs
					{"rhs", String[Symbol]{Terminal("("), NonTerminal("rhs"), Terminal(")")}},        // rhs → "(" rhs ")"
					{"rhs", String[Symbol]{Terminal("["), NonTerminal("rhs"), Terminal("]")}},        // rhs → "[" rhs "]"
					{"rhs", String[Symbol]{Terminal("{"), NonTerminal("rhs"), Terminal("}")}},        // rhs → "{" rhs "}"
					{"rhs", String[Symbol]{Terminal("{{"), NonTerminal("rhs"), Terminal("}}")}},      // rhs → "{{" rhs "}}"
					{"rhs", String[Symbol]{Terminal("IDENT")}},                                       // rhs → IDENT
					{"rhs", String[Symbol]{Terminal("TOKEN")}},                                       // rhs → TOKEN
					{"rhs", String[Symbol]{Terminal("STRING")}},                                      // rhs → STRING
				},
				"grammar",
			),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NoError(t, tc.g.Verify())
			g := tc.g.EliminateCycles()
			assert.NoError(t, g.Verify())
			assert.True(t, g.Equals(tc.expectedGrammar))
		})
	}
}

func TestCFG_EliminateLeftRecursion(t *testing.T) {
	tests := []struct {
		name            string
		g               CFG
		expectedGrammar CFG
	}{
		{
			name: "1st",
			g:    CFGrammars[0],
			expectedGrammar: NewCFG(
				[]Terminal{"0", "1"},
				[]NonTerminal{"S′", "X", "Y"},
				[]Production{
					{"S′", String[Symbol]{NonTerminal("X"), NonTerminal("Y"), NonTerminal("X")}}, // S′ → XYX
					{"S′", String[Symbol]{NonTerminal("X"), NonTerminal("X")}},                   // S′ → XX
					{"S′", String[Symbol]{NonTerminal("X"), NonTerminal("Y")}},                   // S′ → XY
					{"S′", String[Symbol]{NonTerminal("Y"), NonTerminal("X")}},                   // S′ → YX
					{"S′", String[Symbol]{Terminal("0"), NonTerminal("X")}},                      // S′ → 0X
					{"S′", String[Symbol]{Terminal("1"), NonTerminal("Y")}},                      // S′ → 1Y
					{"S′", String[Symbol]{Terminal("0")}},                                        // S′ → 0
					{"S′", String[Symbol]{Terminal("1")}},                                        // S′ → 1
					{"S′", ε},                                                                    // S′ → ε
					{"X", String[Symbol]{Terminal("0"), NonTerminal("X")}},                       // X → 0X
					{"X", String[Symbol]{Terminal("0")}},                                         // X → 0
					{"Y", String[Symbol]{Terminal("1"), NonTerminal("Y")}},                       // Y → 1Y
					{"Y", String[Symbol]{Terminal("1")}},                                         // Y → 1
				},
				"S′",
			),
		},
		{
			name: "2nd",
			g:    CFGrammars[1],
			expectedGrammar: NewCFG(
				[]Terminal{"a", "b"},
				[]NonTerminal{"S′", "S"},
				[]Production{
					{"S′", String[Symbol]{Terminal("a"), NonTerminal("S"), Terminal("b"), NonTerminal("S")}}, // S′ → aSbS
					{"S′", String[Symbol]{Terminal("b"), NonTerminal("S"), Terminal("a"), NonTerminal("S")}}, // S′ → bSaS
					{"S′", String[Symbol]{Terminal("a"), NonTerminal("S"), Terminal("b")}},                   // S′ → aSb
					{"S′", String[Symbol]{Terminal("a"), Terminal("b"), NonTerminal("S")}},                   // S′ → abS
					{"S′", String[Symbol]{Terminal("b"), NonTerminal("S"), Terminal("a")}},                   // S′ → bSa
					{"S′", String[Symbol]{Terminal("b"), Terminal("a"), NonTerminal("S")}},                   // S′ → baS
					{"S′", String[Symbol]{Terminal("a"), Terminal("b")}},                                     // S′ → ab
					{"S′", String[Symbol]{Terminal("b"), Terminal("a")}},                                     // S′ → ba
					{"S′", ε}, // S′ → ε
					{"S", String[Symbol]{Terminal("a"), NonTerminal("S"), Terminal("b"), NonTerminal("S")}}, // S → aSbS
					{"S", String[Symbol]{Terminal("b"), NonTerminal("S"), Terminal("a"), NonTerminal("S")}}, // S → bSaS
					{"S", String[Symbol]{Terminal("a"), NonTerminal("S"), Terminal("b")}},                   // S → aSb
					{"S", String[Symbol]{Terminal("a"), Terminal("b"), NonTerminal("S")}},                   // S → abS
					{"S", String[Symbol]{Terminal("b"), NonTerminal("S"), Terminal("a")}},                   // S → bSa
					{"S", String[Symbol]{Terminal("b"), Terminal("a"), NonTerminal("S")}},                   // S → baS
					{"S", String[Symbol]{Terminal("a"), Terminal("b")}},                                     // S → ab
					{"S", String[Symbol]{Terminal("b"), Terminal("a")}},                                     // S → ba
				},
				"S′",
			),
		},
		{
			name: "3rd",
			g:    CFGrammars[2],
			expectedGrammar: NewCFG(
				[]Terminal{"a", "b"},
				[]NonTerminal{"S", "A", "B"},
				[]Production{
					{"S", String[Symbol]{Terminal("a"), NonTerminal("B"), Terminal("a")}}, // S → aBa
					{"S", String[Symbol]{NonTerminal("A"), Terminal("b")}},                // S → Ab
					{"S", String[Symbol]{Terminal("a"), Terminal("a")}},                   // S → aa
					{"S", String[Symbol]{Terminal("a")}},                                  // S → a
					{"S", String[Symbol]{Terminal("b")}},                                  // S → b
					{"A", String[Symbol]{Terminal("b")}},                                  // A → b
					{"B", String[Symbol]{Terminal("b")}},                                  // B → b
				},
				"S",
			),
		},
		{
			name: "4th",
			g:    CFGrammars[3],
			expectedGrammar: NewCFG(
				[]Terminal{"b", "d", "s"},
				[]NonTerminal{"S"},
				[]Production{
					{"S", String[Symbol]{Terminal("b")}}, // S → b
					{"S", String[Symbol]{Terminal("d")}}, // S → d
					{"S", String[Symbol]{Terminal("s")}}, // S → s
				},
				"S",
			),
		},
		{
			name: "5th",
			g:    CFGrammars[4],
			expectedGrammar: NewCFG(
				[]Terminal{"a", "b"},
				[]NonTerminal{"S", "A", "B"},
				[]Production{
					{"S", String[Symbol]{NonTerminal("A"), NonTerminal("B")}}, // S → AB
					{"A", String[Symbol]{Terminal("a"), NonTerminal("A")}},    // A → aA
					{"A", String[Symbol]{Terminal("a")}},                      // A → a
					{"B", String[Symbol]{Terminal("b"), NonTerminal("B")}},    // B → bB
					{"B", String[Symbol]{Terminal("b")}},                      // B → b
				},
				"S",
			),
		},
		{
			name: "6th",
			g:    CFGrammars[5],
			expectedGrammar: NewCFG(
				[]Terminal{"+", "-", "*", "/", "(", ")", "id"},
				[]NonTerminal{"S", "E", "E′"},
				[]Production{
					{"S", String[Symbol]{NonTerminal("E"), Terminal("+"), NonTerminal("E")}},                 // S → E + E
					{"S", String[Symbol]{NonTerminal("E"), Terminal("-"), NonTerminal("E")}},                 // S → E - E
					{"S", String[Symbol]{NonTerminal("E"), Terminal("*"), NonTerminal("E")}},                 // S → E * E
					{"S", String[Symbol]{NonTerminal("E"), Terminal("/"), NonTerminal("E")}},                 // S → E / E
					{"S", String[Symbol]{Terminal("("), NonTerminal("E"), Terminal(")")}},                    // S → ( E )
					{"S", String[Symbol]{Terminal("-"), NonTerminal("E")}},                                   // S → - E
					{"S", String[Symbol]{Terminal("id")}},                                                    // S → id
					{"E", String[Symbol]{Terminal("("), NonTerminal("E"), Terminal(")"), NonTerminal("E′")}}, // E → ( E ) E′
					{"E", String[Symbol]{Terminal("-"), NonTerminal("E"), NonTerminal("E′")}},                // E → - E E′
					{"E", String[Symbol]{Terminal("id"), NonTerminal("E′")}},                                 // E → id E′
					{"E′", String[Symbol]{Terminal("+"), NonTerminal("E"), NonTerminal("E′")}},               // E′ → + E E′
					{"E′", String[Symbol]{Terminal("-"), NonTerminal("E"), NonTerminal("E′")}},               // E′ → - E E′
					{"E′", String[Symbol]{Terminal("*"), NonTerminal("E"), NonTerminal("E′")}},               // E′ → * E E′
					{"E′", String[Symbol]{Terminal("/"), NonTerminal("E"), NonTerminal("E′")}},               // E′ → / E E′
					{"E′", ε}, // E′ → ε
				},
				"S",
			),
		},
		{
			name: "7th",
			g:    CFGrammars[6],
			expectedGrammar: NewCFG(
				[]Terminal{"+", "-", "*", "/", "(", ")", "id"},
				[]NonTerminal{"S", "E", "E′", "T", "T′", "F"},
				[]Production{
					{"S", String[Symbol]{NonTerminal("E"), Terminal("+"), NonTerminal("T")}},                    // S → E + T
					{"S", String[Symbol]{NonTerminal("E"), Terminal("-"), NonTerminal("T")}},                    // S → E - T
					{"S", String[Symbol]{NonTerminal("T"), Terminal("*"), NonTerminal("F")}},                    // S → T * F
					{"S", String[Symbol]{NonTerminal("T"), Terminal("/"), NonTerminal("F")}},                    // S → T / F
					{"S", String[Symbol]{Terminal("("), NonTerminal("E"), Terminal(")")}},                       // S → ( E )
					{"S", String[Symbol]{Terminal("id")}},                                                       // S → id
					{"E", String[Symbol]{NonTerminal("T"), Terminal("*"), NonTerminal("F"), NonTerminal("E′")}}, // E → T * F E′
					{"E", String[Symbol]{NonTerminal("T"), Terminal("/"), NonTerminal("F"), NonTerminal("E′")}}, // E → T / F E′
					{"E", String[Symbol]{Terminal("("), NonTerminal("E"), Terminal(")"), NonTerminal("E′")}},    // E → ( E ) E′
					{"E", String[Symbol]{Terminal("id"), NonTerminal("E′")}},                                    // E → id E′
					{"E′", String[Symbol]{Terminal("+"), NonTerminal("T"), NonTerminal("E′")}},                  // E′ → + T E′
					{"E′", String[Symbol]{Terminal("-"), NonTerminal("T"), NonTerminal("E′")}},                  // E′ → - T E′
					{"E′", ε}, // E′ → ε
					{"T", String[Symbol]{Terminal("("), NonTerminal("E"), Terminal(")"), NonTerminal("T′")}}, // T → ( E ) T′
					{"T", String[Symbol]{Terminal("id"), NonTerminal("T′")}},                                 // T → id T′
					{"T′", String[Symbol]{Terminal("*"), NonTerminal("F"), NonTerminal("T′")}},               // T′ → * F T′
					{"T′", String[Symbol]{Terminal("/"), NonTerminal("F"), NonTerminal("T′")}},               // T′ → / F T′
					{"T′", ε}, // T′ → ε
					{"F", String[Symbol]{Terminal("("), NonTerminal("E"), Terminal(")")}}, // F → ( E )
					{"F", String[Symbol]{Terminal("id")}},                                 // F → id
				},
				"S",
			),
		},
		{
			name: "8th",
			g:    CFGrammars[7],
			expectedGrammar: NewCFG(
				[]Terminal{"=", "|", "(", ")", "[", "]", "{", "}", "{{", "}}", "GRAMMAR", "IDENT", "TOKEN", "STRING", "REGEX"},
				[]NonTerminal{"grammar", "name", "decls", "decls′", "decl", "lhs", "rhs", "rhs′"},
				[]Production{
					{"grammar", String[Symbol]{NonTerminal("name"), NonTerminal("decls")}},                                  // grammar → name decls
					{"grammar", String[Symbol]{Terminal("GRAMMAR"), Terminal("IDENT")}},                                     // grammar → GRAMMAR IDENT
					{"name", String[Symbol]{Terminal("GRAMMAR"), Terminal("IDENT")}},                                        // name → GRAMMAR IDENT
					{"decls", String[Symbol]{NonTerminal("lhs"), Terminal("="), NonTerminal("rhs"), NonTerminal("decls′")}}, // decls → lhs "=" rhs decls′
					{"decls", String[Symbol]{NonTerminal("lhs"), Terminal("="), NonTerminal("decls′")}},                     // decls → lhs "=" decls′
					{"decls", String[Symbol]{Terminal("TOKEN"), Terminal("="), Terminal("REGEX"), NonTerminal("decls′")}},   // decls → TOKEN "=" REGEX decls′
					{"decls", String[Symbol]{Terminal("TOKEN"), Terminal("="), Terminal("STRING"), NonTerminal("decls′")}},  // decls → TOKEN "=" STRING decls′
					{"decls′", String[Symbol]{NonTerminal("decl"), NonTerminal("decls′")}},                                  // decls′ → decl decls′
					{"decls′", ε}, // decls′ → ε
					{"decl", String[Symbol]{Terminal("IDENT"), Terminal("="), NonTerminal("rhs")}},                   // decl → IDENT "=" rhs
					{"decl", String[Symbol]{Terminal("IDENT"), Terminal("=")}},                                       // decl → IDENT "="
					{"decl", String[Symbol]{Terminal("TOKEN"), Terminal("="), Terminal("REGEX")}},                    // decl → TOKEN "=" REGEX
					{"decl", String[Symbol]{Terminal("TOKEN"), Terminal("="), Terminal("STRING")}},                   // decl → TOKEN "=" STRING
					{"lhs", String[Symbol]{Terminal("IDENT")}},                                                       // lhs → IDENT
					{"rhs", String[Symbol]{Terminal("("), NonTerminal("rhs"), Terminal(")"), NonTerminal("rhs′")}},   // rhs → "(" rhs ")" rhs′
					{"rhs", String[Symbol]{Terminal("["), NonTerminal("rhs"), Terminal("]"), NonTerminal("rhs′")}},   // rhs → "[" rhs "]" rhs′
					{"rhs", String[Symbol]{Terminal("{"), NonTerminal("rhs"), Terminal("}"), NonTerminal("rhs′")}},   // rhs → "{" rhs "}" rhs′
					{"rhs", String[Symbol]{Terminal("{{"), NonTerminal("rhs"), Terminal("}}"), NonTerminal("rhs′")}}, // rhs → "{{" rhs "}}" rhs′
					{"rhs", String[Symbol]{Terminal("IDENT"), NonTerminal("rhs′")}},                                  // rhs → IDENT rhs′
					{"rhs", String[Symbol]{Terminal("TOKEN"), NonTerminal("rhs′")}},                                  // rhs → TOKEN rhs′
					{"rhs", String[Symbol]{Terminal("STRING"), NonTerminal("rhs′")}},                                 // rhs → STRING rhs′
					{"rhs′", String[Symbol]{NonTerminal("rhs"), NonTerminal("rhs′")}},                                // rhs′ → rhs rhs′
					{"rhs′", String[Symbol]{Terminal("|"), NonTerminal("rhs"), NonTerminal("rhs′")}},                 // rhs′ → "|" rhs rhs′
					{"rhs′", ε}, // rhs′ → ε
				},
				"grammar",
			),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NoError(t, tc.g.Verify())
			g := tc.g.EliminateLeftRecursion()
			assert.NoError(t, g.Verify())
			assert.True(t, g.Equals(tc.expectedGrammar))
		})
	}
}

func TestCFG_LeftFactor(t *testing.T) {
	tests := []struct {
		name            string
		g               CFG
		expectedGrammar CFG
	}{
		{
			name:            "1st",
			g:               CFGrammars[0],
			expectedGrammar: CFGrammars[0],
		},
		{
			name:            "2nd",
			g:               CFGrammars[1],
			expectedGrammar: CFGrammars[1],
		},
		{
			name: "3rd",
			g:    CFGrammars[2],
			expectedGrammar: NewCFG(
				[]Terminal{"a", "b"},
				[]NonTerminal{"S", "S′", "A", "B"},
				[]Production{
					{"S", String[Symbol]{Terminal("a"), NonTerminal("S′")}}, // S → aS′
					{"S", String[Symbol]{NonTerminal("A"), Terminal("b")}},  // S → Ab
					{"S′", String[Symbol]{NonTerminal("B"), Terminal("a")}}, // S′ → Ba
					{"S′", ε},                               // S′ → ε
					{"A", String[Symbol]{Terminal("b")}},    // A → b
					{"A", ε},                                // A → ε
					{"B", String[Symbol]{NonTerminal("A")}}, // B → A
					{"B", String[Symbol]{Terminal("b")}},    // B → b
				},
				"S",
			),
		},
		{
			name:            "4th",
			g:               CFGrammars[3],
			expectedGrammar: CFGrammars[3],
		},
		{
			name:            "5th",
			g:               CFGrammars[4],
			expectedGrammar: CFGrammars[4],
		},
		{
			name: "6th",
			g:    CFGrammars[5],
			expectedGrammar: NewCFG(
				[]Terminal{"+", "-", "*", "/", "(", ")", "id"},
				[]NonTerminal{"S", "E", "E′"},
				[]Production{
					{"S", String[Symbol]{NonTerminal("E")}},                               // S → E
					{"E", String[Symbol]{NonTerminal("E"), NonTerminal("E′")}},            // E → EE′
					{"E", String[Symbol]{Terminal("("), NonTerminal("E"), Terminal(")")}}, // E → ( E )
					{"E", String[Symbol]{Terminal("-"), NonTerminal("E")}},                // E → - E
					{"E", String[Symbol]{Terminal("id")}},                                 // E → id
					{"E′", String[Symbol]{Terminal("+"), NonTerminal("E")}},               // E′ → + E
					{"E′", String[Symbol]{Terminal("-"), NonTerminal("E")}},               // E′ → - E
					{"E′", String[Symbol]{Terminal("*"), NonTerminal("E")}},               // E′ → * E
					{"E′", String[Symbol]{Terminal("/"), NonTerminal("E")}},               // E′ → / E
				},
				"S",
			),
		},
		{
			name: "7th",
			g:    CFGrammars[6],
			expectedGrammar: NewCFG(
				[]Terminal{"+", "-", "*", "/", "(", ")", "id"},
				[]NonTerminal{"S", "E", "E′", "T", "T′", "F"},
				[]Production{
					{"S", String[Symbol]{NonTerminal("E")}},                               // S → E
					{"E", String[Symbol]{NonTerminal("E"), NonTerminal("E′")}},            // E → EE′
					{"E", String[Symbol]{NonTerminal("T")}},                               // E → T
					{"E′", String[Symbol]{Terminal("+"), NonTerminal("T")}},               // E′ → + T
					{"E′", String[Symbol]{Terminal("-"), NonTerminal("T")}},               // E′ → - T
					{"T", String[Symbol]{NonTerminal("T"), NonTerminal("T′")}},            // T → TT′
					{"T", String[Symbol]{NonTerminal("F")}},                               // T → F
					{"T′", String[Symbol]{Terminal("*"), NonTerminal("F")}},               // T′ → * F
					{"T′", String[Symbol]{Terminal("/"), NonTerminal("F")}},               // T′ → / F
					{"F", String[Symbol]{Terminal("("), NonTerminal("E"), Terminal(")")}}, // F → ( E )
					{"F", String[Symbol]{Terminal("id")}},                                 // F → id
				},
				"S",
			),
		},
		{
			name: "8th",
			g:    CFGrammars[7],
			expectedGrammar: NewCFG(
				[]Terminal{"=", "|", "(", ")", "[", "]", "{", "}", "{{", "}}", "GRAMMAR", "IDENT", "TOKEN", "STRING", "REGEX"},
				[]NonTerminal{"grammar", "name", "decls", "decl", "token", "rule", "lhs", "rhs", "rhs′", "nonterm", "term"},
				[]Production{
					{"grammar", String[Symbol]{NonTerminal("name"), NonTerminal("decls")}}, // grammar → name decls
					{"name", String[Symbol]{Terminal("GRAMMAR"), Terminal("IDENT")}},       // name → GRAMMAR IDENT
					{"decls", String[Symbol]{NonTerminal("decls"), NonTerminal("decl")}},   // decls → decls decl
					{"decls", ε}, // decls → ε
					{"decl", String[Symbol]{NonTerminal("token")}},                                  // decl → token
					{"decl", String[Symbol]{NonTerminal("rule")}},                                   // decl → rule
					{"token", String[Symbol]{Terminal("TOKEN"), Terminal("="), Terminal("STRING")}}, // token → TOKEN "=" STRING
					{"token", String[Symbol]{Terminal("TOKEN"), Terminal("="), Terminal("REGEX")}},  // token → TOKEN "=" REGEX
					{"rule", String[Symbol]{NonTerminal("lhs"), Terminal("="), NonTerminal("rhs")}}, // rule → lhs "=" rhs
					{"rule", String[Symbol]{NonTerminal("lhs"), Terminal("=")}},                     // rule → lhs "="
					{"lhs", String[Symbol]{NonTerminal("nonterm")}},                                 // lhs → nonterm
					{"rhs", String[Symbol]{NonTerminal("rhs"), NonTerminal("rhs′")}},                // rhs → rhs rhs′
					{"rhs", String[Symbol]{Terminal("("), NonTerminal("rhs"), Terminal(")")}},       // rhs → "(" rhs ")"
					{"rhs", String[Symbol]{Terminal("["), NonTerminal("rhs"), Terminal("]")}},       // rhs → "[" rhs "]"
					{"rhs", String[Symbol]{Terminal("{"), NonTerminal("rhs"), Terminal("}")}},       // rhs → "{" rhs "}"
					{"rhs", String[Symbol]{Terminal("{{"), NonTerminal("rhs"), Terminal("}}")}},     // rhs → "{{" rhs "}}"
					{"rhs", String[Symbol]{NonTerminal("nonterm")}},                                 // rhs → nonterm
					{"rhs", String[Symbol]{NonTerminal("term")}},                                    // rhs → term
					{"rhs′", String[Symbol]{Terminal("|"), NonTerminal("rhs")}},                     // rhs′ → "|" rhs
					{"rhs′", String[Symbol]{NonTerminal("rhs")}},                                    // rhs′ → rhs
					{"nonterm", String[Symbol]{Terminal("IDENT")}},                                  // nonterm → IDENT
					{"term", String[Symbol]{Terminal("TOKEN")}},                                     // term → TOKEN
					{"term", String[Symbol]{Terminal("STRING")}},                                    // term → STRING
				},
				"grammar",
			),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NoError(t, tc.g.Verify())
			g := tc.g.LeftFactor()
			assert.NoError(t, g.Verify())
			assert.True(t, g.Equals(tc.expectedGrammar))
		})
	}
}

func TestGroupByCommonPrefix(t *testing.T) {
	tests := []struct {
		name           string
		prods          set.Set[Production]
		expectedGroups map[string][]string
	}{
		{
			name: "1st",
			prods: set.New(eqProduction,
				Production{"A", String[Symbol]{Terminal("a")}},
				Production{"A", ε},
			),
			expectedGroups: map[string][]string{
				`"a"`: {`ε`},
				`ε`:   {`ε`},
			},
		},
		{
			name: "2nd",
			prods: set.New(eqProduction,
				Production{"stmt", String[Symbol]{NonTerminal("expr")}},
				Production{"stmt", String[Symbol]{Terminal("if"), NonTerminal("expr"), Terminal("then"), NonTerminal("stmt")}},
				Production{"stmt", String[Symbol]{Terminal("if"), NonTerminal("expr"), Terminal("then"), NonTerminal("stmt"), Terminal("else"), NonTerminal("stmt")}},
			),
			expectedGroups: map[string][]string{
				`"if"`: {`expr "then" stmt`, `expr "then" stmt "else" stmt`},
				`expr`: {`ε`},
			},
		},
		{
			name: "3rd",
			prods: set.New(eqProduction,
				Production{"S", String[Symbol]{Terminal("a"), Terminal("b"), Terminal("c"), Terminal("d"), NonTerminal("A"), NonTerminal("B")}},
				Production{"S", String[Symbol]{Terminal("a"), Terminal("b"), Terminal("c"), Terminal("d"), NonTerminal("C"), NonTerminal("D")}},
				Production{"S", String[Symbol]{Terminal("a"), Terminal("b"), Terminal("c"), NonTerminal("E"), NonTerminal("F")}},
				Production{"S", String[Symbol]{Terminal("a"), Terminal("b"), Terminal("c"), NonTerminal("G"), NonTerminal("H")}},
				Production{"S", String[Symbol]{Terminal("a"), Terminal("b"), NonTerminal("I"), NonTerminal("J")}},
				Production{"S", String[Symbol]{Terminal("a"), Terminal("b"), NonTerminal("K"), NonTerminal("L")}},
				Production{"S", String[Symbol]{Terminal("a"), Terminal("b"), NonTerminal("M"), NonTerminal("N")}},
				Production{"S", String[Symbol]{Terminal("a"), NonTerminal("O"), NonTerminal("P")}},
				Production{"S", String[Symbol]{Terminal("a")}},
				Production{"S", String[Symbol]{Terminal("u"), Terminal("v"), NonTerminal("Q"), NonTerminal("R")}},
				Production{"S", String[Symbol]{Terminal("u"), Terminal("v"), Terminal("w"), NonTerminal("S"), NonTerminal("T")}},
				Production{"S", String[Symbol]{Terminal("x"), Terminal("y"), NonTerminal("U"), NonTerminal("V")}},
				Production{"S", String[Symbol]{Terminal("z"), NonTerminal("W")}},
			),
			expectedGroups: map[string][]string{
				`"a"`: {`"b" "c" "d" A B`, `"b" "c" "d" C D`, `"b" "c" E F`, `"b" "c" G H`, `"b" I J`, `"b" K L`, `"b" M N`, `O P`, `ε`},
				`"u"`: {`"v" Q R`, `"v" "w" S T`},
				`"x"`: {`"y" U V`},
				`"z"`: {`W`},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			groups := groupByCommonPrefix(tc.prods)

			for prefix, suffixes := range groups.All() {
				expectedSuffixes, found := tc.expectedGroups[prefix.String()]
				assert.True(t, found, "Prefix %s is not expected", prefix)

				for suffix := range suffixes.All() {
					assert.Contains(t, expectedSuffixes, suffix.String(), "Suffix %s not expected for prefix %s", suffix, prefix)
				}
			}
		})
	}
}

func TestCFG_ChomskyNormalForm(t *testing.T) {
	tests := []struct {
		name            string
		g               CFG
		expectedGrammar CFG
	}{
		{
			name: "1st",
			g:    CFGrammars[0],
			expectedGrammar: NewCFG(
				[]Terminal{"0", "1"},
				[]NonTerminal{
					"S′",
					"X", "Y",
					"S₁",
					"0ₙ", "1ₙ",
				},
				[]Production{
					{"S′", String[Symbol]{NonTerminal("0ₙ"), NonTerminal("X")}}, // S′ → 0ₙX
					{"S′", String[Symbol]{NonTerminal("1ₙ"), NonTerminal("Y")}}, // S′ → 1ₙY
					{"S′", String[Symbol]{NonTerminal("X"), NonTerminal("S₁")}}, // S′ → XS₁
					{"S′", String[Symbol]{NonTerminal("Y"), NonTerminal("X")}},  // S′ → YX
					{"S′", String[Symbol]{Terminal("0")}},                       // S′ → 0
					{"S′", String[Symbol]{Terminal("1")}},                       // S′ → 1
					{"S′", ε},                                                   // S′ → ε
					{"S₁", String[Symbol]{NonTerminal("0ₙ"), NonTerminal("X")}}, // S₁ → 0ₙX
					{"S₁", String[Symbol]{NonTerminal("1ₙ"), NonTerminal("Y")}}, // S₁ → 1ₙY
					{"S₁", String[Symbol]{NonTerminal("Y"), NonTerminal("X")}},  // S₁ → YX
					{"S₁", String[Symbol]{Terminal("0")}},                       // S₁ → 0
					{"S₁", String[Symbol]{Terminal("1")}},                       // S₁ → 1
					{"X", String[Symbol]{NonTerminal("0ₙ"), NonTerminal("X")}},  // X → 0ₙX
					{"X", String[Symbol]{Terminal("0")}},                        // X → 0
					{"Y", String[Symbol]{NonTerminal("1ₙ"), NonTerminal("Y")}},  // Y → 1ₙY
					{"Y", String[Symbol]{Terminal("1")}},                        // Y → 1
					{"0ₙ", String[Symbol]{Terminal("0")}},                       // 0ₙ → 0
					{"1ₙ", String[Symbol]{Terminal("1")}},                       // 1ₙ → 1
				},
				"S′",
			),
		},
		{
			name: "2nd",
			g:    CFGrammars[1],
			expectedGrammar: NewCFG(
				[]Terminal{"a", "b"},
				[]NonTerminal{
					"S″",
					"S",
					"S₁", "S₂", "S₃", "S₄",
					"aₙ", "bₙ",
				},
				[]Production{
					{"S″", String[Symbol]{NonTerminal("aₙ"), NonTerminal("S₁")}}, // S″ → aₙS₁
					{"S″", String[Symbol]{NonTerminal("bₙ"), NonTerminal("S₃")}}, // S″ → bₙS₃
					{"S″", ε}, // S″ → ε
					{"S", String[Symbol]{NonTerminal("aₙ"), NonTerminal("S₁")}}, // S → aₙS₁
					{"S", String[Symbol]{NonTerminal("bₙ"), NonTerminal("S₃")}}, // S → bₙS₃
					{"S₁", String[Symbol]{NonTerminal("S"), NonTerminal("S₂")}}, // S₁ → SS₂
					{"S₁", String[Symbol]{NonTerminal("bₙ"), NonTerminal("S")}}, // S₁ → bₙS
					{"S₁", String[Symbol]{Terminal("b")}},                       // S₁ → b
					{"S₂", String[Symbol]{NonTerminal("bₙ"), NonTerminal("S")}}, // S₂ → bₙS
					{"S₂", String[Symbol]{Terminal("b")}},                       // S₂ → b
					{"S₃", String[Symbol]{NonTerminal("S"), NonTerminal("S₄")}}, // S₃ → SS₄
					{"S₃", String[Symbol]{NonTerminal("aₙ"), NonTerminal("S")}}, // S₃ → aₙS
					{"S₃", String[Symbol]{Terminal("a")}},                       // S₃ → a
					{"S₄", String[Symbol]{NonTerminal("aₙ"), NonTerminal("S")}}, // S₄ → aₙS
					{"S₄", String[Symbol]{Terminal("a")}},                       // S₄ → a
					{"aₙ", String[Symbol]{Terminal("a")}},                       // aₙ → a
					{"bₙ", String[Symbol]{Terminal("b")}},                       // bₙ → b
				},
				"S″",
			),
		},
		{
			name: "3rd",
			g:    CFGrammars[2],
			expectedGrammar: NewCFG(
				[]Terminal{"a", "b"},
				[]NonTerminal{
					"S", "A", "B",
					"S₁",
					"aₙ", "bₙ",
				},
				[]Production{
					{"S", String[Symbol]{NonTerminal("A"), NonTerminal("bₙ")}},  // S → Abₙ
					{"S", String[Symbol]{NonTerminal("aₙ"), NonTerminal("S₁")}}, // S → aₙS₁
					{"S", String[Symbol]{Terminal("a")}},                        // S → a
					{"S", String[Symbol]{Terminal("b")}},                        // S → b
					{"S₁", String[Symbol]{NonTerminal("B"), NonTerminal("aₙ")}}, // S₁ → Baₙ
					{"S₁", String[Symbol]{Terminal("a")}},                       // S₁ → a
					{"A", String[Symbol]{Terminal("b")}},                        // A → b
					{"B", String[Symbol]{Terminal("b")}},                        // B → b
					{"aₙ", String[Symbol]{Terminal("a")}},                       // aₙ → a
					{"bₙ", String[Symbol]{Terminal("b")}},                       // bₙ → b
				},
				"S",
			),
		},
		{
			name: "4th",
			g:    CFGrammars[3],
			expectedGrammar: NewCFG(
				[]Terminal{"b", "d", "s"},
				[]NonTerminal{"S"},
				[]Production{
					{"S", String[Symbol]{Terminal("b")}}, // S → b
					{"S", String[Symbol]{Terminal("d")}}, // S → d
					{"S", String[Symbol]{Terminal("s")}}, // S → s
				},
				"S",
			),
		},
		{
			name: "5th",
			g:    CFGrammars[4],
			expectedGrammar: NewCFG(
				[]Terminal{"a", "b"},
				[]NonTerminal{
					"S", "A", "B",
					"aₙ", "bₙ",
				},
				[]Production{
					{"S", String[Symbol]{NonTerminal("A"), NonTerminal("B")}},  // S → AB
					{"A", String[Symbol]{NonTerminal("aₙ"), NonTerminal("A")}}, // A → aₙA
					{"A", String[Symbol]{Terminal("a")}},                       // A → a
					{"B", String[Symbol]{NonTerminal("bₙ"), NonTerminal("B")}}, // B → bₙB
					{"B", String[Symbol]{Terminal("b")}},                       // B → b
					{"aₙ", String[Symbol]{Terminal("a")}},                      // aₙ → a
					{"bₙ", String[Symbol]{Terminal("b")}},                      // bₙ → b
				},
				"S",
			),
		},
		{
			name: "6th",
			g:    CFGrammars[5],
			expectedGrammar: NewCFG(
				[]Terminal{"+", "-", "*", "/", "(", ")", "id"},
				[]NonTerminal{
					"S", "E",
					"E₁", "E₂", "E₃", "E₄", "E₅",
					"+ₙ", "-ₙ", "*ₙ", "/ₙ", "(ₙ", ")ₙ",
				},
				[]Production{
					{"S", String[Symbol]{NonTerminal("(ₙ"), NonTerminal("E₁")}}, // S → (ₙ E₁
					{"S", String[Symbol]{NonTerminal("E"), NonTerminal("E₂")}},  // S → E E₂
					{"S", String[Symbol]{NonTerminal("E"), NonTerminal("E₃")}},  // S → E E₃
					{"S", String[Symbol]{NonTerminal("E"), NonTerminal("E₄")}},  // S → E E₄
					{"S", String[Symbol]{NonTerminal("E"), NonTerminal("E₅")}},  // S → E E₅
					{"S", String[Symbol]{NonTerminal("-ₙ"), NonTerminal("E")}},  // S → -ₙ E
					{"S", String[Symbol]{Terminal("id")}},                       // S → id
					{"E", String[Symbol]{NonTerminal("(ₙ"), NonTerminal("E₁")}}, // E → (ₙ E₁
					{"E", String[Symbol]{NonTerminal("E"), NonTerminal("E₂")}},  // E → E E₂
					{"E", String[Symbol]{NonTerminal("E"), NonTerminal("E₃")}},  // E → E E₃
					{"E", String[Symbol]{NonTerminal("E"), NonTerminal("E₄")}},  // E → E E₄
					{"E", String[Symbol]{NonTerminal("E"), NonTerminal("E₅")}},  // E → E E₅
					{"E", String[Symbol]{NonTerminal("-ₙ"), NonTerminal("E")}},  // E → -ₙ E
					{"E", String[Symbol]{Terminal("id")}},                       // E → id
					{"E₁", String[Symbol]{NonTerminal("E"), NonTerminal(")ₙ")}}, // E₁ → E )ₙ
					{"E₂", String[Symbol]{NonTerminal("*ₙ"), NonTerminal("E")}}, // E₂ → *ₙ E
					{"E₃", String[Symbol]{NonTerminal("+ₙ"), NonTerminal("E")}}, // E₃ → +ₙ E
					{"E₄", String[Symbol]{NonTerminal("-ₙ"), NonTerminal("E")}}, // E₄ → -ₙ E
					{"E₅", String[Symbol]{NonTerminal("/ₙ"), NonTerminal("E")}}, // E₅ → /ₙ E
					{"+ₙ", String[Symbol]{Terminal("+")}},                       // +ₙ → +
					{"-ₙ", String[Symbol]{Terminal("-")}},                       // -ₙ → -
					{"*ₙ", String[Symbol]{Terminal("*")}},                       // *ₙ → *
					{"/ₙ", String[Symbol]{Terminal("/")}},                       // /ₙ → /
					{"(ₙ", String[Symbol]{Terminal("(")}},                       // (ₙ → (
					{")ₙ", String[Symbol]{Terminal(")")}},                       // )ₙ → )
				},
				"S",
			),
		},
		{
			name: "7th",
			g:    CFGrammars[6],
			expectedGrammar: NewCFG(
				[]Terminal{"+", "-", "*", "/", "(", ")", "id"},
				[]NonTerminal{
					"S", "E", "T", "F",
					"E₁", "E₂", "T₁", "T₂", "F₁",
					"+ₙ", "-ₙ", "*ₙ", "/ₙ", "(ₙ", ")ₙ",
				},
				[]Production{
					{"S", String[Symbol]{NonTerminal("E"), NonTerminal("E₁")}},  // S → E E₁
					{"S", String[Symbol]{NonTerminal("E"), NonTerminal("E₂")}},  // S → E E₂
					{"S", String[Symbol]{NonTerminal("T"), NonTerminal("T₁")}},  // S → T T₁
					{"S", String[Symbol]{NonTerminal("T"), NonTerminal("T₂")}},  // S → T T₂
					{"S", String[Symbol]{NonTerminal("(ₙ"), NonTerminal("F₁")}}, // S → (ₙ F₁
					{"S", String[Symbol]{Terminal("id")}},                       // S → id
					{"E", String[Symbol]{NonTerminal("E"), NonTerminal("E₁")}},  // E → E E₁
					{"E", String[Symbol]{NonTerminal("E"), NonTerminal("E₂")}},  // E → E E₂
					{"E", String[Symbol]{NonTerminal("T"), NonTerminal("T₁")}},  // E → T T₁
					{"E", String[Symbol]{NonTerminal("T"), NonTerminal("T₂")}},  // E → T T₂
					{"E", String[Symbol]{NonTerminal("(ₙ"), NonTerminal("F₁")}}, // E → (ₙ F₁
					{"E", String[Symbol]{Terminal("id")}},                       // E → id
					{"E₁", String[Symbol]{NonTerminal("+ₙ"), NonTerminal("T")}}, // E₁ → +ₙ T
					{"E₂", String[Symbol]{NonTerminal("-ₙ"), NonTerminal("T")}}, // E₂ → -ₙ T
					{"T", String[Symbol]{NonTerminal("T"), NonTerminal("T₁")}},  // T → T T₁
					{"T", String[Symbol]{NonTerminal("T"), NonTerminal("T₂")}},  // T → T T₂
					{"T", String[Symbol]{NonTerminal("(ₙ"), NonTerminal("F₁")}}, // T → (ₙ F₁
					{"T", String[Symbol]{Terminal("id")}},                       // T → id
					{"T₁", String[Symbol]{NonTerminal("*ₙ"), NonTerminal("F")}}, // T₁ → *ₙ F
					{"T₂", String[Symbol]{NonTerminal("/ₙ"), NonTerminal("F")}}, // T₂ → /ₙ F
					{"F", String[Symbol]{NonTerminal("(ₙ"), NonTerminal("F₁")}}, // F → (ₙ F₁
					{"F", String[Symbol]{Terminal("id")}},                       // F → id
					{"F₁", String[Symbol]{NonTerminal("E"), NonTerminal(")ₙ")}}, // F₁ → E )ₙ
					{"+ₙ", String[Symbol]{Terminal("+")}},                       // +ₙ → +
					{"-ₙ", String[Symbol]{Terminal("-")}},                       // -ₙ → -
					{"*ₙ", String[Symbol]{Terminal("*")}},                       // *ₙ → *
					{"/ₙ", String[Symbol]{Terminal("/")}},                       // /ₙ → /
					{"(ₙ", String[Symbol]{Terminal("(")}},                       // (ₙ → (
					{")ₙ", String[Symbol]{Terminal(")")}},                       // )ₙ → )
				},
				"S",
			),
		},
		{
			name: "8th",
			g:    CFGrammars[7],
			expectedGrammar: NewCFG(
				[]Terminal{"=", "|", "(", ")", "[", "]", "{", "}", "{{", "}}", "GRAMMAR", "IDENT", "TOKEN", "STRING", "REGEX"},
				[]NonTerminal{
					"grammar", "name", "decls", "decl", "lhs", "rhs",
					"token₁", "token₂", "rule₁", "rhs₁", "rhs₂", "rhs₃", "rhs₄", "rhs₅",
					"=ₙ", "|ₙ", "(ₙ", ")ₙ", "[ₙ", "]ₙ", "{ₙ", "}ₙ", "{{ₙ", "}}ₙ", "GRAMMARₙ", "IDENTₙ", "TOKENₙ", "STRINGₙ", "REGEXₙ",
				},
				[]Production{
					{"grammar", String[Symbol]{NonTerminal("name"), NonTerminal("decls")}},      // grammar → name decls
					{"grammar", String[Symbol]{NonTerminal("GRAMMARₙ"), NonTerminal("IDENTₙ")}}, // grammar → GRAMMARₙ IDENTₙ
					{"name", String[Symbol]{NonTerminal("GRAMMARₙ"), NonTerminal("IDENTₙ")}},    // name → GRAMMARₙ IDENTₙ
					{"decls", String[Symbol]{NonTerminal("decls"), NonTerminal("decl")}},        // decls → decls decl
					{"decls", String[Symbol]{NonTerminal("TOKENₙ"), NonTerminal("token₁")}},     // decls → TOKENₙ token₁
					{"decls", String[Symbol]{NonTerminal("TOKENₙ"), NonTerminal("token₂")}},     // decls → TOKENₙ token₂
					{"decls", String[Symbol]{NonTerminal("lhs"), NonTerminal("rule₁")}},         // decls → lhs rule₁
					{"decls", String[Symbol]{NonTerminal("lhs"), NonTerminal("=ₙ")}},            // decls → lhs =ₙ
					{"decl", String[Symbol]{NonTerminal("TOKENₙ"), NonTerminal("token₁")}},      // decl → TOKENₙ token₁
					{"decl", String[Symbol]{NonTerminal("TOKENₙ"), NonTerminal("token₂")}},      // decl → TOKENₙ token₂
					{"decl", String[Symbol]{NonTerminal("lhs"), NonTerminal("rule₁")}},          // decl → lhs rule₁
					{"decl", String[Symbol]{NonTerminal("lhs"), NonTerminal("=ₙ")}},             // decl → lhs =ₙ
					{"token₁", String[Symbol]{NonTerminal("=ₙ"), NonTerminal("REGEXₙ")}},        // token₁ → =ₙ REGEXₙ
					{"token₂", String[Symbol]{NonTerminal("=ₙ"), NonTerminal("STRINGₙ")}},       // token₂ → =ₙ STRINGₙ
					{"rule₁", String[Symbol]{NonTerminal("=ₙ"), NonTerminal("rhs")}},            // rule₁ → =ₙ rhs
					{"lhs", String[Symbol]{Terminal("IDENT")}},                                  // lhs → IDENT
					{"rhs", String[Symbol]{NonTerminal("rhs"), NonTerminal("rhs")}},             // rhs → rhs rhs
					{"rhs", String[Symbol]{NonTerminal("(ₙ"), NonTerminal("rhs₁")}},             // rhs → (ₙ rhs₁
					{"rhs", String[Symbol]{NonTerminal("[ₙ"), NonTerminal("rhs₂")}},             // rhs → [ₙ rhs₂
					{"rhs", String[Symbol]{NonTerminal("rhs"), NonTerminal("rhs₃")}},            // rhs → rhs rhs₃
					{"rhs", String[Symbol]{NonTerminal("{{ₙ"), NonTerminal("rhs₄")}},            // rhs → {{ₙ rhs₄
					{"rhs", String[Symbol]{NonTerminal("{ₙ"), NonTerminal("rhs₅")}},             // rhs → {ₙ rhs₅
					{"rhs", String[Symbol]{Terminal("IDENT")}},                                  // rhs → IDENT
					{"rhs", String[Symbol]{Terminal("TOKEN")}},                                  // rhs → TOKEN
					{"rhs", String[Symbol]{Terminal("STRING")}},                                 // rhs → STRING
					{"rhs₁", String[Symbol]{NonTerminal("rhs"), NonTerminal(")ₙ")}},             // rhs₁ → rhs )ₙ
					{"rhs₂", String[Symbol]{NonTerminal("rhs"), NonTerminal("]ₙ")}},             // rhs₂ → rhs ]ₙ
					{"rhs₃", String[Symbol]{NonTerminal("|ₙ"), NonTerminal("rhs")}},             // rhs₃ → |ₙ rhs
					{"rhs₄", String[Symbol]{NonTerminal("rhs"), NonTerminal("}}ₙ")}},            // rhs₄ → rhs }}ₙ
					{"rhs₅", String[Symbol]{NonTerminal("rhs"), NonTerminal("}ₙ")}},             // rhs₅ → rhs }ₙ
					{"=ₙ", String[Symbol]{Terminal("=")}},                                       // =ₙ → =
					{"|ₙ", String[Symbol]{Terminal("|")}},                                       // |ₙ → |
					{"(ₙ", String[Symbol]{Terminal("(")}},                                       // (ₙ → (
					{")ₙ", String[Symbol]{Terminal(")")}},                                       // )ₙ → )
					{"[ₙ", String[Symbol]{Terminal("[")}},                                       // [ₙ → [
					{"]ₙ", String[Symbol]{Terminal("]")}},                                       // ]ₙ → ]
					{"{ₙ", String[Symbol]{Terminal("{")}},                                       // {ₙ → {
					{"}ₙ", String[Symbol]{Terminal("}")}},                                       // }ₙ → }
					{"{{ₙ", String[Symbol]{Terminal("{{")}},                                     // {{ₙ → {{
					{"}}ₙ", String[Symbol]{Terminal("}}")}},                                     // }}ₙ → }}
					{"GRAMMARₙ", String[Symbol]{Terminal("GRAMMAR")}},                           // GRAMMARₙ → GRAMMAR
					{"IDENTₙ", String[Symbol]{Terminal("IDENT")}},                               // IDENTₙ → IDENT
					{"TOKENₙ", String[Symbol]{Terminal("TOKEN")}},                               // TOKENₙ → TOKEN
					{"STRINGₙ", String[Symbol]{Terminal("STRING")}},                             // STRINGₙ → STRING
					{"REGEXₙ", String[Symbol]{Terminal("REGEX")}},                               // REGEXₙ → REGEX
				},
				"grammar",
			),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NoError(t, tc.g.Verify())
			g := tc.g.ChomskyNormalForm()
			assert.NoError(t, g.Verify())
			assert.True(t, g.Equals(tc.expectedGrammar))
		})
	}
}

func TestCFG_ComputeFIRST(t *testing.T) {
	tests := []struct {
		name           string
		g              CFG
		firsts         []String[Symbol]
		expectedFirsts []TerminalsAndEmpty
	}{
		{
			name: "1st",
			g:    CFGrammars[0],
			firsts: []String[Symbol]{
				ε,                                    // ε
				{Terminal("0")},                      // 0
				{Terminal("1")},                      // 1
				{NonTerminal("S")},                   // S
				{NonTerminal("X")},                   // X
				{NonTerminal("Y")},                   // Y
				{Terminal("0"), Terminal("1")},       // 01
				{Terminal("1"), Terminal("0")},       // 10
				{NonTerminal("X"), NonTerminal("Y")}, // XY
				{NonTerminal("Y"), NonTerminal("X")}, // YX
				{NonTerminal("X"), NonTerminal("Y"), NonTerminal("X")}, // XYX
				{NonTerminal("Y"), NonTerminal("X"), NonTerminal("Y")}, // YXY
			},
			expectedFirsts: []TerminalsAndEmpty{
				{Terminals: set.New(eqTerminal), IncludesEmpty: true},           // FIRST(ε)
				{Terminals: set.New(eqTerminal, "0"), IncludesEmpty: false},     // FIRST(0)
				{Terminals: set.New(eqTerminal, "1"), IncludesEmpty: false},     // FIRST(1)
				{Terminals: set.New(eqTerminal, "0", "1"), IncludesEmpty: true}, // FIRST(S)
				{Terminals: set.New(eqTerminal, "0"), IncludesEmpty: true},      // FIRST(X)
				{Terminals: set.New(eqTerminal, "1"), IncludesEmpty: true},      // FIRST(Y)
				{Terminals: set.New(eqTerminal, "0"), IncludesEmpty: false},     // FIRST(01)
				{Terminals: set.New(eqTerminal, "1"), IncludesEmpty: false},     // FIRST(10)
				{Terminals: set.New(eqTerminal, "0", "1"), IncludesEmpty: true}, // FIRST(XY)
				{Terminals: set.New(eqTerminal, "1", "0"), IncludesEmpty: true}, // FIRST(YX)
				{Terminals: set.New(eqTerminal, "0", "1"), IncludesEmpty: true}, // FIRST(XYX)
				{Terminals: set.New(eqTerminal, "1", "0"), IncludesEmpty: true}, // FIRST(YXY)
			},
		},
		{
			name: "2nd",
			g:    CFGrammars[1],
			firsts: []String[Symbol]{
				ε,                              // ε
				{Terminal("a")},                // a
				{Terminal("b")},                // b
				{NonTerminal("S")},             // S
				{Terminal("a"), Terminal("b")}, // ab
				{Terminal("b"), Terminal("a")}, // ba
				{NonTerminal("S"), Terminal("a"), Terminal("b")}, // Sab
				{NonTerminal("S"), Terminal("b"), Terminal("a")}, // Sba
			},
			expectedFirsts: []TerminalsAndEmpty{
				{Terminals: set.New(eqTerminal), IncludesEmpty: true},            // FIRST(ε)
				{Terminals: set.New(eqTerminal, "a"), IncludesEmpty: false},      // FIRST(a)
				{Terminals: set.New(eqTerminal, "b"), IncludesEmpty: false},      // FIRST(b)
				{Terminals: set.New(eqTerminal, "a", "b"), IncludesEmpty: true},  // FIRST(S)
				{Terminals: set.New(eqTerminal, "a"), IncludesEmpty: false},      // FIRST(ab)
				{Terminals: set.New(eqTerminal, "b"), IncludesEmpty: false},      // FIRST(ba)
				{Terminals: set.New(eqTerminal, "a", "b"), IncludesEmpty: false}, // FIRST(Sab)
				{Terminals: set.New(eqTerminal, "b", "a"), IncludesEmpty: false}, // FIRST(Sba)
			},
		},
		{
			name: "3rd",
			g:    CFGrammars[2],
			firsts: []String[Symbol]{
				ε,                                    // ε
				{Terminal("a")},                      // a
				{Terminal("b")},                      // b
				{NonTerminal("S")},                   // S
				{NonTerminal("A")},                   // A
				{NonTerminal("B")},                   // B
				{Terminal("a"), Terminal("b")},       // ab
				{Terminal("b"), Terminal("a")},       // ba
				{NonTerminal("A"), NonTerminal("B")}, // AB
				{NonTerminal("B"), NonTerminal("A")}, // BA
				{NonTerminal("A"), NonTerminal("B"), NonTerminal("A")}, // ABA
				{NonTerminal("B"), NonTerminal("A"), NonTerminal("B")}, // BAB
			},
			expectedFirsts: []TerminalsAndEmpty{
				{Terminals: set.New(eqTerminal), IncludesEmpty: true},            // FIRST(ε)
				{Terminals: set.New(eqTerminal, "a"), IncludesEmpty: false},      // FIRST(a)
				{Terminals: set.New(eqTerminal, "b"), IncludesEmpty: false},      // FIRST(b)
				{Terminals: set.New(eqTerminal, "a", "b"), IncludesEmpty: false}, // FIRST(S)
				{Terminals: set.New(eqTerminal, "b"), IncludesEmpty: true},       // FIRST(A)
				{Terminals: set.New(eqTerminal, "b"), IncludesEmpty: true},       // FIRST(B)
				{Terminals: set.New(eqTerminal, "a"), IncludesEmpty: false},      // FIRST(ab)
				{Terminals: set.New(eqTerminal, "b"), IncludesEmpty: false},      // FIRST(ba)
				{Terminals: set.New(eqTerminal, "b"), IncludesEmpty: true},       // FIRST(AB)
				{Terminals: set.New(eqTerminal, "b"), IncludesEmpty: true},       // FIRST(BA)
				{Terminals: set.New(eqTerminal, "b"), IncludesEmpty: true},       // FIRST(ABA)
				{Terminals: set.New(eqTerminal, "b"), IncludesEmpty: true},       // FIRST(BAB)
			},
		},
		{
			name: "4th",
			g:    CFGrammars[3],
			firsts: []String[Symbol]{
				ε,                                    // ε
				{Terminal("b")},                      // b
				{Terminal("c")},                      // c
				{Terminal("d")},                      // d
				{Terminal("s")},                      // s
				{NonTerminal("S")},                   // S
				{NonTerminal("A")},                   // A
				{NonTerminal("B")},                   // B
				{NonTerminal("C")},                   // C
				{NonTerminal("D")},                   // D
				{NonTerminal("A"), NonTerminal("B")}, // AB
				{NonTerminal("B"), NonTerminal("C")}, // BC
				{NonTerminal("C"), NonTerminal("D")}, // CD
				{NonTerminal("A"), NonTerminal("B"), NonTerminal("C")},                   // ABC
				{NonTerminal("B"), NonTerminal("C"), NonTerminal("D")},                   // BCD
				{NonTerminal("A"), NonTerminal("B"), NonTerminal("C"), NonTerminal("D")}, // ABCD
			},
			expectedFirsts: []TerminalsAndEmpty{
				{Terminals: set.New(eqTerminal), IncludesEmpty: true},                 // FIRST(ε)
				{Terminals: set.New(eqTerminal, "b"), IncludesEmpty: false},           // FIRST(b)
				{Terminals: set.New(eqTerminal, "c"), IncludesEmpty: false},           // FIRST(c)
				{Terminals: set.New(eqTerminal, "d"), IncludesEmpty: false},           // FIRST(d)
				{Terminals: set.New(eqTerminal, "s"), IncludesEmpty: false},           // FIRST(s)
				{Terminals: set.New(eqTerminal, "b", "d", "s"), IncludesEmpty: false}, // FIRST(S)
				{Terminals: set.New(eqTerminal, "d", "b"), IncludesEmpty: false},      // FIRST(A)
				{Terminals: set.New(eqTerminal, "b", "d"), IncludesEmpty: false},      // FIRST(B)
				{Terminals: set.New(eqTerminal, "d"), IncludesEmpty: false},           // FIRST(C)
				{Terminals: set.New(eqTerminal, "d"), IncludesEmpty: false},           // FIRST(D)
				{Terminals: set.New(eqTerminal, "b", "d"), IncludesEmpty: false},      // FIRST(AB)
				{Terminals: set.New(eqTerminal, "b", "d"), IncludesEmpty: false},      // FIRST(BC)
				{Terminals: set.New(eqTerminal, "d"), IncludesEmpty: false},           // FIRST(CD)
				{Terminals: set.New(eqTerminal, "b", "d"), IncludesEmpty: false},      // FIRST(ABC)
				{Terminals: set.New(eqTerminal, "b", "d"), IncludesEmpty: false},      // FIRST(BCD)
				{Terminals: set.New(eqTerminal, "b", "d"), IncludesEmpty: false},      // FIRST(ABCD)
			},
		},
		{
			name: "5th",
			g:    CFGrammars[4],
			firsts: []String[Symbol]{
				ε,                                    // ε
				{Terminal("a")},                      // a
				{Terminal("b")},                      // b
				{Terminal("c")},                      // c
				{Terminal("d")},                      // d
				{NonTerminal("S")},                   // S
				{NonTerminal("A")},                   // A
				{NonTerminal("B")},                   // B
				{NonTerminal("C")},                   // C
				{NonTerminal("D")},                   // D
				{NonTerminal("A"), NonTerminal("B")}, // AB
				{NonTerminal("B"), NonTerminal("C")}, // BC
				{NonTerminal("C"), NonTerminal("D")}, // CD
				{NonTerminal("A"), NonTerminal("B"), NonTerminal("C")},                   // ABC
				{NonTerminal("B"), NonTerminal("C"), NonTerminal("D")},                   // BCD
				{NonTerminal("A"), NonTerminal("B"), NonTerminal("C"), NonTerminal("D")}, // ABCD
			},
			expectedFirsts: []TerminalsAndEmpty{
				{Terminals: set.New(eqTerminal), IncludesEmpty: true},       // FIRST(ε)
				{Terminals: set.New(eqTerminal, "a"), IncludesEmpty: false}, // FIRST(a)
				{Terminals: set.New(eqTerminal, "b"), IncludesEmpty: false}, // FIRST(b)
				{Terminals: set.New(eqTerminal, "c"), IncludesEmpty: false}, // FIRST(c)
				{Terminals: set.New(eqTerminal, "d"), IncludesEmpty: false}, // FIRST(d)
				{Terminals: set.New(eqTerminal, "a"), IncludesEmpty: false}, // FIRST(S)
				{Terminals: set.New(eqTerminal, "a"), IncludesEmpty: false}, // FIRST(A)
				{Terminals: set.New(eqTerminal, "b"), IncludesEmpty: false}, // FIRST(B)
				{Terminals: set.New(eqTerminal, "c"), IncludesEmpty: false}, // FIRST(C)
				{Terminals: set.New(eqTerminal, "d"), IncludesEmpty: false}, // FIRST(D)
				{Terminals: set.New(eqTerminal, "a"), IncludesEmpty: false}, // FIRST(AB)
				{Terminals: set.New(eqTerminal, "b"), IncludesEmpty: false}, // FIRST(BC)
				{Terminals: set.New(eqTerminal, "c"), IncludesEmpty: false}, // FIRST(CD)
				{Terminals: set.New(eqTerminal, "a"), IncludesEmpty: false}, // FIRST(ABC)
				{Terminals: set.New(eqTerminal, "b"), IncludesEmpty: false}, // FIRST(BCD)
				{Terminals: set.New(eqTerminal, "a"), IncludesEmpty: false}, // FIRST(ABCD)
			},
		},
		{
			name: "6th",
			g:    CFGrammars[5],
			firsts: []String[Symbol]{
				ε,                                    // ε
				{Terminal("+")},                      // +
				{Terminal("-")},                      // -
				{Terminal("*")},                      // *
				{Terminal("/")},                      // /
				{Terminal("(")},                      // (
				{Terminal(")")},                      // )
				{Terminal("id")},                     // id
				{NonTerminal("S")},                   // S
				{NonTerminal("E")},                   // E
				{NonTerminal("S"), NonTerminal("E")}, // SE
				{NonTerminal("E"), NonTerminal("E")}, // EE
			},
			expectedFirsts: []TerminalsAndEmpty{
				{Terminals: set.New(eqTerminal), IncludesEmpty: true},                  // FIRST(ε)
				{Terminals: set.New(eqTerminal, "+"), IncludesEmpty: false},            // FIRST(+)
				{Terminals: set.New(eqTerminal, "-"), IncludesEmpty: false},            // FIRST(-)
				{Terminals: set.New(eqTerminal, "*"), IncludesEmpty: false},            // FIRST(*)
				{Terminals: set.New(eqTerminal, "/"), IncludesEmpty: false},            // FIRST(/)
				{Terminals: set.New(eqTerminal, "("), IncludesEmpty: false},            // FIRST(()
				{Terminals: set.New(eqTerminal, ")"), IncludesEmpty: false},            // FIRST())
				{Terminals: set.New(eqTerminal, "id"), IncludesEmpty: false},           // FIRST(id)
				{Terminals: set.New(eqTerminal, "-", "(", "id"), IncludesEmpty: false}, // FIRST(S)
				{Terminals: set.New(eqTerminal, "-", "(", "id"), IncludesEmpty: false}, // FIRST(E)
				{Terminals: set.New(eqTerminal, "-", "(", "id"), IncludesEmpty: false}, // FIRST(SE)
				{Terminals: set.New(eqTerminal, "-", "(", "id"), IncludesEmpty: false}, // FIRST(EE)
			},
		},
		{
			name: "7th",
			g:    CFGrammars[6],
			firsts: []String[Symbol]{
				ε,                                    // ε
				{Terminal("+")},                      // +
				{Terminal("-")},                      // -
				{Terminal("*")},                      // *
				{Terminal("/")},                      // /
				{Terminal("(")},                      // (
				{Terminal(")")},                      // )
				{Terminal("id")},                     // id
				{NonTerminal("S")},                   // S
				{NonTerminal("E")},                   // E
				{NonTerminal("T")},                   // T
				{NonTerminal("F")},                   // F
				{NonTerminal("E"), NonTerminal("T")}, // ET
				{NonTerminal("E"), NonTerminal("F")}, // EF
				{NonTerminal("T"), NonTerminal("F")}, // TF
				{NonTerminal("E"), NonTerminal("T"), NonTerminal("F")}, // ETF
			},
			expectedFirsts: []TerminalsAndEmpty{
				{Terminals: set.New(eqTerminal), IncludesEmpty: true},             // FIRST(ε)
				{Terminals: set.New(eqTerminal, "+"), IncludesEmpty: false},       // FIRST(+)
				{Terminals: set.New(eqTerminal, "-"), IncludesEmpty: false},       // FIRST(-)
				{Terminals: set.New(eqTerminal, "*"), IncludesEmpty: false},       // FIRST(*)
				{Terminals: set.New(eqTerminal, "/"), IncludesEmpty: false},       // FIRST(/)
				{Terminals: set.New(eqTerminal, "("), IncludesEmpty: false},       // FIRST(()
				{Terminals: set.New(eqTerminal, ")"), IncludesEmpty: false},       // FIRST())
				{Terminals: set.New(eqTerminal, "id"), IncludesEmpty: false},      // FIRST(id)
				{Terminals: set.New(eqTerminal, "(", "id"), IncludesEmpty: false}, // FIRST(S)
				{Terminals: set.New(eqTerminal, "(", "id"), IncludesEmpty: false}, // FIRST(E)
				{Terminals: set.New(eqTerminal, "(", "id"), IncludesEmpty: false}, // FIRST(T)
				{Terminals: set.New(eqTerminal, "(", "id"), IncludesEmpty: false}, // FIRST(F)
				{Terminals: set.New(eqTerminal, "(", "id"), IncludesEmpty: false}, // FIRST(ET)
				{Terminals: set.New(eqTerminal, "(", "id"), IncludesEmpty: false}, // FIRST(EF)
				{Terminals: set.New(eqTerminal, "(", "id"), IncludesEmpty: false}, // FIRST(TF)
				{Terminals: set.New(eqTerminal, "(", "id"), IncludesEmpty: false}, // FIRST(ETF)
			},
		},
		{
			name: "8th",
			g:    CFGrammars[7],
			firsts: []String[Symbol]{
				ε,                        // ε
				{Terminal("=")},          // =
				{Terminal("|")},          // |
				{Terminal("(")},          // (
				{Terminal(")")},          // )
				{Terminal("[")},          // [
				{Terminal("]")},          // ]
				{Terminal("{")},          // {
				{Terminal("}")},          // }
				{Terminal("{{")},         // {{
				{Terminal("}}")},         // }}
				{Terminal("GRAMMAR")},    // GRAMMAR
				{Terminal("IDENT")},      // IDENT
				{Terminal("TOKEN")},      // TOKEN
				{Terminal("STRING")},     // STRING
				{Terminal("REGEX")},      // REGEX
				{NonTerminal("grammar")}, // grammar
				{NonTerminal("name")},    // name
				{NonTerminal("decls")},   // decls
				{NonTerminal("decl")},    // decl
				{NonTerminal("token")},   // token
				{NonTerminal("rule")},    // rule
				{NonTerminal("lhs")},     // lhs
				{NonTerminal("rhs")},     // rhs
				{NonTerminal("nonterm")}, // nonterm
				{NonTerminal("term")},    // term
				{Terminal("TOKEN"), Terminal("="), Terminal("REGEX")},                                           // TOKEN "=" REGEX
				{NonTerminal("rhs"), NonTerminal("rhs"), NonTerminal("rhs"), NonTerminal("rhs")},                // rhs rhs rhs
				{NonTerminal("rhs"), Terminal("|"), NonTerminal("rhs"), Terminal("|"), NonTerminal("rhs")},      // rhs "|" rhs "|" rhs
				{NonTerminal("nonterm"), Terminal("="), Terminal("{{"), Terminal("}}"), NonTerminal("nonterm")}, // nonterm "=" "{{" nonterm "}}"
			},
			expectedFirsts: []TerminalsAndEmpty{
				{Terminals: set.New(eqTerminal), IncludesEmpty: true},                                                   // FIRST(ε)
				{Terminals: set.New(eqTerminal, "="), IncludesEmpty: false},                                             // FIRST(=)
				{Terminals: set.New(eqTerminal, "|"), IncludesEmpty: false},                                             // FIRST(|)
				{Terminals: set.New(eqTerminal, "("), IncludesEmpty: false},                                             // FIRST(()
				{Terminals: set.New(eqTerminal, ")"), IncludesEmpty: false},                                             // FIRST())
				{Terminals: set.New(eqTerminal, "["), IncludesEmpty: false},                                             // FIRST([)
				{Terminals: set.New(eqTerminal, "]"), IncludesEmpty: false},                                             // FIRST(])
				{Terminals: set.New(eqTerminal, "{"), IncludesEmpty: false},                                             // FIRST({)
				{Terminals: set.New(eqTerminal, "}"), IncludesEmpty: false},                                             // FIRST(})
				{Terminals: set.New(eqTerminal, "{{"), IncludesEmpty: false},                                            // FIRST({{)
				{Terminals: set.New(eqTerminal, "}}"), IncludesEmpty: false},                                            // FIRST(}})
				{Terminals: set.New(eqTerminal, "GRAMMAR"), IncludesEmpty: false},                                       // FIRST(GRAMMAR)
				{Terminals: set.New(eqTerminal, "IDENT"), IncludesEmpty: false},                                         // FIRST(IDENT)
				{Terminals: set.New(eqTerminal, "TOKEN"), IncludesEmpty: false},                                         // FIRST(TOKEN)
				{Terminals: set.New(eqTerminal, "STRING"), IncludesEmpty: false},                                        // FIRST(STRING)
				{Terminals: set.New(eqTerminal, "REGEX"), IncludesEmpty: false},                                         // FIRST(REGEX)
				{Terminals: set.New(eqTerminal, "GRAMMAR"), IncludesEmpty: false},                                       // FIRST(grammar)
				{Terminals: set.New(eqTerminal, "GRAMMAR"), IncludesEmpty: false},                                       // FIRST(name)
				{Terminals: set.New(eqTerminal, "TOKEN", "IDENT"), IncludesEmpty: true},                                 // FIRST(decls)
				{Terminals: set.New(eqTerminal, "TOKEN", "IDENT"), IncludesEmpty: false},                                // FIRST(decl)
				{Terminals: set.New(eqTerminal, "TOKEN"), IncludesEmpty: false},                                         // FIRST(token)
				{Terminals: set.New(eqTerminal, "IDENT"), IncludesEmpty: false},                                         // FIRST(rule)
				{Terminals: set.New(eqTerminal, "IDENT"), IncludesEmpty: false},                                         // FIRST(lhs)
				{Terminals: set.New(eqTerminal, "(", "[", "{", "{{", "IDENT", "TOKEN", "STRING"), IncludesEmpty: false}, // FIRST(rhs)
				{Terminals: set.New(eqTerminal, "IDENT"), IncludesEmpty: false},                                         // FIRST(nonterm)
				{Terminals: set.New(eqTerminal, "TOKEN", "STRING"), IncludesEmpty: false},                               // FIRST(term)
				{Terminals: set.New(eqTerminal, "TOKEN"), IncludesEmpty: false},                                         // FIRST(TOKEN "=" REGEX)
				{Terminals: set.New(eqTerminal, "(", "[", "{", "{{", "IDENT", "TOKEN", "STRING"), IncludesEmpty: false}, // FIRST(rhs rhs rhs)
				{Terminals: set.New(eqTerminal, "(", "[", "{", "{{", "IDENT", "TOKEN", "STRING"), IncludesEmpty: false}, // FIRST(rhs "|" rhs "|" rhs)
				{Terminals: set.New(eqTerminal, "IDENT"), IncludesEmpty: false},                                         // FIRST(nonterm "=" "{{" nonterm "}}")
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NoError(t, tc.g.Verify())
			first := tc.g.ComputeFIRST()

			for i, s := range tc.firsts {
				t.Run(strconv.Itoa(i), func(t *testing.T) {
					f := first(s)
					assert.True(t, tc.expectedFirsts[i].Terminals.Equals(f.Terminals))
					assert.Equal(t, tc.expectedFirsts[i].IncludesEmpty, f.IncludesEmpty)
				})
			}
		})
	}
}

func TestCFG_ComputeFOLLOW(t *testing.T) {
	tests := []struct {
		name            string
		g               CFG
		follows         []NonTerminal
		expectedFollows []TerminalsAndEndmarker
	}{
		{
			name: "1st",
			g:    CFGrammars[0],
			follows: []NonTerminal{
				NonTerminal("S"), // S
				NonTerminal("X"), // X
				NonTerminal("Y"), // Y
			},
			expectedFollows: []TerminalsAndEndmarker{
				{Terminals: set.New(eqTerminal), IncludesEndmarker: true},           // FOLLOW(S)
				{Terminals: set.New(eqTerminal, "0", "1"), IncludesEndmarker: true}, // FOLLOW(X)
				{Terminals: set.New(eqTerminal, "0"), IncludesEndmarker: true},      // FOLLOW(Y)
			},
		},
		{
			name: "2nd",
			g:    CFGrammars[1],
			follows: []NonTerminal{
				NonTerminal("S"), // S
			},
			expectedFollows: []TerminalsAndEndmarker{
				{Terminals: set.New(eqTerminal, "a", "b"), IncludesEndmarker: true}, // FOLLOW(S)
			},
		},
		{
			name: "3rd",
			g:    CFGrammars[2],
			follows: []NonTerminal{
				NonTerminal("S"), // S
				NonTerminal("A"), // A
				NonTerminal("B"), // B
			},
			expectedFollows: []TerminalsAndEndmarker{
				{Terminals: set.New(eqTerminal), IncludesEndmarker: true},            // FOLLOW(S)
				{Terminals: set.New(eqTerminal, "a", "b"), IncludesEndmarker: false}, // FOLLOW(A)
				{Terminals: set.New(eqTerminal, "a"), IncludesEndmarker: false},      // FOLLOW(B)
			},
		},
		{
			name: "4th",
			g:    CFGrammars[3],
			follows: []NonTerminal{
				NonTerminal("S"), // S
				NonTerminal("A"), // A
				NonTerminal("B"), // B
				NonTerminal("C"), // C
				NonTerminal("D"), // D
			},
			expectedFollows: []TerminalsAndEndmarker{
				{Terminals: set.New(eqTerminal), IncludesEndmarker: true}, // FOLLOW(S)
				{Terminals: set.New(eqTerminal), IncludesEndmarker: true}, // FOLLOW(A)
				{Terminals: set.New(eqTerminal), IncludesEndmarker: true}, // FOLLOW(B)
				{Terminals: set.New(eqTerminal), IncludesEndmarker: true}, // FOLLOW(C)
				{Terminals: set.New(eqTerminal), IncludesEndmarker: true}, // FOLLOW(D)
			},
		},
		{
			name: "5th",
			g:    CFGrammars[4],
			follows: []NonTerminal{
				NonTerminal("S"), // S
				NonTerminal("A"), // A
				NonTerminal("B"), // B
				NonTerminal("C"), // C
				NonTerminal("D"), // D
			},
			expectedFollows: []TerminalsAndEndmarker{
				{Terminals: set.New(eqTerminal), IncludesEndmarker: true},       // FOLLOW(S)
				{Terminals: set.New(eqTerminal, "b"), IncludesEndmarker: false}, // FOLLOW(A)
				{Terminals: set.New(eqTerminal), IncludesEndmarker: true},       // FOLLOW(B)
				{Terminals: set.New(eqTerminal), IncludesEndmarker: false},      // FOLLOW(C)
				{Terminals: set.New(eqTerminal), IncludesEndmarker: false},      // FOLLOW(D)
			},
		},
		{
			name: "6th",
			g:    CFGrammars[5],
			follows: []NonTerminal{
				NonTerminal("S"), // S
				NonTerminal("E"), // E
			},
			expectedFollows: []TerminalsAndEndmarker{
				{Terminals: set.New(eqTerminal), IncludesEndmarker: true},                          // FOLLOW(S)
				{Terminals: set.New(eqTerminal, "+", "-", "*", "/", ")"), IncludesEndmarker: true}, // FOLLOW(E)
			},
		},
		{
			name: "7th",
			g:    CFGrammars[6],
			follows: []NonTerminal{
				NonTerminal("S"), // S
				NonTerminal("E"), // E
				NonTerminal("T"), // T
				NonTerminal("F"), // F
			},
			expectedFollows: []TerminalsAndEndmarker{
				{Terminals: set.New(eqTerminal), IncludesEndmarker: true},                          // FOLLOW(S)
				{Terminals: set.New(eqTerminal, "+", "-", ")"), IncludesEndmarker: true},           // FOLLOW(E)
				{Terminals: set.New(eqTerminal, "+", "-", "*", "/", ")"), IncludesEndmarker: true}, // FOLLOW(T)
				{Terminals: set.New(eqTerminal, "+", "-", "*", "/", ")"), IncludesEndmarker: true}, // FOLLOW(F)
			},
		},
		{
			name: "8th",
			g:    CFGrammars[7],
			follows: []NonTerminal{
				NonTerminal("grammar"), // grammar
				NonTerminal("name"),    // name
				NonTerminal("decls"),   // decls
				NonTerminal("decl"),    // decl
				NonTerminal("token"),   // token
				NonTerminal("rule"),    // rule
				NonTerminal("lhs"),     // lhs
				NonTerminal("rhs"),     // rhs
				NonTerminal("nonterm"), // nonterm
				NonTerminal("term"),    // term
			},
			expectedFollows: []TerminalsAndEndmarker{
				{Terminals: set.New(eqTerminal), IncludesEndmarker: true},                                                                                 // FOLLOW(grammar)
				{Terminals: set.New(eqTerminal, "IDENT", "TOKEN"), IncludesEndmarker: true},                                                               // FOLLOW(name)
				{Terminals: set.New(eqTerminal, "IDENT", "TOKEN"), IncludesEndmarker: true},                                                               // FOLLOW(decls)
				{Terminals: set.New(eqTerminal, "IDENT", "TOKEN"), IncludesEndmarker: true},                                                               // FOLLOW(decl)
				{Terminals: set.New(eqTerminal, "IDENT", "TOKEN"), IncludesEndmarker: true},                                                               // FOLLOW(token)
				{Terminals: set.New(eqTerminal, "IDENT", "TOKEN"), IncludesEndmarker: true},                                                               // FOLLOW(rule)
				{Terminals: set.New(eqTerminal, "="), IncludesEndmarker: false},                                                                           // FOLLOW(lhs)
				{Terminals: set.New(eqTerminal, "|", "(", ")", "[", "]", "{", "}", "{{", "}}", "IDENT", "TOKEN", "STRING"), IncludesEndmarker: true},      // FOLLOW(rhs)
				{Terminals: set.New(eqTerminal, "=", "|", "(", ")", "[", "]", "{", "}", "{{", "}}", "IDENT", "TOKEN", "STRING"), IncludesEndmarker: true}, // FOLLOW(nonterm)
				{Terminals: set.New(eqTerminal, "|", "(", ")", "[", "]", "{", "}", "{{", "}}", "IDENT", "TOKEN", "STRING"), IncludesEndmarker: true},      // FOLLOW(term)
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NoError(t, tc.g.Verify())
			first := tc.g.ComputeFIRST()
			follow := tc.g.ComputeFOLLOW(first)

			for i, n := range tc.follows {
				t.Run(strconv.Itoa(i), func(t *testing.T) {
					f := follow(n)
					assert.True(t, tc.expectedFollows[i].Terminals.Equals(f.Terminals))
					assert.Equal(t, tc.expectedFollows[i].IncludesEndmarker, f.IncludesEndmarker)
				})
			}
		})
	}
}

func TestCFG_addNewNonTerminal(t *testing.T) {
	tests := []struct {
		name                string
		g                   CFG
		prefix              NonTerminal
		suffixes            []string
		expectedNonTerminal NonTerminal
	}{
		{
			name:                "OK",
			g:                   CFGrammars[0],
			prefix:              NonTerminal("S"),
			suffixes:            []string{"_new"},
			expectedNonTerminal: NonTerminal("S_new"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NoError(t, tc.g.Verify())
			nonTerm := tc.g.addNewNonTerminal(tc.prefix, tc.suffixes...)
			assert.Equal(t, tc.expectedNonTerminal, nonTerm)
		})
	}
}

func TestCFG_orderTerminals(t *testing.T) {
	tests := []struct {
		name              string
		g                 CFG
		expectedTerminals String[Terminal]
	}{
		{
			name:              "OK",
			g:                 CFGrammars[4],
			expectedTerminals: String[Terminal]{"a", "b", "c", "d"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NoError(t, tc.g.Verify())
			terms := tc.g.orderTerminals()
			assert.Equal(t, tc.expectedTerminals, terms)
		})
	}
}

func TestCFG_orderNonTerminals(t *testing.T) {
	tests := []struct {
		name                 string
		g                    CFG
		expectedVisited      String[NonTerminal]
		expectedUnvisited    String[NonTerminal]
		expectedNonTerminals String[NonTerminal]
	}{
		{
			name:                 "OK",
			g:                    CFGrammars[4],
			expectedVisited:      String[NonTerminal]{"S", "A", "B"},
			expectedUnvisited:    String[NonTerminal]{"C", "D"},
			expectedNonTerminals: String[NonTerminal]{"S", "A", "B", "C", "D"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NoError(t, tc.g.Verify())
			visited, unvisited, nonTerms := tc.g.orderNonTerminals()
			assert.Equal(t, tc.expectedVisited, visited)
			assert.Equal(t, tc.expectedUnvisited, unvisited)
			assert.Equal(t, tc.expectedNonTerminals, nonTerms)
		})
	}
}
