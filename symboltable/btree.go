package symboltable

import (
	"fmt"

	"github.com/moorara/algo/common"
	"github.com/moorara/algo/internal/graphviz"
)

type btreeEntry[K, V any] struct {
	key  K
	val  V
	next *btreeNode[K, V]
}

type btreeNode[K, V any] struct {
	m        int
	children []*btreeEntry[K, V]
}

type btree[K, V any] struct {
	M      int
	root   *btreeNode[K, V]
	size   int
	height int
	cmpKey common.CompareFunc[K]
}

// NewBTree creates a new B-tree.
//
// B-tree is a self-balancing tree data structure.
// It maintains sorted data and allows searches, sequential access, insertions, and deletions in logarithmic time.
// The B-tree generalizes the binary search tree by allowing up to M-1 key-link pairs per node.
// It is well suited for storage systems that read and write relatively large blocks of data.
//
// In a B-tree of order M:
//  There are internal nodes and external nodes.
//  Each node has at most M-1 keys and M children.
//  The root node has at least one key and two children.
//  All other internal nodes (non-leaves) have at least M/2 children.
//  Internal nodes contain copies of keys to guide search.
//  External nodes (leaves) contain keys and pointers to data.
//
// M is the order of B-tree and MUST be an even number greater than 2.
func NewBTree[K, V any](M int, cmpKey common.CompareFunc[K]) OrderedSymbolTable[K, V] {
	if M < 2 || M%2 != 0 {
		panic("M MUST be an event number greater than 2")
	}

	root := &btreeNode[K, V]{
		m:        0,
		children: make([]*btreeEntry[K, V], M),
	}

	return &btree[K, V]{
		M:      M,
		root:   root,
		cmpKey: cmpKey,
	}
}

func (t *btree[K, V]) verify() bool {
	// TODO:
	return t._isBTree() &&
		t._isSizeOK() &&
		t._isHeightOK() &&
		t._isRankOK()
}

func (t *btree[K, V]) _isBTree() bool {
	// TODO:
	return false
}

func (t *btree[K, V]) _isSizeOK() bool {
	// TODO:
	return false
}

func (t *btree[K, V]) _isHeightOK() bool {
	// TODO:
	return false
}

func (t *btree[K, V]) _isRankOK() bool {
	// TODO:
	return false
}

// split splits a node into two halves.
func (t *btree[K, V]) split(n *btreeNode[K, V]) *btreeNode[K, V] {
	n.m = t.M / 2
	x := &btreeNode[K, V]{
		m:        t.M / 2,
		children: make([]*btreeEntry[K, V], t.M),
	}

	for i := 0; i < t.M/2; i++ {
		x.children[i] = n.children[t.M/2+i]
	}

	return x
}

// Size returns the number of key-value pairs in B-tree.
func (t *btree[K, V]) Size() int {
	return t.size
}

// Height returns the height of B-tree.
func (t *btree[K, V]) Height() int {
	return t.height
}

// IsEmpty returns true if B-tree is empty.
func (t *btree[K, V]) IsEmpty() bool {
	return t.size == 0
}

// Put adds a new key-value pair to B-tree.
func (t *btree[K, V]) Put(key K, val V) {
	x := t._put(t.root, key, val, t.height)
	if x == nil {
		return
	}

	// Split the root node

	var zeroV V

	y := &btreeNode[K, V]{
		m:        2,
		children: make([]*btreeEntry[K, V], t.M),
	}

	y.children[0] = &btreeEntry[K, V]{
		key:  t.root.children[0].key,
		val:  zeroV,
		next: t.root,
	}

	y.children[1] = &btreeEntry[K, V]{
		key:  x.children[0].key,
		val:  zeroV,
		next: x,
	}

	t.root = y
	t.height++
}

func (t *btree[K, V]) _put(n *btreeNode[K, V], key K, val V, h int) *btreeNode[K, V] {
	var zeroV V
	var i int

	e := &btreeEntry[K, V]{
		key: key,
		val: val,
	}

	if h == 0 { // External node
		for i = 0; i < n.m; i++ {
			cmp := t.cmpKey(key, n.children[i].key)
			if cmp == 0 {
				n.children[i].val = val // Update value for the existing key
				return nil
			} else if cmp < 0 {
				break
			}
		}
	} else { // Internal node
		for i = 0; i < n.m; i++ {
			if i+1 == n.m || t.cmpKey(key, n.children[i+1].key) < 0 {
				x := t._put(n.children[i].next, key, val, h-1)
				if x == nil {
					return nil
				}

				e.key = x.children[0].key
				e.val = zeroV
				e.next = x

				i++

				break
			}
		}
	}

	for j := n.m; j > i; j-- {
		n.children[j] = n.children[j-1]
	}

	n.children[i] = e
	n.m++

	if h == 0 {
		t.size++
	}

	if n.m < t.M {
		return nil
	}

	return t.split(n)
}

// Get returns the value of a given key in B-tree.
func (t *btree[K, V]) Get(key K) (V, bool) {
	return t._get(t.root, t.height, key)
}

func (t *btree[K, V]) _get(n *btreeNode[K, V], h int, key K) (V, bool) {
	if h == 0 { // External node
		for i := 0; i < n.m; i++ {
			if t.cmpKey(key, n.children[i].key) == 0 {
				return n.children[i].val, true
			}
		}
	} else { // Internal node
		for i := 0; i < n.m; i++ {
			if i+1 == n.m || t.cmpKey(key, n.children[i+1].key) < 0 {
				return t._get(n.children[i].next, h-1, key)
			}
		}
	}

	var zeroV V
	return zeroV, false
}

// Delete removes a key-value pair from B-tree.
func (t *btree[K, V]) Delete(key K) (V, bool) {
	// TODO:
	var zeroV V
	return zeroV, false
}

func (t *btree[K, V]) _delete(n *btreeNode[K, V], key K, h int) (*btreeNode[K, V], V, bool) {
	// TODO:
	var zeroV V
	return nil, zeroV, false
}

// KeyValues returns all key-value pairs in B-tree.
func (t *btree[K, V]) KeyValues() []KeyValue[K, V] {
	kvs := make([]KeyValue[K, V], 0, t.Size())

	t._traverse(t.root, t.height, Ascending, func(n *btreeNode[K, V], h int) bool {
		if h == 0 {
			for i := 0; i < n.m; i++ {
				kvs = append(kvs, KeyValue[K, V]{
					n.children[i].key,
					n.children[i].val,
				})
			}
		}
		return true
	})

	return kvs
}

// Min returns the minimum key and its value in B-tree.
func (t *btree[K, V]) Min() (K, V, bool) {
	if t.IsEmpty() {
		var zeroK K
		var zeroV V
		return zeroK, zeroV, false
	}

	e := t._min(t.root, t.height)
	return e.key, e.val, true
}

func (t *btree[K, V]) _min(n *btreeNode[K, V], h int) *btreeEntry[K, V] {
	if h == 0 {
		return n.children[0]
	}

	return t._min(n.children[0].next, h-1)
}

// Max returns the maximum key and its value in B-tree.
func (t *btree[K, V]) Max() (K, V, bool) {
	if t.IsEmpty() {
		var zeroK K
		var zeroV V
		return zeroK, zeroV, false
	}

	e := t._max(t.root, t.height)
	return e.key, e.val, true
}

func (t *btree[K, V]) _max(n *btreeNode[K, V], h int) *btreeEntry[K, V] {
	if h == 0 {
		return n.children[n.m-1]
	}

	return t._max(n.children[n.m-1].next, h-1)
}

// Floor returns the largest key in B-tree less than or equal to key.
func (t *btree[K, V]) Floor(key K) (K, V, bool) {
	n := t._floor(t.root, t.height, key)
	if n == nil {
		var zeroK K
		var zeroV V
		return zeroK, zeroV, false
	}

	return n.key, n.val, true
}

func (t *btree[K, V]) _floor(n *btreeNode[K, V], h int, key K) *btreeEntry[K, V] {
	if h > 0 { // Internal node
		for i := 0; i < n.m; i++ {
			if i+1 == n.m || t.cmpKey(key, n.children[i+1].key) < 0 {
				return t._floor(n.children[i].next, h-1, key)
			}
		}
	} else { // External node
		for i := 0; i < n.m; i++ {
			if i+1 == n.m || t.cmpKey(key, n.children[i+1].key) < 0 {
				return n.children[i]
			}
		}
	}

	return nil
}

// Ceiling returns the smallest key in B-tree greater than or equal to key.
func (t *btree[K, V]) Ceiling(key K) (K, V, bool) {
	n := t._ceiling(t.root, t.height, key)
	if n == nil {
		var zeroK K
		var zeroV V
		return zeroK, zeroV, false
	}

	return n.key, n.val, true
}

func (t *btree[K, V]) _ceiling(n *btreeNode[K, V], h int, key K) *btreeEntry[K, V] {
	if h > 0 { // Internal node
		for i := n.m - 1; i >= 0; i-- {
			if i-1 == -1 || t.cmpKey(key, n.children[i-1].key) > 0 {
				return t._ceiling(n.children[i].next, h-1, key)
			}
		}
	} else { // External node
		for i := n.m - 1; i >= 0; i-- {
			if i-1 == -1 || t.cmpKey(key, n.children[i-1].key) > 0 {
				return n.children[i]
			}
		}
	}

	return nil
}

// DeleteMin removes the smallest key and associated value from B-tree.
func (t *btree[K, V]) DeleteMin() (K, V, bool) {
	return t._deleteMin(t.root)
}

func (t *btree[K, V]) _deleteMin(n *btreeNode[K, V]) (K, V, bool) {
	// TODO:
	var zeroK K
	var zeroV V
	return zeroK, zeroV, false
}

// DeleteMax removes the largest key and associated value from B-tree.
func (t *btree[K, V]) DeleteMax() (K, V, bool) {
	return t._deleteMax(t.root)
}

func (t *btree[K, V]) _deleteMax(n *btreeNode[K, V]) (K, V, bool) {
	// TODO:
	var zeroK K
	var zeroV V
	return zeroK, zeroV, false
}

// Select returns the k-th smallest key in B-tree.
func (t *btree[K, V]) Select(rank int) (K, V, bool) {
	// TODO:
	var zeroK K
	var zeroV V
	return zeroK, zeroV, false
}

// Rank returns the number of keys in B-tree less than key.
func (t *btree[K, V]) Rank(key K) int {
	// TODO:
	return -1
}

// RangeSize returns the number of keys in B-tree between two given keys.
func (t *btree[K, V]) RangeSize(lo, hi K) int {
	// TODO:
	return -1
}

// Range returns all keys and associated values in B-tree between two given keys.
func (t *btree[K, V]) Range(lo, hi K) []KeyValue[K, V] {
	// TODO:
	return nil
}

// Traverse is used for visiting all key-value pairs in B-tree.
func (t *btree[K, V]) Traverse(order TraversalOrder, visit VisitFunc[K, V]) {
	t._traverse(t.root, t.height, order, func(n *btreeNode[K, V], h int) bool {
		// Consider only leaf (external) nodes.
		if h > 0 {
			return true
		}

		res := true

		switch order {
		case VLR, LRV, Ascending:
			for i := 0; i < n.m; i++ {
				res = res && visit(n.children[i].key, n.children[i].val)
			}

		case VRL, RLV, Descending:
			for i := n.m - 1; i >= 0; i-- {
				res = res && visit(n.children[i].key, n.children[i].val)
			}
		}

		return res
	})
}

func (t *btree[K, V]) _traverse(n *btreeNode[K, V], h int, order TraversalOrder, visit func(*btreeNode[K, V], int) bool) bool {
	if n == nil {
		return true
	}

	// In a B-tree, non-leaf (internal) nodes only contain keys whereas the leaf (external) nodes contain values.
	// In-order traversal does not make sense in a B-tree.
	switch order {
	case VLR, Ascending:
		res := visit(n, h)
		for i := 0; i < n.m; i++ {
			res = res && t._traverse(n.children[i].next, h-1, order, visit)
		}
		return res

	case VRL, Descending:
		res := visit(n, h)
		for i := n.m - 1; i >= 0; i-- {
			res = res && t._traverse(n.children[i].next, h-1, order, visit)
		}
		return res

	case LRV:
		res := true
		for i := 0; i < n.m; i++ {
			res = res && t._traverse(n.children[i].next, h-1, order, visit)
		}
		return res && visit(n, h)

	case RLV:
		res := true
		for i := n.m - 1; i >= 0; i-- {
			res = res && t._traverse(n.children[i].next, h-1, order, visit)
		}
		return res && visit(n, h)

	default:
		return false
	}
}

func (t *btree[K, V]) Graphviz() string {
	// Create a map of entry --> whether or not is the sentinel
	isSentinel := t._markSentinels(t.root, t.height)

	// Create a map of node --> id
	var id int
	nodeID := map[*btreeNode[K, V]]int{}
	t._traverse(t.root, t.height, VLR, func(n *btreeNode[K, V], h int) bool {
		id++
		nodeID[n] = id
		return true
	})

	graph := graphviz.NewGraph(true, true, false, "B-Tree", graphviz.RankDirTB, "", "", "")

	t._traverse(t.root, t.height, VLR, func(n *btreeNode[K, V], h int) bool {
		name := fmt.Sprintf("%d", nodeID[n])

		if h == 0 { // External node
			rec := graphviz.NewRecord()

			for i, e := range n.children {
				var keyLabel, valLabel string
				if i < n.m {
					keyLabel = fmt.Sprintf("%v", e.key)
					valLabel = fmt.Sprintf("%v", e.val)
				}

				rec.AddField(
					graphviz.NewComplexField(
						graphviz.NewRecord(
							graphviz.NewSimpleField("", keyLabel),
							graphviz.NewSimpleField("", valLabel),
						),
					),
				)
			}

			graph.AddNode(graphviz.NewNode(name, "", rec.Label(), "", graphviz.StyleBold, graphviz.ShapeRecord, "", ""))
		} else { // Internal node
			rec := graphviz.NewRecord()

			for i, e := range n.children {
				var keyLabel, nextName string
				if i < n.m {
					if isSentinel[e] {
						keyLabel = "*"
						nextName = "sentinel"
					} else {
						keyLabel = fmt.Sprintf("%v", e.key)
						nextName = fmt.Sprintf("%v", e.key)
					}

					from := fmt.Sprintf("%s:%s", name, nextName)
					to := fmt.Sprintf("%d", nodeID[e.next])
					graph.AddEdge(graphviz.NewEdge(from, to, graphviz.EdgeTypeDirected, "", "", "", "", "", ""))
				}

				rec.AddField(
					graphviz.NewComplexField(
						graphviz.NewRecord(
							graphviz.NewSimpleField("", keyLabel),
							graphviz.NewSimpleField(nextName, ""),
						),
					),
				)
			}

			graph.AddNode(graphviz.NewNode(name, "", rec.Label(), "", graphviz.StyleSolid, graphviz.ShapeMrecord, "", ""))
		}

		return true
	})

	return graph.DotCode()
}

// Graphviz returns a visualization of B-tree in Graphviz format.
func (t *btree[K, V]) _markSentinels(n *btreeNode[K, V], h int) map[*btreeEntry[K, V]]bool {
	if h == 0 {
		return map[*btreeEntry[K, V]]bool{}
	}

	e := n.children[0]
	m := t._markSentinels(e.next, h-1)
	m[e] = true
	return m
}
