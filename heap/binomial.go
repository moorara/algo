package heap

import (
	"fmt"

	. "github.com/moorara/algo/generic"
	"github.com/moorara/algo/internal/graphviz"
)

// We use the Left-Child-Right-Sibling (LCRS) representation, a.k.a. the
// Binary Tree Representation of General Trees, for implementing the binomial tree.
//
//   - The left child of a node points to its first child.
//   - The right sibling of a node points to its next sibling.
//
// A binomial heap consists of a list of root nodes (root list) of binomial trees sorted by the order of binomial trees.
//
// Examples of binomial heaps in LCRS representation:
//
//	[ 9 ]------[ 1 ]------[ 3 ]
//	             │          │
//	           [ 10 ]     [ 5 ]------[ 4 ]
//	                        │
//	                      [ 12 ]
//
//	[ 4 ]------[ 2 ]
//	             │
//	           [ 3 ]------------------[ 5 ]------[ 8 ]
//	             │                      │
//	           [ 6 ]------[ 12 ]      [ 11 ]
//	             │
//	           [ 14 ]
//
// The LCRS representation uses two pointers per node regardless of the number of children.
type binomialNode[K, V any] struct {
	key            K
	val            V
	order          int // order of the tree rooted at this node
	child, sibling *binomialNode[K, V]
}

// binomial implements a binomial heap tree.
type binomial[K, V any] struct {
	cmpKey CompareFunc[K]
	eqVal  EqualFunc[V]

	n    int                 // number of items on heap
	root *binomialNode[K, V] // root of the root list
}

// NewBinomial creates a new binomial heap that can be used as a priority queue.
//
// Binomial heap is an implementation of the mergeable heap ADT, a priority queue supporting merge operation.
// A binomial heap is implemented as a set of binomial trees that satisfy the binomial heap properties.
//
// A binomial tree Bₖ of order k is defined recursively:
//
//   - A binomial tree B₀ of order 0 is a single node.
//   - A binomial tree Bₖ of order k is formed by linking two binomial trees of orders k-1 together,
//     making the root of one tree a child of the root of the other tree.
//     Equivalently, a binomial tree Bₖ has a root node whose children are roots of binomial trees of orders k-1, k-2, ..., 1, 0.
//
// The number inside each node denotes the order of the binomial sub-tree rooted at that node.
//
//	[ 0 ]      [ 1 ]                 [ 2 ]                                      [ 3 ]
//	             │           ┌─────────┤                    ┌──────────┬──────────┤
//	           [ 0 ]      [ 1 ]      [ 0 ]                [ 2 ]      [ 1 ]      [ 0 ]
//	                        │                     ┌─────────┤          │
//	                      [ 0 ]                 [ 1 ]     [ 0 ]      [ 0 ]
//	                                              │
//	                                            [ 0 ]
//
// Each binomial tree in a binomial heap satisfies the min (max) heap property:
// the key of a parent node is less (greater) than or equal to the keys of its children.
// Additionally, the binomial trees in a binomial heap are structured such that there is at most one binomial tree of each order.
//
// Examples of minimum binomial heaps:
//
//	[ 9 ]──────[ 1 ]──────────────────[ 3 ]
//	             │           ┌──────────┤
//	           [ 10 ]      [ 5 ]      [ 4 ]
//	                         │
//	                       [ 12 ]
//
//	[ 4 ]──────────────────────────────────────────[ 2 ]
//	                           ┌───────────┬─────────┤
//	                         [ 3 ]      [ 5 ]      [ 8 ]
//	               ┌───────────│          │
//	             [ 6 ]       [ 12 ]     [ 11 ]
//	               │
//	             [ 14 ]
//
// Here are some properties of binomial trees:
//
//   - The height of a Bₖ tree is k.
//   - The number of nodes in a Bₖ tree is 2ᵏ.
//   - The root of a Bₖ tree has k children.
//   - The children of the root of a Bₖ tree are the roots of B₀, B₁, ..., Bₖ₋₁ trees.
//   - A binomial tree Bₖ of order k has C(k, d) nodes at depth d, a binomial coefficient.
//
// cmpKey is a function for comparing two keys.
// eqVal is a function for checking the equality of two values.
func NewBinomial[K, V any](cmpKey CompareFunc[K], eqVal EqualFunc[V]) MergeableHeap[K, V] {
	return &binomial[K, V]{
		cmpKey: cmpKey,
		eqVal:  eqVal,
		n:      0,
		root:   nil,
	}
}

// merge merges another heap with the current heap.
func (h *binomial[K, V]) Merge(H MergeableHeap[K, V]) {
	// TODO:
}

// merge merges another binomial heap with the current binomial heap.
func (h *binomial[K, V]) merge(H *binomial[K, V]) *binomial[K, V] {
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

// merge performs a merge sort on two root lists.
// There can be up to two binomial trees of the same order in the merged root list.
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

// link constructs a binomial tree of order k+1 from two binomial trees of order k.
// It does that by attaching one of the root nodes as the left-most child of the other root node.
//
// It accepts two root nodes and assumes the key of the first root comes after
// (greater for min heap and smaller for max) the key of the second root.
// The second root then becomes the new root.
func link[K, V any](r1, r2 *binomialNode[K, V]) {
	r1.sibling = r2.child
	r2.child = r1
	r2.order++
}

// Size returns the number of items on the heap.
func (h *binomial[K, V]) Size() int {
	return h.n
}

// IsEmpty returns true if the heap is empty.
func (h *binomial[K, V]) IsEmpty() bool {
	return h.root == nil
}

// Insert adds a new key-value pair to the heap.
func (h *binomial[K, V]) Insert(key K, val V) {
	H := &binomial[K, V]{
		cmpKey: h.cmpKey,
		eqVal:  h.eqVal,
		n:      1,
		root: &binomialNode[K, V]{
			key:   key,
			val:   val,
			order: 0,
		},
	}

	h.n++
	h.root = h.merge(H).root
}

// Delete removes the extremum (minimum or maximum) key with its value on the heap.
// If the heap is empty, the second return value will be false.
func (h *binomial[K, V]) Delete() (K, V, bool) {
	if h.IsEmpty() {
		var zeroK K
		var zeroV V
		return zeroK, zeroV, false
	}

	var ext, prev, curr *binomialNode[K, V]
	for ext, prev, curr = h.root, nil, h.root; curr.sibling != nil; curr = curr.sibling {
		if h.cmpKey(ext.key, curr.sibling.key) > 0 {
			ext = curr.sibling
			prev = curr
		}
	}

	prev.sibling = ext.sibling
	if ext == h.root {
		h.root = ext.sibling
	}

	h.n--

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
			n:      1,
			root:   n,
		}

		h.root = h.merge(H).root
	}

	return ext.key, ext.val, false
}

// DeleteAll deletes all keys with their values on the heap, leaving it empty.
func (h *binomial[K, V]) DeleteAll() {
	h.n = 0
	h.root = nil
}

// Peek returns the extremum (minimum or maximum) key with its value on the heap without removing it.
// If the heap is empty, the second return value will be false.
func (h *binomial[K, V]) Peek() (K, V, bool) {
	if h.IsEmpty() {
		var zeroK K
		var zeroV V
		return zeroK, zeroV, false
	}

	var ext, curr *binomialNode[K, V]
	for ext, curr = h.root, h.root; curr.sibling != nil; curr = curr.sibling {
		if h.cmpKey(ext.key, curr.sibling.key) > 0 {
			ext = curr
		}
	}

	return ext.key, ext.val, true
}

// ContainsKey returns true if the given key is on the heap.
func (h *binomial[K, V]) ContainsKey(key K) bool {
	if h.IsEmpty() {
		return false
	}

	return h._traverse(h.root, VLR, func(n *binomialNode[K, V]) bool {
		return h.cmpKey(n.key, key) != 0
	})
}

// ContainsValue returns true if the given value is on the heap.
func (h *binomial[K, V]) ContainsValue(val V) bool {
	if h.IsEmpty() {
		return false
	}

	return h._traverse(h.root, VLR, func(n *binomialNode[K, V]) bool {
		return !h.eqVal(n.val, val)
	})
}

// Graphviz returns a visualization of the heap in Graphviz format.
func (h *binomial[K, V]) Graphviz() string {
	// Create a map of node --> id
	var id int
	nodeID := map[*binomialNode[K, V]]int{}
	h._traverse(h.root, VLR, func(n *binomialNode[K, V]) bool {
		id++
		nodeID[n] = id
		return true
	})

	graph := graphviz.NewGraph(true, true, false, "Binomial Heap", "", "", "", graphviz.ShapeMrecord)

	h._traverse(h.root, VLR, func(n *binomialNode[K, V]) bool {
		name := fmt.Sprintf("%d", nodeID[n])

		rec := graphviz.NewRecord(
			graphviz.NewSimpleField("", fmt.Sprintf("%v", n.key)),
			graphviz.NewSimpleField("", fmt.Sprintf("%v", n.val)),
		)

		graph.AddNode(graphviz.NewNode(name, "", rec.Label(), "", "", "", "", ""))

		if n.child != nil {
			child := fmt.Sprintf("%d", nodeID[n.child])
			graph.AddEdge(graphviz.NewEdge(name, child, graphviz.EdgeTypeDirected, "", "", "", "", "", ""))
		}

		if n.sibling != nil {
			sibling := fmt.Sprintf("%d", nodeID[n.sibling])
			graph.AddEdge(graphviz.NewEdge(name, sibling, graphviz.EdgeTypeDirected, "", "", "", graphviz.StyleDashed, "", ""))
		}

		return true
	})

	return graph.DotCode()
}

func (h *binomial[K, V]) _traverse(n *binomialNode[K, V], order TraverseOrder, visit func(*binomialNode[K, V]) bool) bool {
	if n == nil {
		return true
	}

	switch order {
	case VLR:
		return visit(n) && h._traverse(n.child, order, visit) && h._traverse(n.sibling, order, visit)
	case VRL:
		return visit(n) && h._traverse(n.sibling, order, visit) && h._traverse(n.child, order, visit)
	case LVR:
		return h._traverse(n.child, order, visit) && visit(n) && h._traverse(n.sibling, order, visit)
	case RVL:
		return h._traverse(n.sibling, order, visit) && visit(n) && h._traverse(n.child, order, visit)
	case LRV:
		return h._traverse(n.child, order, visit) && h._traverse(n.sibling, order, visit) && visit(n)
	case RLV:
		return h._traverse(n.sibling, order, visit) && h._traverse(n.child, order, visit) && visit(n)
	default:
		return false
	}
}
