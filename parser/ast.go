package parser

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/moorara/algo/generic"
	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/internal/dot"
	"github.com/moorara/algo/lexer"
)

var (
	EqNode = func(lhs, rhs Node) bool {
		return lhs.Equal(rhs)
	}
)

// Node represents a node in an abstract syntax tree (AST)
// derived from an input string allowed based on a context-free grammar.
//
// Each node can be either:
//
//   - An internal node: representing a non-terminal symbol and its associated production rule.
//   - A leaf node: representing a terminal symbol.
type Node interface {
	fmt.Stringer
	generic.Equaler[Node]

	// Symbol returns the grammar symbol associated with this node.
	// For internal nodes, this is a non-terminal symbol
	// (the left-hand side of the production rule represented by the node).
	// For leaf nodes, this is a terminal symbol.
	Symbol() grammar.Symbol

	// Pos returns the leftmost position in the input string that this node represent.
	Pos() lexer.Position

	// Child returns the child node at the specified index (0-based) for this node.
	// The child represents a symbol from the right-hand side of the production rule associated with this node.
	//
	// If the node is internal and the index is out of bounds, the method returns nil and false.
	// This method is not applicable for leaf nodes, and will always return nil and false.
	Child(int) (Node, bool)

	// Annotate associates an annotation with this node.
	// An annotation often represents the node in a different context or type.
	//
	// Common use cases include:
	//
	//   - Storing the result of a type conversion (e.g., converting a string to a number).
	//   - Capturing the outcome of evaluating the right-hand side of a production rule
	//     (e.g., performing an arithmetic operation like addition).
	//   - Associating a reference to a symbol table entry for an identifier.
	//   - Adding metadata or auxiliary information related to the node.
	Annotate(any)

	// Annotation returns the annotation associated with this node.
	// An annotation is a context-specific value of any type, set using the Annotate method.
	//
	// The caller should cast the returned value to the original type used when annotating.
	Annotation() any

	// DOT generates and returns a representation of a node and all its descendants in DOT format.
	// This format is commonly used for visualizing graphs with Graphviz tools.
	DOT() string
}

// Traverse performs a depth-first traversal of an abstract syntax tree (AST),
// starting from the given root node.
// It visits each node according to the specified traversal order
// and passes each node to the provided visit function.
// If the visit function returns false, the traversal is stopped early.
//
// Valid traversal orders for an AST are VLR, VRL, LRV, and RLV.
func Traverse(n Node, order generic.TraverseOrder, visit generic.VisitFunc1[Node]) bool {
	if leaf, ok := n.(*LeafNode); ok {
		return visit(leaf)
	}

	in, ok := n.(*InternalNode)
	if !ok {
		return false
	}

	switch order {
	case generic.VLR:
		res := visit(in)
		for i := range len(in.Children) {
			res = res && Traverse(in.Children[i], order, visit)
		}
		return res

	case generic.VRL:
		res := visit(in)
		for i := len(in.Children) - 1; i >= 0; i-- {
			res = res && Traverse(in.Children[i], order, visit)
		}
		return res

	case generic.LRV:
		res := true
		for i := range len(in.Children) {
			res = res && Traverse(in.Children[i], order, visit)
		}
		return res && visit(in)

	case generic.RLV:
		res := true
		for i := len(in.Children) - 1; i >= 0; i-- {
			res = res && Traverse(in.Children[i], order, visit)
		}
		return res && visit(in)

	default:
		return false
	}
}

// InternalNode represents an internal node in an abstract syntax tree (AST).
// An InternalNode represents a non-terminal symbol and its associated production rule.
type InternalNode struct {
	NonTerminal grammar.NonTerminal
	Production  *grammar.Production
	Children    []Node
	annotation  any
}

// String returns a string representation of an internal node.
func (n *InternalNode) String() string {
	var b bytes.Buffer

	// Find the leftmost leaf of the current node.
	var ll *LeafNode
	var ok bool
	Traverse(n, generic.LRV, func(n Node) bool {
		ll, ok = n.(*LeafNode)
		return false
	})

	if ok {
		fmt.Fprintf(&b, "%s <%s, %s>", n.Production, ll.Lexeme, ll.Position)
	} else {
		fmt.Fprintf(&b, "%s", n.Production)
	}

	return b.String()
}

// Equal determines whether or not two internal nodes are the same.
// Annotations are excluded from the equality check.
func (n *InternalNode) Equal(rhs Node) bool {
	nn, ok := rhs.(*InternalNode)
	if !ok ||
		!n.NonTerminal.Equal(nn.NonTerminal) ||
		!equalProductions(n.Production, nn.Production) ||
		len(n.Children) != len(nn.Children) {
		return false
	}

	for i := range len(n.Children) {
		if !n.Children[i].Equal(nn.Children[i]) {
			return false
		}
	}

	return true
}

func equalProductions(lhs, rhs *grammar.Production) bool {
	if lhs == nil || rhs == nil {
		return lhs == rhs
	}
	return lhs.Equal(rhs)
}

// Symbol returns the non-terminal symbol associated with this internal node
// (the left-hand side of the production rule represented by the node).
func (n *InternalNode) Symbol() grammar.Symbol {
	return n.NonTerminal
}

// Pos returns the position of the first child of this internal node.
func (n *InternalNode) Pos() lexer.Position {
	if len(n.Children) > 0 {
		return n.Children[0].Pos()
	}

	return lexer.Position{}
}

// Child returns the child node at the specified index (0-based) for this internal node.
// The child represents a symbol from the right-hand side of the production rule associated with this node.
//
// If the index is out of bounds, the method returns nil and false.
func (n *InternalNode) Child(i int) (Node, bool) {
	if 0 <= i && i < len(n.Children) {
		return n.Children[i], true
	}

	return nil, false
}

// Annotate associates an annotation with this internal node.
// An annotation often represents the node in a different context or type.
//
// Common use cases include:
//
//   - Storing the result of a type conversion (e.g., converting a string to a number).
//   - Capturing the outcome of evaluating the right-hand side of a production rule
//     (e.g., performing an arithmetic operation like addition).
//   - Adding metadata or auxiliary information related to the node.
func (n *InternalNode) Annotate(val any) {
	n.annotation = val
}

// Annotation returns the annotation associated with this internal node.
// An annotation is a context-specific value of any type, set using the Annotate method.
//
// The caller should cast the returned value to the original type used when annotating.
func (n *InternalNode) Annotation() any {
	return n.annotation
}

// DOT generates and returns a representation of an internal node and all its descendants in DOT format.
// This format is commonly used for visualizing graphs with Graphviz tools.
func (n *InternalNode) DOT() string {
	// Create a map of node --> id
	var id int
	nodeID := map[Node]int{}
	Traverse(n, generic.VLR, func(n Node) bool {
		id++
		nodeID[n] = id
		return true
	})

	graph := dot.NewGraph(true, true, false, "AST", "", "", "", dot.ShapeMrecord)

	Traverse(n, generic.VLR, func(n Node) bool {
		name := fmt.Sprintf("%d", nodeID[n])

		if lf, ok := n.(*LeafNode); ok {
			label := fmt.Sprintf("%s <%s>", lf.Terminal, lf.Lexeme)
			label = strings.ReplaceAll(label, `\t`, `\\t`)
			label = strings.ReplaceAll(label, `\n`, `\\n`)
			label = strings.ReplaceAll(label, `\r`, `\\r`)

			graph.AddNode(dot.NewNode(name, "", label, "", dot.StyleBold, dot.ShapeOval, "", ""))
			return true
		}

		in, ok := n.(*InternalNode)
		if !ok {
			return false
		}

		body := dot.NewRecord()
		for i, X := range in.Production.Body {
			label := X.String()
			label = strings.ReplaceAll(label, `\t`, `\\t`)
			label = strings.ReplaceAll(label, `\n`, `\\n`)
			label = strings.ReplaceAll(label, `\r`, `\\r`)

			body.Fields = append(body.Fields,
				dot.NewSimpleField(fmt.Sprintf("%d", i), label),
			)
		}

		if len(in.Production.Body) == 0 {
			body.Fields = append(body.Fields,
				dot.NewSimpleField("", "ε"),
			)
		}

		rec := dot.NewRecord(
			dot.NewComplexField(
				dot.NewRecord(
					dot.NewSimpleField("", fmt.Sprintf("%s →", in.Production.Head)),
					dot.NewComplexField(body),
				),
			),
		)

		graph.AddNode(dot.NewNode(name, "", rec.Label(), "", "", "", "", ""))

		for i, m := range in.Children {
			from := fmt.Sprintf("%s:%d", name, i)
			to := fmt.Sprintf("%d", nodeID[m])
			graph.AddEdge(dot.NewEdge(from, to, dot.EdgeTypeDirected, "", "", "", "", "", ""))
		}

		return true
	})

	return graph.DOT()
}

// LeafNode represents a leaf node in an abstract syntax tree (AST).
// A LeafNode represents a terminal symbol.
type LeafNode struct {
	Terminal   grammar.Terminal
	Lexeme     string
	Position   lexer.Position
	annotation any
}

// String returns a string representation of a leaf node.
func (n *LeafNode) String() string {
	if n.Position.IsZero() {
		return fmt.Sprintf("%s <%s>", n.Terminal, n.Lexeme)
	}

	return fmt.Sprintf("%s <%s, %s>", n.Terminal, n.Lexeme, n.Position)
}

// Equal determines whether or not two leaf nodes are the same.
// Annotations are excluded from the equality check.
func (n *LeafNode) Equal(rhs Node) bool {
	nn, ok := rhs.(*LeafNode)
	return ok &&
		n.Terminal.Equal(nn.Terminal) &&
		n.Lexeme == nn.Lexeme &&
		n.Position.Equal(nn.Position)
}

// Symbol returns the terminal symbol associated with this leaf node.
func (n *LeafNode) Symbol() grammar.Symbol {
	return n.Terminal
}

// Pos returns the position of the substring represented by this leaf node in the input string.
func (n *LeafNode) Pos() lexer.Position {
	return n.Position
}

// Child method is not applicable for leaf nodes, and will always return nil and false.
func (n *LeafNode) Child(i int) (Node, bool) {
	return nil, false
}

// Annotate associates an annotation with this leaf node.
// An annotation often represents the node in a different context or type.
//
// Common use cases include:
//
//   - Storing the result of a type conversion (e.g., converting a string to a number).
//   - Associating a reference to a symbol table entry for an identifier.
//   - Adding metadata or auxiliary information related to the node.
func (n *LeafNode) Annotate(val any) {
	n.annotation = val
}

// Annotation returns the annotation associated with this leaf node.
// An annotation is a context-specific value of any type, set using the Annotate method.
//
// The caller should cast the returned value to the original type used when annotating.
func (n *LeafNode) Annotation() any {
	return n.annotation
}

// DOT generates and returns a representation of a leaf node in DOT format.
// This format is commonly used for visualizing graphs with Graphviz tools.
func (n *LeafNode) DOT() string {
	label := fmt.Sprintf("%s <%s>", n.Terminal, n.Lexeme)
	label = strings.ReplaceAll(label, "\t", `\\t`)
	label = strings.ReplaceAll(label, "\n", `\\n`)
	label = strings.ReplaceAll(label, "\r", `\\r`)

	graph := dot.NewGraph(false, false, false, "", "", "", "", "")
	graph.AddNode(dot.NewNode("1", "", label, "", dot.StyleBold, "", "", ""))

	return graph.DOT()
}
