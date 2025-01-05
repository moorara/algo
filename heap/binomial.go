package heap

import (
	"fmt"

	. "github.com/moorara/algo/generic"
	"github.com/moorara/algo/internal/graphviz"
)

// The Left-Child-Right-Sibling (LCRS) representation, a.k.a. the Binary Representation of N-ary Tree,
// is used for implementing the binomial tree.
//
//   - The left child of a node points to its first child.
//   - The right sibling of a node points to its next sibling.
//
// A binomial heap consists of a list of root nodes (root list) of binomial trees
// sorted by the increasing order of binomial trees.
//
//   - Each root node represents the root of a binomial tree in the heap.
//   - The root nodes of binomial trees are linked together using the right sibling pointers.
//   - The children of each node, including root nodes, are linked using the left child and right sibling pointers.
//
// Examples of binomial heaps in LCRS representation:
//
//	[ 9 ]──────[ 1 ]──────────[ 3 ]                    [ 9 ]----[ 1 ]----[ 3 ]
//	             │         ┌────┴────┐                            │        │
//	           [ 10 ]    [ 5 ]     [ 4 ]       --->             [ 10 ]   [ 5 ]----[ 4 ]
//	                       │                                               │
//	                     [ 12 ]                                          [ 12 ]
//
//	[ 4 ]────────────────[ 2 ]                         [ 4 ]----[ 2 ]
//	         ┌─────────────┼─────────────┐                        │
//	       [ 3 ]         [ 5 ]        [ 8 ]                     [ 3 ]-------------[ 5 ]----[ 8 ]
//	    ┌────┴────┐        │                   --->               │                 │
//	  [ 6 ]      [ 12 ]  [ 11 ]                                 [ 6 ]----[ 12 ]   [ 11 ]
//	    │                                                         │
//	  [ 14 ]                                                    [ 14 ]
//
// The LCRS representation uses two pointers per node regardless of the number of children.
type binomialNode[K, V any] struct {
	key            K
	val            V
	order          int // order of the binomial tree rooted at this node
	child, sibling *binomialNode[K, V]
}

// binomial implements a binomial heap.
// A binomial heap is implemented as a list of root nodes (root list) of binomial trees
// sorted by the increasing order of binomial trees.
type binomial[K, V any] struct {
	cmpKey CompareFunc[K]
	eqVal  EqualFunc[V]

	n    int                 // number of items on heap
	head *binomialNode[K, V] // head of the root list
}

// NewBinomial creates a new binomial heap that can be used as a mergeable priority queue.
//
// Binomial heap is an implementation of the mergeable heap ADT, a priority queue supporting merge operation.
// A binomial heap is implemented as a set of binomial trees that satisfy the binomial heap properties:
//
//   - Heap Property: Every binomial tree in a binomial heap satisfies the min-heap or max-heap property.
//   - Structural Property: The heap contains at most one binomial tree of any given order.
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
//	[ 0 ]    [ 1 ]          [ 2 ]                          [ 3 ]
//	           │         ┌────┴────┐             ┌───────────┼───────────┐
//	         [ 0 ]    [ 1 ]      [ 0 ]         [ 2 ]       [ 1 ]       [ 0 ]
//	                    │                   ┌────┴────┐      │
//	                  [ 0 ]               [ 1 ]     [ 0 ]  [ 0 ]
//	                                        │
//	                                      [ 0 ]
//
// Examples of minimum binomial heaps:
//
//	[ 9 ]────[ 1 ]────────[ 3 ]
//	           │       ┌────┴────┐
//	         [ 10 ]  [ 5 ]     [ 4 ]
//	                   │
//	                 [ 12 ]
//
//	[ 4 ]───────────────────[ 2 ]
//	             ┌────────────┼────────────┐
//	           [ 3 ]        [ 5 ]        [ 8 ]
//	        ┌────┴────┐       │
//	      [ 6 ]     [ 12 ]  [ 11 ]
//	        │
//	      [ 14 ]
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
		head:   nil,
	}
}

// findExtremum traverses the sibling linked list starting from the given node n
// and finds the extremum (minimum in min heap or maximum in max heap) node.
//
// It returns the predecessor of the extremum node and the extremum node itself.
// If the list is empty, it returns nil for both values.
func (h *binomial[K, V]) findExtremum(n *binomialNode[K, V]) (*binomialNode[K, V], *binomialNode[K, V]) {
	if n == nil {
		return nil, nil
	}

	var prev, ext, curr *binomialNode[K, V]
	for prev, ext, curr = nil, n, n; curr.sibling != nil; curr = curr.sibling {
		if h.cmpKey(curr.sibling.key, ext.key) < 0 {
			prev, ext = curr, curr.sibling
		}
	}

	return prev, ext
}

// link constructs a binomial tree of order k+1 from two binomial trees of order k.
// It attaches one of the root nodes as the left-most child of the other.
//
// It accepts two root nodes and assumes the key of the first root comes after
// the key of the second root (greater in min heap or smaller in max heap).
// The second root then becomes the new root.
//
// General tree representation:
//
//	     [ 3 ]                   [ 2 ]                                  [ 2 ]
//	  ┌────┴────┐             ┌────┴────┐                     ┌───────────┼───────────┐
//	[ 7 ]     [ 5 ]    +    [ 6 ]     [ 4 ]    --->         [ 3 ]       [ 6 ]       [ 4 ]
//	  │                       │                          ┌────┴────┐      │
//	[ 9 ]                   [ 8 ]                      [ 7 ]     [ 5 ]  [ 8 ]
//	                                                     │
//	                                                   [ 8 ]
//
// LCRS representation:
//
//	[ 3 ]                  [ 2 ]                     [ 2 ]
//	  │                      │                         │
//	[ 7 ]----[ 5 ]    +    [ 6 ]----[ 4 ]    --->    [ 3 ]-----------[ 6 ]----[ 4 ]
//	  │                      │                         │               │
//	[ 9 ]                  [ 8 ]                     [ 7 ]----[ 5 ]  [ 8 ]
//	                                                   │
//	                                                 [ 8 ]
//
// This method is defined on the binomial struct to prevent name clashes
// with other similar implementations in this package.
func (_ *binomial[K, V]) link(child, parent *binomialNode[K, V]) {
	child.sibling = parent.child
	parent.child = child
	parent.order++
}

// merge performs a merge sort on two root lists of binomial trees.
//
// The resulting root list is sorted in monotonically increasing order of binomial tree orders,
// and it may contain up to two binomial trees of the same order.
// These trees are combined later during the consolidate operation, which follows this merge operation.
//
// Below are examples illustrating all possible cases when merging two root lists.
// The number inside each node denotes the order of the binomial tree rooted at that node.
//
// Case 1: h1 = nil, h2 = nil
//
//	nil    +    nil    =    nil
//
// Case 2: h1 = nil, h2 ≠ nil
//
//	nil    +    [ 0 ]----[ 1 ]    =    [ 0 ]----[ 1 ]
//	                       │                      │
//	                     [ 0 ]                  [ 0 ]
//
// Case 3: h1 ≠ nil, h2 = nil
//
//	[ 0 ]----[ 1 ]    +    nil    =    [ 0 ]----[ 1 ]
//	           │                                  │
//	         [ 0 ]                              [ 0 ]
//
// Case 4: h1.order < h2.order
//
//	[ 0 ]----[ 2 ]                  [ 1 ]----[ 3 ]                                  [ 0 ]----[ 1 ]----[ 2 ]-------------[ 3 ]
//	           │                      │        │                                               │        │                 │
//	         [ 1 ]----[ 0 ]    +    [ 0 ]    [ 2 ]-----------[ 1 ]----[ 0 ]    =             [ 0 ]    [ 1 ]----[ 0 ]    [ 2 ]-----------[ 1 ]----[ 0 ]
//	           │                               │               │                                        │                 │               │
//	         [ 0 ]                           [ 1 ]----[ 0 ]  [ 0 ]                                    [ 0 ]             [ 1 ]----[ 0 ]  [ 0 ]
//	                                           │                                                                          │
//	                                         [ 0 ]                                                                      [ 0 ]
//
// Case 5: h1.order ≥ h2.order
//
//	[ 1 ]----[ 3 ]                                  [ 0 ]----[ 2 ]                  [ 0 ]----[ 1 ]----[ 2 ]-------------[ 3 ]
//	  │        │                                               │                               │        │                 │
//	[ 0 ]    [ 2 ]-----------[ 1 ]----[ 0 ]    +             [ 1 ]----[ 0 ]    +             [ 0 ]    [ 1 ]----[ 0 ]    [ 2 ]-----------[ 1 ]----[ 0 ]
//	           │               │                               │                                        │                 │               │
//	         [ 1 ]----[ 0 ]  [ 0 ]                           [ 0 ]                                    [ 0 ]             [ 1 ]----[ 0 ]  [ 0 ]
//	           │                                                                                                          │
//	         [ 0 ]                                                                                                      [ 0 ]
//
// This method is defined on the binomial struct to prevent name clashes
// with other similar implementations in this package.
func (h *binomial[K, V]) merge(h1, h2 *binomialNode[K, V]) *binomialNode[K, V] {
	switch {
	case h1 == nil:
		return h2
	case h2 == nil:
		return h1
	case h1.order < h2.order:
		h1.sibling = h.merge(h1.sibling, h2)
		return h1
	default:
		h2.sibling = h.merge(h1, h2.sibling)
		return h2
	}
}

// consolidate consolidates the root list of a binomial heap by
// iterating through the root list and merging binomial trees of the same order.
// It ensures after consolidation, no two binomial trees in the root list have the same order.
func (h *binomial[K, V]) consolidate(head *binomialNode[K, V]) *binomialNode[K, V] {
	if head == nil {
		return nil
	}

	// Scanning the root list of binomial trees.
	var prev, curr, next *binomialNode[K, V]
	for prev, curr, next = nil, head, head.sibling; next != nil; next = curr.sibling {
		if curr.order != next.order || (next.sibling != nil && next.sibling.order == curr.order) {
			/*
			 * Case 1: The current and the next binomal trees have different orders.
			 *         There is no conflict in this case and we proceed with scanning.
			 *
			 * Case 2: The current and the next binomal trees have the same order,
			 *         but the next has a sibling with the same order as the current.
			 *
			 *         This case can occur when there are two binmal trees of order k-1
			 *           followed by two binomia tress of order k in the root list.
			 *         After merging the two binomial trees of order k-1,
			 *           we end up with three binomial trees of order k.
			 *         If we merge the first two binomial trees of order k immediately,
			 *           it would result in a binomail tree of order k+1 positioned before a binomial tree of order k.
			 *         This violates the structural properties of the binomial heap.
			 *         Instead, we defer merging of the two binomial trees of order k to avoid creating additional conflicts.
			 */
			prev, curr = curr, next
		} else if h.cmpKey(next.key, curr.key) > 0 {
			// Case 3: The next binomial tree should become the child of current one.
			curr.sibling = next.sibling
			h.link(next, curr)
		} else {
			// Case 4: The current binomial tree should become the child of next one.
			if prev == nil {
				head = next
			} else {
				prev.sibling = next
			}
			h.link(curr, next)
			curr = next
		}
	}

	return head
}

// union merges two binomial heaps into a single binomial heap.
func (h *binomial[K, V]) union(h1, h2 *binomialNode[K, V]) *binomialNode[K, V] {
	head := h.merge(h1, h2)
	head = h.consolidate(head)
	return head
}

// childrenToRootList reverses the order of a binomial tree node's children,
// converting the child list into a reversed root list.
func (h *binomial[K, V]) childrenToRootList(n *binomialNode[K, V]) *binomialNode[K, V] {
	var prev, curr, next *binomialNode[K, V]
	for prev, curr = nil, n.child; curr != nil; prev, curr = curr, next {
		next = curr.sibling
		curr.sibling = prev
	}

	// prev is now the head of the reversed root list.
	return prev
}

// traverse performs a depth-first traversal of the binomial heap,
// visiting each node according to the specified traversal order.
func (h *binomial[K, V]) traverse(n *binomialNode[K, V], order TraverseOrder, visit func(*binomialNode[K, V]) bool) bool {
	if n == nil {
		return true
	}

	switch order {
	case VLR:
		return visit(n) && h.traverse(n.child, order, visit) && h.traverse(n.sibling, order, visit)
	case VRL:
		return visit(n) && h.traverse(n.sibling, order, visit) && h.traverse(n.child, order, visit)
	case LVR:
		return h.traverse(n.child, order, visit) && visit(n) && h.traverse(n.sibling, order, visit)
	case RVL:
		return h.traverse(n.sibling, order, visit) && visit(n) && h.traverse(n.child, order, visit)
	case LRV:
		return h.traverse(n.child, order, visit) && h.traverse(n.sibling, order, visit) && visit(n)
	case RLV:
		return h.traverse(n.sibling, order, visit) && h.traverse(n.child, order, visit) && visit(n)
	default:
		return false
	}
}

// Size returns the number of items on the heap.
func (h *binomial[K, V]) Size() int {
	return h.n
}

// IsEmpty returns true if the heap is empty.
func (h *binomial[K, V]) IsEmpty() bool {
	return h.head == nil
}

// Insert adds a new key-value pair to the heap.
func (h *binomial[K, V]) Insert(key K, val V) {
	n := &binomialNode[K, V]{
		key:   key,
		val:   val,
		order: 0,
	}

	h.head = h.union(h.head, n)
	h.n++
}

// Merge merges the current heap with another heap.
// The new heap must have the same underlying type as the current one.
func (h *binomial[K, V]) Merge(H MergeableHeap[K, V]) {
	if hh, ok := H.(*binomial[K, V]); ok {
		h.head = h.union(h.head, hh.head)
		h.n += hh.n
	}
}

// Delete removes the extremum (minimum or maximum) key with its value on the heap.
// If the heap is empty, the second return value will be false.
func (h *binomial[K, V]) Delete() (K, V, bool) {
	if h.IsEmpty() {
		var zeroK K
		var zeroV V
		return zeroK, zeroV, false
	}

	// Find the extremum (minimum in min heap or maximum in max heap) node in the root list.
	prev, ext := h.findExtremum(h.head)

	// Remove the extremum node from the root list.
	if prev == nil { // ext == h.head
		h.head = ext.sibling
	} else {
		prev.sibling = ext.sibling
	}

	// Convert the deleted root's children into a root list and merge it with the current heap.
	head := h.childrenToRootList(ext)
	h.head = h.union(h.head, head)
	h.n--

	return ext.key, ext.val, true
}

// DeleteAll deletes all keys with their values on the heap, leaving it empty.
func (h *binomial[K, V]) DeleteAll() {
	h.n = 0
	h.head = nil
}

// Peek returns the extremum (minimum or maximum) key with its value on the heap without removing it.
// If the heap is empty, the second return value will be false.
func (h *binomial[K, V]) Peek() (K, V, bool) {
	if h.IsEmpty() {
		var zeroK K
		var zeroV V
		return zeroK, zeroV, false
	}

	// Find the extremum (minimum in min heap or maximum in max heap) node in the root list.
	_, ext := h.findExtremum(h.head)

	return ext.key, ext.val, true
}

// ContainsKey returns true if the given key is on the heap.
func (h *binomial[K, V]) ContainsKey(key K) bool {
	if h.IsEmpty() {
		return false
	}

	// A false result indicates a match was found.
	return !h.traverse(h.head, VLR, func(n *binomialNode[K, V]) bool {
		// If a match is found, stop the traversal by returning false.
		return h.cmpKey(n.key, key) != 0
	})
}

// ContainsValue returns true if the given value is on the heap.
func (h *binomial[K, V]) ContainsValue(val V) bool {
	if h.IsEmpty() {
		return false
	}

	// A false result indicates a match was found.
	return !h.traverse(h.head, VLR, func(n *binomialNode[K, V]) bool {
		// If a match is found, stop the traversal by returning false.
		return !h.eqVal(n.val, val)
	})
}

// Graphviz returns a visualization of the heap in Graphviz format.
func (h *binomial[K, V]) Graphviz() string {
	// Create a map of node --> id
	var id int
	nodeID := map[*binomialNode[K, V]]int{}
	h.traverse(h.head, VLR, func(n *binomialNode[K, V]) bool {
		id++
		nodeID[n] = id
		return true
	})

	graph := graphviz.NewGraph(true, true, false, "Binomial Heap", "", "", "", graphviz.ShapeMrecord)

	h.traverse(h.head, VLR, func(n *binomialNode[K, V]) bool {
		name := fmt.Sprintf("%d", nodeID[n])

		rec := graphviz.NewRecord(
			graphviz.NewSimpleField("", fmt.Sprintf("%v", n.key)),
			graphviz.NewSimpleField("", fmt.Sprintf("%v", n.val)),
		)

		graph.AddNode(graphviz.NewNode(name, "", rec.Label(), "", "", "", "", ""))

		if n.child != nil {
			child := fmt.Sprintf("%d", nodeID[n.child])
			graph.AddEdge(graphviz.NewEdge(name, child, graphviz.EdgeTypeDirected, "", "", graphviz.ColorBlue, "", "", ""))
		}

		if n.sibling != nil {
			sibling := fmt.Sprintf("%d", nodeID[n.sibling])
			graph.AddEdge(graphviz.NewEdge(name, sibling, graphviz.EdgeTypeDirected, "", "", graphviz.ColorRed, graphviz.StyleDashed, "", ""))
		}

		return true
	})

	return graph.DotCode()
}
