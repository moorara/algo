package trie

import (
	"fmt"

	"github.com/moorara/algo/common"
	"github.com/moorara/algo/internal/graphviz"
)

type radixNode[V any] struct {
	label    []byte
	val      V
	term     bool
	children []*radixNode[V]
}

// search looks for a child node that is a prefix of the given key.
func (n *radixNode[V]) search(key []byte) (int, int) {
	for i, e := range n.children {
		// Children (edges) are sorted
		if key[0] < e.label[0] {
			return i, 0
		}

		j := 0
		for j < len(key) && j < len(e.label) && key[j] == e.label[j] {
			j++
		}

		if j > 0 {
			return i, j
		}
	}

	return len(n.children), 0
}

type radix[V any] struct {
	size int
	root *radixNode[V]
}

// NewRadix creates a new Radix tree.
//
// A Radix tree is an ordered tree that represents a space-optimized trie (prefix tree).
// It is derived from a regular trie by merging each node, that is the only child of its parent, with its parent.
//
// Like regular tries,
//  The root node is always associated with the empty string.
//  Keys are stored on a path from the root node to any arbitrary node (keys are stored on edges).
//  All the descendants of a node have the same common prefix of the string associated with that node.
//
// Unlike regular tries, edges can be labeled with sequences of digits as well as single digits.
func NewRadix[V any]() Trie[V] {
	return &radix[V]{
		size: 0,
		root: new(radixNode[V]),
	}
}

func (t *radix[V]) verify() bool {
	// TODO:
	return false
}

// Size returns the number of key-value pairs in Radix tree.
func (t *radix[V]) Size() int {
	return t.size
}

// Height returns the height of Radix tree.
func (t *radix[V]) Height() int {
	return t._height(t.root)
}

func (t *radix[V]) _height(n *radixNode[V]) int {
	if n == nil {
		return 0
	}

	h := 0
	for _, c := range n.children {
		h = common.Max[int](h, t._height(c))
	}

	return 1 + h
}

// IsEmpty returns true if Radix tree is empty.
func (t *radix[V]) IsEmpty() bool {
	return t.size == 0
}

// Put adds a new key-value pair to Radix tree.
func (t *radix[V]) Put(key string, val V) {
	k := []byte(key)
	last := t.root

	for l := 0; l < len(k); {
		i, j := last.search(k[l:])
		if j < len(last.children[i].label) {
			break
		}

		last = last.children[i]
		l += len(last.label)
	}

	if last.term {
		last.val = val // Update value for the existing key
		return
	}

	// TODO:
}

// Get returns the value of a given key in Radix tree.
func (t *radix[V]) Get(key string) (V, bool) {
	var zeroV V

	k := []byte(key)
	curr := t.root

	for l := 0; l < len(k); {
		i, j := curr.search(k[l:])
		if j < len(curr.children[i].label) {
			return zeroV, false
		}

		curr = curr.children[i]
		l += len(curr.label)
	}

	if curr.term {
		return curr.val, true
	}

	return zeroV, false
}

// Delete removes a key-value pair from Radix tree.
func (t *radix[V]) Delete(key string) (val V, ok bool) {
	// TODO:
	var zeroV V
	return zeroV, false
}

// KeyValues returns all key-value pairs in Radix tree.
func (t *radix[V]) KeyValues() []KeyValue[V] {
	// TODO:
	return nil
}

// Min returns the minimum key and its value in Radix tree.
func (t *radix[V]) Min() (string, V, bool) {
	// TODO:
	var zeroV V
	return "", zeroV, false
}

// Max returns the maximum key and its value in Radix tree.
func (t *radix[V]) Max() (string, V, bool) {
	// TODO:
	var zeroV V
	return "", zeroV, false
}

// Floor returns the largest key in Radix tree less than or equal to key.
func (t *radix[V]) Floor(key string) (string, V, bool) {
	// TODO:
	var zeroV V
	return "", zeroV, false
}

// Ceiling returns the smallest key in Radix tree greater than or equal to key.
func (t *radix[V]) Ceiling(key string) (string, V, bool) {
	// TODO:
	var zeroV V
	return "", zeroV, false
}

// DeleteMin removes the smallest key and associated value from Radix tree.
func (t *radix[V]) DeleteMin() (string, V, bool) {
	// TODO:
	var zeroV V
	return "", zeroV, false
}

// DeleteMax removes the largest key and associated value from Radix tree.
func (t *radix[V]) DeleteMax() (string, V, bool) {
	// TODO:
	var zeroV V
	return "", zeroV, false
}

// Select returns the k-th smallest key in Radix tree.
func (t *radix[V]) Select(rank int) (string, V, bool) {
	// TODO:
	var zeroV V
	return "", zeroV, false
}

// Rank returns the number of keys in Radix tree less than key.
func (t *radix[V]) Rank(key string) int {
	// TODO:
	return -1
}

// RangeSize returns the number of keys in Radix tree between two given keys.
func (t *radix[V]) RangeSize(lo, hi string) int {
	// TODO:
	return -1
}

// Range returns all keys and associated values in Radix tree between two given keys.
func (t *radix[V]) Range(lo, hi string) []KeyValue[V] {
	// TODO:
	return nil
}

// Traverse is used for visiting all key-value pairs in Radix tree.
func (t *radix[V]) Traverse(order TraversalOrder, visit VisitFunc[V]) {
	t._traverse("", t.root, order, func(_ string, n *radixNode[V]) bool {
		return visit(string(n.label), n.val)
	})
}

func (t *radix[V]) _traverse(prefix string, n *radixNode[V], order TraversalOrder, visit func(string, *radixNode[V]) bool) bool {
	if n == nil {
		return true
	}

	prefix = prefix + string(n.label)

	switch order {
	case VLR, Ascending:
		res := visit(prefix, n)
		for i := 0; i < len(n.children); i++ {
			res = res && t._traverse(prefix, n.children[i], order, visit)
		}
		return res

	case VRL:
		res := visit(prefix, n)
		for i := len(n.children) - 1; i >= 0; i-- {
			res = res && t._traverse(prefix, n.children[i], order, visit)
		}
		return res

	case LRV:
		res := true
		for i := 0; i < len(n.children); i++ {
			res = res && t._traverse(prefix, n.children[i], order, visit)
		}
		return res && visit(prefix, n)

	case RLV, Descending:
		res := true
		for i := len(n.children) - 1; i >= 0; i-- {
			res = res && t._traverse(prefix, n.children[i], order, visit)
		}
		return res && visit(prefix, n)

	default:
		return false
	}
}

// Graphviz returns a visualization of Radix tree in Graphviz format.
func (t *radix[V]) Graphviz() string {
	// Create a map of node --> id
	var id int
	nodeID := map[*radixNode[V]]int{}
	t._traverse("", t.root, VLR, func(_ string, n *radixNode[V]) bool {
		id++
		nodeID[n] = id
		return true
	})

	graph := graphviz.NewGraph(true, true, false, "Radix Tree", graphviz.RankDirTB, "", "", graphviz.ShapeMrecord)

	t._traverse("", t.root, VLR, func(_ string, n *radixNode[V]) bool {
		name := fmt.Sprintf("%d", nodeID[n])

		var label string
		var style graphviz.Style
		var topField graphviz.Field

		switch {
		case n == t.root:
			label = "•"
			topField = graphviz.NewSimpleField("", label)
		case !n.term:
			label = string(n.label)
			topField = graphviz.NewSimpleField("", label)
		default:
			label = string(n.label)
			style = graphviz.StyleBold
			topField = graphviz.NewComplexField(
				graphviz.NewRecord(
					graphviz.NewSimpleField("", label),
					graphviz.NewSimpleField("", fmt.Sprintf("%v", n.val)),
				),
			)
		}

		bottomRec := graphviz.NewRecord()
		for _, e := range n.children {
			edgeName := string(e.label)
			from := fmt.Sprintf("%s:%s", name, edgeName)
			to := fmt.Sprintf("%d", nodeID[e])

			bottomRec.AddField(graphviz.NewSimpleField(edgeName, ""))
			graph.AddEdge(graphviz.NewEdge(from, to, graphviz.EdgeTypeDirected, "", "", "", "", "", ""))
		}

		rec := graphviz.NewRecord(
			graphviz.NewComplexField(
				graphviz.NewRecord(
					topField,
					graphviz.NewComplexField(bottomRec),
				),
			),
		)

		graph.AddNode(graphviz.NewNode(name, "", rec.Label(), "", style, "", "", ""))

		return true
	})

	return graph.DotCode()
}

// Match returns all the keys and associated values in Radix tree that match s where * matches any character.
func (t *radix[V]) Match(pattern string) []KeyValue[V] {
	// TODO:
	return nil
}

// WithPrefix returns all the keys and associated values in Radix tree having s as a prefix.
func (t *radix[V]) WithPrefix(prefix string) []KeyValue[V] {
	// TODO:
	return nil
}

// LongestPrefix returns the longest key and associated value that is a prefix of s from Radix tree.
func (t *radix[V]) LongestPrefix(prefix string) (string, V, bool) {
	// TODO:
	var zeroV V
	return "", zeroV, false
}
