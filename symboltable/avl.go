package symboltable

import (
	"fmt"

	"github.com/moorara/algo/compare"
	"github.com/moorara/algo/pkg/graphviz"
)

type avlNode struct {
	key    interface{}
	value  interface{}
	left   *avlNode
	right  *avlNode
	size   int
	height int
}

type avl struct {
	root   *avlNode
	cmpKey compare.Func
}

// NewAVL creates a new AVL tree.
//
// AVL tree is a self-balancing binary search tree.
// In an AVL tree, the heights of the left and right subtrees of any node differ by at most 1.
func NewAVL(cmpKey compare.Func) OrderedSymbolTable {
	return &avl{
		root:   nil,
		cmpKey: cmpKey,
	}
}

func (t *avl) isBST(n *avlNode, min, max interface{}) bool {
	if n == nil {
		return true
	}

	if (min != nil && t.cmpKey(n.key, min) <= 0) ||
		(max != nil && t.cmpKey(n.key, max) >= 0) {
		return false
	}

	return t.isBST(n.left, min, n.key) && t.isBST(n.right, n.key, max)
}

func (t *avl) isSizeOK(n *avlNode) bool {
	if n == nil {
		return true
	}

	if n.size != 1+t.size(n.left)+t.size(n.right) {
		return false
	}

	return t.isSizeOK(n.left) && t.isSizeOK(n.right)
}

func (t *avl) isAVL(n *avlNode) bool {
	if n == nil {
		return true
	}

	bf := t.balanceFactor(n)
	if bf > 1 || bf < -1 {
		return false
	}

	return t.isAVL(n.left) && t.isAVL(n.right)
}

func (t *avl) verify() bool {
	return t.isBST(t.root, nil, nil) &&
		t.isSizeOK(t.root) &&
		t.isAVL(t.root)
}

func (t *avl) size(n *avlNode) int {
	if n == nil {
		return 0
	}

	return n.size
}

func (t *avl) height(n *avlNode) int {
	if n == nil {
		return 0
	}

	return n.height
}

func (t *avl) balanceFactor(n *avlNode) int {
	return t.height(n.left) - t.height(n.right)
}

func (t *avl) rotateLeft(n *avlNode) *avlNode {
	r := n.right
	n.right = r.left
	r.left = n

	r.size = n.size
	n.size = 1 + t.size(n.left) + t.size(n.right)
	n.height = 1 + max(t.height(n.left), t.height(n.right))
	r.height = 1 + max(t.height(r.left), t.height(r.right))

	return r
}

func (t *avl) rotateRight(n *avlNode) *avlNode {
	l := n.left
	n.left = l.right
	l.right = n

	l.size = n.size
	n.size = 1 + t.size(n.left) + t.size(n.right)
	n.height = 1 + max(t.height(n.left), t.height(n.right))
	l.height = 1 + max(t.height(l.left), t.height(l.right))

	return l
}

func (t *avl) balance(n *avlNode) *avlNode {
	if t.balanceFactor(n) == 2 {
		if t.balanceFactor(n) == -1 {
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

// Size returns the number of key-value pairs in AVL tree.
func (t *avl) Size() int {
	return t.size(t.root)
}

// Height returns the height of AVL tree.
func (t *avl) Height() int {
	return t.height(t.root)
}

// IsEmpty returns true if AVL tree is empty.
func (t *avl) IsEmpty() bool {
	return t.root == nil
}

func (t *avl) _put(n *avlNode, key, value interface{}) *avlNode {
	if n == nil {
		return &avlNode{
			key:    key,
			value:  value,
			size:   1,
			height: 1,
		}
	}

	cmp := t.cmpKey(key, n.key)
	switch {
	case cmp < 0:
		n.left = t._put(n.left, key, value)
	case cmp > 0:
		n.right = t._put(n.right, key, value)
	default:
		n.value = value
		return n
	}

	n.size = 1 + t.size(n.left) + t.size(n.right)
	n.height = 1 + max(t.height(n.left), t.height(n.right))
	return t.balance(n)
}

// Put adds a new key-value pair to AVL tree.
func (t *avl) Put(key, value interface{}) {
	if key == nil {
		return
	}

	t.root = t._put(t.root, key, value)
}

func (t *avl) _get(n *avlNode, key interface{}) (interface{}, bool) {
	if n == nil || key == nil {
		return nil, false
	}

	cmp := t.cmpKey(key, n.key)
	switch {
	case cmp < 0:
		return t._get(n.left, key)
	case cmp > 0:
		return t._get(n.right, key)
	default:
		return n.value, true
	}
}

// Get returns the value of a given key in AVL tree.
func (t *avl) Get(key interface{}) (interface{}, bool) {
	return t._get(t.root, key)
}

func (t *avl) _delete(n *avlNode, key interface{}) (*avlNode, interface{}, bool) {
	if n == nil || key == nil {
		return n, nil, false
	}

	var ok bool
	var value interface{}

	cmp := t.cmpKey(key, n.key)
	if cmp < 0 {
		n.left, value, ok = t._delete(n.left, key)
	} else if cmp > 0 {
		n.right, value, ok = t._delete(n.right, key)
	} else {
		ok = true
		value = n.value

		if n.left == nil {
			return n.right, value, ok
		} else if n.right == nil {
			return n.left, value, ok
		} else {
			m := n
			n = t._min(m.right)
			n.right, _ = t._deleteMin(m.right)
			n.left = m.left
		}
	}

	n.size = 1 + t.size(n.left) + t.size(n.right)
	n.height = 1 + max(t.height(n.left), t.height(n.right))
	return t.balance(n), value, ok
}

// Delete removes a key-value pair from AVL tree.
func (t *avl) Delete(key interface{}) (value interface{}, ok bool) {
	t.root, value, ok = t._delete(t.root, key)
	return value, ok
}

// KeyValues returns all key-value pairs in AVL tree.
func (t *avl) KeyValues() []KeyValue {
	i := 0
	kvs := make([]KeyValue, t.Size())

	t._traverse(t.root, InOrder, func(n *avlNode) bool {
		kvs[i] = KeyValue{n.key, n.value}
		i++
		return true
	})
	return kvs
}

func (t *avl) _min(n *avlNode) *avlNode {
	if n.left == nil {
		return n
	}
	return t._min(n.left)
}

// Min returns the minimum key and its value in AVL tree.
func (t *avl) Min() (interface{}, interface{}) {
	if t.root == nil {
		return nil, nil
	}

	n := t._min(t.root)
	return n.key, n.value
}

func (t *avl) _max(n *avlNode) *avlNode {
	if n.right == nil {
		return n
	}
	return t._max(n.right)
}

// Max returns the maximum key and its value in AVL tree.
func (t *avl) Max() (interface{}, interface{}) {
	if t.root == nil {
		return nil, nil
	}

	n := t._max(t.root)
	return n.key, n.value
}

func (t *avl) _floor(n *avlNode, key interface{}) *avlNode {
	if n == nil || key == nil {
		return nil
	}

	cmp := t.cmpKey(key, n.key)
	if cmp == 0 {
		return n
	} else if cmp < 0 {
		return t._floor(n.left, key)
	}

	m := t._floor(n.right, key)
	if m != nil {
		return m
	}
	return n
}

// Floor returns the largest key in AVL tree less than or equal to key.
func (t *avl) Floor(key interface{}) (interface{}, interface{}) {
	n := t._floor(t.root, key)
	if n == nil {
		return nil, nil
	}
	return n.key, n.value
}

func (t *avl) _ceiling(n *avlNode, key interface{}) *avlNode {
	if n == nil || key == nil {
		return nil
	}

	cmp := t.cmpKey(key, n.key)
	if cmp == 0 {
		return n
	} else if cmp > 0 {
		return t._ceiling(n.right, key)
	}

	m := t._ceiling(n.left, key)
	if m != nil {
		return m
	}
	return n
}

// Ceiling returns the smallest key in AVL tree greater than or equal to key.
func (t *avl) Ceiling(key interface{}) (interface{}, interface{}) {
	n := t._ceiling(t.root, key)
	if n == nil {
		return nil, nil
	}
	return n.key, n.value
}

func (t *avl) _rank(n *avlNode, key interface{}) int {
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

// Rank returns the number of keys in AVL tree less than key.
func (t *avl) Rank(key interface{}) int {
	if key == nil {
		return -1
	}

	return t._rank(t.root, key)
}

func (t *avl) _select(n *avlNode, rank int) *avlNode {
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

// Select return the k-th smallest key in AVL tree.
func (t *avl) Select(rank int) (interface{}, interface{}) {
	if rank < 0 || rank >= t.Size() {
		return nil, nil
	}

	n := t._select(t.root, rank)
	return n.key, n.value
}

func (t *avl) _deleteMin(n *avlNode) (*avlNode, *avlNode) {
	if n.left == nil {
		return n.right, n
	}

	var min *avlNode
	n.left, min = t._deleteMin(n.left)
	n.size = 1 + t.size(n.left) + t.size(n.right)
	n.height = 1 + max(t.height(n.left), t.height(n.right))
	return t.balance(n), min
}

// DeleteMin removes the smallest key and associated value from AVL tree.
func (t *avl) DeleteMin() (interface{}, interface{}) {
	if t.root == nil {
		return nil, nil
	}

	var min *avlNode
	t.root, min = t._deleteMin(t.root)
	return min.key, min.value
}

func (t *avl) _deleteMax(n *avlNode) (*avlNode, *avlNode) {
	if n.right == nil {
		return n.left, n
	}

	var max *avlNode
	n.right, max = t._deleteMax(n.right)
	n.size = 1 + t.size(n.left) + t.size(n.right)
	return t.balance(n), max
}

// DeleteMax removes the largest key and associated value from AVL tree.
func (t *avl) DeleteMax() (interface{}, interface{}) {
	if t.root == nil {
		return nil, nil
	}

	var max *avlNode
	t.root, max = t._deleteMax(t.root)
	return max.key, max.value
}

// RangeSize returns the number of keys in AVL tree between two given keys.
func (t *avl) RangeSize(lo, hi interface{}) int {
	if lo == nil || hi == nil {
		return -1
	}

	if t.cmpKey(lo, hi) > 0 {
		return 0
	} else if _, found := t.Get(hi); found {
		return 1 + t.Rank(hi) - t.Rank(lo)
	} else {
		return t.Rank(hi) - t.Rank(lo)
	}
}

func (t *avl) _range(n *avlNode, kvs *[]KeyValue, lo, hi interface{}) int {
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
		*kvs = append(*kvs, KeyValue{n.key, n.value})
		len++
	}
	if cmpHi > 0 {
		len += t._range(n.right, kvs, lo, hi)
	}

	return len
}

// Range returns all keys and associated values in AVL tree between two given keys.
func (t *avl) Range(lo, hi interface{}) []KeyValue {
	if lo == nil || hi == nil {
		return nil
	}

	kvs := make([]KeyValue, 0)
	len := t._range(t.root, &kvs, lo, hi)
	return kvs[0:len]
}

func (t *avl) _traverse(n *avlNode, order TraversalOrder, visit func(*avlNode) bool) bool {
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

// Traverse is used for visiting all key-value pairs in AVL tree.
func (t *avl) Traverse(order TraversalOrder, visit VisitFunc) {
	if order != PreOrder && order != InOrder && order != PostOrder {
		return
	}

	t._traverse(t.root, order, func(n *avlNode) bool {
		return visit(n.key, n.value)
	})
}

// Graphviz returns a visualization of AVL tree in Graphviz format.
func (t *avl) Graphviz() string {
	var parent, left, right, label string
	graph := graphviz.NewGraph(true, true, "AVL", "", "", "", graphviz.ShapeOval)

	t._traverse(t.root, PreOrder, func(n *avlNode) bool {
		parent = fmt.Sprintf("%v", n.key)
		label = fmt.Sprintf("%v,%v", n.key, n.value)
		graph.AddNode(graphviz.NewNode(parent, "", label, "", "", "", "", ""))
		if n.left != nil {
			left = fmt.Sprintf("%v", n.left.key)
			graph.AddEdge(graphviz.NewEdge(parent, left, graphviz.EdgeTypeDirected, "", "", "", "", ""))
		}
		if n.right != nil {
			right = fmt.Sprintf("%v", n.right.key)
			graph.AddEdge(graphviz.NewEdge(parent, right, graphviz.EdgeTypeDirected, "", "", "", "", ""))
		}
		return true
	})

	return graph.DotCode()
}
