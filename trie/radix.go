package trie

import (
	"fmt"
	"iter"
	"strings"

	"github.com/moorara/algo/dot"
	"github.com/moorara/algo/generic"
)

type radixNode[V any] struct {
	label    []byte
	val      V
	term     bool
	children []*radixNode[V]
}

// search looks for the child node (edge) whose label shares a prefix with the given key.
// It returns the index of the child node and the length of the matching prefix.
// If no match is found, it returns the index where the key would fit and 0 as the prefix length.
func (n *radixNode[V]) search(key []byte) (int, int) {
	for i, c := range n.children {
		// Children (edges) are sorted by their first byte.
		if key[0] < c.label[0] {
			return i, 0
		}

		// Compare the key with the child's label.
		l := 0
		for l < len(key) && l < len(c.label) && key[l] == c.label[l] {
			l++
		}

		if l > 0 {
			return i, l // Prefix match found
		}
	}

	return len(n.children), 0
}

type radix[V any] struct {
	size  int
	root  *radixNode[V]
	eqVal generic.EqualFunc[V]
}

// NewRadix creates a new radix trie (a.k.a radix tree and compact prefix tree).
//
// A radix trie is an ordered tree that represents a space-optimized trie (prefix tree).
// It is derived from a regular trie by merging each node, that is the only child of its parent, with its parent.
//
// Like regular tries,
//
//   - The root node is always associated with the empty string.
//   - Keys are stored on a path from the root node to any arbitrary node (keys are stored on edges).
//   - All the descendants of a node have the same common prefix of the string associated with that node.
//
// Unlike regular tries, edges can be labeled with sequences of digits as well as single digits.
func NewRadix[V any](eqVal generic.EqualFunc[V]) Trie[V] {
	return &radix[V]{
		size:  0,
		root:  new(radixNode[V]),
		eqVal: eqVal,
	}
}

func (t *radix[V]) verify() bool {
	// TODO:
	return false
}

// Size returns the number of key-value pairs in the radix trie.
func (t *radix[V]) Size() int {
	return t.size
}

// Height returns the height of the radix trie.
func (t *radix[V]) Height() int {
	return t._height(t.root)
}

func (t *radix[V]) _height(n *radixNode[V]) int {
	if n == nil {
		return 0
	}

	maxH := 0
	for _, c := range n.children {
		maxH = max(maxH, t._height(c))
	}

	return 1 + maxH
}

// IsEmpty returns true if the radix trie is empty.
func (t *radix[V]) IsEmpty() bool {
	return t.size == 0
}

// Put adds a new key-value pair to the radix trie.
func (t *radix[V]) Put(key string, val V) {
	k := []byte(key)
	n := t.root

	for j := 0; j < len(k); {
		i, l := n.search(k[j:])
		if i >= len(n.children) || l < len(n.children[i].label) {
			break
		}

		n = n.children[i]
		j += len(n.label)
	}

	if n.term {
		n.val = val // Update value for the existing key
		return
	}

	// TODO:
}

// Get returns the value of a given key in the radix trie.
func (t *radix[V]) Get(key string) (V, bool) {
	var zeroV V

	k := []byte(key)
	n := t.root

	for j := 0; j < len(k); {
		i, l := n.search(k[j:])
		if i >= len(n.children) || l < len(n.children[i].label) {
			return zeroV, false
		}

		n = n.children[i]
		j += len(n.label)
	}

	if n.term {
		return n.val, true
	}

	return zeroV, false
}

// Delete removes a key-value pair from the radix trie.
func (t *radix[V]) Delete(key string) (val V, ok bool) {
	// TODO:
	var zeroV V
	return zeroV, false
}

// DeleteAll deletes all key-values from the radix trie, leaving it empty.
func (t *radix[V]) DeleteAll() {
	t.size = 0
	t.root = new(radixNode[V])
}

// Min returns the minimum key and its value in the radix trie.
func (t *radix[V]) Min() (string, V, bool) {
	// TODO:
	var zeroV V
	return "", zeroV, false
}

// Max returns the maximum key and its value in the radix trie.
func (t *radix[V]) Max() (string, V, bool) {
	// TODO:
	var zeroV V
	return "", zeroV, false
}

// Floor returns the largest key in the radix trie less than or equal to key.
func (t *radix[V]) Floor(key string) (string, V, bool) {
	// TODO:
	var zeroV V
	return "", zeroV, false
}

// Ceiling returns the smallest key in the radix trie greater than or equal to key.
func (t *radix[V]) Ceiling(key string) (string, V, bool) {
	// TODO:
	var zeroV V
	return "", zeroV, false
}

// DeleteMin removes the smallest key and associated value from the radix trie.
func (t *radix[V]) DeleteMin() (string, V, bool) {
	// TODO:
	var zeroV V
	return "", zeroV, false
}

// DeleteMax removes the largest key and associated value from the radix trie.
func (t *radix[V]) DeleteMax() (string, V, bool) {
	// TODO:
	var zeroV V
	return "", zeroV, false
}

// Select returns the k-th smallest key in the radix trie.
func (t *radix[V]) Select(rank int) (string, V, bool) {
	// TODO:
	var zeroV V
	return "", zeroV, false
}

// Rank returns the number of keys in the radix trie less than key.
func (t *radix[V]) Rank(key string) int {
	// TODO:
	return -1
}

// Range returns all keys and associated values in the radix trie between two given keys.
func (t *radix[V]) Range(lo, hi string) []generic.KeyValue[string, V] {
	// TODO:
	return nil
}

// RangeSize returns the number of keys in the radix trie between two given keys.
func (t *radix[V]) RangeSize(lo, hi string) int {
	// TODO:
	return -1
}

// Match returns all the keys and associated values in the radix trie
// that match the given pattern in which * matches any character.
func (t *radix[V]) Match(pattern string) []generic.KeyValue[string, V] {
	// TODO:
	return nil
}

// WithPrefix returns all the keys and associated values in the radix trie having s as a prefix.
func (t *radix[V]) WithPrefix(prefix string) []generic.KeyValue[string, V] {
	// TODO:
	return nil
}

// LongestPrefixOf returns the key and associated value in the radix trie
// that is the longest prefix of the given key.
func (t *radix[V]) LongestPrefixOf(prefix string) (string, V, bool) {
	// TODO:
	var zeroV V
	return "", zeroV, false
}

// String returns a string representation of the radix trie.
func (t *radix[V]) String() string {
	i := 0
	pairs := make([]string, t.Size())

	t._traverse(t.root, "", generic.Ascending, func(k string, n *radixNode[V]) bool {
		if n.term {
			pairs[i] = fmt.Sprintf("<%v:%v>", k, n.val)
			i++
		}
		return true
	})

	return fmt.Sprintf("{%s}", strings.Join(pairs, " "))
}

// Equal determines whether or not two radix tries have the same key-value pairs.
func (t *radix[V]) Equal(rhs Trie[V]) bool {
	t2, ok := rhs.(*radix[V])
	if !ok {
		return false
	}

	return t._traverse(t.root, "", generic.Ascending, func(k string, n *radixNode[V]) bool { // t ⊂ t2
		if n.term {
			val, ok := t2.Get(k)
			return ok && t.eqVal(n.val, val)
		}
		return true
	}) && t2._traverse(t2.root, "", generic.Ascending, func(k string, n *radixNode[V]) bool { // t2 ⊂ t
		if n.term {
			val, ok := t.Get(k)
			return ok && t.eqVal(n.val, val)
		}
		return true
	})
}

// All returns an iterator sequence containing all the key-value pairs in the radix trie.
func (t *radix[V]) All() iter.Seq2[string, V] {
	return func(yield func(string, V) bool) {
		t._traverse(t.root, "", generic.Ascending, func(k string, n *radixNode[V]) bool {
			return !n.term || yield(k, n.val)
		})
	}
}

// AnyMatch returns true if at least one key-value pair in the radix trie satisfies the provided predicate.
func (t *radix[V]) AnyMatch(p generic.Predicate2[string, V]) bool {
	return !t._traverse(t.root, "", generic.VLR, func(key string, n *radixNode[V]) bool {
		return !n.term || !p(key, n.val)
	})
}

// AllMatch returns true if all key-value pairs in the radix trie satisfy the provided predicate.
// If the radix trie is empty, it returns true.
func (t *radix[V]) AllMatch(p generic.Predicate2[string, V]) bool {
	return t._traverse(t.root, "", generic.VLR, func(key string, n *radixNode[V]) bool {
		return !n.term || p(key, n.val)
	})
}

// FirstMatch returns the first key-value in the radix trie that satisfies the given predicate.
// If no match is found, it returns the zero values of K and V, along with false.
func (t *radix[V]) FirstMatch(p generic.Predicate2[string, V]) (string, V, bool) {
	var k string
	var v V
	var ok bool

	// TODO:

	return k, v, ok
}

// SelectMatch selects a subset of key-values from the radix trie that satisfy the given predicate.
// It returns a new radix trie containing the matching key-values, of the same type as the original radix trie.
func (t *radix[V]) SelectMatch(p generic.Predicate2[string, V]) generic.Collection2[string, V] {
	newT := NewRadix[V](t.eqVal).(*radix[V])

	t._traverse(t.root, "", generic.VLR, func(key string, n *radixNode[V]) bool {
		if n.term && p(key, n.val) {
			newT.Put(key, n.val)
		}
		return true
	})

	return newT
}

// PartitionMatch partitions the key-values in the radix trie
// into two separate radix tries based on the provided predicate.
// The first radix trie contains the key-values that satisfy the predicate (matched key-values),
// while the second radix trie contains those that do not satisfy the predicate (unmatched key-values).
// Both radix tries are of the same type as the original radix trie.
func (t *radix[V]) PartitionMatch(p generic.Predicate2[string, V]) (generic.Collection2[string, V], generic.Collection2[string, V]) {
	matched := NewBinary[V](t.eqVal)
	ummatched := NewBinary[V](t.eqVal)

	// TODO:

	return matched, ummatched
}

// Traverse performs a traversal of the radix trie using the specified traversal order
// and yields the key-value pair of each node to the provided VisitFunc2 function.
//
// If the function returns false, the traversal is halted.
func (t *radix[V]) Traverse(order generic.TraverseOrder, visit generic.VisitFunc2[string, V]) {
	t._traverse(t.root, "", order, func(_ string, n *radixNode[V]) bool {
		return visit(string(n.label), n.val)
	})
}

func (t *radix[V]) _traverse(n *radixNode[V], prefix string, order generic.TraverseOrder, visit func(string, *radixNode[V]) bool) bool {
	if n == nil {
		return true
	}

	prefix = prefix + string(n.label)

	switch order {
	case generic.VLR, generic.Ascending:
		res := visit(prefix, n)
		for i := 0; i < len(n.children); i++ {
			res = res && t._traverse(n.children[i], prefix, order, visit)
		}
		return res

	case generic.VRL:
		res := visit(prefix, n)
		for i := len(n.children) - 1; i >= 0; i-- {
			res = res && t._traverse(n.children[i], prefix, order, visit)
		}
		return res

	case generic.LRV:
		res := true
		for i := 0; i < len(n.children); i++ {
			res = res && t._traverse(n.children[i], prefix, order, visit)
		}
		return res && visit(prefix, n)

	case generic.RLV, generic.Descending:
		res := true
		for i := len(n.children) - 1; i >= 0; i-- {
			res = res && t._traverse(n.children[i], prefix, order, visit)
		}
		return res && visit(prefix, n)

	default:
		return false
	}
}

// DOT generates a representation of the radix trie in DOT format.
// This format is commonly used for visualizing graphs with Graphviz tools.
func (t *radix[V]) DOT() string {
	// Create a map of node --> id
	var id int
	nodeID := map[*radixNode[V]]int{}
	t._traverse(t.root, "", generic.VLR, func(_ string, n *radixNode[V]) bool {
		id++
		nodeID[n] = id
		return true
	})

	graph := dot.NewGraph(true, true, false, "Radix Trie", dot.RankDirTB, "", "", dot.ShapeMrecord)

	t._traverse(t.root, "", generic.VLR, func(_ string, n *radixNode[V]) bool {
		name := fmt.Sprintf("%d", nodeID[n])

		var label string
		var style dot.Style
		var topField dot.Field

		switch {
		case n == t.root:
			label = "•"
			topField = dot.NewSimpleField("", label)
		case !n.term:
			label = string(n.label)
			topField = dot.NewSimpleField("", label)
		default:
			label = string(n.label)
			style = dot.StyleBold
			topField = dot.NewComplexField(
				dot.NewRecord(
					dot.NewSimpleField("", label),
					dot.NewSimpleField("", fmt.Sprintf("%v", n.val)),
				),
			)
		}

		bottomRec := dot.NewRecord()
		for _, e := range n.children {
			edgeName := string(e.label)
			from := fmt.Sprintf("%s:%s", name, edgeName)
			to := fmt.Sprintf("%d", nodeID[e])

			bottomRec.AddField(dot.NewSimpleField(edgeName, ""))
			graph.AddEdge(dot.NewEdge(from, to, dot.EdgeTypeDirected, "", "", "", "", "", ""))
		}

		rec := dot.NewRecord(
			dot.NewComplexField(
				dot.NewRecord(
					topField,
					dot.NewComplexField(bottomRec),
				),
			),
		)

		graph.AddNode(dot.NewNode(name, "", rec.Label(), "", style, "", "", ""))

		return true
	})

	return graph.DOT()
}
