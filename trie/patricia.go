package trie

import (
	"fmt"

	"github.com/moorara/algo/common"
	"github.com/moorara/algo/internal/graphviz"
)

type patriciaNode[V any] struct {
	bitPos      int
	key         *bitString
	val         V
	left, right *patriciaNode[V]
}

type patricia[V any] struct {
	size int
	root *patriciaNode[V]
}

// NewPatricia creates a new Patricia tree.
//
// A Patricia tree is a space-optimized variation of tries (prefix trees).
// Patricia tree is a special case of Radix trie with a radix of two (r = 2^x, x = 1).
// As a result, each node has only two children, like a binary tree.
//
// The root node always only has a left child.
// Keys are sequence of bits and stored in nodes (the number of nodes equals the number of keys).
// Patricia is a threaded tree in which nil links are also utilized.
//
// Decent implementations of Patricia tree can often outperform balanced binary trees, and even hash tables.
// Patricia tree performs admirably when its bit-testing loops are well tuned.
func NewPatricia[V any]() Trie[V] {
	return &patricia[V]{
		size: 0,
		root: nil,
	}
}

func (t *patricia[V]) verify() bool {
	if t.root == nil {
		return true
	}

	return t.root.right == nil &&
		t._isPatricia(t.root, t.root.left, empty) &&
		t._isSizeOK() &&
		t._isRankOK()
}

func (t *patricia[V]) _isPatricia(prev, curr *patriciaNode[V], prefix *bitString) bool {
	// Ensure the current node key has the given prefix
	if !curr.key.HasPrefix(prefix) {
		return false
	}

	if curr.bitPos <= prev.bitPos {
		return true
	}

	// Determine the new prefix for children
	prefix = curr.key.Sub(1, curr.bitPos-1)

	return t._isPatricia(curr, curr.left, prefix.Concat(zero)) &&
		t._isPatricia(curr, curr.right, prefix.Concat(one))
}

func (t *patricia[V]) _isSizeOK() bool {
	size := 0

	if t.root != nil {
		t._traverse(t.root.left, Ascending, func(n *patriciaNode[V]) bool {
			size++
			return true
		})
	}

	return t.size == size
}

func (t *patricia[V]) _isRankOK() bool {
	for i := 0; i < t.Size(); i++ {
		k, _, _ := t.Select(i)
		if t.Rank(k) != i {
			return false
		}
	}

	for _, kv := range t.KeyValues() {
		k, _, _ := t.Select(t.Rank(kv.key))
		if kv.key != k {
			return false
		}
	}

	return true
}

// search looks for a key in the tree and returns the last node visited during the search.
func (t *patricia[V]) search(bitKey *bitString) *patriciaNode[V] {
	if t.root == nil {
		return nil
	}

	// Always take the left child of the root
	curr := t.root.left

	for prev := t.root; curr.bitPos > prev.bitPos; {
		prev = curr
		if bitKey.Bit(curr.bitPos) == 0 {
			curr = curr.left
		} else {
			curr = curr.right
		}
	}

	return curr
}

// remove removes a given node from the tree.
//
//	z is the node to delete (target)
//	y is the node pointing to z with a thread (referrer)
//	x is the parent node of y (referrer parent)
//	p is the node pointing to z with a link (parent)
func (t *patricia[V]) remove(z, y, x, p *patriciaNode[V]) {
	if z == y { // Case 1: remove a leaf node
		var c *patriciaNode[V]
		if z != t.root && z.key.Bit(z.bitPos) == 0 {
			c = z.right
		} else {
			c = z.left
		}

		if p == t.root || z.key.Bit(p.bitPos) == 0 {
			p.left = c
		} else {
			p.right = c
		}
	} else { // Case 2: remove a non-leaf node
		var c *patriciaNode[V]
		if y != t.root && z.key.Bit(y.bitPos) == 0 {
			c = y.right
		} else {
			c = y.left
		}

		if x == t.root || z.key.Bit(x.bitPos) == 0 {
			x.left = c
		} else {
			x.right = c
		}

		if p == t.root || z.key.Bit(p.bitPos) == 0 {
			p.left = y
		} else {
			p.right = y
		}

		y.bitPos, y.left, y.right = z.bitPos, z.left, z.right
	}

	if t.size--; t.size == 0 {
		t.root = nil
	}
}

// Size returns the number of key-value pairs in Patricia tree.
func (t *patricia[V]) Size() int {
	return t.size
}

// Height returns the height of Patricia tree.
func (t *patricia[V]) Height() int {
	if t.root == nil {
		return 0
	}

	return t._height(t.root, t.root.left)
}

func (t *patricia[V]) _height(prev, curr *patriciaNode[V]) int {
	if curr.bitPos <= prev.bitPos {
		return 0
	}

	return 1 + common.Max[int](t._height(curr, curr.left), t._height(curr, curr.right))
}

// IsEmpty returns true if Patricia tree is empty.
func (t *patricia[V]) IsEmpty() bool {
	return t.size == 0
}

// Put adds a new key-value pair to Patricia tree.
func (t *patricia[V]) Put(key string, val V) {
	bitKey := newBitString(key)

	if t.root == nil {
		t.root = &patriciaNode[V]{
			bitPos: 0,
			key:    bitKey,
			val:    val,
		}
		t.root.left = t.root
		t.size = 1
		return
	}

	last := t.search(bitKey)
	if last.key.Equals(bitKey) {
		last.val = val // Update value for the existing key
		return
	}

	diffPos := last.key.DiffPos(bitKey)
	prev, next := t.root, t.root.left
	for next.bitPos > prev.bitPos && next.bitPos < diffPos {
		prev = next
		if bitKey.Bit(next.bitPos) == 0 {
			next = next.left
		} else {
			next = next.right
		}
	}

	new := &patriciaNode[V]{
		bitPos: diffPos,
		key:    bitKey,
		val:    val,
	}

	if bitKey.Bit(diffPos) == 0 {
		new.left, new.right = new, next
	} else {
		new.left, new.right = next, new
	}

	if prev.left == next {
		prev.left = new
	} else {
		prev.right = new
	}

	t.size++
}

// Get returns the value of a given key in Patricia tree.
func (t *patricia[V]) Get(key string) (V, bool) {
	bitKey := newBitString(key)
	if n := t.search(bitKey); n != nil && n.key.Equals(bitKey) {
		return n.val, true
	}

	var zeroV V
	return zeroV, false
}

// Delete removes a key-value pair from Patricia tree.
func (t *patricia[V]) Delete(key string) (V, bool) {
	bitKey := newBitString(key)

	if t.root == nil {
		var zeroV V
		return zeroV, false
	}

	// Find the node to delete (z) along side its two preceding nodes (x and y)
	var x, y, z *patriciaNode[V]
	for x, y, z = t.root, t.root, t.root.left; y.bitPos < z.bitPos; {
		x, y = y, z
		if bitKey.Bit(z.bitPos) == 0 {
			z = z.left
		} else {
			z = z.right
		}
	}

	if !z.key.Equals(bitKey) {
		var zeroV V
		return zeroV, false
	}

	// Find the node to delete (q) along side its parent node (p)
	var p, q *patriciaNode[V]
	for p, q = t.root, t.root.left; q != z; {
		p = q
		if bitKey.Bit(q.bitPos) == 0 {
			q = q.left
		} else {
			q = q.right
		}
	}

	t.remove(z, y, x, p)

	return z.val, true
}

// KeyValues returns all key-value pairs in Patricia tree.
func (t *patricia[V]) KeyValues() []KeyValue[V] {
	kvs := make([]KeyValue[V], 0, t.Size())
	t._traverse(t.root, Ascending, func(n *patriciaNode[V]) bool {
		kvs = append(kvs, KeyValue[V]{n.key.String(), n.val})
		return true
	})

	return kvs
}

// Min returns the minimum key and its value in Patricia tree.
func (t *patricia[V]) Min() (string, V, bool) {
	return t._min(t.root)
}

func (t *patricia[V]) _min(n *patriciaNode[V]) (string, V, bool) {
	if n == nil {
		var zeroV V
		return "", zeroV, false
	}

	if n.left.bitPos <= n.bitPos {
		return n.left.key.String(), n.left.val, true
	}

	return t._min(n.left)
}

// Max returns the maximum key and its value in Patricia tree.
func (t *patricia[V]) Max() (string, V, bool) {
	return t._max(t.root)
}

func (t *patricia[V]) _max(n *patriciaNode[V]) (string, V, bool) {
	if n == nil {
		var zeroV V
		return "", zeroV, false
	}

	var next *patriciaNode[V]
	if n == t.root {
		next = n.left
	} else {
		next = n.right
	}

	if next.bitPos <= n.bitPos {
		return next.key.String(), next.val, true
	}

	return t._max(next)
}

// Floor returns the largest key in Patricia tree less than or equal to key.
func (t *patricia[V]) Floor(key string) (string, V, bool) {
	var lastKey string
	var lastVal V
	var ok bool

	t._traverse(t.root, Ascending, func(n *patriciaNode[V]) bool {
		if key < n.key.String() {
			return false
		}
		lastKey, lastVal, ok = n.key.String(), n.val, true
		return true
	})

	return lastKey, lastVal, ok
}

// Ceiling returns the smallest key in Patricia tree greater than or equal to key.
func (t *patricia[V]) Ceiling(key string) (string, V, bool) {
	var lastKey string
	var lastVal V
	var ok bool

	t._traverse(t.root, Descending, func(n *patriciaNode[V]) bool {
		if n.key.String() < key {
			return false
		}
		lastKey, lastVal, ok = n.key.String(), n.val, true
		return true
	})

	return lastKey, lastVal, ok
}

// DeleteMin removes the smallest key and associated value from Patricia tree.
func (t *patricia[V]) DeleteMin() (string, V, bool) {
	if t.root == nil {
		var zeroV V
		return "", zeroV, false
	}

	// Find the node to delete (z) along side its two preceding nodes (x and y).
	var x, y, z *patriciaNode[V]
	for x, y, z = t.root, t.root, t.root.left; y.bitPos < z.bitPos; {
		x, y, z = y, z, z.left
	}

	// Find the node to delete (q) along side its parent node (p).
	var p, q *patriciaNode[V]
	for p, q = t.root, t.root.left; q != z; {
		p, q = q, q.left
	}

	t.remove(z, y, x, p)

	return z.key.String(), z.val, true
}

// DeleteMax removes the largest key and associated value from Patricia tree.
func (t *patricia[V]) DeleteMax() (string, V, bool) {
	if t.root == nil {
		var zeroV V
		return "", zeroV, false
	}

	// Find the node to delete (z) along side its two preceding nodes (x and y).
	var x, y, z *patriciaNode[V]
	for x, y, z = t.root, t.root, t.root.left; y.bitPos < z.bitPos; {
		x, y, z = y, z, z.right
	}

	// Find the node to delete (q) along side its parent node (p).
	var p, q *patriciaNode[V]
	for p, q = t.root, t.root.left; q != z; {
		p, q = q, q.right
	}

	t.remove(z, y, x, p)

	return z.key.String(), z.val, true
}

// Select returns the k-th smallest key in Patricia tree.
func (t *patricia[V]) Select(rank int) (string, V, bool) {
	var lastKey string
	var lastVal V
	var ok bool

	if t.root == nil || rank < 0 || rank >= t.Size() {
		return lastKey, lastVal, false
	}

	i := 0
	t._traverse(t.root.left, Ascending, func(n *patriciaNode[V]) bool {
		if i == rank {
			lastKey, lastVal, ok = n.key.String(), n.val, true
			return false
		}

		i++
		return true
	})

	return lastKey, lastVal, ok
}

// Rank returns the number of keys in Patricia tree less than key.
func (t *patricia[V]) Rank(key string) int {
	i := 0

	if t.root != nil {
		t._traverse(t.root.left, Ascending, func(n *patriciaNode[V]) bool {
			if n.key.String() == key {
				return false
			}

			i++
			return true
		})
	}

	return i
}

// RangeSize returns the number of keys in Patricia tree between two given keys.
func (t *patricia[V]) RangeSize(lo, hi string) int {
	i := 0

	if t.root != nil {
		t._traverse(t.root.left, Ascending, func(n *patriciaNode[V]) bool {
			if lo <= n.key.String() && n.key.String() <= hi {
				i++
			} else if n.key.String() > hi {
				return false
			}

			return true
		})
	}

	return i
}

// Range returns all keys and associated values in Patricia tree between two given keys.
func (t *patricia[V]) Range(lo, hi string) []KeyValue[V] {
	kvs := []KeyValue[V]{}

	if t.root != nil {
		t._traverse(t.root.left, Ascending, func(n *patriciaNode[V]) bool {
			if lo <= n.key.String() && n.key.String() <= hi {
				kvs = append(kvs, KeyValue[V]{n.key.String(), n.val})
			} else if n.key.String() > hi {
				return false
			}

			return true
		})
	}

	return kvs
}

// Traverse is used for visiting all key-value pairs in Patricia tree.
func (t *patricia[V]) Traverse(order TraversalOrder, visit VisitFunc[V]) {
	t._traverse(t.root, order, func(n *patriciaNode[V]) bool {
		return visit(n.key.String(), n.val)
	})
}

func (t *patricia[V]) _traverse(n *patriciaNode[V], order TraversalOrder, visit func(*patriciaNode[V]) bool) bool {
	if n == nil {
		return true
	}

	isLeftThread := n.left.bitPos <= n.bitPos                  // left links are never nil
	isRightThread := n != t.root && n.right.bitPos <= n.bitPos // Only the root node has a nil right

	switch order {
	case VLR:
		return visit(n) &&
			(isLeftThread || t._traverse(n.left, order, visit)) &&
			(isRightThread || t._traverse(n.right, order, visit))

	case VRL:
		return visit(n) &&
			(isRightThread || t._traverse(n.right, order, visit)) &&
			(isLeftThread || t._traverse(n.left, order, visit))

	case LVR:
		return (isLeftThread || t._traverse(n.left, order, visit)) &&
			visit(n) &&
			(isRightThread || t._traverse(n.right, order, visit))

	case RVL:
		return (isRightThread || t._traverse(n.right, order, visit)) &&
			visit(n) &&
			(isLeftThread || t._traverse(n.left, order, visit))

	case LRV:
		return (isLeftThread || t._traverse(n.left, order, visit)) &&
			(isRightThread || t._traverse(n.right, order, visit)) &&
			visit(n)

	case RLV:
		return (isRightThread || t._traverse(n.right, order, visit)) &&
			(isLeftThread || t._traverse(n.left, order, visit)) &&
			visit(n)

	case Ascending:
		return (!isLeftThread || visit(n.left)) &&
			(isLeftThread || t._traverse(n.left, order, visit)) &&
			(!isRightThread || visit(n.right)) &&
			(isRightThread || t._traverse(n.right, order, visit))

	case Descending:
		return (!isRightThread || visit(n.right)) &&
			(isRightThread || t._traverse(n.right, order, visit)) &&
			(!isLeftThread || visit(n.left)) &&
			(isLeftThread || t._traverse(n.left, order, visit))

	default:
		return false
	}
}

// Graphviz returns a visualization of Patricia tree in Graphviz format.
func (t *patricia[V]) Graphviz() string {
	// Create a map of node --> id
	var id int
	nodeID := map[*patriciaNode[V]]int{}
	t._traverse(t.root, VLR, func(n *patriciaNode[V]) bool {
		id++
		nodeID[n] = id
		return true
	})

	graph := graphviz.NewGraph(true, true, false, "Patricia", graphviz.RankDirTB, "", "", graphviz.ShapeMrecord)

	t._traverse(t.root, VLR, func(n *patriciaNode[V]) bool {
		name := fmt.Sprintf("%d", nodeID[n])

		rec := graphviz.NewRecord(
			graphviz.NewComplexField(
				graphviz.NewRecord(
					graphviz.NewSimpleField("", fmt.Sprintf("%s,%v", n.key, n.val)),
					graphviz.NewComplexField(
						graphviz.NewRecord(
							graphviz.NewSimpleField("l", "•"),
							graphviz.NewSimpleField("", fmt.Sprintf("%d", n.bitPos)),
							graphviz.NewSimpleField("", n.key.BitString()),
							graphviz.NewSimpleField("r", "•"),
						),
					),
				),
			),
		)

		graph.AddNode(graphviz.NewNode(name, "", rec.Label(), "", "", "", "", ""))

		from := fmt.Sprintf("%s:l", name)
		left := fmt.Sprintf("%d", nodeID[n.left])

		var color graphviz.Color
		var style graphviz.Style

		if n.left.bitPos > n.bitPos {
			color = graphviz.ColorBlue
		} else {
			color = graphviz.ColorRed
			style = graphviz.StyleDashed
		}

		graph.AddEdge(graphviz.NewEdge(from, left, graphviz.EdgeTypeDirected, "", "", color, style, "", ""))

		if n != t.root {
			from := fmt.Sprintf("%s:r", name)
			right := fmt.Sprintf("%d", nodeID[n.right])

			var color graphviz.Color
			var style graphviz.Style

			if n.right.bitPos > n.bitPos {
				color = graphviz.ColorBlue
			} else {
				color = graphviz.ColorRed
				style = graphviz.StyleDashed
			}

			graph.AddEdge(graphviz.NewEdge(from, right, graphviz.EdgeTypeDirected, "", "", color, style, "", ""))
		}

		return true
	})

	return graph.DotCode()
}

// Match returns all the keys and associated values in Patricia tree that match s where * matches any character.
func (t *patricia[V]) Match(pattern string) []KeyValue[V] {
	// TODO:
	return nil
}

func (t *patricia[V]) _match(pattern string) []KeyValue[V] {
	// TODO:
	return nil
}

// WithPrefix returns all the keys and associated values in Patricia tree having s as a prefix.
func (t *patricia[V]) WithPrefix(prefix string) []KeyValue[V] {
	// TODO:
	return nil
}

// LongestPrefix returns the longest key and associated value that is a prefix of s from Patricia tree.
func (t *patricia[V]) LongestPrefix(prefix string) (string, V, bool) {
	// TODO:
	var zeroV V
	return "", zeroV, false
}
