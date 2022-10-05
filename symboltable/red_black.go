package symboltable

import (
	"fmt"

	"github.com/moorara/algo/common"
	"github.com/moorara/algo/internal/graphviz"
)

const (
	red   = true
	black = false
)

type rbNode[K, V any] struct {
	key         K
	val         V
	left, right *rbNode[K, V]
	size        int
	color       bool
}

// redBlack is a left-leaning Red-Black tree.
type redBlack[K, V any] struct {
	root   *rbNode[K, V]
	cmpKey common.CompareFunc[K]
}

// NewRedBlack creates a new Red-Black tree.
//
// A Red-Black tree is 2-3 Tree represented as a binary search tree.
// In a left-leaning Red-Black tree, left-leaning red links are used to construct 3-nodes.
// A left-leaning Red-Black tree is a BST such that:
//
//	Red links lean left.
//	No node has two red links connect to it.
//	Every path from root to null link has the same number of black links.
func NewRedBlack[K, V any](cmpKey common.CompareFunc[K]) OrderedSymbolTable[K, V] {
	return &redBlack[K, V]{
		root:   nil,
		cmpKey: cmpKey,
	}
}

func (t *redBlack[K, V]) verify() bool {
	return t._isBST(t.root, nil, nil) &&
		t._isRedBlack(t.root) &&
		t._isSizeOK(t.root) &&
		t._isRankOK() &&
		t._isBalanced()
}

func (t *redBlack[K, V]) _isBST(n *rbNode[K, V], min, max *K) bool {
	if n == nil {
		return true
	}

	if (min != nil && t.cmpKey(n.key, *min) <= 0) ||
		(max != nil && t.cmpKey(n.key, *max) >= 0) {
		return false
	}

	return t._isBST(n.left, min, &n.key) && t._isBST(n.right, &n.key, max)
}

// A Red-Black tree should have no red right links, and at most one left red links in a row on any path.
func (t *redBlack[K, V]) _isRedBlack(n *rbNode[K, V]) bool {
	if n == nil {
		return true
	}

	if t.isRed(n.right) ||
		n != t.root && t.isRed(n) && t.isRed(n.left) {
		return false
	}

	return true
}

func (t *redBlack[K, V]) _isSizeOK(n *rbNode[K, V]) bool {
	if n == nil {
		return true
	}

	if n.size != 1+t._size(n.left)+t._size(n.right) {
		return false
	}

	return t._isSizeOK(n.left) && t._isSizeOK(n.right)
}

func (t *redBlack[K, V]) _isRankOK() bool {
	for i := 0; i < t.Size(); i++ {
		k, _, _ := t.Select(i)
		if t.Rank(k) != i {
			return false
		}
	}

	for _, kv := range t.KeyValues() {
		k, _, _ := t.Select(t.Rank(kv.Key))
		if t.cmpKey(kv.Key, k) != 0 {
			return false
		}
	}

	return true
}

// All paths from root to leaf should have same number of black edges.
func (t *redBlack[K, V]) _isBalanced() bool {
	count := 0
	var n *rbNode[K, V]
	for n = t.root; n != nil; n = n.left {
		if !t.isRed(n) {
			count++
		}
	}

	return t._isBalancedAt(t.root, count)
}

func (t *redBlack[K, V]) _isBalancedAt(n *rbNode[K, V], count int) bool {
	if n == nil {
		return count == 0
	}

	if !t.isRed(n) {
		count--
	}

	return t._isBalancedAt(n.left, count) && t._isBalancedAt(n.right, count)
}

func (t *redBlack[K, V]) isRed(n *rbNode[K, V]) bool {
	if n == nil {
		return black
	}

	return n.color == red
}

// Assuming n is not nil.
func (t *redBlack[K, V]) balance(n *rbNode[K, V]) *rbNode[K, V] {
	// assert n != nil

	if t.isRed(n.right) {
		n = t.rotateLeft(n)
	}

	if t.isRed(n.left) && t.isRed(n.left.left) {
		n = t.rotateRight(n)
	}

	if t.isRed(n.left) && t.isRed(n.right) {
		t.flipColors(n)
	}

	n.size = 1 + t._size(n.left) + t._size(n.right)

	return n
}

func (t *redBlack[K, V]) rotateLeft(n *rbNode[K, V]) *rbNode[K, V] {
	r := n.right
	n.right = r.left
	r.left = n

	r.color = r.left.color
	r.left.color = red
	r.size = n.size
	n.size = 1 + t._size(n.left) + t._size(n.right)

	return r
}

func (t *redBlack[K, V]) rotateRight(n *rbNode[K, V]) *rbNode[K, V] {
	l := n.left
	n.left = l.right
	l.right = n

	l.color = l.right.color
	l.right.color = red
	l.size = n.size
	n.size = 1 + t._size(n.left) + t._size(n.right)

	return l
}

func (t *redBlack[K, V]) flipColors(n *rbNode[K, V]) {
	n.color = !n.color
	n.left.color = !n.left.color
	n.right.color = !n.right.color
}

// Assuming n is red and both n.left and n.left.left are black, make n.left or one of its children red.
func (t *redBlack[K, V]) moveRedLeft(n *rbNode[K, V]) *rbNode[K, V] {
	// assert n != nil
	// assert t.isRed(n) && !t.isRed(n.left) && !t.isRed(n.left.left)

	t.flipColors(n)
	if t.isRed(n.right.left) {
		n.right = t.rotateRight(n.right)
		n = t.rotateLeft(n)
		t.flipColors(n)
	}

	return n
}

// Assuming n is red and both n.right and n.right.left are black, make n.right or one of its children red.
func (t *redBlack[K, V]) moveRedRight(n *rbNode[K, V]) *rbNode[K, V] {
	// assert n != nil
	// assert t.isRed(n) && !t.isRed(n.right) && !t.isRed(n.right.left)

	t.flipColors(n)
	if t.isRed(n.left.left) {
		n = t.rotateRight(n)
		t.flipColors(n)
	}

	return n
}

// Size returns the number of key-value pairs in Red-Black tree.
func (t *redBlack[K, V]) Size() int {
	return t._size(t.root)
}

func (t *redBlack[K, V]) _size(n *rbNode[K, V]) int {
	if n == nil {
		return 0
	}

	return n.size
}

// Height returns the height of Red-Black tree.
func (t *redBlack[K, V]) Height() int {
	return t._height(t.root)
}

func (t *redBlack[K, V]) _height(n *rbNode[K, V]) int {
	if n == nil {
		return 0
	}

	return 1 + common.Max[int](t._height(n.left), t._height(n.right))
}

// IsEmpty returns true if Red-Black tree is empty.
func (t *redBlack[K, V]) IsEmpty() bool {
	return t.root == nil
}

// Put adds a new key-value pair to Red-Black tree.
func (t *redBlack[K, V]) Put(key K, val V) {
	t.root = t._put(t.root, key, val)
	t.root.color = black
}

func (t *redBlack[K, V]) _put(n *rbNode[K, V], key K, val V) *rbNode[K, V] {
	if n == nil {
		return &rbNode[K, V]{
			key:   key,
			val:   val,
			size:  1,
			color: red,
		}
	}

	cmp := t.cmpKey(key, n.key)
	switch {
	case cmp < 0:
		n.left = t._put(n.left, key, val)
	case cmp > 0:
		n.right = t._put(n.right, key, val)
	default:
		n.val = val
	}

	// fix-up any right-leaning links
	if t.isRed(n.right) && !t.isRed(n.left) {
		n = t.rotateLeft(n)
	}
	if t.isRed(n.left) && t.isRed(n.left.left) {
		n = t.rotateRight(n)
	}
	if t.isRed(n.left) && t.isRed(n.right) {
		t.flipColors(n)
	}

	n.size = 1 + t._size(n.left) + t._size(n.right)

	return n
}

// Get returns the value of a given key in Red-Black tree.
func (t *redBlack[K, V]) Get(key K) (V, bool) {
	return t._get(t.root, key)
}

func (t *redBlack[K, V]) _get(n *rbNode[K, V], key K) (V, bool) {
	if n == nil {
		var zeroV V
		return zeroV, false
	}

	cmp := t.cmpKey(key, n.key)
	switch {
	case cmp < 0:
		return t._get(n.left, key)
	case cmp > 0:
		return t._get(n.right, key)
	default:
		return n.val, true
	}
}

// Delete removes a key-value pair from Red-Black tree.
func (t *redBlack[K, V]) Delete(key K) (val V, ok bool) {
	if t.root == nil {
		var zeroV V
		return zeroV, false
	}

	if !t.isRed(t.root.left) && !t.isRed(t.root.right) {
		t.root.color = red
	}

	t.root, val, ok = t._delete(t.root, key)
	if t.root != nil {
		t.root.color = black
	}

	return val, ok
}

func (t *redBlack[K, V]) _delete(n *rbNode[K, V], key K) (*rbNode[K, V], V, bool) {
	var ok bool
	var val V

	if t.cmpKey(key, n.key) < 0 {
		if !t.isRed(n.left) && !t.isRed(n.left.left) {
			n = t.moveRedLeft(n)
		}
		n.left, val, ok = t._delete(n.left, key)
	} else {
		if t.isRed(n.left) {
			n = t.rotateRight(n)
		}

		if t.cmpKey(key, n.key) == 0 && n.right == nil {
			return nil, n.val, true
		}

		if !t.isRed(n.right) && !t.isRed(n.right.left) {
			n = t.moveRedRight(n)
		}

		if t.cmpKey(key, n.key) == 0 {
			var min *rbNode[K, V]
			val, ok = n.val, true
			n.right, min = t._deleteMin(n.right)
			n.key, n.val = min.key, min.val
		} else {
			n.right, val, ok = t._delete(n.right, key)
		}
	}

	return t.balance(n), val, ok
}

// KeyValues returns all key-value pairs in Red-Black tree.
func (t *redBlack[K, V]) KeyValues() []KeyValue[K, V] {
	kvs := make([]KeyValue[K, V], 0, t.Size())
	t._traverse(t.root, Ascending, func(n *rbNode[K, V]) bool {
		kvs = append(kvs, KeyValue[K, V]{n.key, n.val})
		return true
	})

	return kvs
}

// Min returns the minimum key and its value in Red-Black tree.
func (t *redBlack[K, V]) Min() (K, V, bool) {
	if t.root == nil {
		var zeroK K
		var zeroV V
		return zeroK, zeroV, false
	}

	n := t._min(t.root)
	return n.key, n.val, true
}

func (t *redBlack[K, V]) _min(n *rbNode[K, V]) *rbNode[K, V] {
	if n.left == nil {
		return n
	}

	return t._min(n.left)
}

// Max returns the maximum key and its value in Red-Black tree.
func (t *redBlack[K, V]) Max() (K, V, bool) {
	if t.root == nil {
		var zeroK K
		var zeroV V
		return zeroK, zeroV, false
	}

	n := t._max(t.root)
	return n.key, n.val, true
}

func (t *redBlack[K, V]) _max(n *rbNode[K, V]) *rbNode[K, V] {
	if n.right == nil {
		return n
	}

	return t._max(n.right)
}

// Floor returns the largest key in Red-Black tree less than or equal to key.
func (t *redBlack[K, V]) Floor(key K) (K, V, bool) {
	n := t._floor(t.root, key)
	if n == nil {
		var zeroK K
		var zeroV V
		return zeroK, zeroV, false
	}

	return n.key, n.val, true
}

func (t *redBlack[K, V]) _floor(n *rbNode[K, V], key K) *rbNode[K, V] {
	if n == nil {
		return nil
	}

	if cmp := t.cmpKey(key, n.key); cmp == 0 {
		return n
	} else if cmp < 0 {
		return t._floor(n.left, key)
	}

	if m := t._floor(n.right, key); m != nil {
		return m
	}

	return n
}

// Ceiling returns the smallest key in Red-Black tree greater than or equal to key.
func (t *redBlack[K, V]) Ceiling(key K) (K, V, bool) {
	n := t._ceiling(t.root, key)
	if n == nil {
		var zeroK K
		var zeroV V
		return zeroK, zeroV, false
	}

	return n.key, n.val, true
}

func (t *redBlack[K, V]) _ceiling(n *rbNode[K, V], key K) *rbNode[K, V] {
	if n == nil {
		return nil
	}

	if cmp := t.cmpKey(key, n.key); cmp == 0 {
		return n
	} else if cmp > 0 {
		return t._ceiling(n.right, key)
	}

	if m := t._ceiling(n.left, key); m != nil {
		return m
	}

	return n
}

// DeleteMin removes the smallest key and associated value from Red-Black tree.
func (t *redBlack[K, V]) DeleteMin() (K, V, bool) {
	if t.root == nil {
		var zeroK K
		var zeroV V
		return zeroK, zeroV, false
	}

	if !t.isRed(t.root.left) && !t.isRed(t.root.right) {
		t.root.color = red
	}

	var min *rbNode[K, V]
	t.root, min = t._deleteMin(t.root)
	if t.root != nil {
		t.root.color = black
	}

	return min.key, min.val, true
}

func (t *redBlack[K, V]) _deleteMin(n *rbNode[K, V]) (*rbNode[K, V], *rbNode[K, V]) {
	if n.left == nil {
		return n.right, n
	}

	if !t.isRed(n.left) && !t.isRed(n.left.left) {
		n = t.moveRedLeft(n)
	}

	var min *rbNode[K, V]
	n.left, min = t._deleteMin(n.left)
	return t.balance(n), min
}

// DeleteMax removes the largest key and associated value from Red-Black tree.
func (t *redBlack[K, V]) DeleteMax() (K, V, bool) {
	if t.root == nil {
		var zeroK K
		var zeroV V
		return zeroK, zeroV, false
	}

	if !t.isRed(t.root.left) && !t.isRed(t.root.right) {
		t.root.color = red
	}

	var max *rbNode[K, V]
	t.root, max = t._deleteMax(t.root)
	if t.root != nil {
		t.root.color = black
	}

	return max.key, max.val, true
}

func (t *redBlack[K, V]) _deleteMax(n *rbNode[K, V]) (*rbNode[K, V], *rbNode[K, V]) {
	if t.isRed(n.left) {
		n = t.rotateRight(n)
	}

	if n.right == nil {
		return n.left, n
	}

	if !t.isRed(n.right) && !t.isRed(n.right.left) {
		n = t.moveRedRight(n)
	}

	var max *rbNode[K, V]
	n.right, max = t._deleteMax(n.right)
	return t.balance(n), max
}

// Select returns the k-th smallest key in Red-Black tree.
func (t *redBlack[K, V]) Select(rank int) (K, V, bool) {
	if rank < 0 || rank >= t.Size() {
		var zeroK K
		var zeroV V
		return zeroK, zeroV, false
	}

	n := t._select(t.root, rank)
	return n.key, n.val, true
}

func (t *redBlack[K, V]) _select(n *rbNode[K, V], rank int) *rbNode[K, V] {
	if n == nil {
		return nil
	}

	s := t._size(n.left)
	switch {
	case rank < s:
		return t._select(n.left, rank)
	case rank > s:
		return t._select(n.right, rank-s-1)
	default:
		return n
	}
}

// Rank returns the number of keys in Red-Black tree less than key.
func (t *redBlack[K, V]) Rank(key K) int {
	return t._rank(t.root, key)
}

func (t *redBlack[K, V]) _rank(n *rbNode[K, V], key K) int {
	if n == nil {
		return 0
	}

	cmp := t.cmpKey(key, n.key)
	switch {
	case cmp < 0:
		return t._rank(n.left, key)
	case cmp > 0:
		return 1 + t._size(n.left) + t._rank(n.right, key)
	default:
		return t._size(n.left)
	}
}

// RangeSize returns the number of keys in Red-Black tree between two given keys.
func (t *redBlack[K, V]) RangeSize(lo, hi K) int {
	if t.cmpKey(lo, hi) > 0 {
		return 0
	} else if _, found := t.Get(hi); found {
		return 1 + t.Rank(hi) - t.Rank(lo)
	} else {
		return t.Rank(hi) - t.Rank(lo)
	}
}

// Range returns all keys and associated values in Red-Black tree between two given keys.
func (t *redBlack[K, V]) Range(lo, hi K) []KeyValue[K, V] {
	kvs := make([]KeyValue[K, V], 0)
	len := t._range(t.root, &kvs, lo, hi)
	return kvs[0:len]
}

func (t *redBlack[K, V]) _range(n *rbNode[K, V], kvs *[]KeyValue[K, V], lo, hi K) int {
	if n == nil {
		return 0
	}

	len := 0
	cmpLo := t.cmpKey(lo, n.key)
	cmpHi := t.cmpKey(hi, n.key)

	if cmpLo < 0 {
		len += t._range(n.left, kvs, lo, hi)
	}
	if cmpLo <= 0 && cmpHi >= 0 {
		*kvs = append(*kvs, KeyValue[K, V]{n.key, n.val})
		len++
	}
	if cmpHi > 0 {
		len += t._range(n.right, kvs, lo, hi)
	}

	return len
}

// Traverse is used for visiting all key-value pairs in Red-Black tree.
func (t *redBlack[K, V]) Traverse(order TraversalOrder, visit VisitFunc[K, V]) {
	t._traverse(t.root, order, func(n *rbNode[K, V]) bool {
		return visit(n.key, n.val)
	})
}

func (t *redBlack[K, V]) _traverse(n *rbNode[K, V], order TraversalOrder, visit func(*rbNode[K, V]) bool) bool {
	if n == nil {
		return true
	}

	switch order {
	case VLR:
		return visit(n) && t._traverse(n.left, order, visit) && t._traverse(n.right, order, visit)
	case VRL:
		return visit(n) && t._traverse(n.right, order, visit) && t._traverse(n.left, order, visit)
	case LVR, Ascending:
		return t._traverse(n.left, order, visit) && visit(n) && t._traverse(n.right, order, visit)
	case RVL, Descending:
		return t._traverse(n.right, order, visit) && visit(n) && t._traverse(n.left, order, visit)
	case LRV:
		return t._traverse(n.left, order, visit) && t._traverse(n.right, order, visit) && visit(n)
	case RLV:
		return t._traverse(n.right, order, visit) && t._traverse(n.left, order, visit) && visit(n)
	default:
		return false
	}
}

// Graphviz returns a visualization of Red-Black tree in Graphviz format.
func (t *redBlack[K, V]) Graphviz() string {
	// Create a map of node --> id
	var id int
	nodeID := map[*rbNode[K, V]]int{}
	t._traverse(t.root, VLR, func(n *rbNode[K, V]) bool {
		id++
		nodeID[n] = id
		return true
	})

	graph := graphviz.NewGraph(true, true, false, "Red-Black", "", "", graphviz.StyleFilled, graphviz.ShapeOval)

	t._traverse(t.root, VLR, func(n *rbNode[K, V]) bool {
		var nodeColor, fontColor, edgeColor graphviz.Color

		name := fmt.Sprintf("%d", nodeID[n])
		label := fmt.Sprintf("%v,%v", n.key, n.val)

		if t.isRed(n) {
			nodeColor = graphviz.ColorRed
			fontColor = graphviz.ColorWhite
		} else {
			nodeColor = graphviz.ColorBlack
			fontColor = graphviz.ColorWhite
		}

		graph.AddNode(graphviz.NewNode(name, "", label, nodeColor, "", "", fontColor, ""))

		if n.left != nil {
			left := fmt.Sprintf("%d", nodeID[n.left])
			if t.isRed(n.left) {
				edgeColor = graphviz.ColorRed
			} else {
				edgeColor = graphviz.ColorBlack
			}
			graph.AddEdge(graphviz.NewEdge(name, left, graphviz.EdgeTypeDirected, "", "", edgeColor, "", "", ""))
		}

		if n.right != nil {
			right := fmt.Sprintf("%d", nodeID[n.right])
			graph.AddEdge(graphviz.NewEdge(name, right, graphviz.EdgeTypeDirected, "", "", "", "", "", ""))
		}

		return true
	})

	return graph.DotCode()
}
