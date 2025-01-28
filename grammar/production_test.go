package grammar

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/generic"
	"github.com/moorara/algo/set"
)

func getTestProductions() []*Productions {
	p0 := NewProductions()

	p1 := NewProductions()
	p1.Add(&Production{"S", String[Symbol]{Terminal("a"), NonTerminal("S"), Terminal("b"), NonTerminal("S")}}) // S → aSbS
	p1.Add(&Production{"S", String[Symbol]{Terminal("b"), NonTerminal("S"), Terminal("a"), NonTerminal("S")}}) // S → bSaS
	p1.Add(&Production{"S", E})                                                                                // S → ε

	p2 := NewProductions()
	p2.Add(&Production{"S", String[Symbol]{NonTerminal("E")}})                                  // S → E
	p2.Add(&Production{"E", String[Symbol]{NonTerminal("E"), Terminal("+"), NonTerminal("T")}}) // E → E + T
	p2.Add(&Production{"E", String[Symbol]{NonTerminal("E"), Terminal("-"), NonTerminal("T")}}) // E → E - T
	p2.Add(&Production{"E", String[Symbol]{NonTerminal("T")}})                                  // E → T
	p2.Add(&Production{"T", String[Symbol]{NonTerminal("T"), Terminal("*"), NonTerminal("F")}}) // T → T * F
	p2.Add(&Production{"T", String[Symbol]{NonTerminal("T"), Terminal("/"), NonTerminal("F")}}) // T → T / F
	p2.Add(&Production{"T", String[Symbol]{NonTerminal("F")}})                                  // T → F
	p2.Add(&Production{"F", String[Symbol]{Terminal("("), NonTerminal("E"), Terminal(")")}})    // F → ( E )
	p2.Add(&Production{"F", String[Symbol]{Terminal("id")}})                                    // F → id

	p3 := NewProductions()
	p3.Add(&Production{"S", String[Symbol]{NonTerminal("E")}})                                  // S → E
	p3.Add(&Production{"E", String[Symbol]{NonTerminal("E"), Terminal("+"), NonTerminal("E")}}) // E → E + E
	p3.Add(&Production{"E", String[Symbol]{NonTerminal("E"), Terminal("-"), NonTerminal("E")}}) // E → E - E
	p3.Add(&Production{"E", String[Symbol]{NonTerminal("E"), Terminal("*"), NonTerminal("E")}}) // E → E * E
	p3.Add(&Production{"E", String[Symbol]{NonTerminal("E"), Terminal("/"), NonTerminal("E")}}) // E → E / E
	p3.Add(&Production{"E", String[Symbol]{Terminal("("), NonTerminal("E"), Terminal(")")}})    // E → ( E )
	p3.Add(&Production{"E", String[Symbol]{Terminal("-"), NonTerminal("E")}})                   // E → - E
	p3.Add(&Production{"E", String[Symbol]{Terminal("id")}})                                    // E → id

	return []*Productions{p0, p1, p2, p3}
}

func TestProduction(t *testing.T) {
	tests := []struct {
		name                    string
		p                       *Production
		expectedString          string
		expectedIsEmpty         bool
		expectedIsSingle        bool
		expectedIsLeftRecursive bool
		expectedIsCNFBinary     bool
		expectedIsCNFTerminal   bool
	}{
		{
			name:                    "1st",
			p:                       &Production{"S", E},
			expectedString:          `S → ε`,
			expectedIsEmpty:         true,
			expectedIsSingle:        false,
			expectedIsLeftRecursive: false,
			expectedIsCNFBinary:     false,
			expectedIsCNFTerminal:   false,
		},
		{
			name:                    "2nd",
			p:                       &Production{"A", String[Symbol]{Terminal("a")}},
			expectedString:          `A → "a"`,
			expectedIsEmpty:         false,
			expectedIsSingle:        false,
			expectedIsLeftRecursive: false,
			expectedIsCNFBinary:     false,
			expectedIsCNFTerminal:   true,
		},
		{
			name:                    "3rd",
			p:                       &Production{"A", String[Symbol]{NonTerminal("A")}},
			expectedString:          `A → A`,
			expectedIsEmpty:         false,
			expectedIsSingle:        true,
			expectedIsLeftRecursive: true,
			expectedIsCNFBinary:     false,
			expectedIsCNFTerminal:   false,
		},
		{
			name:                    "4th",
			p:                       &Production{"A", String[Symbol]{NonTerminal("B")}},
			expectedString:          `A → B`,
			expectedIsEmpty:         false,
			expectedIsSingle:        true,
			expectedIsLeftRecursive: false,
			expectedIsCNFBinary:     false,
			expectedIsCNFTerminal:   false,
		},
		{
			name:                    "5th",
			p:                       &Production{"A", String[Symbol]{NonTerminal("A"), Terminal("a")}},
			expectedString:          `A → A "a"`,
			expectedIsEmpty:         false,
			expectedIsSingle:        false,
			expectedIsLeftRecursive: true,
			expectedIsCNFBinary:     false,
			expectedIsCNFTerminal:   false,
		},
		{
			name:                    "6th",
			p:                       &Production{"A", String[Symbol]{NonTerminal("A"), NonTerminal("B")}},
			expectedString:          `A → A B`,
			expectedIsEmpty:         false,
			expectedIsSingle:        false,
			expectedIsLeftRecursive: true,
			expectedIsCNFBinary:     true,
			expectedIsCNFTerminal:   false,
		},
		{
			name:                    "7th",
			p:                       &Production{"stmt", String[Symbol]{Terminal("if"), NonTerminal("expr"), Terminal("then"), NonTerminal("stmt")}},
			expectedString:          `stmt → "if" expr "then" stmt`,
			expectedIsEmpty:         false,
			expectedIsSingle:        false,
			expectedIsLeftRecursive: false,
		},
	}

	notEqual := &Production{"😐", String[Symbol]{Terminal("🙂"), NonTerminal("🙃")}}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedString, tc.p.String())
			assert.True(t, tc.p.Equal(tc.p))
			assert.False(t, tc.p.Equal(notEqual))
			assert.Equal(t, tc.expectedIsEmpty, tc.p.IsEmpty())
			assert.Equal(t, tc.expectedIsSingle, tc.p.IsSingle())
			assert.Equal(t, tc.expectedIsLeftRecursive, tc.p.IsLeftRecursive())

			isBinary, isTerminal := tc.p.IsCNF()
			assert.Equal(t, tc.expectedIsCNFBinary, isBinary)
			assert.Equal(t, tc.expectedIsCNFTerminal, isTerminal)
		})
	}
}

func TestNewProductions(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		p := NewProductions()
		assert.NotNil(t, p)
	})
}

func TestProductions_String(t *testing.T) {
	p := getTestProductions()

	tests := []struct {
		name               string
		p                  *Productions
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
		p    *Productions
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
			assert.True(t, newP.Equal(tc.p))
		})
	}
}

func TestProductions_Equal(t *testing.T) {
	p := getTestProductions()

	tests := []struct {
		name          string
		p             *Productions
		rhs           *Productions
		expectedEqual bool
	}{
		{
			name:          "Equal",
			p:             p[2],
			rhs:           p[2],
			expectedEqual: true,
		},
		{
			name:          "NotEqual",
			p:             p[2],
			rhs:           p[3],
			expectedEqual: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedEqual, tc.p.Equal(tc.rhs))
		})
	}
}

func TestProductions_Add(t *testing.T) {
	p := getTestProductions()

	tests := []struct {
		name                string
		p                   *Productions
		ps                  []*Production
		expectedProductions *Productions
	}{
		{
			name: "1st",
			p:    p[1],
			ps: []*Production{
				{"S", String[Symbol]{Terminal("a"), NonTerminal("S"), Terminal("b"), NonTerminal("S")}}, // S → aSbS
				{"S", String[Symbol]{Terminal("b"), NonTerminal("S"), Terminal("a"), NonTerminal("S")}}, // S → bSaS
				{"S", E}, // S → ε
			},
			expectedProductions: p[1],
		},
		{
			name: "2nd",
			p:    p[2],
			ps: []*Production{
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
			ps: []*Production{
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
			assert.True(t, tc.p.Equal(tc.expectedProductions))
		})
	}
}

func TestProductions_Remove(t *testing.T) {
	p := getTestProductions()

	tests := []struct {
		name                string
		p                   *Productions
		ps                  []*Production
		expectedProductions *Productions
	}{
		{
			name: "1st",
			p:    p[1],
			ps: []*Production{
				{"S", String[Symbol]{Terminal("a"), NonTerminal("S"), Terminal("b"), NonTerminal("S")}}, // S → aSbS
				{"S", String[Symbol]{Terminal("b"), NonTerminal("S"), Terminal("a"), NonTerminal("S")}}, // S → bSaS
				{"S", E}, // S → ε
			},
			expectedProductions: p[0],
		},
		{
			name: "2nd",
			p:    p[2],
			ps: []*Production{
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
			ps: []*Production{
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
			assert.True(t, tc.p.Equal(tc.expectedProductions))
		})
	}
}

func TestProductions_RemoveAll(t *testing.T) {
	p := getTestProductions()

	tests := []struct {
		name                string
		p                   *Productions
		heads               []NonTerminal
		expectedProductions *Productions
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
			assert.True(t, tc.p.Equal(tc.expectedProductions))
		})
	}
}

func TestProductions_Get(t *testing.T) {
	p := getTestProductions()

	s1 := set.New(EqProduction,
		&Production{"S", String[Symbol]{Terminal("a"), NonTerminal("S"), Terminal("b"), NonTerminal("S")}}, // S → aSbS
		&Production{"S", String[Symbol]{Terminal("b"), NonTerminal("S"), Terminal("a"), NonTerminal("S")}}, // S → bSaS
		&Production{"S", E}, // S → ε
	)

	s2 := set.New(EqProduction,
		&Production{"T", String[Symbol]{NonTerminal("T"), Terminal("*"), NonTerminal("F")}}, // T → T * F
		&Production{"T", String[Symbol]{NonTerminal("T"), Terminal("/"), NonTerminal("F")}}, // T → T / F
		&Production{"T", String[Symbol]{NonTerminal("F")}},                                  // T → F
	)

	s3 := set.New(EqProduction,
		&Production{"E", String[Symbol]{NonTerminal("E"), Terminal("+"), NonTerminal("E")}}, // E → E + E
		&Production{"E", String[Symbol]{NonTerminal("E"), Terminal("-"), NonTerminal("E")}}, // E → E - E
		&Production{"E", String[Symbol]{NonTerminal("E"), Terminal("*"), NonTerminal("E")}}, // E → E * E
		&Production{"E", String[Symbol]{NonTerminal("E"), Terminal("/"), NonTerminal("E")}}, // E → E / E
		&Production{"E", String[Symbol]{Terminal("("), NonTerminal("E"), Terminal(")")}},    // E → ( E )
		&Production{"E", String[Symbol]{Terminal("-"), NonTerminal("E")}},                   // E → - E
		&Production{"E", String[Symbol]{Terminal("id")}},                                    // E → id
	)

	tests := []struct {
		name                string
		p                   *Productions
		head                NonTerminal
		expectedProductions set.Set[*Production]
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
				assert.True(t, prods.Equal(tc.expectedProductions))
			}
		})
	}
}

func TestProductions_All(t *testing.T) {
	p := getTestProductions()

	tests := []struct {
		name                string
		p                   *Productions
		expectedProductions []*Production
	}{
		{
			name: "1st",
			p:    p[1],
			expectedProductions: []*Production{
				{"S", String[Symbol]{Terminal("a"), NonTerminal("S"), Terminal("b"), NonTerminal("S")}}, // S → aSbS
				{"S", String[Symbol]{Terminal("b"), NonTerminal("S"), Terminal("a"), NonTerminal("S")}}, // S → bSaS
				{"S", E}, // S → ε
			},
		},
		{
			name: "2nd",
			p:    p[2],
			expectedProductions: []*Production{
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
			expectedProductions: []*Production{
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
		p                   *Productions
		expectedProductions []*Production
	}{
		{
			name: "1st",
			p:    p[1],
			expectedProductions: []*Production{
				{"S", String[Symbol]{Terminal("a"), NonTerminal("S"), Terminal("b"), NonTerminal("S")}}, // S → aSbS
				{"S", String[Symbol]{Terminal("b"), NonTerminal("S"), Terminal("a"), NonTerminal("S")}}, // S → bSaS
				{"S", E}, // S → ε
			},
		},
		{
			name: "2nd",
			p:    p[2],
			expectedProductions: []*Production{
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
			expectedProductions: []*Production{
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
					assert.True(t, p.Head.Equal(head))
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
		p                *Productions
		pred             generic.Predicate1[*Production]
		expectedAnyMatch bool
	}{
		{
			name:             "OK",
			p:                p[2],
			pred:             func(p *Production) bool { return p.IsSingle() },
			expectedAnyMatch: true,
		},
		{
			name:             "NotOK",
			p:                p[2],
			pred:             func(p *Production) bool { return p.IsEmpty() },
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
		p                *Productions
		pred             generic.Predicate1[*Production]
		expectedAllMatch bool
	}{
		{
			name:             "OK",
			p:                p[2],
			pred:             func(p *Production) bool { return !p.IsEmpty() },
			expectedAllMatch: true,
		},
		{
			name:             "NotOK",
			p:                p[2],
			pred:             func(p *Production) bool { return !p.IsSingle() },
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

	q1 := NewProductions()
	q1.Add(&Production{"S", String[Symbol]{NonTerminal("E")}}) // S → E
	q1.Add(&Production{"E", String[Symbol]{NonTerminal("T")}}) // E → T
	q1.Add(&Production{"T", String[Symbol]{NonTerminal("F")}}) // T → F

	tests := []struct {
		name                string
		p                   *Productions
		pred                generic.Predicate1[*Production]
		expectedSelectMatch *Productions
	}{
		{
			name:                "OK",
			p:                   p[2],
			pred:                func(p *Production) bool { return p.IsSingle() },
			expectedSelectMatch: q1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			selectMatch := tc.p.SelectMatch(tc.pred)
			assert.True(t, selectMatch.Equal(tc.expectedSelectMatch))
		})
	}
}

func TestOrderProductionSet(t *testing.T) {
	p := getTestProductions()

	tests := []struct {
		name                string
		set                 set.Set[*Production]
		expectedProductions []*Production
	}{
		{
			name: "1st",
			set:  p[1].Get("S"),
			expectedProductions: []*Production{
				{"S", String[Symbol]{Terminal("a"), NonTerminal("S"), Terminal("b"), NonTerminal("S")}}, // S → aSbS
				{"S", String[Symbol]{Terminal("b"), NonTerminal("S"), Terminal("a"), NonTerminal("S")}}, // S → bSaS
				{"S", E}, // S → ε
			},
		},
		{
			name: "2nd",
			set:  p[2].Get("T"),
			expectedProductions: []*Production{
				{"T", String[Symbol]{NonTerminal("T"), Terminal("*"), NonTerminal("F")}}, // T → T * F
				{"T", String[Symbol]{NonTerminal("T"), Terminal("/"), NonTerminal("F")}}, // T → T / F
				{"T", String[Symbol]{NonTerminal("F")}},                                  // T → F
			},
		},
		{
			name: "3rd",
			set:  p[3].Get("E"),
			expectedProductions: []*Production{
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
			prods := OrderProductionSet(tc.set)
			assert.Equal(t, tc.expectedProductions, prods)
		})
	}
}
