package grammar

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/moorara/algo/generic"
	"github.com/moorara/algo/set"
)

func getTestProductions() []*cfProductions {
	p0 := NewCFProductions().(*cfProductions)

	p1 := NewCFProductions().(*cfProductions)
	p1.Add(CFProduction{"S", String[Symbol]{Terminal("a"), NonTerminal("S"), Terminal("b"), NonTerminal("S")}}) // S → aSbS
	p1.Add(CFProduction{"S", String[Symbol]{Terminal("b"), NonTerminal("S"), Terminal("a"), NonTerminal("S")}}) // S → bSaS
	p1.Add(CFProduction{"S", ε})                                                                                // S → ε

	p2 := NewCFProductions().(*cfProductions)
	p2.Add(CFProduction{"S", String[Symbol]{NonTerminal("E")}})                                  // S → E
	p2.Add(CFProduction{"E", String[Symbol]{NonTerminal("E"), Terminal("+"), NonTerminal("T")}}) // E → E + T
	p2.Add(CFProduction{"E", String[Symbol]{NonTerminal("E"), Terminal("-"), NonTerminal("T")}}) // E → E - T
	p2.Add(CFProduction{"E", String[Symbol]{NonTerminal("T")}})                                  // E → T
	p2.Add(CFProduction{"T", String[Symbol]{NonTerminal("T"), Terminal("*"), NonTerminal("F")}}) // T → T * F
	p2.Add(CFProduction{"T", String[Symbol]{NonTerminal("T"), Terminal("/"), NonTerminal("F")}}) // T → T / F
	p2.Add(CFProduction{"T", String[Symbol]{NonTerminal("F")}})                                  // T → F
	p2.Add(CFProduction{"F", String[Symbol]{Terminal("("), NonTerminal("E"), Terminal(")")}})    // F → ( E )
	p2.Add(CFProduction{"F", String[Symbol]{Terminal("id")}})                                    // F → id

	p3 := NewCFProductions().(*cfProductions)
	p3.Add(CFProduction{"S", String[Symbol]{NonTerminal("E")}})                                  // S → E
	p3.Add(CFProduction{"E", String[Symbol]{NonTerminal("E"), Terminal("+"), NonTerminal("E")}}) // E → E + E
	p3.Add(CFProduction{"E", String[Symbol]{NonTerminal("E"), Terminal("-"), NonTerminal("E")}}) // E → E - E
	p3.Add(CFProduction{"E", String[Symbol]{NonTerminal("E"), Terminal("*"), NonTerminal("E")}}) // E → E * E
	p3.Add(CFProduction{"E", String[Symbol]{NonTerminal("E"), Terminal("/"), NonTerminal("E")}}) // E → E / E
	p3.Add(CFProduction{"E", String[Symbol]{Terminal("("), NonTerminal("E"), Terminal(")")}})    // E → ( E )
	p3.Add(CFProduction{"E", String[Symbol]{Terminal("-"), NonTerminal("E")}})                   // E → - E
	p3.Add(CFProduction{"E", String[Symbol]{Terminal("id")}})                                    // E → id

	return []*cfProductions{p0, p1, p2, p3}
}

func TestProduction(t *testing.T) {
	tests := []struct {
		name                    string
		p                       CFProduction
		expectedString          string
		expectedIsEmpty         bool
		expectedIsSingle        bool
		expectedIsLeftRecursive bool
		expectedIsCNFBinary     bool
		expectedIsCNFTerminal   bool
	}{
		{
			name:                    "1st",
			p:                       CFProduction{"S", ε},
			expectedString:          `S → ε`,
			expectedIsEmpty:         true,
			expectedIsSingle:        false,
			expectedIsLeftRecursive: false,
			expectedIsCNFBinary:     false,
			expectedIsCNFTerminal:   false,
		},
		{
			name:                    "2nd",
			p:                       CFProduction{"A", String[Symbol]{Terminal("a")}},
			expectedString:          `A → "a"`,
			expectedIsEmpty:         false,
			expectedIsSingle:        false,
			expectedIsLeftRecursive: false,
			expectedIsCNFBinary:     false,
			expectedIsCNFTerminal:   true,
		},
		{
			name:                    "3rd",
			p:                       CFProduction{"A", String[Symbol]{NonTerminal("A")}},
			expectedString:          `A → A`,
			expectedIsEmpty:         false,
			expectedIsSingle:        true,
			expectedIsLeftRecursive: true,
			expectedIsCNFBinary:     false,
			expectedIsCNFTerminal:   false,
		},
		{
			name:                    "4th",
			p:                       CFProduction{"A", String[Symbol]{NonTerminal("B")}},
			expectedString:          `A → B`,
			expectedIsEmpty:         false,
			expectedIsSingle:        true,
			expectedIsLeftRecursive: false,
			expectedIsCNFBinary:     false,
			expectedIsCNFTerminal:   false,
		},
		{
			name:                    "5th",
			p:                       CFProduction{"A", String[Symbol]{NonTerminal("A"), Terminal("a")}},
			expectedString:          `A → A "a"`,
			expectedIsEmpty:         false,
			expectedIsSingle:        false,
			expectedIsLeftRecursive: true,
			expectedIsCNFBinary:     false,
			expectedIsCNFTerminal:   false,
		},
		{
			name:                    "6th",
			p:                       CFProduction{"A", String[Symbol]{NonTerminal("A"), NonTerminal("B")}},
			expectedString:          `A → A B`,
			expectedIsEmpty:         false,
			expectedIsSingle:        false,
			expectedIsLeftRecursive: true,
			expectedIsCNFBinary:     true,
			expectedIsCNFTerminal:   false,
		},
		{
			name:                    "7th",
			p:                       CFProduction{"stmt", String[Symbol]{Terminal("if"), NonTerminal("expr"), Terminal("then"), NonTerminal("stmt")}},
			expectedString:          `stmt → "if" expr "then" stmt`,
			expectedIsEmpty:         false,
			expectedIsSingle:        false,
			expectedIsLeftRecursive: false,
		},
	}

	notEqual := CFProduction{"😐", String[Symbol]{Terminal("🙂"), NonTerminal("🙃")}}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedString, tc.p.String())
			assert.True(t, tc.p.Equals(tc.p))
			assert.False(t, tc.p.Equals(notEqual))
			assert.Equal(t, tc.expectedIsEmpty, tc.p.IsEmpty())
			assert.Equal(t, tc.expectedIsSingle, tc.p.IsSingle())
			assert.Equal(t, tc.expectedIsLeftRecursive, tc.p.IsLeftRecursive())

			isBinary, isTerminal := tc.p.IsCNF()
			assert.Equal(t, tc.expectedIsCNFBinary, isBinary)
			assert.Equal(t, tc.expectedIsCNFTerminal, isTerminal)
		})
	}
}

func TestNewCFProductions(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		p := NewCFProductions()
		assert.NotNil(t, p)
	})
}

func TestProductions_String(t *testing.T) {
	p := getTestProductions()

	tests := []struct {
		name               string
		p                  *cfProductions
		expectedSubstrings []string
	}{
		{
			name: "1st",
			p:    p[1],
			expectedSubstrings: []string{
				`S → "a" S "b" S | "b" S "a" S | ε`,
			},
		},
		{
			name: "2nd",
			p:    p[2],
			expectedSubstrings: []string{
				`S → E`,
				`E → E "+" T | E "-" T | T`,
				`T → T "*" F | T "/" F | F`,
				`F → "(" E ")" | "id"`,
			},
		},
		{
			name: "3rd",
			p:    p[3],
			expectedSubstrings: []string{
				`S → E`,
				`E → E "*" E | E "+" E | E "-" E | E "/" E | "(" E ")" | "-" E | "id"`,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s := tc.p.String()

			for _, expectedSubstring := range tc.expectedSubstrings {
				assert.Contains(t, s, expectedSubstring)
			}
		})
	}
}

func TestProductions_Clone(t *testing.T) {
	p := getTestProductions()

	tests := []struct {
		name string
		p    *cfProductions
	}{
		{
			name: "1st",
			p:    p[1],
		},
		{
			name: "2nd",
			p:    p[2],
		},
		{
			name: "3rd",
			p:    p[3],
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			newP := tc.p.Clone()
			assert.False(t, newP == tc.p)
			assert.True(t, newP.Equals(tc.p))
		})
	}
}

func TestProductions_Equals(t *testing.T) {
	p := getTestProductions()

	tests := []struct {
		name           string
		p              *cfProductions
		rhs            CFProductions
		expectedEquals bool
	}{
		{
			name:           "Equal",
			p:              p[2],
			rhs:            p[2],
			expectedEquals: true,
		},
		{
			name:           "NotEqual",
			p:              p[2],
			rhs:            p[3],
			expectedEquals: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedEquals, tc.p.Equals(tc.rhs))
		})
	}
}

func TestProductions_Add(t *testing.T) {
	p := getTestProductions()

	tests := []struct {
		name                string
		p                   *cfProductions
		ps                  []CFProduction
		expectedProductions *cfProductions
	}{
		{
			name: "1st",
			p:    p[1],
			ps: []CFProduction{
				{"S", String[Symbol]{Terminal("a"), NonTerminal("S"), Terminal("b"), NonTerminal("S")}}, // S → aSbS
				{"S", String[Symbol]{Terminal("b"), NonTerminal("S"), Terminal("a"), NonTerminal("S")}}, // S → bSaS
				{"S", ε}, // S → ε
			},
			expectedProductions: p[1],
		},
		{
			name: "2nd",
			p:    p[2],
			ps: []CFProduction{
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
			expectedProductions: p[2],
		},
		{
			name: "3rd",
			p:    p[3],
			ps: []CFProduction{
				{"S", String[Symbol]{NonTerminal("E")}},                                  // S → E
				{"E", String[Symbol]{NonTerminal("E"), Terminal("+"), NonTerminal("E")}}, // E → E + E
				{"E", String[Symbol]{NonTerminal("E"), Terminal("-"), NonTerminal("E")}}, // E → E - E
				{"E", String[Symbol]{NonTerminal("E"), Terminal("*"), NonTerminal("E")}}, // E → E * E
				{"E", String[Symbol]{NonTerminal("E"), Terminal("/"), NonTerminal("E")}}, // E → E / E
				{"E", String[Symbol]{Terminal("("), NonTerminal("E"), Terminal(")")}},    // E → ( E )
				{"E", String[Symbol]{Terminal("-"), NonTerminal("E")}},                   // E → - E
				{"E", String[Symbol]{Terminal("id")}},                                    // E → id
			},
			expectedProductions: p[3],
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.p.Add(tc.ps...)
			assert.True(t, tc.p.Equals(tc.expectedProductions))
		})
	}
}

func TestProductions_Remove(t *testing.T) {
	p := getTestProductions()

	tests := []struct {
		name                string
		p                   *cfProductions
		ps                  []CFProduction
		expectedProductions *cfProductions
	}{
		{
			name: "1st",
			p:    p[1],
			ps: []CFProduction{
				{"S", String[Symbol]{Terminal("a"), NonTerminal("S"), Terminal("b"), NonTerminal("S")}}, // S → aSbS
				{"S", String[Symbol]{Terminal("b"), NonTerminal("S"), Terminal("a"), NonTerminal("S")}}, // S → bSaS
				{"S", ε}, // S → ε
			},
			expectedProductions: p[0],
		},
		{
			name: "2nd",
			p:    p[2],
			ps: []CFProduction{
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
			expectedProductions: p[0],
		},
		{
			name: "3rd",
			p:    p[3],
			ps: []CFProduction{
				{"S", String[Symbol]{NonTerminal("E")}},                                  // S → E
				{"E", String[Symbol]{NonTerminal("E"), Terminal("+"), NonTerminal("E")}}, // E → E + E
				{"E", String[Symbol]{NonTerminal("E"), Terminal("-"), NonTerminal("E")}}, // E → E - E
				{"E", String[Symbol]{NonTerminal("E"), Terminal("*"), NonTerminal("E")}}, // E → E * E
				{"E", String[Symbol]{NonTerminal("E"), Terminal("/"), NonTerminal("E")}}, // E → E / E
				{"E", String[Symbol]{Terminal("("), NonTerminal("E"), Terminal(")")}},    // E → ( E )
				{"E", String[Symbol]{Terminal("-"), NonTerminal("E")}},                   // E → - E
				{"E", String[Symbol]{Terminal("id")}},                                    // E → id
			},
			expectedProductions: p[0],
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.p.Remove(tc.ps...)
			assert.True(t, tc.p.Equals(tc.expectedProductions))
		})
	}
}

func TestProductions_RemoveAll(t *testing.T) {
	p := getTestProductions()

	tests := []struct {
		name                string
		p                   *cfProductions
		heads               []NonTerminal
		expectedProductions *cfProductions
	}{
		{
			name:                "1st",
			p:                   p[1],
			heads:               []NonTerminal{"S"},
			expectedProductions: p[0],
		},
		{
			name:                "2nd",
			p:                   p[2],
			heads:               []NonTerminal{"S", "E", "T", "F"},
			expectedProductions: p[0],
		},
		{
			name:                "3rd",
			p:                   p[3],
			heads:               []NonTerminal{"S", "E"},
			expectedProductions: p[0],
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.p.RemoveAll(tc.heads...)
			assert.True(t, tc.p.Equals(tc.expectedProductions))
		})
	}
}

func TestProductions_Get(t *testing.T) {
	p := getTestProductions()

	s1 := set.New[CFProduction](eqCFProduction,
		CFProduction{"S", String[Symbol]{Terminal("a"), NonTerminal("S"), Terminal("b"), NonTerminal("S")}}, // S → aSbS
		CFProduction{"S", String[Symbol]{Terminal("b"), NonTerminal("S"), Terminal("a"), NonTerminal("S")}}, // S → bSaS
		CFProduction{"S", ε}, // S → ε
	)

	s2 := set.New[CFProduction](eqCFProduction,
		CFProduction{"T", String[Symbol]{NonTerminal("T"), Terminal("*"), NonTerminal("F")}}, // T → T * F
		CFProduction{"T", String[Symbol]{NonTerminal("T"), Terminal("/"), NonTerminal("F")}}, // T → T / F
		CFProduction{"T", String[Symbol]{NonTerminal("F")}},                                  // T → F
	)

	s3 := set.New[CFProduction](eqCFProduction,
		CFProduction{"E", String[Symbol]{NonTerminal("E"), Terminal("+"), NonTerminal("E")}}, // E → E + E
		CFProduction{"E", String[Symbol]{NonTerminal("E"), Terminal("-"), NonTerminal("E")}}, // E → E - E
		CFProduction{"E", String[Symbol]{NonTerminal("E"), Terminal("*"), NonTerminal("E")}}, // E → E * E
		CFProduction{"E", String[Symbol]{NonTerminal("E"), Terminal("/"), NonTerminal("E")}}, // E → E / E
		CFProduction{"E", String[Symbol]{Terminal("("), NonTerminal("E"), Terminal(")")}},    // E → ( E )
		CFProduction{"E", String[Symbol]{Terminal("-"), NonTerminal("E")}},                   // E → - E
		CFProduction{"E", String[Symbol]{Terminal("id")}},                                    // E → id
	)

	tests := []struct {
		name                string
		p                   *cfProductions
		head                NonTerminal
		expectedProductions set.Set[CFProduction]
	}{
		{
			name:                "Nil",
			p:                   p[0],
			head:                NonTerminal("E"),
			expectedProductions: nil,
		},
		{
			name:                "1st",
			p:                   p[1],
			head:                NonTerminal("S"),
			expectedProductions: s1,
		},
		{
			name:                "2nd",
			p:                   p[2],
			head:                NonTerminal("T"),
			expectedProductions: s2,
		},
		{
			name:                "3rd",
			p:                   p[3],
			head:                NonTerminal("E"),
			expectedProductions: s3,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			prods := tc.p.Get(tc.head)

			if tc.expectedProductions == nil {
				assert.Nil(t, prods)
			} else {
				assert.True(t, prods.Equals(tc.expectedProductions))
			}
		})
	}
}

func TestProductions_Order(t *testing.T) {
	p := getTestProductions()

	tests := []struct {
		name                string
		p                   *cfProductions
		head                NonTerminal
		expectedProductions []CFProduction
	}{
		{
			name: "1st",
			p:    p[1],
			head: NonTerminal("S"),
			expectedProductions: []CFProduction{
				{"S", String[Symbol]{Terminal("a"), NonTerminal("S"), Terminal("b"), NonTerminal("S")}}, // S → aSbS
				{"S", String[Symbol]{Terminal("b"), NonTerminal("S"), Terminal("a"), NonTerminal("S")}}, // S → bSaS
				{"S", ε}, // S → ε
			},
		},
		{
			name: "2nd",
			p:    p[2],
			head: NonTerminal("T"),
			expectedProductions: []CFProduction{
				{"T", String[Symbol]{NonTerminal("T"), Terminal("*"), NonTerminal("F")}}, // T → T * F
				{"T", String[Symbol]{NonTerminal("T"), Terminal("/"), NonTerminal("F")}}, // T → T / F
				{"T", String[Symbol]{NonTerminal("F")}},                                  // T → F
			},
		},
		{
			name: "3rd",
			p:    p[3],
			head: NonTerminal("E"),
			expectedProductions: []CFProduction{
				{"E", String[Symbol]{NonTerminal("E"), Terminal("*"), NonTerminal("E")}}, // E → E * E
				{"E", String[Symbol]{NonTerminal("E"), Terminal("+"), NonTerminal("E")}}, // E → E + E
				{"E", String[Symbol]{NonTerminal("E"), Terminal("-"), NonTerminal("E")}}, // E → E - E
				{"E", String[Symbol]{NonTerminal("E"), Terminal("/"), NonTerminal("E")}}, // E → E / E
				{"E", String[Symbol]{Terminal("("), NonTerminal("E"), Terminal(")")}},    // E → ( E )
				{"E", String[Symbol]{Terminal("-"), NonTerminal("E")}},                   // E → - E
				{"E", String[Symbol]{Terminal("id")}},                                    // E → id
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			prods := tc.p.Order(tc.head)
			assert.Equal(t, tc.expectedProductions, prods)
		})
	}
}

func TestProductions_All(t *testing.T) {
	p := getTestProductions()

	tests := []struct {
		name                string
		p                   *cfProductions
		expectedProductions []CFProduction
	}{
		{
			name: "1st",
			p:    p[1],
			expectedProductions: []CFProduction{
				{"S", String[Symbol]{Terminal("a"), NonTerminal("S"), Terminal("b"), NonTerminal("S")}}, // S → aSbS
				{"S", String[Symbol]{Terminal("b"), NonTerminal("S"), Terminal("a"), NonTerminal("S")}}, // S → bSaS
				{"S", ε}, // S → ε
			},
		},
		{
			name: "2nd",
			p:    p[2],
			expectedProductions: []CFProduction{
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
		},
		{
			name: "3rd",
			p:    p[3],
			expectedProductions: []CFProduction{
				{"S", String[Symbol]{NonTerminal("E")}},                                  // S → E
				{"E", String[Symbol]{NonTerminal("E"), Terminal("+"), NonTerminal("E")}}, // E → E + E
				{"E", String[Symbol]{NonTerminal("E"), Terminal("-"), NonTerminal("E")}}, // E → E - E
				{"E", String[Symbol]{NonTerminal("E"), Terminal("*"), NonTerminal("E")}}, // E → E * E
				{"E", String[Symbol]{NonTerminal("E"), Terminal("/"), NonTerminal("E")}}, // E → E / E
				{"E", String[Symbol]{Terminal("("), NonTerminal("E"), Terminal(")")}},    // E → ( E )
				{"E", String[Symbol]{Terminal("-"), NonTerminal("E")}},                   // E → - E
				{"E", String[Symbol]{Terminal("id")}},                                    // E → id
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			for p := range tc.p.All() {
				assert.Contains(t, tc.expectedProductions, p)
			}
		})
	}
}

func TestProductions_AllByHead(t *testing.T) {
	p := getTestProductions()

	tests := []struct {
		name                string
		p                   *cfProductions
		expectedProductions []CFProduction
	}{
		{
			name: "1st",
			p:    p[1],
			expectedProductions: []CFProduction{
				{"S", String[Symbol]{Terminal("a"), NonTerminal("S"), Terminal("b"), NonTerminal("S")}}, // S → aSbS
				{"S", String[Symbol]{Terminal("b"), NonTerminal("S"), Terminal("a"), NonTerminal("S")}}, // S → bSaS
				{"S", ε}, // S → ε
			},
		},
		{
			name: "2nd",
			p:    p[2],
			expectedProductions: []CFProduction{
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
		},
		{
			name: "3rd",
			p:    p[3],
			expectedProductions: []CFProduction{
				{"S", String[Symbol]{NonTerminal("E")}},                                  // S → E
				{"E", String[Symbol]{NonTerminal("E"), Terminal("+"), NonTerminal("E")}}, // E → E + E
				{"E", String[Symbol]{NonTerminal("E"), Terminal("-"), NonTerminal("E")}}, // E → E - E
				{"E", String[Symbol]{NonTerminal("E"), Terminal("*"), NonTerminal("E")}}, // E → E * E
				{"E", String[Symbol]{NonTerminal("E"), Terminal("/"), NonTerminal("E")}}, // E → E / E
				{"E", String[Symbol]{Terminal("("), NonTerminal("E"), Terminal(")")}},    // E → ( E )
				{"E", String[Symbol]{Terminal("-"), NonTerminal("E")}},                   // E → - E
				{"E", String[Symbol]{Terminal("id")}},                                    // E → id
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			for head, list := range tc.p.AllByHead() {
				for p := range list.All() {
					assert.True(t, p.Head.Equals(head))
					assert.Contains(t, tc.expectedProductions, p)
				}
			}
		})
	}
}

func TestProductions_AnyMatch(t *testing.T) {
	p := getTestProductions()

	tests := []struct {
		name             string
		p                *cfProductions
		pred             Predicate1[CFProduction]
		expectedAnyMatch bool
	}{
		{
			name:             "OK",
			p:                p[2],
			pred:             func(p CFProduction) bool { return p.IsSingle() },
			expectedAnyMatch: true,
		},
		{
			name:             "NotOK",
			p:                p[2],
			pred:             func(p CFProduction) bool { return p.IsEmpty() },
			expectedAnyMatch: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			anyMatch := tc.p.AnyMatch(tc.pred)
			assert.Equal(t, tc.expectedAnyMatch, anyMatch)
		})
	}
}

func TestProductions_AllMatch(t *testing.T) {
	p := getTestProductions()

	tests := []struct {
		name             string
		p                *cfProductions
		pred             Predicate1[CFProduction]
		expectedAllMatch bool
	}{
		{
			name:             "OK",
			p:                p[2],
			pred:             func(p CFProduction) bool { return !p.IsEmpty() },
			expectedAllMatch: true,
		},
		{
			name:             "NotOK",
			p:                p[2],
			pred:             func(p CFProduction) bool { return !p.IsSingle() },
			expectedAllMatch: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			allMatch := tc.p.AllMatch(tc.pred)
			assert.Equal(t, tc.expectedAllMatch, allMatch)
		})
	}
}

func TestProductions_SelectMatch(t *testing.T) {
	p := getTestProductions()

	q1 := NewCFProductions().(*cfProductions)
	q1.Add(CFProduction{"S", String[Symbol]{NonTerminal("E")}}) // S → E
	q1.Add(CFProduction{"E", String[Symbol]{NonTerminal("T")}}) // E → T
	q1.Add(CFProduction{"T", String[Symbol]{NonTerminal("F")}}) // T → F

	tests := []struct {
		name                string
		p                   *cfProductions
		pred                Predicate1[CFProduction]
		expectedSelectMatch *cfProductions
	}{
		{
			name:                "OK",
			p:                   p[2],
			pred:                func(p CFProduction) bool { return p.IsSingle() },
			expectedSelectMatch: q1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			selectMatch := tc.p.SelectMatch(tc.pred)
			assert.True(t, selectMatch.Equals(tc.expectedSelectMatch))
		})
	}
}
