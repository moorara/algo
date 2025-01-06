package heap

import (
	"fmt"

	. "github.com/moorara/algo/generic"
	"github.com/moorara/algo/internal/dot"
)

// The Left-Child-Right-Sibling (LCRS) representation, a.k.a. the Binary Representation of N-ary Tree,
// is used for implementing the indexed binomial tree.
//
//   - The left child of a node points to its first child.
//   - The right sibling of a node points to its next sibling.
//
// The LCRS representation uses two pointers per node regardless of the number of children.
type indexedBinomialNode[K, V any] struct {
	index                  int // index associated with the key
	key                    K
	val                    V
	order                  int // order of the tree rooted at this node
	parent, child, sibling *indexedBinomialNode[K, V]
}

// indexedBinomial implements an indexed binomial heap.
// An indexed binomial heap is implemented as a list of root nodes (root list) of binomial trees
// sorted by the increasing order of binomial trees as well as a mapping between indices and nodes.
type indexedBinomial[K, V any] struct {
	cmpKey CompareFunc[K]
	eqVal  EqualFunc[V]

	n     int                          // current number of items on heap
	head  *indexedBinomialNode[K, V]   // head of the root list
	nodes []*indexedBinomialNode[K, V] // map of indices to nodes
}

// NewIndexedBinomial creates a new indexed binomial heap that can be used as a priority queue.
//
// An indexed heap (priority queue) associates an index with each key-value pair.
// It allows changing the key (priority) of an index, deleting by index, and looking up by index.
// The size of an indexed binary heap is fixed.
//
// The indexed binomial heap does not support the merge operation,
// as doing so would cause conflicts between indices.
//
// cap is the maximum number of items on the heap.
// cmpKey is a function for comparing two keys.
// eqVal is a function for checking the equality of two values.
func NewIndexedBinomial[K, V any](cap int, cmpKey CompareFunc[K], eqVal EqualFunc[V]) IndexedHeap[K, V] {
	return &indexedBinomial[K, V]{
		cmpKey: cmpKey,
		eqVal:  eqVal,
		n:      0,
		head:   nil,
		nodes:  make([]*indexedBinomialNode[K, V], cap),
	}
}

// This method verifies the integrity of an indexed binomial heap.
func (h *indexedBinomial[K, V]) verify() bool {
	if h.head == nil {
		return true
	}

	// Verify the structural property:
	// Ensure the root list is sorted in monotonically increasing order of binomial tree orders.
	for curr, next := h.head, h.head.sibling; next != nil; curr, next = next, next.sibling {
		if curr.order >= next.order {
			return false
		}
	}

	// Verify the properties of each binomial tree in the root list.
	for curr := h.head; curr != nil; curr = curr.sibling {
		if !h.verifyBinomialTree(curr) {
			return false
		}
	}

	return true
}

// verifyBinomialTree verifies the properties of a binomial tree rooted at the given node.
func (h *indexedBinomial[K, V]) verifyBinomialTree(n *indexedBinomialNode[K, V]) bool {
	// Verifry the index map for the current node.
	if h.nodes[n.index] != n {
		return false
	}

	for i, curr := 1, n.child; curr != nil; i, curr = i+1, curr.sibling {
		// In a min-heap, each node's key must be smaller than or equal to its children's keys.
		// In a max-heap, each node's key must be greater than or equal to its children's keys.
		if h.cmpKey(n.key, curr.key) > 0 {
			return false
		}

		// A binomial node of order k has children with orders k-1, k-2, ..., 0 from left to right.
		if curr.order != n.order-i {
			return false
		}

		// Recursively, verify each binomial tree root at the current child.
		if !h.verifyBinomialTree(curr) {
			return false
		}
	}

	return true
}

// swap exchanges the indices, keys, and values between the given child and parent nodes.
// It also updates the index map to reflect the changes.
// This operation preserves the structural relationships within the heap.
//
// This method is defined on the indexedBinomial struct to prevent name clashes
// with other similar implementations in this package.
func (h *indexedBinomial[K, V]) swap(child, parent *indexedBinomialNode[K, V]) {
	// Swap the indices, keys, and values between the current node and its parent,
	// while preserving the parent-child and sibling relationships.
	child.index, parent.index = parent.index, child.index
	child.key, parent.key = parent.key, child.key
	child.val, parent.val = parent.val, child.val

	// Update the index map (swapping the references).
	h.nodes[child.index], h.nodes[parent.index] = h.nodes[parent.index], h.nodes[child.index]
}

// promote promotes a node n in the binomial tree by repeatedly swapping it with its parent
// while its key is smaller than the parent's key, respecting the heap property.
// It stops when the node is in the correct position and the heap property is restored.
func (h *indexedBinomial[K, V]) promote(n *indexedBinomialNode[K, V]) {
	for n.parent != nil && h.cmpKey(n.parent.key, n.key) > 0 {
		h.swap(n, n.parent)
		n = n.parent
	}
}

// demote demotes a node n in the binomial tree by repeatedly swapping it with the extremum of its children
// while the child's key is smaller than the node's key, respecting the heap property.
// It stops when the node is in the correct position and the heap property is restored.
func (h *indexedBinomial[K, V]) demote(n *indexedBinomialNode[K, V]) {
	_, child := h.findExtremum(n.child)
	for child != nil && h.cmpKey(child.key, n.key) < 0 {
		h.swap(child, n)
		n = child
		_, child = h.findExtremum(child.child)
	}
}

// findExtremum traverses the sibling linked list starting from the given node n
// and finds the extremum (minimum in min heap or maximum in max heap) node.
//
// It returns the predecessor of the extremum node and the extremum node itself.
// If the list is empty, it returns nil for both values.
func (h *indexedBinomial[K, V]) findExtremum(n *indexedBinomialNode[K, V]) (*indexedBinomialNode[K, V], *indexedBinomialNode[K, V]) {
	if n == nil {
		return nil, nil
	}

	var prev, ext, curr *indexedBinomialNode[K, V]
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
// This method is defined on the indexedBinomial struct to prevent name clashes
// with other similar implementations in this package.
func (_ *indexedBinomial[K, V]) link(child, parent *indexedBinomialNode[K, V]) {
	child.parent = parent
	child.sibling = parent.child
	parent.child = child
	parent.order++
}

// merge performs a merge sort on two root lists of binomial trees.
//
// The resulting root list is sorted in monotonically non-decreasing order of binomial tree orders,
// and it may contain up to two binomial trees of the same order.
// These trees are combined later during the consolidate operation, which follows this merge operation.
//
// This method is defined on the indexedBinomial struct to prevent name clashes
// with other similar implementations in this package.
func (h *indexedBinomial[K, V]) merge(h1, h2 *indexedBinomialNode[K, V]) *indexedBinomialNode[K, V] {
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
func (h *indexedBinomial[K, V]) consolidate(head *indexedBinomialNode[K, V]) *indexedBinomialNode[K, V] {
	if head == nil {
		return nil
	}

	// Scanning the root list of binomial trees.
	var prev, curr, next *indexedBinomialNode[K, V]
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

// union merges two indexed binomial heaps into a single indexed binomial heap.
func (h *indexedBinomial[K, V]) union(h1, h2 *indexedBinomialNode[K, V]) *indexedBinomialNode[K, V] {
	head := h.merge(h1, h2)
	head = h.consolidate(head)
	return head
}

// childrenToRootList reverses the order of a binomial tree node's children,
// converting the child list into a reversed root list.
func (h *indexedBinomial[K, V]) childrenToRootList(n *indexedBinomialNode[K, V]) *indexedBinomialNode[K, V] {
	var prev, curr, next *indexedBinomialNode[K, V]
	for prev, curr = nil, n.child; curr != nil; prev, curr = curr, next {
		next = curr.sibling
		curr.parent = nil
		curr.sibling = prev
	}

	// prev is now the head of the reversed root list.
	return prev
}

// traverse performs a depth-first traversal of the binomial heap,
// visiting each node according to the specified traversal order.
func (h *indexedBinomial[K, V]) traverse(n *indexedBinomialNode[K, V], order TraverseOrder, visit func(*indexedBinomialNode[K, V]) bool) bool {
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
func (h *indexedBinomial[K, V]) Size() int {
	return h.n
}

// IsEmpty returns true if the heap is empty.
func (h *indexedBinomial[K, V]) IsEmpty() bool {
	return h.head == nil
}

// Insert adds a new key-value pair to the heap.
func (h *indexedBinomial[K, V]) Insert(i int, key K, val V) bool {
	// ContainsIndex validates the index too.
	if h.ContainsIndex(i) {
		return false
	}

	n := &indexedBinomialNode[K, V]{
		index: i,
		key:   key,
		val:   val,
		order: 0,
	}

	h.head = h.union(h.head, n)
	h.nodes[i] = n
	h.n++

	return true
}

// ChangeKey changes the key associated with an index.
func (h *indexedBinomial[K, V]) ChangeKey(i int, key K) bool {
	// ContainsIndex validates the index too.
	if !h.ContainsIndex(i) {
		return false
	}

	n := h.nodes[i]
	n.key = key
	h.promote(n)
	h.demote(n)

	return true
}

// Delete removes the extremum (minimum or maximum) key with its value on the heap.
// If the heap is empty, the second return value will be false.
func (h *indexedBinomial[K, V]) Delete() (int, K, V, bool) {
	if h.IsEmpty() {
		var zeroK K
		var zeroV V
		return -1, zeroK, zeroV, false
	}

	// Find the extremum (minimum in min heap or maximum in max heap) node in the root list.
	prev, ext := h.findExtremum(h.head)

	// Remove the extremum node from the root list.
	if prev == nil { // ext == h.head
		h.head = ext.sibling
	} else {
		prev.sibling = ext.sibling
	}

	// Remove the extremum node from the index map.
	h.nodes[ext.index] = nil

	// Convert the deleted root's children into a root list and merge it with the current heap.
	head := h.childrenToRootList(ext)
	h.head = h.union(h.head, head)
	h.n--

	return ext.index, ext.key, ext.val, true
}

// DeleteIndex removes a key-value pair and its associated index from the heap.
// If the index is not valid or not on the heap, the second return value will be false.
func (h *indexedBinomial[K, V]) DeleteIndex(i int) (K, V, bool) {
	// ContainsIndex validates the index too.
	if !h.ContainsIndex(i) {
		var zeroK K
		var zeroV V
		return zeroK, zeroV, false
	}

	// Move the node at index i to the root of its binomial tree.
	// This simulates changing the key of the node to the minimum value.
	// Since the minimum of a generic type K is unknown at implementation time,
	// we achieve this by moving the node up the tree.
	var n *indexedBinomialNode[K, V]
	for n = h.nodes[i]; n.parent != nil; n = n.parent {
		h.swap(n, n.parent)
	}

	// Find the target node in the root list.
	var prev, curr *indexedBinomialNode[K, V]
	for prev, curr = nil, h.head; curr != n; prev, curr = curr, curr.sibling {
	}

	// Remove the target node from the root list.
	if prev == nil { // ext == h.head
		h.head = curr.sibling
	} else {
		prev.sibling = curr.sibling
	}

	// Remove the extremum node from the index map.
	h.nodes[n.index] = nil

	// Convert the deleted root's children into a root list and merge it with the current heap.
	head := h.childrenToRootList(n)
	h.head = h.union(h.head, head)
	h.n--

	return n.key, n.val, true
}

// DeleteAll deletes all keys with their values and indices on the heap, leaving it empty.
func (h *indexedBinomial[K, V]) DeleteAll() {
	h.n = 0
	h.head = nil
	h.nodes = make([]*indexedBinomialNode[K, V], len(h.nodes))
}

// Peek returns the extremum (minimum or maximum) key with its value on the heap without removing it.
// If the heap is empty, the second return value will be false.
func (h *indexedBinomial[K, V]) Peek() (int, K, V, bool) {
	if h.IsEmpty() {
		var zeroK K
		var zeroV V
		return -1, zeroK, zeroV, false
	}

	// Find the extremum (minimum in min heap or maximum in max heap) node in the root list.
	_, ext := h.findExtremum(h.head)

	return ext.index, ext.key, ext.val, true
}

// PeekIndex returns a key-value pair on the heap by its associated index without removing it.
// If the index is not valid or not on the heap, the second return value will be false.
func (h *indexedBinomial[K, V]) PeekIndex(i int) (K, V, bool) {
	// ContainsIndex validates the index too.
	if !h.ContainsIndex(i) {
		var zeroK K
		var zeroV V
		return zeroK, zeroV, false
	}

	return h.nodes[i].key, h.nodes[i].val, true
}

// ContainsIndex returns true if a given index is on the heap.
func (h *indexedBinomial[K, V]) ContainsIndex(i int) bool {
	return 0 <= i && i < len(h.nodes) && h.nodes[i] != nil
}

// ContainsKey returns true if the given key is on the heap.
func (h *indexedBinomial[K, V]) ContainsKey(key K) bool {
	for i := 0; i < h.n; i++ {
		if h.nodes[i] != nil && h.cmpKey(h.nodes[i].key, key) == 0 {
			return true
		}
	}

	return false
}

// ContainsValue returns true if the given value is on the heap.
func (h *indexedBinomial[K, V]) ContainsValue(val V) bool {
	for i := 0; i < h.n; i++ {
		if h.nodes[i] != nil && h.eqVal(h.nodes[i].val, val) {
			return true
		}
	}

	return false
}

// DOT generates a DOT representation of the heap.
func (h *indexedBinomial[K, V]) DOT() string {
	graph := dot.NewGraph(true, true, false, "Indexed Binomial Heap", "", "", "", dot.ShapeMrecord)

	h.traverse(h.head, VLR, func(n *indexedBinomialNode[K, V]) bool {
		name := fmt.Sprintf("%d", n.index)

		rec := dot.NewRecord(
			dot.NewComplexField(
				dot.NewRecord(
					dot.NewSimpleField("", fmt.Sprintf("%v", n.index)),
					dot.NewComplexField(
						dot.NewRecord(
							dot.NewSimpleField("", fmt.Sprintf("%v", n.key)),
							dot.NewSimpleField("", fmt.Sprintf("%v", n.val)),
						),
					),
				),
			),
		)

		graph.AddNode(dot.NewNode(name, "", rec.Label(), "", "", "", "", ""))

		if n.parent != nil {
			parent := fmt.Sprintf("%d", n.parent.index)
			graph.AddEdge(dot.NewEdge(name, parent, dot.EdgeTypeDirected, "", "", dot.ColorTurquoise, dot.StyleDashed, "", ""))
		}

		if n.child != nil {
			child := fmt.Sprintf("%d", n.child.index)
			graph.AddEdge(dot.NewEdge(name, child, dot.EdgeTypeDirected, "", "", dot.ColorBlue, "", "", ""))
		}

		if n.sibling != nil {
			sibling := fmt.Sprintf("%d", n.sibling.index)
			graph.AddEdge(dot.NewEdge(name, sibling, dot.EdgeTypeDirected, "", "", dot.ColorRed, dot.StyleDashed, "", ""))
		}

		return true
	})

	return graph.DOT()
}
