package heap

import (
	"fmt"
	"math"

	"github.com/moorara/algo/dot"
	"github.com/moorara/algo/generic"
)

// The Left-Child-Right-Sibling (LCRS) representation, a.k.a. the Binary Representation of N-ary Tree,
// is used for implementing the indexed Fibonacci tree.
//
//   - The left child of a node points to its first child.
//   - The right sibling of a node points to its next sibling.
//
// The siblings nodes are doubly linked in a circular manner using an additional previous pointer.
//
// The LCRS representation uses three pointers per node regardless of the number of children.
type indexedFibonacciNode[K, V any] struct {
	index                     int // index associated with the key
	key                       K
	val                       V
	degree                    int  // the number of direct children of this node
	mark                      bool // indicates if this node already lost a child
	parent, child, prev, next *indexedFibonacciNode[K, V]
}

// indexedFibonacci implements an indexed Fibonacci heap.
// A Fibonacci heap is implemented as a circular doubly linked list of root nodes (root list) of trees
// as well as a mapping between indices and nodes.
type indexedFibonacci[K, V any] struct {
	cmpKey generic.CompareFunc[K]
	eqVal  generic.EqualFunc[V]

	n     int                           // current number of items on heap
	ext   *indexedFibonacciNode[K, V]   // extremum (minimum or maximum) node of the root list
	nodes []*indexedFibonacciNode[K, V] // map of indices to nodes
}

// NewIndexedFibonacci creates a new indexed Fibonacci heap that can be used as a priority queue.
//
// An indexed heap (priority queue) associates an index with each key-value pair.
// It allows changing the key (priority) of an index, deleting by index, and looking up by index.
// The size of an indexed binary heap is fixed.
//
// The indexed Fibonacci heap does not support the merge operation,
// as doing so would cause conflicts between indices.
//
// Parameters:
//
//   - cap is the maximum number of items on the heap.
//   - cmpKey is a function for comparing two keys.
//   - eqVal is a function for checking the equality of two values.
func NewIndexedFibonacci[K, V any](cap int, cmpKey generic.CompareFunc[K], eqVal generic.EqualFunc[V]) IndexedHeap[K, V] {
	return &indexedFibonacci[K, V]{
		cmpKey: cmpKey,
		eqVal:  eqVal,
		n:      0,
		ext:    nil,
		nodes:  make([]*indexedFibonacciNode[K, V], cap),
	}
}

// nolint: unused
// This method verifies the integrity of a Fibonacci heap.
func (h *indexedFibonacci[K, V]) verify() bool {
	if h.ext == nil {
		return true
	}

	maxD := h.maxDegree()
	var ext *indexedFibonacciNode[K, V]

	// Verify the properties of each tree in the root list.
	for stop, curr := h.ext, h.ext; ; {
		ext = h.pickExt(ext, curr)

		if !h.verifyTree(curr, maxD) {
			return false
		}

		if curr = curr.next; curr == stop {
			break
		}
	}

	return h.ext == ext
}

// nolint: unused
// verifyTree verifies the properties of a tree rooted at the given node.
func (h *indexedFibonacci[K, V]) verifyTree(n *indexedFibonacciNode[K, V], maxD int) bool {
	// Verify the index map for the current node.
	if h.nodes[n.index] != n {
		return false
	}

	// Verify the degree of each node is at most logᵩn.
	if n.degree > maxD {
		return false
	}

	for i, stop, child := 1, n.child, n.child; child != nil; i++ {
		// In a min-heap, each node's key must be smaller than or equal to its children's keys.
		// In a max-heap, each node's key must be greater than or equal to its children's keys.
		if h.cmpKey(n.key, child.key) > 0 {
			return false
		}

		// Verify the child degree.
		if child.degree > n.degree {
			return false
		}

		// Verify the parent link.
		if child.parent != n {
			return false
		}

		// Recursively, verify each tree root at the current child.
		if !h.verifyTree(child, maxD) {
			return false
		}

		if child = child.next; child == stop {
			break
		}
	}

	return true
}

// pickExt compares the keys of two nodes and returns the node with extremum key
// (minimum in min-heap and maximum in max-heap).
// If one of the nodes is nil, the other node is returned.
func (h *indexedFibonacci[K, V]) pickExt(a, b *indexedFibonacciNode[K, V]) *indexedFibonacciNode[K, V] {
	if a == nil {
		return b
	} else if b == nil {
		return a
	} else if h.cmpKey(a.key, b.key) <= 0 {
		return a
	} else {
		return b
	}
}

// insert adds a new node to a circular doubly linked list.
//
// This method is defined on the fibonacci struct to prevent name clashes
// with other similar implementations in this package.
func (_ *indexedFibonacci[K, V]) insert(head, n *indexedFibonacciNode[K, V]) *indexedFibonacciNode[K, V] {
	if head == nil {
		n.prev, n.next = n, n
	} else {
		head.prev.next = n
		n.prev = head.prev
		head.prev = n
		n.next = head
	}

	return n
}

// cut removes a tree from a circular doubly linked list.
//
// This method is defined on the fibonacci struct to prevent name clashes
// with other similar implementations in this package.
func (_ *indexedFibonacci[K, V]) cut(head, n *indexedFibonacciNode[K, V]) *indexedFibonacciNode[K, V] {
	// n is the only node in the circular root list.
	if n.next == n && n.prev == n {
		n.prev, n.next = nil, nil
		return nil
	}

	n.prev.next = n.next
	n.next.prev = n.prev

	if head == n {
		head = n.next
	}

	n.prev, n.next = nil, nil

	return head
}

// cutAndCascade removes a tree from a circular doubly linked list and insert it into the root list.
//
// If the parent of the cut node already lost a child, the parent node itslef will be cut-and-cascaded.
// This is crucial to ensure the logᵩn upper bound (O(logn)) on the degree of any node
// and the amortized running times of operations in a Fibonacci heap.
func (h *indexedFibonacci[K, V]) cutAndCascade(n *indexedFibonacciNode[K, V]) {
	if n.parent == nil {
		return
	}

	parent := n.parent
	n.parent = nil
	parent.child = h.cut(parent.child, n)
	parent.degree--
	h.insert(h.ext, n)

	if parent.mark = !parent.mark; !parent.mark {
		h.cutAndCascade(parent)
	}
}

// meld merges two circular doubly linked lists.
//
// This method is defined on the fibonacci struct to prevent name clashes
// with other similar implementations in this package.
func (_ *indexedFibonacci[K, V]) meld(h1, h2 *indexedFibonacciNode[K, V]) *indexedFibonacciNode[K, V] {
	switch {
	case h1 == nil:
		return h2
	case h2 == nil:
		return h1
	default:
		h1.prev.next = h2.next
		h2.next.prev = h1.prev
		h1.prev = h2
		h2.next = h1
		return h1
	}
}

// consolidate is a lazy operation run after deleting the extremum node from the heap.
func (h *indexedFibonacci[K, V]) consolidate() {
	// Calculate an upper bound on the maximum degree of nodes.
	maxD := h.maxDegree()

	// Create a slice for mapping node degrees to root nodes.
	roots := make([]*indexedFibonacciNode[K, V], maxD)

	// This loop scans the circular root list.
	for stop, curr := h.ext, h.ext; ; {
		x := curr

		// This loop combines the trees of the same degrees until there is no two tree with the same degrees.
		// y != x is necessary since we may have visited the x previosuly and x might be already in the roots map.
		for y := roots[x.degree]; y != nil && y != x; y = roots[x.degree] {
			roots[x.degree] = nil

			/*
			 * h.ext serves as an entry point to the circular root list too.
			 * It must be updated after cutting a node in case the cut node was h.ext itself.
			 * The correct extremum node will be identified and set later.
			 */

			if h.cmpKey(x.key, y.key) > 0 {
				// x needs to be removed from the circular root list and becomes a child of y.
				h.ext = h.cut(h.ext, x)
				h.link(x, y)
				x = y // x should always point to the root of combined trees.
			} else {
				// y needs to be removed from the circular root list and becomes a child of x.
				h.ext = h.cut(h.ext, y)
				h.link(y, x)
			}

			// We need to continue scanning the circular root list
			// from this root and stop until reaching here again.
			stop, curr = x, x
		}

		roots[x.degree] = x

		if curr = curr.next; curr == stop {
			break
		}
	}

	// Identify and set the correct extremum node.
	for _, r := range roots {
		if r != nil {
			h.ext = h.pickExt(h.ext, r)
		}
	}
}

// maxDegree calculates an upper bound on the maximum degree of nodes in a Fibonacci heap.
func (h *indexedFibonacci[K, V]) maxDegree() int {
	/*
	 * The Fibonacci sequence is defined as:
	 *
	 *	F₀ = 0           n = 0
	 *	F₁ = 1           n = 1
	 *	Fₙ = Fₙ₋₁ + Fₙ₋₂  n ≥ 2
	 *
	 * As the Fibonacci sequence progresses,
	 * the ratio of consecutive Fibonacci nubmers (Fₙ₊₁ / Fₙ) approaches the golden ratio.
	 *
	 *	As n → ∞, (Fₙ₊₁/Fₙ) → φ = (1 + √5) / 2
	 *
	 * Fibonacci numbers have a closed-form expression known as Binet's Formula,
	 * which involves the golden ratio:
	 *
	 *	Fₙ = (φⁿ - (-φ)⁻ⁿ) / √5
	 *
	 * For large values of n, the term (-φ)⁻ⁿ becomes negligible,
	 * so the Fibonacci numbers can be approximated by:
	 *
	 *	Fₙ ≈ φⁿ / √5
	 *
	 * In a Fibonacci heap,
	 * the size of the subtree of a node with degree k is related to Fibonacci numbers:
	 *
	 *	n = Fₖ₊₂ ≈ φᵏ⁺² / √5
	 *
	 * Rearranging to solve for k in terms of n:
	 *
	 *	k ≈ logᵩ(n √5) - 2 ≈ logᵩ(n)
	 *
	 * Using this approximation,
	 * we can calculate an upper bound on the maximum degree of nodes in a Fibonacci heap.
	 * This value determines the size of slice used during the consolidation operation.
	 */

	φ := (1 + math.Sqrt(5)) / 2 // Golden ratio
	maxD := int(math.Log(float64(h.n))/math.Log(φ)) + 1

	return maxD
}

// link constructs a new tree from two existing trees.
// It inserts one of the root nodes as the child of the other.
//
// It accepts two root nodes and assumes the key of the first root comes after
// the key of the second root (greater in min-heap or smaller in max-heap).
// The second root then becomes the new root.
func (h *indexedFibonacci[K, V]) link(child, parent *indexedFibonacciNode[K, V]) {
	child.parent = parent
	parent.child = h.insert(parent.child, child)
	parent.degree++
}

// Size returns the number of items on the heap.
func (h *indexedFibonacci[K, V]) Size() int {
	return h.n
}

// IsEmpty returns true if the heap is empty.
func (h *indexedFibonacci[K, V]) IsEmpty() bool {
	return h.ext == nil
}

// Insert adds a new key-value pair to the heap.
func (h *indexedFibonacci[K, V]) Insert(i int, key K, val V) bool {
	// ContainsIndex validates the index too.
	if h.ContainsIndex(i) {
		return false
	}

	n := &indexedFibonacciNode[K, V]{
		index:  i,
		key:    key,
		val:    val,
		degree: 0,
	}

	h.insert(h.ext, n)
	h.ext = h.pickExt(h.ext, n)
	h.nodes[i] = n
	h.n++

	return true
}

// ChangeKey changes the key associated with an index.
func (h *indexedFibonacci[K, V]) ChangeKey(i int, key K) bool {
	// ContainsIndex validates the index too.
	if !h.ContainsIndex(i) {
		return false
	}

	n := h.nodes[i]

	if cmp := h.cmpKey(key, n.key); cmp < 0 {
		// Decrease Key
		n.key = key
		if n.parent != nil && h.cmpKey(n.parent.key, n.key) > 0 {
			h.cutAndCascade(n)
		}
		h.ext = h.pickExt(h.ext, n)
	} else if cmp > 0 {
		// Increase Key
		h.DeleteIndex(i)
		h.Insert(i, key, n.val)
	}

	return true
}

// Delete removes the extremum (minimum or maximum) key with its value on the heap.
// If the heap is empty, the second return value will be false.
func (h *indexedFibonacci[K, V]) Delete() (int, K, V, bool) {
	if h.IsEmpty() {
		var zeroK K
		var zeroV V
		return -1, zeroK, zeroV, false
	}

	ext := h.ext

	// Remove the extremum (minimum or maximum) node from the circular root list.
	// h.ext serves as an entry point to the circular root list too.
	// It must be updated after cutting a node in case the cut node was h.ext itself.
	// The correct extremum node will be identified and set later.
	h.ext = h.cut(h.ext, ext)

	// Merge the children of the deleted node with the root list.
	if c := ext.child; c != nil {
		// Set the parent pointers of the deleted node's children to nil,
		// preparing them for merging with the root list.
		for {
			c.parent = nil
			if c = c.next; c == ext.child {
				break
			}
		}

		// The correct extremum node will be identified and set later.
		h.ext = h.meld(h.ext, ext.child)
		ext.child = nil
	}

	h.nodes[ext.index] = nil
	h.n--

	if h.IsEmpty() {
		h.ext = nil
	} else {
		h.consolidate()
	}

	return ext.index, ext.key, ext.val, true
}

// DeleteIndex removes a key-value pair and its associated index from the heap.
// If the index is not valid or not on the heap, the second return value will be false.
func (h *indexedFibonacci[K, V]) DeleteIndex(i int) (K, V, bool) {
	// ContainsIndex validates the index too.
	if !h.ContainsIndex(i) {
		var zeroK K
		var zeroV V
		return zeroK, zeroV, false
	}

	n := h.nodes[i]
	h.cutAndCascade(n)

	// Remove the node from the circular root list.
	// h.ext serves as an entry point to the circular root list too.
	// It must be updated after cutting a node in case the cut node was h.ext itself.
	// The correct extremum node will be identified and set later.
	h.ext = h.cut(h.ext, n)

	// Merge the children of the deleted node with the root list.
	if c := n.child; c != nil {
		// Set the parent pointers of the deleted node's children to nil,
		// preparing them for merging with the root list.
		for {
			c.parent = nil
			if c = c.next; c == n.child {
				break
			}
		}

		// The correct extremum node will be identified and set later.
		h.ext = h.meld(h.ext, n.child)
		n.child = nil
	}

	h.nodes[n.index] = nil
	h.n--

	if h.IsEmpty() {
		h.ext = nil
	} else {
		h.consolidate()
	}

	return n.key, n.val, true
}

// DeleteAll deletes all keys with their values and indices on the heap, leaving it empty.
func (h *indexedFibonacci[K, V]) DeleteAll() {
	h.n = 0
	h.ext = nil
	h.nodes = make([]*indexedFibonacciNode[K, V], len(h.nodes))
}

// Peek returns the extremum (minimum or maximum) key with its value on the heap without removing it.
// If the heap is empty, the second return value will be false.
func (h *indexedFibonacci[K, V]) Peek() (int, K, V, bool) {
	if h.IsEmpty() {
		var zeroK K
		var zeroV V
		return -1, zeroK, zeroV, false
	}

	return h.ext.index, h.ext.key, h.ext.val, true
}

// PeekIndex returns a key-value pair on the heap by its associated index without removing it.
// If the index is not valid or not on the heap, the second return value will be false.
func (h *indexedFibonacci[K, V]) PeekIndex(i int) (K, V, bool) {
	// ContainsIndex validates the index too.
	if !h.ContainsIndex(i) {
		var zeroK K
		var zeroV V
		return zeroK, zeroV, false
	}

	return h.nodes[i].key, h.nodes[i].val, true
}

// ContainsIndex returns true if a given index is on the heap.
func (h *indexedFibonacci[K, V]) ContainsIndex(i int) bool {
	return 0 <= i && i < len(h.nodes) && h.nodes[i] != nil
}

// ContainsKey returns true if the given key is on the heap.
func (h *indexedFibonacci[K, V]) ContainsKey(key K) bool {
	for i := 0; i < h.n; i++ {
		if h.nodes[i] != nil && h.cmpKey(h.nodes[i].key, key) == 0 {
			return true
		}
	}

	return false
}

// ContainsValue returns true if the given value is on the heap.
func (h *indexedFibonacci[K, V]) ContainsValue(val V) bool {
	for i := 0; i < h.n; i++ {
		if h.nodes[i] != nil && h.eqVal(h.nodes[i].val, val) {
			return true
		}
	}

	return false
}

// DOT generates a DOT representation of the heap.
func (h *indexedFibonacci[K, V]) DOT() string {
	// Create a map of node --> id
	var id int
	nodeID := map[*indexedFibonacciNode[K, V]]int{}
	h.traverse(h.ext, generic.VLR, func(n *indexedFibonacciNode[K, V]) bool {
		id++
		nodeID[n] = id
		return true
	})

	graph := dot.NewGraph(true, true, false, "Fibonacci Heap", "", "", "", dot.ShapeMrecord)

	h.traverse(h.ext, generic.VLR, func(n *indexedFibonacciNode[K, V]) bool {
		name := fmt.Sprintf("%d", nodeID[n])

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

		var color dot.Color
		var style dot.Style

		if n == h.ext {
			color = dot.ColorLimeGreen
			style = dot.StyleFilled
		}

		graph.AddNode(dot.NewNode(name, "", rec.Label(), color, style, "", "", ""))

		if n.child != nil {
			child := fmt.Sprintf("%d", nodeID[n.child])
			graph.AddEdge(dot.NewEdge(name, child, dot.EdgeTypeDirected, "", "", dot.ColorBlue, "", "", ""))
		}

		if n.next != nil {
			next := fmt.Sprintf("%d", nodeID[n.next])
			graph.AddEdge(dot.NewEdge(name, next, dot.EdgeTypeDirected, "", "", dot.ColorRed, dot.StyleDashed, "", ""))
		}

		if n.prev != nil {
			prev := fmt.Sprintf("%d", nodeID[n.prev])
			graph.AddEdge(dot.NewEdge(name, prev, dot.EdgeTypeDirected, "", "", dot.ColorOrange, dot.StyleDashed, "", ""))
		}

		return true
	})

	return graph.DOT()
}

// traverse performs a depth-first traversal of the Fibonacci heap,
// visiting each node according to the specified traversal order.
func (h *indexedFibonacci[K, V]) traverse(n *indexedFibonacciNode[K, V], order generic.TraverseOrder, visit func(*indexedFibonacciNode[K, V]) bool) bool {
	if n == nil {
		return true
	}

	visited := map[*indexedFibonacciNode[K, V]]bool{}

	var traverse func(n *indexedFibonacciNode[K, V], order generic.TraverseOrder) bool
	traverse = func(n *indexedFibonacciNode[K, V], order generic.TraverseOrder) bool {
		if n == nil || visited[n] {
			return true
		}

		visited[n] = true

		switch order {
		case generic.VLR:
			return visit(n) && traverse(n.child, order) && traverse(n.next, order)
		case generic.VRL:
			return visit(n) && traverse(n.next, order) && traverse(n.child, order)
		case generic.LVR:
			return traverse(n.child, order) && visit(n) && traverse(n.next, order)
		case generic.RVL:
			return traverse(n.next, order) && visit(n) && traverse(n.child, order)
		case generic.LRV:
			return traverse(n.child, order) && traverse(n.next, order) && visit(n)
		case generic.RLV:
			return traverse(n.next, order) && traverse(n.child, order) && visit(n)
		default:
			return false
		}
	}

	return traverse(n, order)
}
