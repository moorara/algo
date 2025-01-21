package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/generic"
	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/lexer"
)

func getTestInternalNodes() []*InternalNode {
	n0 := &InternalNode{
		nonTerminal: "E",
		production: &grammar.Production{
			Head: "E",
			Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("E"), grammar.Terminal("+"), grammar.NonTerminal("E")},
		},
		children: []Node{
			&InternalNode{
				nonTerminal: "E",
				production: &grammar.Production{
					Head: "E",
					Body: grammar.String[grammar.Symbol]{grammar.Terminal("id")},
				},
				children: []Node{
					&LeafNode{
						terminal: "id",
						lexeme:   "fee",
						pos: &lexer.Position{
							Filename: "test",
							Offset:   2,
							Line:     1,
							Column:   3,
						},
					},
				},
			},
			&LeafNode{
				terminal: "+",
				lexeme:   "+",
				pos: &lexer.Position{
					Filename: "test",
					Offset:   8,
					Line:     1,
					Column:   9,
				},
			},
			&InternalNode{
				nonTerminal: "E",
				production: &grammar.Production{
					Head: "E",
					Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("E"), grammar.Terminal("*"), grammar.NonTerminal("E")},
				},
				children: []Node{
					&InternalNode{
						nonTerminal: "E",
						production: &grammar.Production{
							Head: "E",
							Body: grammar.String[grammar.Symbol]{grammar.Terminal("id")},
						},
						children: []Node{
							&LeafNode{
								terminal: "id",
								lexeme:   "count",
								pos: &lexer.Position{
									Filename: "test",
									Offset:   10,
									Line:     1,
									Column:   11,
								},
							},
						},
					},
					&LeafNode{
						terminal: "*",
						lexeme:   "*",
						pos: &lexer.Position{
							Filename: "test",
							Offset:   18,
							Line:     1,
							Column:   19,
						},
					},
					&InternalNode{
						nonTerminal: "E",
						production: &grammar.Production{
							Head: "E",
							Body: grammar.String[grammar.Symbol]{grammar.Terminal("id")},
						},
						children: []Node{
							&LeafNode{
								terminal: "id",
								lexeme:   "price",
								pos: &lexer.Position{
									Filename: "test",
									Offset:   20,
									Line:     1,
									Column:   21,
								},
							},
						},
					},
				},
			},
		},
		annotation: 50,
	}

	n1 := &InternalNode{
		nonTerminal: "E",
		production: &grammar.Production{
			Head: "E",
			Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("E"), grammar.Terminal("*"), grammar.NonTerminal("E")},
		},
		children: []Node{
			&InternalNode{
				nonTerminal: "E",
				production: &grammar.Production{
					Head: "E",
					Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("E"), grammar.Terminal("+"), grammar.NonTerminal("E")},
				},
				children: []Node{
					&InternalNode{
						nonTerminal: "E",
						production: &grammar.Production{
							Head: "E",
							Body: grammar.String[grammar.Symbol]{grammar.Terminal("id")},
						},
						children: []Node{
							&LeafNode{
								terminal: "id",
								lexeme:   "fee",
								pos: &lexer.Position{
									Filename: "test",
									Offset:   2,
									Line:     1,
									Column:   3,
								},
							},
						},
					},
					&LeafNode{
						terminal: "+",
						lexeme:   "+",
						pos: &lexer.Position{
							Filename: "test",
							Offset:   8,
							Line:     1,
							Column:   9,
						},
					},
					&InternalNode{
						nonTerminal: "E",
						production: &grammar.Production{
							Head: "E",
							Body: grammar.String[grammar.Symbol]{grammar.Terminal("id")},
						},
						children: []Node{
							&LeafNode{
								terminal: "id",
								lexeme:   "count",
								pos: &lexer.Position{
									Filename: "test",
									Offset:   10,
									Line:     1,
									Column:   11,
								},
							},
						},
					},
				},
			},
			&LeafNode{
				terminal: "*",
				lexeme:   "*",
				pos: &lexer.Position{
					Filename: "test",
					Offset:   18,
					Line:     1,
					Column:   19,
				},
			},
			&InternalNode{
				nonTerminal: "E",
				production: &grammar.Production{
					Head: "E",
					Body: grammar.String[grammar.Symbol]{grammar.Terminal("id")},
				},
				children: []Node{
					&LeafNode{
						terminal: "id",
						lexeme:   "price",
						pos: &lexer.Position{
							Filename: "test",
							Offset:   20,
							Line:     1,
							Column:   21,
						},
					},
				},
			},
		},
		annotation: 240,
	}

	n2 := &InternalNode{
		nonTerminal: "E",
		production: &grammar.Production{
			Head: "E",
			Body: grammar.String[grammar.Symbol]{grammar.Terminal("id")},
		},
		children: []Node{
			&LeafNode{
				terminal: "id",
				lexeme:   "fee",
				pos: &lexer.Position{
					Filename: "test",
					Offset:   2,
					Line:     1,
					Column:   3,
				},
			},
		},
	}

	n3 := &InternalNode{
		nonTerminal: "E",
		production: &grammar.Production{
			Head: "E",
			Body: grammar.String[grammar.Symbol]{grammar.Terminal("id")},
		},
		children: []Node{
			&LeafNode{
				terminal: "id",
				lexeme:   "count",
				pos: &lexer.Position{
					Filename: "test",
					Offset:   10,
					Line:     1,
					Column:   11,
				},
			},
		},
	}

	n4 := &InternalNode{
		nonTerminal: "E",
		production: &grammar.Production{
			Head: "E",
			Body: grammar.String[grammar.Symbol]{grammar.Terminal("id")},
		},
		children: []Node{
			&LeafNode{
				terminal: "id",
				lexeme:   "price",
				pos: &lexer.Position{
					Filename: "test",
					Offset:   20,
					Line:     1,
					Column:   21,
				},
			},
		},
	}

	return []*InternalNode{n0, n1, n2, n3, n4}
}

func getTestLeafNodes() []*LeafNode {
	n0 := &LeafNode{
		terminal: "id",
		lexeme:   "fee",
		pos: &lexer.Position{
			Filename: "test",
			Offset:   2,
			Line:     1,
			Column:   3,
		},
		annotation: 10,
	}

	n1 := &LeafNode{
		terminal: "+",
		lexeme:   "+",
		pos: &lexer.Position{
			Filename: "test",
			Offset:   8,
			Line:     1,
			Column:   9,
		},
	}

	n2 := &LeafNode{
		terminal: "id",
		lexeme:   "count",
		pos: &lexer.Position{
			Filename: "test",
			Offset:   10,
			Line:     1,
			Column:   11,
		},
		annotation: 2,
	}

	n3 := &LeafNode{
		terminal: "*",
		lexeme:   "*",
		pos: &lexer.Position{
			Filename: "test",
			Offset:   18,
			Line:     1,
			Column:   19,
		},
	}

	n4 := &LeafNode{
		terminal: "id",
		lexeme:   "price",
		pos: &lexer.Position{
			Filename: "test",
			Offset:   20,
			Line:     1,
			Column:   21,
		},
		annotation: 20,
	}

	return []*LeafNode{n0, n1, n2, n3, n4}
}

func TestTraverse(t *testing.T) {
	n := getTestInternalNodes()

	tests := []struct {
		name           string
		n              Node
		order          generic.TraverseOrder
		expectedVisits grammar.String[grammar.Symbol]
	}{
		{
			name:  "VLR",
			n:     n[0],
			order: generic.VLR,
			expectedVisits: grammar.String[grammar.Symbol]{
				grammar.NonTerminal("E"),
				grammar.NonTerminal("E"),
				grammar.Terminal("id"),
				grammar.Terminal("+"),
				grammar.NonTerminal("E"),
				grammar.NonTerminal("E"),
				grammar.Terminal("id"),
				grammar.Terminal("*"),
				grammar.NonTerminal("E"),
				grammar.Terminal("id"),
			},
		},
		{
			name:  "VRL",
			n:     n[0],
			order: generic.VRL,
			expectedVisits: grammar.String[grammar.Symbol]{
				grammar.NonTerminal("E"),
				grammar.NonTerminal("E"),
				grammar.NonTerminal("E"),
				grammar.Terminal("id"),
				grammar.Terminal("*"),
				grammar.NonTerminal("E"),
				grammar.Terminal("id"),
				grammar.Terminal("+"),
				grammar.NonTerminal("E"),
				grammar.Terminal("id"),
			},
		},
		{
			name:  "LRV",
			n:     n[0],
			order: generic.LRV,
			expectedVisits: grammar.String[grammar.Symbol]{
				grammar.Terminal("id"),
				grammar.NonTerminal("E"),
				grammar.Terminal("+"),
				grammar.Terminal("id"),
				grammar.NonTerminal("E"),
				grammar.Terminal("*"),
				grammar.Terminal("id"),
				grammar.NonTerminal("E"),
				grammar.NonTerminal("E"),
				grammar.NonTerminal("E"),
			},
		},
		{
			name:  "RLV",
			n:     n[0],
			order: generic.RLV,
			expectedVisits: grammar.String[grammar.Symbol]{
				grammar.Terminal("id"),
				grammar.NonTerminal("E"),
				grammar.Terminal("*"),
				grammar.Terminal("id"),
				grammar.NonTerminal("E"),
				grammar.NonTerminal("E"),
				grammar.Terminal("+"),
				grammar.Terminal("id"),
				grammar.NonTerminal("E"),
				grammar.NonTerminal("E"),
			},
		},
		{
			name:           "InvalidOrder",
			n:              n[0],
			order:          generic.RVL,
			expectedVisits: grammar.String[grammar.Symbol]{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var visits grammar.String[grammar.Symbol]
			Traverse(tc.n, tc.order, func(n Node) bool {
				visits = append(visits, n.Symbol())
				return true
			})

			assert.True(t, visits.Equals(tc.expectedVisits))
		})
	}
}

func TestNewInternalNode(t *testing.T) {
	tests := []struct {
		name        string
		nonTerminal grammar.NonTerminal
		production  *grammar.Production
		children    []Node
	}{
		{
			name:        "OK",
			nonTerminal: grammar.NonTerminal("E"),
			production: &grammar.Production{
				Head: "E",
				Body: grammar.String[grammar.Symbol]{grammar.Terminal("id")},
			},
			children: []Node{
				&LeafNode{
					terminal: "id",
					lexeme:   "count",
					pos: &lexer.Position{
						Filename: "test",
						Offset:   10,
						Line:     1,
						Column:   11,
					},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			n := NewInternalNode(tc.nonTerminal, tc.production, tc.children...)

			assert.NotNil(t, n)
			assert.Equal(t, tc.nonTerminal, n.nonTerminal)
			assert.Equal(t, tc.production, n.production)
			assert.Equal(t, tc.children, n.children)
		})
	}
}

func TestInternalNode_String(t *testing.T) {
	n := getTestInternalNodes()

	tests := []struct {
		name           string
		n              *InternalNode
		expectedString string
	}{
		{
			name:           "Zero",
			n:              &InternalNode{},
			expectedString: `<nil> <<nil>>`,
		},
		{
			name:           "OK",
			n:              n[0],
			expectedString: `E → E "+" E <fee, test:1:3>`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedString, tc.n.String())
		})
	}
}

func TestInternalNode_Equals(t *testing.T) {
	n := getTestInternalNodes()

	tests := []struct {
		name           string
		n              *InternalNode
		rhs            Node
		expectedEquals bool
	}{
		{
			name:           "Equal",
			n:              n[0],
			rhs:            n[0],
			expectedEquals: true,
		},
		{
			name:           "ProductionsNotEqual",
			n:              n[0],
			rhs:            n[1],
			expectedEquals: false,
		},
		{
			name:           "ChildrenNotEqual",
			n:              n[2],
			rhs:            n[3],
			expectedEquals: false,
		},
		{
			name: "NilProduction",
			n:    n[4],
			rhs: &InternalNode{
				nonTerminal: "E",
			},
			expectedEquals: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedEquals, tc.n.Equals(tc.rhs))
		})
	}
}

func TestInternalNode_Symbol(t *testing.T) {
	n := getTestInternalNodes()

	tests := []struct {
		name           string
		n              *InternalNode
		expectedSymbol grammar.Symbol
	}{
		{
			name:           "OK",
			n:              n[0],
			expectedSymbol: grammar.NonTerminal("E"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedSymbol, tc.n.Symbol())
		})
	}
}

func TestInternalNode_Pos(t *testing.T) {
	n := getTestInternalNodes()

	tests := []struct {
		name        string
		n           *InternalNode
		expectedPos *lexer.Position
	}{
		{
			name: "OK",
			n:    n[0],
			expectedPos: &lexer.Position{
				Filename: "test",
				Offset:   2,
				Line:     1,
				Column:   3,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			pos := tc.n.Pos()

			if tc.expectedPos == nil {
				assert.Nil(t, pos)
			} else {
				assert.True(t, pos.Equals(*tc.expectedPos))
			}
		})
	}
}

func TestInternalNode_Child(t *testing.T) {
	n := getTestInternalNodes()

	tests := []struct {
		name          string
		n             *InternalNode
		i             int
		expectedChild Node
		expectedOK    bool
	}{
		{
			name:          "InvalidIndex",
			n:             n[0],
			i:             3,
			expectedChild: nil,
			expectedOK:    false,
		},
		{
			name: "OK",
			n:    n[0],
			i:    1,
			expectedChild: &LeafNode{
				terminal: "+",
				lexeme:   "+",
				pos: &lexer.Position{
					Filename: "test",
					Offset:   8,
					Line:     1,
					Column:   9,
				},
			},
			expectedOK: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			child, ok := tc.n.Child(tc.i)

			assert.Equal(t, tc.expectedChild, child)
			assert.Equal(t, tc.expectedOK, ok)
		})
	}
}

func TestInternalNode_Annotate(t *testing.T) {
	n := getTestInternalNodes()

	tests := []struct {
		name string
		n    *InternalNode
		key  string
		val  any
	}{
		{
			name: "OK",
			n:    n[0],
			val:  50 * 1.13,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.n.Annotate(tc.val)
			assert.Equal(t, tc.val, tc.n.annotation)
		})
	}
}

func TestInternalNode_GetAnnotation(t *testing.T) {
	n := getTestInternalNodes()

	tests := []struct {
		name          string
		n             *InternalNode
		expectedValue any
	}{
		{
			name:          "OK",
			n:             n[0],
			expectedValue: 50,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			val := tc.n.Annotation()
			assert.Equal(t, tc.expectedValue, val)
		})
	}
}

func TestInternalNode_DOT(t *testing.T) {
	n := getTestInternalNodes()

	tests := []struct {
		name        string
		n           *InternalNode
		expectedDOT string
	}{
		{
			name: "OK",
			n:    n[0],
			expectedDOT: `strict digraph "AST" {
  concentrate=false;
  node [shape=Mrecord];

  1 [label="{ E → | { <0>E | <1>\"+\" | <2>E } }"];
  2 [label="{ E → | { <0>\"id\" } }"];
  3 [label="\"id\" <fee>", style=bold, shape=oval];
  4 [label="\"+\" <+>", style=bold, shape=oval];
  5 [label="{ E → | { <0>E | <1>\"*\" | <2>E } }"];
  6 [label="{ E → | { <0>\"id\" } }"];
  7 [label="\"id\" <count>", style=bold, shape=oval];
  8 [label="\"*\" <*>", style=bold, shape=oval];
  9 [label="{ E → | { <0>\"id\" } }"];
  10 [label="\"id\" <price>", style=bold, shape=oval];

  1:0 -> 2 [];
  1:1 -> 4 [];
  1:2 -> 5 [];
  2:0 -> 3 [];
  5:0 -> 6 [];
  5:1 -> 8 [];
  5:2 -> 9 [];
  6:0 -> 7 [];
  9:0 -> 10 [];
}`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedDOT, tc.n.DOT())
		})
	}
}

func TestNewLeafNode(t *testing.T) {
	tests := []struct {
		name     string
		terminal grammar.Terminal
		lexeme   string
		pos      *lexer.Position
	}{
		{
			name:     "OK",
			terminal: grammar.Terminal("id"),
			lexeme:   "count",
			pos: &lexer.Position{
				Filename: "test",
				Offset:   10,
				Line:     1,
				Column:   11,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			n := NewLeafNode(tc.terminal, tc.lexeme, tc.pos)

			assert.NotNil(t, n)
			assert.Equal(t, tc.terminal, n.terminal)
			assert.Equal(t, tc.lexeme, n.lexeme)
			assert.Equal(t, tc.pos, n.pos)
		})
	}
}

func TestLeafNode_String(t *testing.T) {
	n := getTestLeafNodes()

	tests := []struct {
		name           string
		n              *LeafNode
		expectedString string
	}{
		{
			name:           "Zero",
			n:              &LeafNode{},
			expectedString: `"" <, <nil>>`,
		},
		{
			name:           "OK",
			n:              n[0],
			expectedString: `"id" <fee, test:1:3>`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedString, tc.n.String())
		})
	}
}

func TestLeafNode_Equals(t *testing.T) {
	n := getTestLeafNodes()

	tests := []struct {
		name           string
		n              *LeafNode
		rhs            Node
		expectedEquals bool
	}{
		{
			name:           "Equal",
			n:              n[0],
			rhs:            n[0],
			expectedEquals: true,
		},
		{
			name:           "NotEqual",
			n:              n[0],
			rhs:            n[1],
			expectedEquals: false,
		},
		{
			name: "NilPosition",
			n:    n[0],
			rhs: &LeafNode{
				terminal: "id",
				lexeme:   "fee",
			},
			expectedEquals: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedEquals, tc.n.Equals(tc.rhs))
		})
	}
}

func TestLeafNode_Symbol(t *testing.T) {
	n := getTestLeafNodes()

	tests := []struct {
		name           string
		n              *LeafNode
		expectedSymbol grammar.Symbol
	}{
		{
			name:           "OK",
			n:              n[0],
			expectedSymbol: grammar.Terminal("id"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedSymbol, tc.n.Symbol())
		})
	}
}

func TestLeafNode_Pos(t *testing.T) {
	n := getTestLeafNodes()

	tests := []struct {
		name        string
		n           *LeafNode
		expectedPos *lexer.Position
	}{
		{
			name: "OK",
			n:    n[0],
			expectedPos: &lexer.Position{
				Filename: "test",
				Offset:   2,
				Line:     1,
				Column:   3,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			pos := tc.n.Pos()

			if tc.expectedPos == nil {
				assert.Nil(t, pos)
			} else {
				assert.True(t, pos.Equals(*tc.expectedPos))
			}
		})
	}
}

func TestLeafNode_Child(t *testing.T) {
	tests := []struct {
		name          string
		n             *LeafNode
		i             int
		expectedChild Node
		expectedOK    bool
	}{
		{
			name:          "OK",
			n:             &LeafNode{},
			i:             0,
			expectedChild: nil,
			expectedOK:    false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			child, ok := tc.n.Child(tc.i)

			assert.Equal(t, tc.expectedChild, child)
			assert.Equal(t, tc.expectedOK, ok)
		})
	}
}

func TestLeafNode_Annotate(t *testing.T) {
	n := getTestLeafNodes()

	tests := []struct {
		name string
		n    *LeafNode
		val  any
	}{
		{
			name: "OK",
			n:    n[0],
			val:  10.50,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.n.Annotate(tc.val)

			assert.Equal(t, tc.val, tc.n.annotation)
		})
	}
}

func TestLeafNode_Annotation(t *testing.T) {
	n := getTestLeafNodes()

	tests := []struct {
		name          string
		n             *LeafNode
		key           string
		expectedValue any
	}{
		{
			name:          "OK",
			n:             n[0],
			expectedValue: 10,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			val := tc.n.Annotation()
			assert.Equal(t, tc.expectedValue, val)
		})
	}
}

func TestLeafNode_DOT(t *testing.T) {
	n := getTestLeafNodes()

	tests := []struct {
		name        string
		n           *LeafNode
		expectedDOT string
	}{
		{
			name: "OK",
			n:    n[0],
			expectedDOT: `graph {
  concentrate=false;
  node [];

  1 [label="\"id\" <fee>", style=bold];
}`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedDOT, tc.n.DOT())
		})
	}
}
