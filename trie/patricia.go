package trie

import (
	"fmt"

	"github.com/moorara/algo/generic"
	"github.com/moorara/algo/internal/graphviz"
)

type patriciaNode[V any] struct {
	bp          int // bit pos
	key         *bitString
	val         V
	left, right *patriciaNode[V]
}

type patricia[V any] struct {
	size  int
	root  *patriciaNode[V]
	eqVal generic.EqualFunc[V]
}

// NewPatricia creates a new Patricia tree.
//
// A Patricia tree is a space-optimized variation of tries (prefix trees).
// Patricia tree is a special case of Radix trie with a radix of two (r = 2^x, x = 1).
// As a result, each node has only two children, like a binary tree.
//
// The root node's bit position is always zero and it always only has a left child.
// Keys are a sequence of bits stored in nodes (the number of nodes equals the number of keys).
// Patricia is a threaded tree in which nil links are also utilized.
//
// A Patricia tree is derived from a digital search tree in few steps.
//
// A digital search tree is a binary tree in which keys are a sequence of bits stored in nodes.
// The i'th level in a DST corresponds to the i'th bit of the keys.
// At any given node at the i'th level, the left sub-tree includes all the keys with i'th bit set to zero
// and the right sub-tree includes all the keys with i'th bit set to one.
//
//	               ┌──<1000>──┐
//	               │          │
//	         ┌──<0010>      <1001>──┐
//	         │                      │
//	   ┌──<0011>──┐               <1100>
//	   │          │
//	<0000>      <0011>
//
// For searching (or inserting) a key in a DST, we start from the root node.
// At a given node at the i'th level, we compare the search key with the key in the node.
// If keys are not equal, we look at the i'th bit of the search key;
// If zero we continue the search with the left sub-tree, and if one we continue the search with the right sub-tree.
//
// To reduce the number of compares required, we introduce the binary trie.
// A binary trie, in this context, is a binary tree with two kinds of nodes: internal nodes and leaf nodes.
// Internals nodes are only for guiding the search and leaf nodes contain the keys.
// Similar to digital search trees, the i'th level corresponds to the i'th bit of the keys.
//
//	                     ┌───────────[    ]───────────┐
//	                     │                            │
//	               ┌──[    ]                     ┌──[    ]──┐
//	               │                             │          │
//	         ┌──[    ]──┐                  ┌──[    ]      <1100>
//	         │          │                  │
//	   ┌──[    ]──┐   <0010>         ┌──[    ]──┐
//	   │          │                  │          │
//	<0000>      <0001>            <1000>      <1001>
//
// We can further optimize a binary trie and build a compressed binary trie.
// To do so, we merge the internal nodes with only one child and add a bit position field to each internal node.
// The bit position at a given internal node determines which i'th bit of the keys should be tested at that node.
//
//	               ┌────────[ 1  ]────────┐
//	               │                      │
//	         ┌──[  3 ]──┐            ┌──[ 2  ]──┐
//	         │          │            │          │
//	   ┌──[  4 ]──┐   <0010>   ┌──[  4 ]──┐   <1100>
//	   │          │            │          │
//	<0000>      <0001>      <1000>      <1001>
//
// We can finally derive a Patricia tree from a compressed binary trie.
// To this end, we substitute every internal node with a Patricia node.
// Since the number of internal nodes is one less than the number of leaf nodes, we add an extra Patricia node.
// This extra Patricia node is the root of the tree, its bit position is always zero,
// its left child points to the rest of the tree, and its right child is always nil.
// We move keys from leaf nodes to Patricia nodes in such a way that
// the bit number in Patricia nodes is equal to or less than the bit number in the parent node of the leaf node.
// Pointers from internal nodes to leaf nodes become thread pointers in the Patricia tree.
//
//	.......................... ┌────────(0|1100)...
//	:                        : │                  :
//	:             ┌────────(1|0000)────────┐      :
//	:             │                        │      :
//	:      ┌──(3|0010)...            ┌──(2|1001)..:
//	:      │      :.....:            │     :
//	:..(4|0001)...            ...(4|1000)..:
//	       :.....:            :.....:
//
// Decent implementations of Patricia tree can often outperform balanced binary trees, and even hash tables.
// Patricia tree performs admirably when its bit-testing loops are well tuned.
//
// The second parameter (eqVal) is needed only if you want to use the Equals method.
func NewPatricia[V any](eqVal generic.EqualFunc[V]) Trie[V] {
	return &patricia[V]{
		size:  0,
		root:  nil,
		eqVal: eqVal,
	}
}

// nolint: unused
func (t *patricia[V]) verify() bool {
	if t.root == nil {
		return true
	}

	return t.root.right == nil &&
		t._isPatricia(t.root, t.root.left, empty) &&
		t._isSizeOK() &&
		t._isRankOK()
}

// nolint: unused
func (t *patricia[V]) _isPatricia(prev, curr *patriciaNode[V], prefix *bitString) bool {
	// Ensure the current node key has the given prefix
	if !curr.key.HasPrefix(prefix) {
		return false
	}

	if curr.bp <= prev.bp {
		return true
	}

	// Determine the new prefix for children
	prefix = curr.key.Sub(1, curr.bp-1)

	return t._isPatricia(curr, curr.left, prefix.Concat(zero)) &&
		t._isPatricia(curr, curr.right, prefix.Concat(one))
}

// nolint: unused
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

// nolint: unused
func (t *patricia[V]) _isRankOK() bool {
	for i := 0; i < t.Size(); i++ {
		k, _, _ := t.Select(i)
		if t.Rank(k) != i {
			return false
		}
	}

	for _, kv := range t.KeyValues() {
		k, _, _ := t.Select(t.Rank(kv.Key))
		if kv.Key != k {
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

	for prev := t.root; curr.bp > prev.bp; {
		prev = curr
		if bitKey.Bit(curr.bp) {
			curr = curr.right
		} else {
			curr = curr.left
		}
	}

	return curr
}

// remove removes a given node from the tree.
//
//	n is the node to delete (target)
//	r is the node pointing to z with a thread (referrer)
//	rp is the parent node of y (referrer parent)
//	np is the node pointing to z with a link (parent)
func (t *patricia[V]) remove(n, r, rp, np *patriciaNode[V]) {
	var c *patriciaNode[V] // the other child of the referrer
	if r == t.root || n.key.Bit(r.bp) {
		c = r.left
	} else {
		c = r.right
	}

	if n == r { // Case 1: remove a leaf node
		if np != t.root && n.key.Bit(np.bp) {
			np.right = c
		} else {
			np.left = c
		}
	} else { // Case 2: remove a non-leaf node
		if rp != t.root && n.key.Bit(rp.bp) {
			rp.right = c
		} else {
			rp.left = c
		}

		if np != t.root && n.key.Bit(np.bp) {
			np.right = r
		} else {
			np.left = r
		}

		if n == t.root {
			t.root = np
		}

		r.bp = n.bp
		r.left, r.right = n.left, n.right
	}

	if t.size--; t.size == 0 {
		t.root = nil
	}
}

// Size returns the number of key-value pairs in the Patricia Trie.
func (t *patricia[V]) Size() int {
	return t.size
}

// Height returns the height of the Patricia Trie.
func (t *patricia[V]) Height() int {
	if t.root == nil {
		return 0
	}

	return t._height(t.root, t.root.left)
}

func (t *patricia[V]) _height(prev, curr *patriciaNode[V]) int {
	if curr.bp <= prev.bp {
		return 0
	}

	return 1 + generic.Max[int](t._height(curr, curr.left), t._height(curr, curr.right))
}

// IsEmpty returns true if the Patricia Trie is empty.
func (t *patricia[V]) IsEmpty() bool {
	return t.size == 0
}

// Put adds a new key-value pair to the Patricia Trie.
func (t *patricia[V]) Put(key string, val V) {
	t._put(newBitString(key), val)
}

func (t *patricia[V]) _put(key *bitString, val V) {
	if t.root == nil {
		t.root = &patriciaNode[V]{
			bp:  0,
			key: key,
			val: val,
		}
		t.root.left = t.root
		t.size = 1
		return
	}

	last := t.search(key)
	if last.key.Equals(key) {
		last.val = val // Update value for the existing key
		return
	}

	diffPos := last.key.DiffPos(key)
	prev, next := t.root, t.root.left
	for next.bp > prev.bp && next.bp < diffPos {
		prev = next
		if key.Bit(next.bp) {
			next = next.right
		} else {
			next = next.left
		}
	}

	new := &patriciaNode[V]{
		bp:  diffPos,
		key: key,
		val: val,
	}

	if key.Bit(diffPos) {
		new.left, new.right = next, new
	} else {
		new.left, new.right = new, next
	}

	if prev.left == next {
		prev.left = new
	} else {
		prev.right = new
	}

	t.size++
}

// Get returns the value of a given key in the Patricia Trie.
func (t *patricia[V]) Get(key string) (V, bool) {
	return t._get(newBitString(key))
}

func (t *patricia[V]) _get(key *bitString) (V, bool) {
	if n := t.search(key); n != nil && n.key.Equals(key) {
		return n.val, true
	}

	var zeroV V
	return zeroV, false
}

// Delete removes a key-value pair from the Patricia Trie.
func (t *patricia[V]) Delete(key string) (V, bool) {
	return t._delete(newBitString(key))
}

func (t *patricia[V]) _delete(key *bitString) (V, bool) {
	if t.root == nil {
		var zeroV V
		return zeroV, false
	}

	// Find the node to delete (z) along side its two preceding nodes (x and y)
	var rp, r, n *patriciaNode[V]
	for rp, r, n = t.root, t.root, t.root.left; r.bp < n.bp; {
		rp, r = r, n
		if key.Bit(n.bp) {
			n = n.right
		} else {
			n = n.left
		}
	}

	if !n.key.Equals(key) {
		var zeroV V
		return zeroV, false
	}

	// Find the node to delete (q) along side its parent node (p)
	var np, m *patriciaNode[V]
	for np, m = t.root, t.root.left; m != n; {
		np = m
		if key.Bit(m.bp) {
			m = m.right
		} else {
			m = m.left
		}
	}

	t.remove(n, r, rp, np)

	return n.val, true
}

// KeyValues returns all key-value pairs in the Patricia Trie.
func (t *patricia[V]) KeyValues() []KeyValue[V] {
	kvs := make([]KeyValue[V], 0, t.Size())
	t._traverse(t.root, Ascending, func(n *patriciaNode[V]) bool {
		kvs = append(kvs, KeyValue[V]{n.key.String(), n.val})
		return true
	})

	return kvs
}

// Equals determines whether or not two Patricia Tries have the same key-value pairs.
func (t *patricia[V]) Equals(u Trie[V]) bool {
	tt, ok := u.(*patricia[V])
	if !ok {
		return false
	}

	return t._traverse(t.root, Ascending, func(n *patriciaNode[V]) bool { // t ⊂ tt
		val, ok := tt._get(n.key)
		return ok && t.eqVal(n.val, val)
	}) && tt._traverse(tt.root, Ascending, func(n *patriciaNode[V]) bool { // tt ⊂ t
		val, ok := t._get(n.key)
		return ok && t.eqVal(n.val, val)
	})
}

// Min returns the minimum key and its value in the Patricia Trie.
func (t *patricia[V]) Min() (string, V, bool) {
	return t._min(t.root)
}

func (t *patricia[V]) _min(n *patriciaNode[V]) (string, V, bool) {
	if n == nil {
		var zeroV V
		return "", zeroV, false
	}

	if n.left.bp <= n.bp {
		return n.left.key.String(), n.left.val, true
	}

	return t._min(n.left)
}

// Max returns the maximum key and its value in the Patricia Trie.
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

	if next.bp <= n.bp {
		return next.key.String(), next.val, true
	}

	return t._max(next)
}

// Floor returns the largest key in the Patricia Trie less than or equal to key.
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

// Ceiling returns the smallest key in the Patricia Trie greater than or equal to key.
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

// DeleteMin removes the smallest key and associated value from the Patricia Trie.
func (t *patricia[V]) DeleteMin() (string, V, bool) {
	if t.root == nil {
		var zeroV V
		return "", zeroV, false
	}

	// Find the node to delete (z) along side its two preceding nodes (x and y).
	var rp, r, n *patriciaNode[V]
	for rp, r, n = t.root, t.root, t.root.left; r.bp < n.bp; {
		rp, r, n = r, n, n.left
	}

	// Find the node to delete (q) along side its parent node (p).
	var np, m *patriciaNode[V]
	for np, m = t.root, t.root.left; m != n; {
		np, m = m, m.left
	}

	t.remove(n, r, rp, np)

	return n.key.String(), n.val, true
}

// DeleteMax removes the largest key and associated value from the Patricia Trie.
func (t *patricia[V]) DeleteMax() (string, V, bool) {
	if t.root == nil {
		var zeroV V
		return "", zeroV, false
	}

	// Find the node to delete (z) along side its two preceding nodes (x and y).
	var rp, r, n *patriciaNode[V]
	for rp, r, n = t.root, t.root, t.root.left; r.bp < n.bp; {
		rp, r, n = r, n, n.right
	}

	// Find the node to delete (q) along side its parent node (p).
	var np, m *patriciaNode[V]
	for np, m = t.root, t.root.left; m != n; {
		np, m = m, m.right
	}

	t.remove(n, r, rp, np)

	return n.key.String(), n.val, true
}

// Select returns the k-th smallest key in the Patricia Trie.
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

// Rank returns the number of keys in the Patricia Trie less than key.
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

// RangeSize returns the number of keys in the Patricia Trie between two given keys.
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

// Range returns all keys and associated values in the Patricia Trie between two given keys.
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

// Traverse is used for visiting all key-value pairs in the Patricia Trie.
func (t *patricia[V]) Traverse(order TraversalOrder, visit VisitFunc[V]) {
	t._traverse(t.root, order, func(n *patriciaNode[V]) bool {
		return visit(n.key.String(), n.val)
	})
}

func (t *patricia[V]) _traverse(n *patriciaNode[V], order TraversalOrder, visit func(*patriciaNode[V]) bool) bool {
	if n == nil {
		return true
	}

	isLeftThread := n.left.bp <= n.bp                  // left links are never nil
	isRightThread := n != t.root && n.right.bp <= n.bp // Only the root node has a nil right

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
		return (!isLeftThread || visit(n.left)) && // visit the left child only if the left link is threaded (leaf node)
			(isLeftThread || t._traverse(n.left, order, visit)) && // visit the left sub-tree if the left link is not threaded (internal node)
			(!isRightThread || visit(n.right)) && // visit the right child only if the right link is threaded (leaf node)
			(isRightThread || t._traverse(n.right, order, visit)) // visit the right sub-tree if the right link is not threaded (internal node)

	case Descending:
		return (!isRightThread || visit(n.right)) && // visit the right child only if the right link is threaded (leaf node)
			(isRightThread || t._traverse(n.right, order, visit)) && // visit the right sub-tree if the right link is not threaded (internal node)
			(!isLeftThread || visit(n.left)) && // visit the left child only if the left link is threaded (leaf node)
			(isLeftThread || t._traverse(n.left, order, visit)) // visit the left sub-tree if the left link is not threaded (internal node)

	default:
		return false
	}
}

// Graphviz returns a visualization of the Patricia Trie in Graphviz format.
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
							graphviz.NewSimpleField("", fmt.Sprintf("%d", n.bp)),
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

		if n.left.bp > n.bp {
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

			if n.right.bp > n.bp {
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

// Match returns all the keys and associated values in the Patricia Trie
// that match the given pattern in which * matches any character.
func (t *patricia[V]) Match(pattern string) []KeyValue[V] {
	kvs := []KeyValue[V]{}
	t._match(t.root, t.root.left, newBitPattern(pattern), func(n *patriciaNode[V]) {
		kvs = append(kvs, KeyValue[V]{n.key.String(), n.val})
	})

	return kvs
}

func (t *patricia[V]) _match(prev, curr *patriciaNode[V], pattern *bitPattern, visit func(n *patriciaNode[V])) {
	if prev.bp >= curr.bp {
		if curr.key.Len() == pattern.Len() {
			visit(curr)
		}
		return
	}

	switch pattern.Bit(curr.bp) {
	case '0':
		t._match(curr, curr.left, pattern, visit)
	case '1':
		t._match(curr, curr.right, pattern, visit)
	case '*':
		t._match(curr, curr.left, pattern, visit)
		t._match(curr, curr.right, pattern, visit)
	}
}

// WithPrefix returns all the keys and associated values in the Patricia Trie with the given prefix.
func (t *patricia[V]) WithPrefix(key string) []KeyValue[V] {
	kvs := []KeyValue[V]{}
	bitKey := newBitString(key)

	if n := t.search(bitKey); n != nil && n.key.Equals(bitKey) {
		kvs = append(kvs, KeyValue[V]{n.key.String(), n.val})
	} else {
		t._traverse(n, Ascending, func(n *patriciaNode[V]) bool {
			kvs = append(kvs, KeyValue[V]{n.key.String(), n.val})
			return true
		})
	}

	return kvs
}

// LongestPrefix returns the key and associated value in the Patricia Trie
// that is the longest prefix of the given key.
func (t *patricia[V]) LongestPrefixOf(key string) (string, V, bool) {
	bitKey := newBitString(key)
	if n := t.search(bitKey); n != nil && bitKey.HasPrefix(n.key) {
		return n.key.String(), n.val, true
	}

	var zeroV V
	return "", zeroV, false
}
