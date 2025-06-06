package symboltable

import (
	"fmt"
	"iter"
	"strings"

	"github.com/moorara/algo/dot"
	"github.com/moorara/algo/generic"
)

type avlNode[K, V any] struct {
	key          K
	val          V
	left, right  *avlNode[K, V]
	size, height int
}

type avl[K, V any] struct {
	root   *avlNode[K, V]
	cmpKey generic.CompareFunc[K]
	eqVal  generic.EqualFunc[V]
}

// NewAVL creates a new AVL tree.
//
// AVL tree is a self-balancing binary search tree.
// In an AVL tree, the heights of the left and right subtrees of any node differ by at most 1.
//
// The second parameter (eqVal) is needed only if you want to use the Equal method.
func NewAVL[K, V any](cmpKey generic.CompareFunc[K], eqVal generic.EqualFunc[V]) OrderedSymbolTable[K, V] {
	return &avl[K, V]{
		root:   nil,
		cmpKey: cmpKey,
		eqVal:  eqVal,
	}
}

// nolint: unused
func (t *avl[K, V]) verify() bool {
	return t._isBST(t.root, nil, nil) &&
		t._isAVL(t.root) &&
		t._isSizeOK(t.root) &&
		t._isRankOK()
}

// nolint: unused
func (t *avl[K, V]) _isBST(n *avlNode[K, V], min, max *K) bool {
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
func (t *avl[K, V]) _isAVL(n *avlNode[K, V]) bool {
	if n == nil {
		return true
	}

	bf := t.balanceFactor(n)
	if bf < -1 || 1 < bf {
		return false
	}

	return t._isAVL(n.left) && t._isAVL(n.right)
}

// nolint: unused
func (t *avl[K, V]) _isSizeOK(n *avlNode[K, V]) bool {
	if n == nil {
		return true
	}

	if n.size != 1+t._size(n.left)+t._size(n.right) {
		return false
	}

	return t._isSizeOK(n.left) && t._isSizeOK(n.right)
}

// nolint: unused
func (t *avl[K, V]) _isRankOK() bool {
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

func (t *avl[K, V]) balance(n *avlNode[K, V]) *avlNode[K, V] {
	if t.balanceFactor(n) == 2 {
		if t.balanceFactor(n.left) == -1 {
			n.left = t.rotateLeft(n.left)
		}
		n = t.rotateRight(n)
	} else if t.balanceFactor(n) == -2 {
		if t.balanceFactor(n.right) == 1 {
			n.right = t.rotateRight(n.right)
		}
		n = t.rotateLeft(n)
	}

	return n
}

func (t *avl[K, V]) balanceFactor(n *avlNode[K, V]) int {
	return t._height(n.left) - t._height(n.right)
}

func (t *avl[K, V]) rotateLeft(n *avlNode[K, V]) *avlNode[K, V] {
	r := n.right
	n.right = r.left
	r.left = n

	r.size = n.size
	n.size = 1 + t._size(n.left) + t._size(n.right)
	n.height = 1 + max(t._height(n.left), t._height(n.right))
	r.height = 1 + max(t._height(r.left), t._height(r.right))

	return r
}

func (t *avl[K, V]) rotateRight(n *avlNode[K, V]) *avlNode[K, V] {
	l := n.left
	n.left = l.right
	l.right = n

	l.size = n.size
	n.size = 1 + t._size(n.left) + t._size(n.right)
	n.height = 1 + max(t._height(n.left), t._height(n.right))
	l.height = 1 + max(t._height(l.left), t._height(l.right))

	return l
}

// Size returns the number of key-values in the AVL tree.
func (t *avl[K, V]) Size() int {
	return t._size(t.root)
}

func (t *avl[K, V]) _size(n *avlNode[K, V]) int {
	if n == nil {
		return 0
	}

	return n.size
}

// Height returns the height of the AVL tree.
func (t *avl[K, V]) Height() int {
	return t._height(t.root)
}

func (t *avl[K, V]) _height(n *avlNode[K, V]) int {
	if n == nil {
		return 0
	}

	return n.height
}

// IsEmpty returns true if the AVL tree is empty.
func (t *avl[K, V]) IsEmpty() bool {
	return t.root == nil
}

// Put adds a new key-value to the AVL tree.
func (t *avl[K, V]) Put(key K, val V) {
	t.root = t._put(t.root, key, val)
}

func (t *avl[K, V]) _put(n *avlNode[K, V], key K, val V) *avlNode[K, V] {
	if n == nil {
		return &avlNode[K, V]{
			key:    key,
			val:    val,
			size:   1,
			height: 1,
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
		return n
	}

	n.size = 1 + t._size(n.left) + t._size(n.right)
	n.height = 1 + max(t._height(n.left), t._height(n.right))

	return t.balance(n)
}

// Get returns the value of a given key in the AVL tree.
func (t *avl[K, V]) Get(key K) (V, bool) {
	return t._get(t.root, key)
}

func (t *avl[K, V]) _get(n *avlNode[K, V], key K) (V, bool) {
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

// Delete deletes a key-value from the AVL tree.
func (t *avl[K, V]) Delete(key K) (val V, ok bool) {
	t.root, val, ok = t._delete(t.root, key)
	return val, ok
}

func (t *avl[K, V]) _delete(n *avlNode[K, V], key K) (*avlNode[K, V], V, bool) {
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
	n.height = 1 + max(t._height(n.left), t._height(n.right))
	return t.balance(n), val, ok
}

// DeleteAll deletes all key-values from the AVL tree, leaving it empty.
func (t *avl[K, V]) DeleteAll() {
	t.root = nil
}

// Min returns the minimum key and its value in the AVL tree.
func (t *avl[K, V]) Min() (K, V, bool) {
	if t.root == nil {
		var zeroK K
		var zeroV V
		return zeroK, zeroV, false
	}

	n := t._min(t.root)
	return n.key, n.val, true
}

func (t *avl[K, V]) _min(n *avlNode[K, V]) *avlNode[K, V] {
	if n.left == nil {
		return n
	}

	return t._min(n.left)
}

// Max returns the maximum key and its value in the AVL tree.
func (t *avl[K, V]) Max() (K, V, bool) {
	if t.root == nil {
		var zeroK K
		var zeroV V
		return zeroK, zeroV, false
	}

	n := t._max(t.root)
	return n.key, n.val, true
}

func (t *avl[K, V]) _max(n *avlNode[K, V]) *avlNode[K, V] {
	if n.right == nil {
		return n
	}

	return t._max(n.right)
}

// Floor returns the largest key in the AVL tree less than or equal to key.
func (t *avl[K, V]) Floor(key K) (K, V, bool) {
	n := t._floor(t.root, key)
	if n == nil {
		var zeroK K
		var zeroV V
		return zeroK, zeroV, false
	}

	return n.key, n.val, true
}

func (t *avl[K, V]) _floor(n *avlNode[K, V], key K) *avlNode[K, V] {
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

// Ceiling returns the smallest key in the AVL tree greater than or equal to key.
func (t *avl[K, V]) Ceiling(key K) (K, V, bool) {
	n := t._ceiling(t.root, key)
	if n == nil {
		var zeroK K
		var zeroV V
		return zeroK, zeroV, false
	}

	return n.key, n.val, true
}

func (t *avl[K, V]) _ceiling(n *avlNode[K, V], key K) *avlNode[K, V] {
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

// DeleteMin deletes the smallest key and associated value from the AVL tree.
func (t *avl[K, V]) DeleteMin() (K, V, bool) {
	if t.root == nil {
		var zeroK K
		var zeroV V
		return zeroK, zeroV, false
	}

	var min *avlNode[K, V]
	t.root, min = t._deleteMin(t.root)
	return min.key, min.val, true
}

func (t *avl[K, V]) _deleteMin(n *avlNode[K, V]) (*avlNode[K, V], *avlNode[K, V]) {
	if n.left == nil {
		return n.right, n
	}

	var min *avlNode[K, V]
	n.left, min = t._deleteMin(n.left)
	n.size = 1 + t._size(n.left) + t._size(n.right)
	n.height = 1 + max(t._height(n.left), t._height(n.right))
	return t.balance(n), min
}

// DeleteMax deletes the largest key and associated value from the AVL tree.
func (t *avl[K, V]) DeleteMax() (K, V, bool) {
	if t.root == nil {
		var zeroK K
		var zeroV V
		return zeroK, zeroV, false
	}

	var max *avlNode[K, V]
	t.root, max = t._deleteMax(t.root)
	return max.key, max.val, true
}

func (t *avl[K, V]) _deleteMax(n *avlNode[K, V]) (*avlNode[K, V], *avlNode[K, V]) {
	if n.right == nil {
		return n.left, n
	}

	var max *avlNode[K, V]
	n.right, max = t._deleteMax(n.right)
	n.size = 1 + t._size(n.left) + t._size(n.right)
	return t.balance(n), max
}

// Select returns the k-th smallest key in the AVL tree.
func (t *avl[K, V]) Select(rank int) (K, V, bool) {
	if rank < 0 || rank >= t.Size() {
		var zeroK K
		var zeroV V
		return zeroK, zeroV, false
	}

	n := t._select(t.root, rank)
	return n.key, n.val, true
}

func (t *avl[K, V]) _select(n *avlNode[K, V], rank int) *avlNode[K, V] {
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

// Rank returns the number of keys in the AVL tree less than key.
func (t *avl[K, V]) Rank(key K) int {
	return t._rank(t.root, key)
}

func (t *avl[K, V]) _rank(n *avlNode[K, V], key K) int {
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

// Range returns all keys and associated values in the AVL tree between two given keys.
func (t *avl[K, V]) Range(lo, hi K) []generic.KeyValue[K, V] {
	kvs := make([]generic.KeyValue[K, V], 0)
	len := t._range(t.root, &kvs, lo, hi)
	return kvs[0:len]
}

func (t *avl[K, V]) _range(n *avlNode[K, V], kvs *[]generic.KeyValue[K, V], lo, hi K) int {
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
		*kvs = append(*kvs, generic.KeyValue[K, V]{Key: n.key, Val: n.val})
		len++
	}
	if cmpHi > 0 {
		len += t._range(n.right, kvs, lo, hi)
	}

	return len
}

// RangeSize returns the number of keys in the AVL tree between two given keys.
func (t *avl[K, V]) RangeSize(lo, hi K) int {
	if t.cmpKey(lo, hi) > 0 {
		return 0
	} else if _, found := t.Get(hi); found {
		return 1 + t.Rank(hi) - t.Rank(lo)
	} else {
		return t.Rank(hi) - t.Rank(lo)
	}
}

// String returns a string representation of the AVL tree.
func (t *avl[K, V]) String() string {
	i := 0
	pairs := make([]string, t.Size())

	t._traverse(t.root, generic.Ascending, func(n *avlNode[K, V]) bool {
		pairs[i] = fmt.Sprintf("<%v:%v>", n.key, n.val)
		i++
		return true
	})

	return fmt.Sprintf("{%s}", strings.Join(pairs, " "))
}

// Equal determines whether or not two AVLs have the same key-values.
func (t *avl[K, V]) Equal(rhs SymbolTable[K, V]) bool {
	t2, ok := rhs.(*avl[K, V])
	if !ok {
		return false
	}

	return t._traverse(t.root, generic.Ascending, func(n *avlNode[K, V]) bool { // t ⊂ t2
		val, ok := t2.Get(n.key)
		return ok && t.eqVal(n.val, val)
	}) && t2._traverse(t2.root, generic.Ascending, func(n *avlNode[K, V]) bool { // t2 ⊂ t
		val, ok := t.Get(n.key)
		return ok && t.eqVal(n.val, val)
	})
}

// All returns an iterator sequence containing all the key-values in the AVL tree.
func (t *avl[K, V]) All() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		t._traverse(t.root, generic.Ascending, func(n *avlNode[K, V]) bool {
			return yield(n.key, n.val)
		})
	}
}

// AnyMatch returns true if at least one key-value in the AVL tree satisfies the provided predicate.
func (t *avl[K, V]) AnyMatch(p generic.Predicate2[K, V]) bool {
	return !t._traverse(t.root, generic.VLR, func(n *avlNode[K, V]) bool {
		return !p(n.key, n.val)
	})
}

// AllMatch returns true if all key-values in the AVL tree satisfy the provided predicate.
// If the AVL tree is empty, it returns true.
func (t *avl[K, V]) AllMatch(p generic.Predicate2[K, V]) bool {
	return t._traverse(t.root, generic.VLR, func(n *avlNode[K, V]) bool {
		return p(n.key, n.val)
	})
}

// FirstMatch returns the first key-value in the AVL tree that satisfies the given predicate.
// If no match is found, it returns the zero values of K and V, along with false.
func (t *avl[K, V]) FirstMatch(p generic.Predicate2[K, V]) (K, V, bool) {
	var k K
	var v V
	var ok bool

	t._traverse(t.root, generic.VLR, func(n *avlNode[K, V]) bool {
		if p(n.key, n.val) {
			k, v, ok = n.key, n.val, true
			return false
		}
		return true
	})

	return k, v, ok
}

// SelectMatch selects a subset of key-values from the AVL tree that satisfy the given predicate.
// It returns a new AVL tree containing the matching key-values, of the same type as the original AVL tree.
func (t *avl[K, V]) SelectMatch(p generic.Predicate2[K, V]) generic.Collection2[K, V] {
	newST := NewAVL[K, V](t.cmpKey, t.eqVal)

	t._traverse(t.root, generic.VLR, func(n *avlNode[K, V]) bool {
		if p(n.key, n.val) {
			newST.Put(n.key, n.val)
		}
		return true
	})

	return newST
}

// PartitionMatch partitions the key-values in the AVL tree
// into two separate AVL trees based on the provided predicate.
// The first AVL tree contains the key-values that satisfy the predicate (matched key-values),
// while the second AVL tree contains those that do not satisfy the predicate (unmatched key-values).
// Both AVL trees are of the same type as the original AVL tree.
func (t *avl[K, V]) PartitionMatch(p generic.Predicate2[K, V]) (generic.Collection2[K, V], generic.Collection2[K, V]) {
	matched := NewAVL[K, V](t.cmpKey, t.eqVal)
	unmatched := NewAVL[K, V](t.cmpKey, t.eqVal)

	t._traverse(t.root, generic.VLR, func(n *avlNode[K, V]) bool {
		if p(n.key, n.val) {
			matched.Put(n.key, n.val)
		} else {
			unmatched.Put(n.key, n.val)
		}
		return true
	})

	return matched, unmatched
}

// Traverse performs a traversal of the AVL tree using the specified traversal order
// and yields the key-value of each node to the provided VisitFunc2 function.
//
// If the function returns false, the traversal is halted.
func (t *avl[K, V]) Traverse(order generic.TraverseOrder, visit generic.VisitFunc2[K, V]) {
	t._traverse(t.root, order, func(n *avlNode[K, V]) bool {
		return visit(n.key, n.val)
	})
}

func (t *avl[K, V]) _traverse(n *avlNode[K, V], order generic.TraverseOrder, visit func(*avlNode[K, V]) bool) bool {
	if n == nil {
		return true
	}

	switch order {
	case generic.VLR:
		return visit(n) && t._traverse(n.left, order, visit) && t._traverse(n.right, order, visit)
	case generic.VRL:
		return visit(n) && t._traverse(n.right, order, visit) && t._traverse(n.left, order, visit)
	case generic.LVR, generic.Ascending:
		return t._traverse(n.left, order, visit) && visit(n) && t._traverse(n.right, order, visit)
	case generic.RVL, generic.Descending:
		return t._traverse(n.right, order, visit) && visit(n) && t._traverse(n.left, order, visit)
	case generic.LRV:
		return t._traverse(n.left, order, visit) && t._traverse(n.right, order, visit) && visit(n)
	case generic.RLV:
		return t._traverse(n.right, order, visit) && t._traverse(n.left, order, visit) && visit(n)
	default:
		return false
	}
}

// DOT generates a representation of the AVL tree in DOT format.
// This format is commonly used for visualizing graphs with Graphviz tools.
func (t *avl[K, V]) DOT() string {
	// Create a map of node --> id
	var id int
	nodeID := map[*avlNode[K, V]]int{}
	t._traverse(t.root, generic.VLR, func(n *avlNode[K, V]) bool {
		id++
		nodeID[n] = id
		return true
	})

	graph := dot.NewGraph(true, true, false, "AVL", "", "", "", dot.ShapeOval)

	t._traverse(t.root, generic.VLR, func(n *avlNode[K, V]) bool {
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
