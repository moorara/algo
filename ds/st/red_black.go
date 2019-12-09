package st

import (
	"fmt"

	"github.com/moorara/algo/pkg/graphviz"
)

const (
	red   = true
	black = false
)

type rbNode struct {
	key   interface{}
	value interface{}
	left  *rbNode
	right *rbNode
	size  int
	color bool
}

// redBlack is a left-leaning Red-Black tree.
type redBlack struct {
	root   *rbNode
	cmpKey CompareFunc
}

// NewRedBlack creates a new Red-Black tree.
//
// A Red-Black tree is 2-3 Tree represented as a binary search tree.
// In a left-leaning Red-Black tree, left-leaning red links are used to construct 3-nodes.
// A left-leaning Red-Black tree is a BST such that:
//   Red links lean lef.
//   No node has two red links connect to it.
//   Every path from root to null link has the same number of black links.
func NewRedBlack(cmpKey CompareFunc) OrderedSymbolTable {
	return &redBlack{
		root:   nil,
		cmpKey: cmpKey,
	}
}

func (t *redBlack) isBST(n *rbNode, min, max interface{}) bool {
	if n == nil {
		return true
	}

	if (min != nil && t.cmpKey(n.key, min) <= 0) ||
		(max != nil && t.cmpKey(n.key, max) >= 0) {
		return false
	}

	return t.isBST(n.left, min, n.key) && t.isBST(n.right, n.key, max)
}

func (t *redBlack) isSizeOK(n *rbNode) bool {
	if n == nil {
		return true
	}

	if n.size != 1+t.size(n.left)+t.size(n.right) {
		return false
	}

	return t.isSizeOK(n.left) && t.isSizeOK(n.right)
}

// A Red-Black tree should have no red right links, and at most one left red links in a row on any path.
func (t *redBlack) isRedBlack(n *rbNode) bool {
	if n == nil {
		return true
	}

	if t.isRed(n.right) ||
		n != t.root && t.isRed(n) && t.isRed(n.left) {
		return false
	}

	return true
}

func (t *redBlack) _isBalanced(n *rbNode, count int) bool {
	if n == nil {
		return count == 0
	}

	if !t.isRed(n) {
		count--
	}
	return t._isBalanced(n.left, count) && t._isBalanced(n.right, count)
}

// All paths from root to leaf should have same number of black edges.
func (t *redBlack) isBalanced() bool {
	count := 0
	var n *rbNode
	for n = t.root; n != nil; n = n.left {
		if !t.isRed(n) {
			count++
		}
	}

	return t._isBalanced(t.root, count)
}

func (t *redBlack) verify() bool {
	return t.isBST(t.root, nil, nil) &&
		t.isSizeOK(t.root) &&
		t.isRedBlack(t.root) &&
		t.isBalanced()
}

func (t *redBlack) size(n *rbNode) int {
	if n == nil {
		return 0
	}

	return n.size
}

func (t *redBlack) height(n *rbNode) int {
	if n == nil {
		return 0
	}

	return 1 + max(t.height(n.left), t.height(n.right))
}

func (t *redBlack) isRed(n *rbNode) bool {
	if n == nil {
		return black
	}

	return n.color == red
}

func (t *redBlack) rotateLeft(n *rbNode) *rbNode {
	r := n.right
	n.right = r.left
	r.left = n

	r.color = r.left.color
	r.left.color = red
	r.size = n.size
	n.size = 1 + t.size(n.left) + t.size(n.right)

	return r
}

func (t *redBlack) rotateRight(n *rbNode) *rbNode {
	l := n.left
	n.left = l.right
	l.right = n

	l.color = l.right.color
	l.right.color = red
	l.size = n.size
	n.size = 1 + t.size(n.left) + t.size(n.right)

	return l
}

func (t *redBlack) flipColors(n *rbNode) {
	n.color = !n.color
	n.left.color = !n.left.color
	n.right.color = !n.right.color
}

// Assuming n is red and both n.left and n.left.left are black, make n.left or one of its children red.
func (t *redBlack) moveRedLeft(n *rbNode) *rbNode {
	/* if n == nil || !t.isRed(n) || t.isRed(n.left) || t.isRed(n.left.left) {
		return nil
	} */

	t.flipColors(n)
	if t.isRed(n.right.left) {
		n.right = t.rotateRight(n.right)
		n = t.rotateLeft(n)
		t.flipColors(n)
	}

	return n
}

// Assuming n is red and both n.right and n.right.left are black, make n.right or one of its children red.
func (t *redBlack) moveRedRight(n *rbNode) *rbNode {
	/* if n == nil || !t.isRed(n) || t.isRed(n.right) || t.isRed(n.right.left) {
		return nil
	} */

	t.flipColors(n)
	if t.isRed(n.left.left) {
		n = t.rotateRight(n)
		t.flipColors(n)
	}

	return n
}

// Assuming n is not nil.
func (t *redBlack) balance(n *rbNode) *rbNode {
	/* if n == nil {
		return nil
	} */

	if t.isRed(n.right) {
		n = t.rotateLeft(n)
	}
	if t.isRed(n.left) && t.isRed(n.left.left) {
		n = t.rotateRight(n)
	}
	if t.isRed(n.left) && t.isRed(n.right) {
		t.flipColors(n)
	}

	n.size = 1 + t.size(n.left) + t.size(n.right)
	return n
}

// Size returns the number of key-value pairs in Red-Black tree.
func (t *redBlack) Size() int {
	return t.size(t.root)
}

// Height returns the height of Red-Black tree.
func (t *redBlack) Height() int {
	return t.height(t.root)
}

// IsEmpty returns true if Red-Black tree is empty.
func (t *redBlack) IsEmpty() bool {
	return t.root == nil
}

func (t *redBlack) _put(n *rbNode, key, value interface{}) *rbNode {
	if n == nil {
		return &rbNode{
			key:   key,
			value: value,
			size:  1,
			color: red,
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

	n.size = 1 + t.size(n.left) + t.size(n.right)
	return n
}

// Put adds a new key-value pair to Red-Black tree.
func (t *redBlack) Put(key, value interface{}) {
	if key == nil {
		return
	}

	t.root = t._put(t.root, key, value)
	t.root.color = black
}

func (t *redBlack) _get(n *rbNode, key interface{}) (interface{}, bool) {
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

// Get returns the value of a given key in Red-Black tree.
func (t *redBlack) Get(key interface{}) (interface{}, bool) {
	return t._get(t.root, key)
}

func (t *redBlack) _delete(n *rbNode, key interface{}) (*rbNode, interface{}, bool) {
	var ok bool
	var value interface{}

	if t.cmpKey(key, n.key) < 0 {
		if !t.isRed(n.left) && !t.isRed(n.left.left) {
			n = t.moveRedLeft(n)
		}
		n.left, value, ok = t._delete(n.left, key)
	} else {
		if t.isRed(n.left) {
			n = t.rotateRight(n)
		}

		if t.cmpKey(key, n.key) == 0 && n.right == nil {
			return nil, n.value, true
		}

		if !t.isRed(n.right) && !t.isRed(n.right.left) {
			n = t.moveRedRight(n)
		}

		if t.cmpKey(key, n.key) == 0 {
			var min *rbNode
			value, ok = n.value, true
			n.right, min = t._deleteMin(n.right)
			n.key, n.value = min.key, min.value
		} else {
			n.right, value, ok = t._delete(n.right, key)
		}
	}

	return t.balance(n), value, ok
}

// Delete removes a key-value pair from Red-Black tree.
func (t *redBlack) Delete(key interface{}) (value interface{}, ok bool) {
	if t.root == nil || key == nil {
		return nil, false
	}

	if !t.isRed(t.root.left) && !t.isRed(t.root.right) {
		t.root.color = red
	}

	t.root, value, ok = t._delete(t.root, key)
	if t.root != nil {
		t.root.color = black
	}
	return value, ok
}

// KeyValues returns all key-value pairs in Red-Black tree.
func (t *redBlack) KeyValues() []KeyValue {
	i := 0
	kvs := make([]KeyValue, t.Size())

	t._traverse(t.root, InOrder, func(n *rbNode) bool {
		kvs[i] = KeyValue{n.key, n.value}
		i++
		return true
	})
	return kvs
}

func (t *redBlack) _min(n *rbNode) *rbNode {
	if n.left == nil {
		return n
	}
	return t._min(n.left)
}

// Min returns the minimum key and its value in Red-Black tree.
func (t *redBlack) Min() (interface{}, interface{}) {
	if t.root == nil {
		return nil, nil
	}

	n := t._min(t.root)
	return n.key, n.value
}

func (t *redBlack) _max(n *rbNode) *rbNode {
	if n.right == nil {
		return n
	}
	return t._max(n.right)
}

// Max returns the maximum key and its value in Red-Black tree.
func (t *redBlack) Max() (interface{}, interface{}) {
	if t.root == nil {
		return nil, nil
	}

	n := t._max(t.root)
	return n.key, n.value
}

func (t *redBlack) _floor(n *rbNode, key interface{}) *rbNode {
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

// Floor returns the largest key in Red-Black tree less than or equal to key.
func (t *redBlack) Floor(key interface{}) (interface{}, interface{}) {
	n := t._floor(t.root, key)
	if n == nil {
		return nil, nil
	}
	return n.key, n.value
}

func (t *redBlack) _ceiling(n *rbNode, key interface{}) *rbNode {
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

// Ceiling returns the smallest key in Red-Black tree greater than or equal to key.
func (t *redBlack) Ceiling(key interface{}) (interface{}, interface{}) {
	n := t._ceiling(t.root, key)
	if n == nil {
		return nil, nil
	}
	return n.key, n.value
}

func (t *redBlack) _rank(n *rbNode, key interface{}) int {
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

// Rank returns the number of keys in Red-Black tree less than key.
func (t *redBlack) Rank(key interface{}) int {
	if key == nil {
		return -1
	}

	return t._rank(t.root, key)
}

func (t *redBlack) _select(n *rbNode, rank int) *rbNode {
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

// Select return the k-th smallest key in Red-Black tree.
func (t *redBlack) Select(rank int) (interface{}, interface{}) {
	if rank < 0 || rank >= t.Size() {
		return nil, nil
	}

	n := t._select(t.root, rank)
	return n.key, n.value
}

func (t *redBlack) _deleteMin(n *rbNode) (*rbNode, *rbNode) {
	if n.left == nil {
		return n.right, n
	}

	if !t.isRed(n.left) && !t.isRed(n.left.left) {
		n = t.moveRedLeft(n)
	}

	var min *rbNode
	n.left, min = t._deleteMin(n.left)
	return t.balance(n), min
}

// DeleteMin removes the smallest key and associated value from Red-Black tree.
func (t *redBlack) DeleteMin() (interface{}, interface{}) {
	if t.root == nil {
		return nil, nil
	}

	if !t.isRed(t.root.left) && !t.isRed(t.root.right) {
		t.root.color = red
	}

	var min *rbNode
	t.root, min = t._deleteMin(t.root)
	if t.root != nil {
		t.root.color = black
	}
	return min.key, min.value
}

func (t *redBlack) _deleteMax(n *rbNode) (*rbNode, *rbNode) {
	if t.isRed(n.left) {
		n = t.rotateRight(n)
	}

	if n.right == nil {
		return n.left, n
	}

	if !t.isRed(n.right) && !t.isRed(n.right.left) {
		n = t.moveRedRight(n)
	}

	var max *rbNode
	n.right, max = t._deleteMax(n.right)
	return t.balance(n), max
}

// DeleteMax removes the largest key and associated value from Red-Black tree.
func (t *redBlack) DeleteMax() (interface{}, interface{}) {
	if t.root == nil {
		return nil, nil
	}

	if !t.isRed(t.root.left) && !t.isRed(t.root.right) {
		t.root.color = red
	}

	var max *rbNode
	t.root, max = t._deleteMax(t.root)
	if t.root != nil {
		t.root.color = black
	}
	return max.key, max.value
}

// RangeSize returns the number of keys in Red-Black tree between two given keys.
func (t *redBlack) RangeSize(lo, hi interface{}) int {
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

func (t *redBlack) _range(n *rbNode, kvs *[]KeyValue, lo, hi interface{}) int {
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

// Range returns all keys and associated values in Red-Black tree between two given keys.
func (t *redBlack) Range(lo, hi interface{}) []KeyValue {
	if lo == nil || hi == nil {
		return nil
	}

	kvs := make([]KeyValue, 0)
	len := t._range(t.root, &kvs, lo, hi)
	return kvs[0:len]
}

func (t *redBlack) _traverse(n *rbNode, order TraverseOrder, visit func(*rbNode) bool) bool {
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

// Traverse is used for visiting all key-value pairs in Red-Black tree.
func (t *redBlack) Traverse(order TraverseOrder, visit VisitFunc) {
	if order != PreOrder && order != InOrder && order != PostOrder {
		return
	}

	t._traverse(t.root, order, func(n *rbNode) bool {
		return visit(n.key, n.value)
	})
}

// Graphviz returns a visualization of Red-Black tree in Graphviz format.
func (t *redBlack) Graphviz() string {
	var parent, left, right, label, nodeColor, fontColor, edgeColor string
	graph := graphviz.NewGraph(true, true, "RedBlack", "", "", graphviz.StyleFilled, graphviz.ShapeOval)

	t._traverse(t.root, PreOrder, func(n *rbNode) bool {
		parent = fmt.Sprintf("%v", n.key)
		label = fmt.Sprintf("%v,%v", n.key, n.value)
		if t.isRed(n) {
			nodeColor = graphviz.ColorRed
			fontColor = graphviz.ColorWhite
		} else {
			nodeColor = graphviz.ColorBlack
			fontColor = graphviz.ColorWhite
		}
		graph.AddNode(graphviz.NewNode(parent, "", label, nodeColor, "", "", fontColor, ""))
		if n.left != nil {
			left = fmt.Sprintf("%v", n.left.key)
			if t.isRed(n.left) {
				edgeColor = graphviz.ColorRed
			} else {
				edgeColor = graphviz.ColorBlack
			}
			graph.AddEdge(graphviz.NewEdge(parent, left, graphviz.EdgeTypeDirected, "", "", edgeColor, "", ""))
		}
		if n.right != nil {
			right = fmt.Sprintf("%v", n.right.key)
			graph.AddEdge(graphviz.NewEdge(parent, right, graphviz.EdgeTypeDirected, "", "", "", "", ""))
		}
		return true
	})

	return graph.DotCode()
}
