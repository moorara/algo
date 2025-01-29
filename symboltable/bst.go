package symboltable

import (
	"fmt"
	"iter"
	"strings"

	. "github.com/moorara/algo/generic"
	"github.com/moorara/algo/internal/dot"
)

type bstNode[K, V any] struct {
	key         K
	val         V
	left, right *bstNode[K, V]
	size        int
}

type bst[K, V any] struct {
	root   *bstNode[K, V]
	cmpKey CompareFunc[K]
	eqVal  EqualFunc[V]
}

// NewBST creates a new binary search tree.
//
// A binary search tree (BST) is a binary tree in symmetric order.
// Every node's key is:
//
//	Larger than all keys in its left sub-tree.
//	Smaller than all keys in its right sub-tree.
//
// The second parameter (eqVal) is needed only if you want to use the Equal method.
func NewBST[K, V any](cmpKey CompareFunc[K], eqVal EqualFunc[V]) OrderedSymbolTable[K, V] {
	return &bst[K, V]{
		root:   nil,
		cmpKey: cmpKey,
		eqVal:  eqVal,
	}
}

// nolint: unused
func (t *bst[K, V]) verify() bool {
	return t._isBST(t.root, nil, nil) &&
		t._isSizeOK(t.root) &&
		t._isRankOK()
}

// nolint: unused
func (t *bst[K, V]) _isBST(n *bstNode[K, V], min, max *K) bool {
	if n == nil {
		return true
	}

	if (min != nil && t.cmpKey(n.key, *min) <= 0) ||
		(max != nil && t.cmpKey(n.key, *max) >= 0) {
		return false
	}

	return t._isBST(n.left, min, &n.key) && t._isBST(n.right, &n.key, max)
}

// nolint: unused
func (t *bst[K, V]) _isSizeOK(n *bstNode[K, V]) bool {
	if n == nil {
		return true
	}

	if n.size != 1+t._size(n.left)+t._size(n.right) {
		return false
	}

	return t._isSizeOK(n.left) && t._isSizeOK(n.right)
}

// nolint: unused
func (t *bst[K, V]) _isRankOK() bool {
	for i := 0; i < t.Size(); i++ {
		k, _, _ := t.Select(i)
		if t.Rank(k) != i {
			return false
		}
	}

	for key := range t.All() {
		k, _, _ := t.Select(t.Rank(key))
		if t.cmpKey(key, k) != 0 {
			return false
		}
	}

	return true
}

// Size returns the number of key-values in the BST.
func (t *bst[K, V]) Size() int {
	return t._size(t.root)
}

func (t *bst[K, V]) _size(n *bstNode[K, V]) int {
	if n == nil {
		return 0
	}

	return n.size
}

// Height returns the height of the BST.
func (t *bst[K, V]) Height() int {
	return t._height(t.root)
}

func (t *bst[K, V]) _height(n *bstNode[K, V]) int {
	if n == nil {
		return 0
	}

	return 1 + max(t._height(n.left), t._height(n.right))
}

// IsEmpty returns true if the BST is empty.
func (t *bst[K, V]) IsEmpty() bool {
	return t.root == nil
}

// Put adds a new key-value to the BST.
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

	n.size = 1 + t._size(n.left) + t._size(n.right)

	return n
}

// Get returns the value of a given key in the BST.
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

// Delete deletes a key-value from the BST.
func (t *bst[K, V]) Delete(key K) (val V, ok bool) {
	t.root, val, ok = t._delete(t.root, key)
	return val, ok
}

func (t *bst[K, V]) _delete(n *bstNode[K, V], key K) (*bstNode[K, V], V, bool) {
	if n == nil {
		var zeroV V
		return nil, zeroV, false
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

	n.size = 1 + t._size(n.left) + t._size(n.right)
	return n, val, ok
}

// DeleteAll deletes all key-values from the BST, leaving it empty.
func (t *bst[K, V]) DeleteAll() {
	t.root = nil
}

// Min returns the minimum key and its value in the BST.
func (t *bst[K, V]) Min() (K, V, bool) {
	if t.root == nil {
		var zeroK K
		var zeroV V
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

// Max returns the maximum key and its value in the BST.
func (t *bst[K, V]) Max() (K, V, bool) {
	if t.root == nil {
		var zeroK K
		var zeroV V
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

// Floor returns the largest key in the BST less than or equal to key.
func (t *bst[K, V]) Floor(key K) (K, V, bool) {
	n := t._floor(t.root, key)
	if n == nil {
		var zeroK K
		var zeroV V
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

// Ceiling returns the smallest key in the BST greater than or equal to key.
func (t *bst[K, V]) Ceiling(key K) (K, V, bool) {
	n := t._ceiling(t.root, key)
	if n == nil {
		var zeroK K
		var zeroV V
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

// DeleteMin deletes the smallest key and associated value from the BST.
func (t *bst[K, V]) DeleteMin() (K, V, bool) {
	if t.root == nil {
		var zeroK K
		var zeroV V
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
	n.size = 1 + t._size(n.left) + t._size(n.right)
	return n, min
}

// DeleteMax deletes the largest key and associated value from the BST.
func (t *bst[K, V]) DeleteMax() (K, V, bool) {
	if t.root == nil {
		var zeroK K
		var zeroV V
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
	n.size = 1 + t._size(n.left) + t._size(n.right)
	return n, max
}

// Select returns the k-th smallest key in the BST.
func (t *bst[K, V]) Select(rank int) (K, V, bool) {
	if rank < 0 || rank >= t.Size() {
		var zeroK K
		var zeroV V
		return zeroK, zeroV, false
	}

	n := t._select(t.root, rank)
	return n.key, n.val, true
}

func (t *bst[K, V]) _select(n *bstNode[K, V], rank int) *bstNode[K, V] {
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

// Rank returns the number of keys in the BST less than key.
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
		return 1 + t._size(n.left) + t._rank(n.right, key)
	default:
		return t._size(n.left)
	}
}

// Range returns all keys and associated values in the BST between two given keys.
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
		*kvs = append(*kvs, KeyValue[K, V]{Key: n.key, Val: n.val})
		len++
	}
	if cmpHi > 0 {
		len += t._range(n.right, kvs, lo, hi)
	}

	return len
}

// RangeSize returns the number of keys in the BST between two given keys.
func (t *bst[K, V]) RangeSize(lo, hi K) int {
	if t.cmpKey(lo, hi) > 0 {
		return 0
	} else if _, found := t.Get(hi); found {
		return 1 + t.Rank(hi) - t.Rank(lo)
	} else {
		return t.Rank(hi) - t.Rank(lo)
	}
}

// String returns a string representation of the BST.
func (t *bst[K, V]) String() string {
	i := 0
	pairs := make([]string, t.Size())

	t._traverse(t.root, Ascending, func(n *bstNode[K, V]) bool {
		pairs[i] = fmt.Sprintf("<%v:%v>", n.key, n.val)
		i++
		return true
	})

	return fmt.Sprintf("{%s}", strings.Join(pairs, " "))
}

// Equal determines whether or not two BSTs have the same key-values.
func (t *bst[K, V]) Equal(rhs SymbolTable[K, V]) bool {
	t2, ok := rhs.(*bst[K, V])
	if !ok {
		return false
	}

	return t._traverse(t.root, Ascending, func(n *bstNode[K, V]) bool { // t ⊂ t2
		val, ok := t2.Get(n.key)
		return ok && t.eqVal(n.val, val)
	}) && t2._traverse(t2.root, Ascending, func(n *bstNode[K, V]) bool { // t2 ⊂ t
		val, ok := t.Get(n.key)
		return ok && t.eqVal(n.val, val)
	})
}

// All returns an iterator sequence containing all the key-values in the BST.
func (t *bst[K, V]) All() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		t._traverse(t.root, Ascending, func(n *bstNode[K, V]) bool {
			return yield(n.key, n.val)
		})
	}
}

// AnyMatch returns true if at least one key-value in the BST satisfies the provided predicate.
func (t *bst[K, V]) AnyMatch(p Predicate2[K, V]) bool {
	return !t._traverse(t.root, VLR, func(n *bstNode[K, V]) bool {
		return !p(n.key, n.val)
	})
}

// AllMatch returns true if all key-values in the BST satisfy the provided predicate.
// If the BST is empty, it returns true.
func (t *bst[K, V]) AllMatch(p Predicate2[K, V]) bool {
	return t._traverse(t.root, VLR, func(n *bstNode[K, V]) bool {
		return p(n.key, n.val)
	})
}

// FirstMatch returns the first key-value in the BST that satisfies the given predicate.
// If no match is found, it returns the zero values of K and V, along with false.
func (t *bst[K, V]) FirstMatch(p Predicate2[K, V]) (K, V, bool) {
	var k K
	var v V
	var ok bool

	t._traverse(t.root, VLR, func(n *bstNode[K, V]) bool {
		if p(n.key, n.val) {
			k, v, ok = n.key, n.val, true
			return false
		}
		return true
	})

	return k, v, ok
}

// SelectMatch selects a subset of key-values from the BST that satisfy the given predicate.
// It returns a new BST containing the matching key-values, of the same type as the original BST.
func (t *bst[K, V]) SelectMatch(p Predicate2[K, V]) Collection2[K, V] {
	newST := NewBST[K, V](t.cmpKey, t.eqVal)

	t._traverse(t.root, VLR, func(n *bstNode[K, V]) bool {
		if p(n.key, n.val) {
			newST.Put(n.key, n.val)
		}
		return true
	})

	return newST
}

// Traverse performs a traversal of the BST using the specified traversal order
// and yields the key-value of each node to the provided VisitFunc2 function.
//
// If the function returns false, the traversal is halted.
func (t *bst[K, V]) Traverse(order TraverseOrder, visit VisitFunc2[K, V]) {
	t._traverse(t.root, order, func(n *bstNode[K, V]) bool {
		return visit(n.key, n.val)
	})
}

func (t *bst[K, V]) _traverse(n *bstNode[K, V], order TraverseOrder, visit func(*bstNode[K, V]) bool) bool {
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

// DOT generates a representation of the BST in DOT format.
// This format is commonly used for visualizing graphs with Graphviz tools.
func (t *bst[K, V]) DOT() string {
	// Create a map of node --> id
	var id int
	nodeID := map[*bstNode[K, V]]int{}
	t._traverse(t.root, VLR, func(n *bstNode[K, V]) bool {
		id++
		nodeID[n] = id
		return true
	})

	graph := dot.NewGraph(true, true, false, "BST", "", "", "", dot.ShapeOval)

	t._traverse(t.root, VLR, func(n *bstNode[K, V]) bool {
		name := fmt.Sprintf("%d", nodeID[n])
		label := fmt.Sprintf("%v,%v", n.key, n.val)

		graph.AddNode(dot.NewNode(name, "", label, "", "", "", "", ""))

		if n.left != nil {
			left := fmt.Sprintf("%d", nodeID[n.left])
			graph.AddEdge(dot.NewEdge(name, left, dot.EdgeTypeDirected, "", "", "", "", "", ""))
		}

		if n.right != nil {
			right := fmt.Sprintf("%d", nodeID[n.right])
			graph.AddEdge(dot.NewEdge(name, right, dot.EdgeTypeDirected, "", "", "", "", "", ""))
		}

		return true
	})

	return graph.DOT()
}
