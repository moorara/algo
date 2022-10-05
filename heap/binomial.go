package heap

import (
	"github.com/moorara/algo/generic"
	"github.com/moorara/algo/internal/graphviz"
)

type binomialNode[K, V any] struct {
	key            K
	val            V
	order          int // order of the tree rooted at this node
	child, sibling *binomialNode[K, V]
}

// TODO:
//
// link constructs a Binomial tree of order k+1 from two Binomial trees of order k.
// It does that by attaching one of the root nodes as the left-most child of the other root node.
//
// It accepts two root nodes and assumes
// the key of the first root comes after (greater for min heap and smaller for max) the key of the second root.
// The second root then becomes the new root
func link[K, V any](r1, r2 *binomialNode[K, V]) {
	r1.sibling = r2.child
	r2.child = r1
	r2.order++
}

// TODO:
//
// merge performs a merge sort on two root lists.
// There can be up to two Binomial trees of the same order in the merged root list.
func merge[K, V any](h, x, y *binomialNode[K, V]) *binomialNode[K, V] {
	if x == nil && y == nil {
		return h
	} else if x == nil {
		h.sibling = merge(y, nil, y.sibling)
	} else if y == nil {
		h.sibling = merge(x, x.sibling, nil)
	} else if x.order < y.order {
		h.sibling = merge(x, x.sibling, y)
	} else {
		h.sibling = merge(y, x, y.sibling)
	}

	return h
}

// binomial implements a Binomial heap.
type binomial[K, V any] struct {
	cmpKey generic.CompareFunc[K]
	eqVal  generic.EqualFunc[V]

	root *binomialNode[K, V]
	size int
}

// NewBinomial creates a new Binomial heap that can be used as a priority queue.
//
// Binomial heap is an implementation of the mergeable heap ADT, which is a priority queue supporting merge operation.
// A Binomial heap is implemented as a set of Binomial trees that satisfy the Binomial heap properties.
//
// A Binomial tree of order 0 is a single node.
// A Binomial tree of order k has a root node whose children are roots of Binomial trees of orders k-1, k-2, ..., 1, 0 respectively.
//
//	The number inside each node denotes the order of the Binomial sub-tree rooted at that node.
//
//	[  0 ]      [  1 ]                  [ 2  ]                                          [ 3  ]
//	               │           ┌──────────┤                        ┌───────────┬──────────┤
//	            [  0 ]      [  1 ]      [ 0  ]                  [  2 ]      [  1 ]      [ 0  ]
//	                           │                       ┌───────────┤           │
//	                        [  0 ]                  [  1 ]      [  0 ]      [  0 ]
//	                                                   │
//	                                                [  0 ]
//
// Each Binomial tree in a heap follows the heap property:
// The key of a node is greater than or equal to the key of its parent.
// There can be at most one Binomial tree for each order, including zero order.
//
//	Examples of minimum Binomal heaps:
//
//	[  9 ]──────[  1 ]──────────────────[  3 ]
//	               │           ┌───────────┤
//	            [ 10 ]      [  5 ]      [  4 ]
//	                           │
//	                        [ 12 ]
//
//	[  4 ]──────────────────────────────────────────[  2 ]
//	                           ┌───────────┬───────────┤
//	                        [  3 ]      [  5 ]      [  8 ]
//	               ┌───────────│           │
//	            [  6 ]      [ 12 ]      [ 11 ]
//	               │
//	            [ 14 ]
//
// A Binomial tree of order k has k children, 2^k nodes, and height k.
// A Binomial tree of order k has k has C(k, d) nodes at depth d, a Binomial coefficient.
// A Binomial tree of order k can be constructed from two trees of order k-1
// by attaching one of them as the left-most child of the root of the other tree.
//
// We use the Left-Child-Right-Sibling (LCRS) representation for implementing a multi-way Binomial tree.
// A Binomial heap consists of a list of root nodes (root list) of Binomial trees sorted by the order of Binomial trees.
//
//	Examples of Binomial heapd in LCRS representation:
//
//	[  9 ]------[  1 ]------[  3 ]
//	               │           │
//	            [ 10 ]      [  5 ]------[  4 ]
//	                           │
//	                        [ 12 ]
//
//	[  4 ]------[  2 ]
//	               │
//	            [  3 ]------------------[  5 ]------[  8 ]
//	               │                       │
//	            [  6 ]------[ 12 ]      [ 11 ]
//	               │
//	            [ 14 ]
//
// cmpKey is a function for comparing two keys.
// eqVal is a function for checking the equality of two values.
func NewBinomial[K, V any](cmpKey generic.CompareFunc[K], eqVal generic.EqualFunc[V]) Heap[K, V] {
	return &binomial[K, V]{
		cmpKey: cmpKey,
		eqVal:  eqVal,

		root: nil,
	}
}

// TODO:
//
// union merges another Binomial heap with the current Binomial heap.
func (h *binomial[K, V]) union(H *binomial[K, V]) *binomial[K, V] {
	if H == nil {
		panic("Binomial heap is nil")
	}

	n := new(binomialNode[K, V])
	h.root = merge(n, h.root, H.root).sibling

	var prev, curr, next *binomialNode[K, V]
	for prev, curr, next = nil, h.root, h.root.sibling; next != nil; next = curr.sibling {
		if curr.order < next.order || (next.sibling != nil && next.sibling.order == curr.order) {
			prev, curr = curr, next
		} else if h.cmpKey(next.key, curr.key) > 0 {
			curr.sibling = next.sibling
			link(next, curr)
		} else {
			if prev == nil {
				h.root = next
			} else {
				prev.sibling = next
			}

			link(curr, next)
			curr = next
		}
	}

	return h
}

// Size returns the number of items on the heap.
func (h *binomial[K, V]) Size() int {
	return h.size
}

// IsEmpty returns true if the heap is empty.
func (h *binomial[K, V]) IsEmpty() bool {
	return h.root == nil
}

// TODO:
//
// Insert adds a new key-value pair to the heap.
func (h *binomial[K, V]) Insert(key K, val V) {
	H := &binomial[K, V]{
		cmpKey: h.cmpKey,
		eqVal:  h.eqVal,
		root: &binomialNode[K, V]{
			key:   key,
			val:   val,
			order: 0,
		},
		size: 1,
	}

	h.root = h.union(H).root
	h.size++
}

// TODO:
//
// Delete removes the extremum (minimum or maximum) key with its value on the heap.
// If the heap is empty, the second return value will be false.
func (h *binomial[K, V]) Delete() (K, V, bool) {
	if h.root == nil {
		var zeroK K
		var zeroV V
		return zeroK, zeroV, false
	}

	// Remove the extremum (minimum or maximum) key with its value on the heap
	var ext, prev, curr *binomialNode[K, V]
	for ext, prev, curr = h.root, nil, h.root; curr.sibling != nil; curr = curr.sibling {
		if h.cmpKey(ext.key, curr.sibling.key) > 0 {
			ext, prev = curr.sibling, curr
		}
	}

	prev.sibling = ext.sibling
	if ext == h.root {
		h.root = ext.sibling
	}

	h.size--

	//
	var n *binomialNode[K, V]
	if ext.child == nil {
		n = ext
	} else {
		n = ext.child
	}

	if ext.child != nil {
		ext.child = nil

		var prev, next *binomialNode[K, V]
		for prev, next = nil, n.sibling; next != nil; next = next.sibling {
			n.sibling = prev
			prev = n
			n = next
		}

		n.sibling = prev
		H := &binomial[K, V]{
			cmpKey: h.cmpKey,
			eqVal:  h.eqVal,
			root:   n,
			size:   1,
		}

		h.root = h.union(H).root
	}

	return ext.key, ext.val, false
}

// TODO:
//
// Peek returns the extremum (minimum or maximum) key with its value on the heap without removing it.
// If the heap is empty, the second return value will be false.
func (h *binomial[K, V]) Peek() (K, V, bool) {
	if h.root == nil {
		var zeroK K
		var zeroV V
		return zeroK, zeroV, false
	}

	var ext, curr *binomialNode[K, V]
	for ext, curr = h.root, h.root; curr.sibling != nil; {
		if h.cmpKey(ext.key, curr.sibling.key) > 0 {
			ext = curr
		}
		curr = curr.sibling
	}

	return ext.key, ext.val, true
}

// ContainsKey returns true if the given key is on the heap.
func (h *binomial[K, V]) ContainsKey(key K) bool {
	// TODO:
	return false
}

// ContainsValue returns true if the given value is on the heap.
func (h *binomial[K, V]) ContainsValue(val V) bool {
	// TODO:
	return false
}

// Graphviz returns a visualization of ... in Graphviz format.
func (h *binomial[K, V]) Graphviz() string {
	graph := graphviz.NewGraph(true, true, false, "Binomial Heap", "", "", "", graphviz.ShapeOval)

	// TODO:

	return graph.DotCode()
}
