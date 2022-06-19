package symboltable

import (
	"fmt"

	"github.com/moorara/algo/common"
	"github.com/moorara/algo/internal/graphviz"
)

type bstNode[K, V any] struct {
	key   K
	val   V
	left  *bstNode[K, V]
	right *bstNode[K, V]
	size  int
}

type bst[K, V any] struct {
	root   *bstNode[K, V]
	cmpKey common.CompareFunc[K]
}

// NewBST creates a new binary search tree.
//
// A binary search tree (BST) is a binary tree in symmetric order.
// Every node's key is:
//   Larger than all keys in its left sub-tree.
//   Smaller than all keys in its right sub-tree.
func NewBST[K, V any](cmpKey common.CompareFunc[K]) OrderedSymbolTable[K, V] {
	return &bst[K, V]{
		root:   nil,
		cmpKey: cmpKey,
	}
}

func (t *bst[K, V]) verify() bool {
	return t.isBST(t.root, nil, nil) &&
		t.isSizeOK(t.root) &&
		t.isRankOK()
}

func (t *bst[K, V]) isBST(n *bstNode[K, V], min, max *K) bool {
	if n == nil {
		return true
	}

	if (min != nil && t.cmpKey(n.key, *min) <= 0) ||
		(max != nil && t.cmpKey(n.key, *max) >= 0) {
		return false
	}

	return t.isBST(n.left, min, &n.key) && t.isBST(n.right, &n.key, max)
}

func (t *bst[K, V]) isSizeOK(n *bstNode[K, V]) bool {
	if n == nil {
		return true
	}

	if n.size != 1+t.size(n.left)+t.size(n.right) {
		return false
	}

	return t.isSizeOK(n.left) && t.isSizeOK(n.right)
}

func (t *bst[K, V]) isRankOK() bool {
	for i := 0; i < t.Size(); i++ {
		k, _, _ := t.Select(i)
		if t.Rank(k) != i {
			return false
		}
	}

	for _, kv := range t.KeyValues() {
		k, _, _ := t.Select(t.Rank(kv.key))
		if t.cmpKey(kv.key, k) != 0 {
			return false
		}
	}

	return true
}

// Size returns the number of key-value pairs in BST.
func (t *bst[K, V]) Size() int {
	return t.size(t.root)
}

func (t *bst[K, V]) size(n *bstNode[K, V]) int {
	if n == nil {
		return 0
	}

	return n.size
}

// Height returns the height of BST.
func (t *bst[K, V]) Height() int {
	return t.height(t.root)
}

func (t *bst[K, V]) height(n *bstNode[K, V]) int {
	if n == nil {
		return 0
	}

	return 1 + max(t.height(n.left), t.height(n.right))
}

// IsEmpty returns true if BST is empty.
func (t *bst[K, V]) IsEmpty() bool {
	return t.root == nil
}

// Put adds a new key-value pair to BST.
func (t *bst[K, V]) Put(key K, val V) {
	t.root = t._put(t.root, key, val)
}

func (t *bst[K, V]) _put(n *bstNode[K, V], key K, val V) *bstNode[K, V] {
	if n == nil {
		return &bstNode[K, V]{
			key:  key,
			val:  val,
			size: 1,
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

	n.size = 1 + t.size(n.left) + t.size(n.right)

	return n
}

// Get returns the value of a given key in BST.
func (t *bst[K, V]) Get(key K) (V, bool) {
	return t._get(t.root, key)
}

func (t *bst[K, V]) _get(n *bstNode[K, V], key K) (V, bool) {
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

// Delete removes a key-value pair from BST.
func (t *bst[K, V]) Delete(key K) (val V, ok bool) {
	t.root, val, ok = t._delete(t.root, key)
	return val, ok
}

func (t *bst[K, V]) _delete(n *bstNode[K, V], key K) (*bstNode[K, V], V, bool) {
	if n == nil {
		var zeroV V
		return n, zeroV, false
	}

	var ok bool
	var val V

	cmp := t.cmpKey(key, n.key)
	if cmp < 0 {
		n.left, val, ok = t._delete(n.left, key)
	} else if cmp > 0 {
		n.right, val, ok = t._delete(n.right, key)
	} else {
		ok = true
		val = n.val

		if n.left == nil {
			return n.right, val, ok
		} else if n.right == nil {
			return n.left, val, ok
		} else {
			m := n
			n = t._min(m.right)
			n.right, _ = t._deleteMin(m.right)
			n.left = m.left
		}
	}

	n.size = 1 + t.size(n.left) + t.size(n.right)
	return n, val, ok
}

// KeyValues returns all key-value pairs in BST.
func (t *bst[K, V]) KeyValues() []KeyValue[K, V] {
	i := 0
	kvs := make([]KeyValue[K, V], t.Size())

	t._traverse(t.root, InOrder, func(n *bstNode[K, V]) bool {
		kvs[i] = KeyValue[K, V]{n.key, n.val}
		i++
		return true
	})

	return kvs
}

// Min returns the minimum key and its value in BST.
func (t *bst[K, V]) Min() (K, V, bool) {
	var zeroK K
	var zeroV V

	if t.root == nil {
		return zeroK, zeroV, false
	}

	n := t._min(t.root)
	return n.key, n.val, true
}

func (t *bst[K, V]) _min(n *bstNode[K, V]) *bstNode[K, V] {
	if n.left == nil {
		return n
	}

	return t._min(n.left)
}

// Max returns the maximum key and its value in BST.
func (t *bst[K, V]) Max() (K, V, bool) {
	var zeroK K
	var zeroV V

	if t.root == nil {
		return zeroK, zeroV, false
	}

	n := t._max(t.root)
	return n.key, n.val, true
}

func (t *bst[K, V]) _max(n *bstNode[K, V]) *bstNode[K, V] {
	if n.right == nil {
		return n
	}

	return t._max(n.right)
}

// Floor returns the largest key in BST less than or equal to key.
func (t *bst[K, V]) Floor(key K) (K, V, bool) {
	var zeroK K
	var zeroV V

	n := t._floor(t.root, key)
	if n == nil {
		return zeroK, zeroV, false
	}

	return n.key, n.val, true
}

func (t *bst[K, V]) _floor(n *bstNode[K, V], key K) *bstNode[K, V] {
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

// Ceiling returns the smallest key in BST greater than or equal to key.
func (t *bst[K, V]) Ceiling(key K) (K, V, bool) {
	var zeroK K
	var zeroV V

	n := t._ceiling(t.root, key)
	if n == nil {
		return zeroK, zeroV, false
	}

	return n.key, n.val, true
}

func (t *bst[K, V]) _ceiling(n *bstNode[K, V], key K) *bstNode[K, V] {
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

// DeleteMin removes the smallest key and associated value from BST.
func (t *bst[K, V]) DeleteMin() (K, V, bool) {
	var zeroK K
	var zeroV V

	if t.root == nil {
		return zeroK, zeroV, false
	}

	var min *bstNode[K, V]
	t.root, min = t._deleteMin(t.root)

	return min.key, min.val, true
}

func (t *bst[K, V]) _deleteMin(n *bstNode[K, V]) (*bstNode[K, V], *bstNode[K, V]) {
	if n.left == nil {
		return n.right, n
	}

	var min *bstNode[K, V]
	n.left, min = t._deleteMin(n.left)
	n.size = 1 + t.size(n.left) + t.size(n.right)
	return n, min
}

// DeleteMax removes the largest key and associated value from BST.
func (t *bst[K, V]) DeleteMax() (K, V, bool) {
	var zeroK K
	var zeroV V

	if t.root == nil {
		return zeroK, zeroV, false
	}

	var max *bstNode[K, V]
	t.root, max = t._deleteMax(t.root)

	return max.key, max.val, true
}

func (t *bst[K, V]) _deleteMax(n *bstNode[K, V]) (*bstNode[K, V], *bstNode[K, V]) {
	if n.right == nil {
		return n.left, n
	}

	var max *bstNode[K, V]
	n.right, max = t._deleteMax(n.right)
	n.size = 1 + t.size(n.left) + t.size(n.right)
	return n, max
}

// Select return the k-th smallest key in BST.
func (t *bst[K, V]) Select(rank int) (K, V, bool) {
	var zeroK K
	var zeroV V

	if rank < 0 || rank >= t.Size() {
		return zeroK, zeroV, false
	}

	n := t._select(t.root, rank)

	return n.key, n.val, true
}

func (t *bst[K, V]) _select(n *bstNode[K, V], rank int) *bstNode[K, V] {
	if n == nil {
		return nil
	}

	s := t.size(n.left)
	switch {
	case rank < s:
		return t._select(n.left, rank)
	case rank > s:
		return t._select(n.right, rank-s-1)
	default:
		return n
	}
}

// Rank returns the number of keys in BST less than key.
func (t *bst[K, V]) Rank(key K) int {
	return t._rank(t.root, key)
}

func (t *bst[K, V]) _rank(n *bstNode[K, V], key K) int {
	if n == nil {
		return 0
	}

	cmp := t.cmpKey(key, n.key)
	switch {
	case cmp < 0:
		return t._rank(n.left, key)
	case cmp > 0:
		return 1 + t.size(n.left) + t._rank(n.right, key)
	default:
		return t.size(n.left)
	}
}

// RangeSize returns the number of keys in BST between two given keys.
func (t *bst[K, V]) RangeSize(lo, hi K) int {
	if t.cmpKey(lo, hi) > 0 {
		return 0
	} else if _, found := t.Get(hi); found {
		return 1 + t.Rank(hi) - t.Rank(lo)
	} else {
		return t.Rank(hi) - t.Rank(lo)
	}
}

// Range returns all keys and associated values in BST between two given keys.
func (t *bst[K, V]) Range(lo, hi K) []KeyValue[K, V] {
	kvs := make([]KeyValue[K, V], 0)
	len := t._range(t.root, &kvs, lo, hi)
	return kvs[0:len]
}

func (t *bst[K, V]) _range(n *bstNode[K, V], kvs *[]KeyValue[K, V], lo, hi K) int {
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

// Traverse is used for visiting all key-value pairs in BST.
func (t *bst[K, V]) Traverse(order TraversalOrder, visit VisitFunc[K, V]) {
	if order != PreOrder && order != InOrder && order != PostOrder {
		return
	}

	t._traverse(t.root, order, func(n *bstNode[K, V]) bool {
		return visit(n.key, n.val)
	})
}

func (t *bst[K, V]) _traverse(n *bstNode[K, V], order TraversalOrder, visit func(*bstNode[K, V]) bool) bool {
	if n == nil {
		return true
	}

	switch order {
	case PreOrder:
		return visit(n) &&
			t._traverse(n.left, order, visit) &&
			t._traverse(n.right, order, visit)
	case InOrder:
		return t._traverse(n.left, order, visit) &&
			visit(n) &&
			t._traverse(n.right, order, visit)
	case PostOrder:
		return t._traverse(n.left, order, visit) &&
			t._traverse(n.right, order, visit) &&
			visit(n)
	default:
		return false
	}
}

// Graphviz returns a visualization of BST in Graphviz format.
func (t *bst[K, V]) Graphviz() string {
	var node, label, left, right string

	graph := graphviz.NewGraph(true, true, "BST", "", "", "", graphviz.ShapeOval)

	t._traverse(t.root, PreOrder, func(n *bstNode[K, V]) bool {
		node = fmt.Sprintf("%d", t.Rank(n.key))
		label = fmt.Sprintf("%v,%v", n.key, n.val)

		graph.AddNode(graphviz.NewNode(node, "", label, "", "", "", "", ""))

		if n.left != nil {
			left = fmt.Sprintf("%d", t.Rank(n.left.key))
			graph.AddEdge(graphviz.NewEdge(node, left, graphviz.EdgeTypeDirected, "", "", "", "", ""))
		}

		if n.right != nil {
			right = fmt.Sprintf("%d", t.Rank(n.right.key))
			graph.AddEdge(graphviz.NewEdge(node, right, graphviz.EdgeTypeDirected, "", "", "", "", ""))
		}

		return true
	})

	return graph.DotCode()
}
