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
		NonTerminal: "E",
		Production: &grammar.Production{
			Head: "E",
			Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("E"), grammar.Terminal("+"), grammar.NonTerminal("E")},
		},
		Children: []Node{
			&InternalNode{
				NonTerminal: "E",
				Production: &grammar.Production{
					Head: "E",
					Body: grammar.String[grammar.Symbol]{grammar.Terminal("id")},
				},
				Children: []Node{
					&LeafNode{
						Terminal: "id",
						Lexeme:   "fee",
						Position: lexer.Position{
							Filename: "test",
							Offset:   2,
							Line:     1,
							Column:   3,
						},
					},
				},
			},
			&LeafNode{
				Terminal: "+",
				Lexeme:   "+",
				Position: lexer.Position{
					Filename: "test",
					Offset:   8,
					Line:     1,
					Column:   9,
				},
			},
			&InternalNode{
				NonTerminal: "E",
				Production: &grammar.Production{
					Head: "E",
					Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("E"), grammar.Terminal("*"), grammar.NonTerminal("E")},
				},
				Children: []Node{
					&InternalNode{
						NonTerminal: "E",
						Production: &grammar.Production{
							Head: "E",
							Body: grammar.String[grammar.Symbol]{grammar.Terminal("id")},
						},
						Children: []Node{
							&LeafNode{
								Terminal: "id",
								Lexeme:   "count",
								Position: lexer.Position{
									Filename: "test",
									Offset:   10,
									Line:     1,
									Column:   11,
								},
							},
						},
					},
					&LeafNode{
						Terminal: "*",
						Lexeme:   "*",
						Position: lexer.Position{
							Filename: "test",
							Offset:   18,
							Line:     1,
							Column:   19,
						},
					},
					&InternalNode{
						NonTerminal: "E",
						Production: &grammar.Production{
							Head: "E",
							Body: grammar.String[grammar.Symbol]{grammar.Terminal("id")},
						},
						Children: []Node{
							&LeafNode{
								Terminal: "id",
								Lexeme:   "price",
								Position: lexer.Position{
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
		NonTerminal: "E",
		Production: &grammar.Production{
			Head: "E",
			Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("E"), grammar.Terminal("*"), grammar.NonTerminal("E")},
		},
		Children: []Node{
			&InternalNode{
				NonTerminal: "E",
				Production: &grammar.Production{
					Head: "E",
					Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("E"), grammar.Terminal("+"), grammar.NonTerminal("E")},
				},
				Children: []Node{
					&InternalNode{
						NonTerminal: "E",
						Production: &grammar.Production{
							Head: "E",
							Body: grammar.String[grammar.Symbol]{grammar.Terminal("id")},
						},
						Children: []Node{
							&LeafNode{
								Terminal: "id",
								Lexeme:   "fee",
								Position: lexer.Position{
									Filename: "test",
									Offset:   2,
									Line:     1,
									Column:   3,
								},
							},
						},
					},
					&LeafNode{
						Terminal: "+",
						Lexeme:   "+",
						Position: lexer.Position{
							Filename: "test",
							Offset:   8,
							Line:     1,
							Column:   9,
						},
					},
					&InternalNode{
						NonTerminal: "E",
						Production: &grammar.Production{
							Head: "E",
							Body: grammar.String[grammar.Symbol]{grammar.Terminal("id")},
						},
						Children: []Node{
							&LeafNode{
								Terminal: "id",
								Lexeme:   "count",
								Position: lexer.Position{
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
				Terminal: "*",
				Lexeme:   "*",
				Position: lexer.Position{
					Filename: "test",
					Offset:   18,
					Line:     1,
					Column:   19,
				},
			},
			&InternalNode{
				NonTerminal: "E",
				Production: &grammar.Production{
					Head: "E",
					Body: grammar.String[grammar.Symbol]{grammar.Terminal("id")},
				},
				Children: []Node{
					&LeafNode{
						Terminal: "id",
						Lexeme:   "price",
						Position: lexer.Position{
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
		NonTerminal: "E",
		Production: &grammar.Production{
			Head: "E",
			Body: grammar.String[grammar.Symbol]{grammar.Terminal("id")},
		},
		Children: []Node{
			&LeafNode{
				Terminal: "id",
				Lexeme:   "fee",
				Position: lexer.Position{
					Filename: "test",
					Offset:   2,
					Line:     1,
					Column:   3,
				},
			},
		},
	}

	n3 := &InternalNode{
		NonTerminal: "E",
		Production: &grammar.Production{
			Head: "E",
			Body: grammar.String[grammar.Symbol]{grammar.Terminal("id")},
		},
		Children: []Node{
			&LeafNode{
				Terminal: "id",
				Lexeme:   "count",
				Position: lexer.Position{
					Filename: "test",
					Offset:   10,
					Line:     1,
					Column:   11,
				},
			},
		},
	}

	n4 := &InternalNode{
		NonTerminal: "E",
		Production: &grammar.Production{
			Head: "E",
			Body: grammar.String[grammar.Symbol]{grammar.Terminal("id")},
		},
		Children: []Node{
			&LeafNode{
				Terminal: "id",
				Lexeme:   "price",
				Position: lexer.Position{
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
		Terminal: "id",
		Lexeme:   "fee",
		Position: lexer.Position{
			Filename: "test",
			Offset:   2,
			Line:     1,
			Column:   3,
		},
		annotation: 10,
	}

	n1 := &LeafNode{
		Terminal: "+",
		Lexeme:   "+",
		Position: lexer.Position{
			Filename: "test",
			Offset:   8,
			Line:     1,
			Column:   9,
		},
	}

	n2 := &LeafNode{
		Terminal: "id",
		Lexeme:   "count",
		Position: lexer.Position{
			Filename: "test",
			Offset:   10,
			Line:     1,
			Column:   11,
		},
		annotation: 2,
	}

	n3 := &LeafNode{
		Terminal: "*",
		Lexeme:   "*",
		Position: lexer.Position{
			Filename: "test",
			Offset:   18,
			Line:     1,
			Column:   19,
		},
	}

	n4 := &LeafNode{
		Terminal: "id",
		Lexeme:   "price",
		Position: lexer.Position{
			Filename: "test",
			Offset:   20,
			Line:     1,
			Column:   21,
		},
		annotation: 20,
	}

	return []*LeafNode{n0, n1, n2, n3, n4}
}

func TestEqNode(t *testing.T) {
	in := getTestInternalNodes()
	ln := getTestLeafNodes()

	tests := []struct {
		name          string
		lhs           Node
		rhs           Node
		expectedEqual bool
	}{
		{
			name:          "BothInternal_Equal",
			lhs:           in[0],
			rhs:           in[0],
			expectedEqual: true,
		},
		{
			name:          "BothInternal_NotEqual",
			lhs:           in[0],
			rhs:           in[1],
			expectedEqual: false,
		},
		{
			name:          "BothLeaf_Equal",
			lhs:           ln[0],
			rhs:           ln[0],
			expectedEqual: true,
		},
		{
			name:          "BothLeaf_NotEqual",
			lhs:           ln[0],
			rhs:           ln[1],
			expectedEqual: false,
		},
		{
			name:          "InternalAndLeaf_NotEqual",
			lhs:           in[0],
			rhs:           ln[0],
			expectedEqual: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedEqual, EqNode(tc.lhs, tc.rhs))
		})
	}
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

			assert.True(t, visits.Equal(tc.expectedVisits))
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
			expectedString: `<nil>`,
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

func TestInternalNode_Equal(t *testing.T) {
	n := getTestInternalNodes()

	tests := []struct {
		name          string
		n             *InternalNode
		rhs           Node
		expectedEqual bool
	}{
		{
			name:          "Equal",
			n:             n[0],
			rhs:           n[0],
			expectedEqual: true,
		},
		{
			name:          "ProductionsNotEqual",
			n:             n[0],
			rhs:           n[1],
			expectedEqual: false,
		},
		{
			name:          "ChildrenNotEqual",
			n:             n[2],
			rhs:           n[3],
			expectedEqual: false,
		},
		{
			name: "NilProduction",
			n:    n[4],
			rhs: &InternalNode{
				NonTerminal: "E",
			},
			expectedEqual: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedEqual, tc.n.Equal(tc.rhs))
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
		expectedPos lexer.Position
	}{
		{
			name: "OK",
			n:    n[0],
			expectedPos: lexer.Position{
				Filename: "test",
				Offset:   2,
				Line:     1,
				Column:   3,
			},
		},
		{
			name: "EmptyProduction",
			n: &InternalNode{
				NonTerminal: "E",
				Production: &grammar.Production{
					Head: "E",
					Body: grammar.E,
				},
			},
			expectedPos: lexer.Position{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			pos := tc.n.Pos()

			assert.True(t, pos.Equal(tc.expectedPos))
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
				Terminal: "+",
				Lexeme:   "+",
				Position: lexer.Position{
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
			expectedString: `"" <>`,
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

func TestLeafNode_Equal(t *testing.T) {
	n := getTestLeafNodes()

	tests := []struct {
		name          string
		n             *LeafNode
		rhs           Node
		expectedEqual bool
	}{
		{
			name:          "Equal",
			n:             n[0],
			rhs:           n[0],
			expectedEqual: true,
		},
		{
			name:          "NotEqual",
			n:             n[0],
			rhs:           n[1],
			expectedEqual: false,
		},
		{
			name: "NilPosition",
			n:    n[0],
			rhs: &LeafNode{
				Terminal: "id",
				Lexeme:   "fee",
			},
			expectedEqual: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedEqual, tc.n.Equal(tc.rhs))
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
		expectedPos lexer.Position
	}{
		{
			name: "OK",
			n:    n[0],
			expectedPos: lexer.Position{
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

			assert.True(t, pos.Equal(tc.expectedPos))
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
