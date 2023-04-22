package trie

import (
	"fmt"

	"github.com/moorara/algo/generic"
	"github.com/moorara/algo/internal/graphviz"
)

type binaryNode[V any] struct {
	char        byte
	val         V
	term        bool
	left, right *binaryNode[V]
}

type binary[V any] struct {
	size  int
	root  *binaryNode[V]
	eqVal generic.EqualFunc[V]
}

// NewBinaryTrie creates a new Binary Trie tree.
//
// A trie, or prefix tree, is an ordered tree uses the digits in the keys where the keys are usually strings.
// We can use any radix to decompose the keys into digits.
// We shall choose radixes so that the digits are decimal digits, English letters, ASCII characters, etc.
//
// In a trie, keys are stored on a path from the root node to any arbitrary node.
// In a sense, keys are stored on edges and not in nodes and the nodes themselves do not store the keys.
// The root node is always associated with the empty string.
// All the descendants of a node have the same common prefix of the string associated with that node.
//
// A binary trie is a binary tree in which the left child represents the next character in a string
// and the right child represents an alternative character in the same position as the current character.
// The root node is always a sentinel node with a nil right child.
//
//	_
//	|
//	b-------------d
//	|             |
//	a---------o   a
//	|         |   |
//	b--d*--n  x*  d*--n
//	|      |          |
//	y*     k*         c
//	                  |
//	                  e*
//
// Includes words baby, bad, bank, box, dad, and dance.
//
// The second parameter (eqVal) is needed only if you want to use the Equals method.
func NewBinary[V any](eqVal generic.EqualFunc[V]) Trie[V] {
	return &binary[V]{
		size:  0,
		root:  new(binaryNode[V]),
		eqVal: eqVal,
	}
}

// nolint: unused
func (t *binary[V]) verify() bool {
	return t.root != nil && t.root.right == nil &&
		t._isTrie(t.root.left) &&
		t._isSizeOK() &&
		t._isRankOK()
}

// nolint: unused
func (t *binary[V]) _isTrie(n *binaryNode[V]) bool {
	if n == nil {
		return true
	}

	return t._isTrie(n.left) && t._isTrie(n.right) &&
		(n.right == nil || n.right.char > n.char)
}

// nolint: unused
func (t *binary[V]) _isSizeOK() bool {
	size := 0
	t._traverse(t.root.left, "", VLR, func(_ string, n *binaryNode[V]) bool {
		if n.term {
			size++
		}
		return true
	})

	return t.size == size
}

// nolint: unused
func (t *binary[V]) _isRankOK() bool {
	for i := 0; i < t.Size(); i++ {
		k, _, _ := t.Select(i)
		if t.Rank(k) != i {
			return false
		}
	}

	for _, kv := range t.KeyValues() {
		k, _, _ := t.Select(t.Rank(kv.Key))
		if kv.Key != k {
			return false
		}
	}

	return true
}

// Size returns the number of key-value pairs in the Binary Trie.
func (t *binary[V]) Size() int {
	return t.size
}

// Height returns the height of the Binary Trie.
func (t *binary[V]) Height() int {
	return t._height(t.root.left)
}

func (t *binary[V]) _height(n *binaryNode[V]) int {
	if n == nil {
		return 0
	}

	return 1 + generic.Max[int](t._height(n.left), t._height(n.right))
}

// IsEmpty returns true if the Binary Trie is empty.
func (t *binary[V]) IsEmpty() bool {
	return t.size == 0
}

// Put adds a new key-value pair to the Binary Trie.
func (t *binary[V]) Put(key string, val V) {
	// Special case of empty string
	if key == "" {
		panic("Trie does not allow empty string key")
	}

	t.root.left = t._put(t.root.left, key, val)
}

func (t *binary[V]) _put(n *binaryNode[V], key string, val V) *binaryNode[V] {
	if n == nil {
		n = &binaryNode[V]{
			char: key[0],
		}
	} else if n.char > key[0] { // Keep right links sorted
		n = &binaryNode[V]{
			char:  key[0],
			right: n,
		}
	}

	if n.char == key[0] {
		if len(key) == 1 {
			if !n.term {
				t.size++
			}
			n.val, n.term = val, true
		} else {
			n.left = t._put(n.left, key[1:], val)
		}
	} else {
		n.right = t._put(n.right, key, val)
	}

	return n
}

// Get returns the value of a given key in the Binary Trie.
func (t *binary[V]) Get(key string) (V, bool) {
	// Special case of empty string
	if key == "" {
		panic("Trie does not allow empty string key")
	}

	return t._get(t.root.left, key)
}

func (t *binary[V]) _get(n *binaryNode[V], key string) (V, bool) {
	if n == nil || len(key) == 0 || n.char > key[0] { // right links are sorted
		var zeroV V
		return zeroV, false
	}

	if n.char == key[0] {
		if n.term && len(key) == 1 {
			return n.val, true
		}
		return t._get(n.left, key[1:])
	}

	return t._get(n.right, key)
}

// Delete removes a key-value pair from the Binary Trie.
func (t *binary[V]) Delete(key string) (val V, ok bool) {
	// Special case of empty string
	if key == "" {
		panic("Trie does not allow empty string key")
	}

	t.root.left, val, ok = t._delete(t.root.left, key)
	return val, ok
}

func (t *binary[V]) _delete(n *binaryNode[V], key string) (*binaryNode[V], V, bool) {
	var zeroV, val V
	var ok bool

	if n == nil || n.char > key[0] { // right links are sorted
		return nil, zeroV, false
	}

	if n.char == key[0] {
		if len(key) == 1 {
			t.size--
			val, ok = n.val, true
			n.val, n.term = zeroV, false
		} else {
			n.left, val, ok = t._delete(n.left, key[1:])
		}
		if n.left == nil {
			n = n.right
		}
	} else {
		n.right, val, ok = t._delete(n.right, key)
	}

	return n, val, ok
}

// KeyValues returns all key-value pairs in the Binary Trie.
func (t *binary[V]) KeyValues() []KeyValue[V] {
	kvs := make([]KeyValue[V], 0, t.Size())
	t._traverse(t.root.left, "", Ascending, func(k string, n *binaryNode[V]) bool {
		if n.term {
			kvs = append(kvs, KeyValue[V]{k, n.val})
		}
		return true
	})

	return kvs
}

// Equals determines whether or not two Binary Tries have the same key-value pairs.
func (t *binary[V]) Equals(u Trie[V]) bool {
	tt, ok := u.(*binary[V])
	if !ok {
		return false
	}

	return t._traverse(t.root.left, "", Ascending, func(k string, n *binaryNode[V]) bool { // t ⊂ tt
		if n.term {
			val, ok := tt.Get(k)
			return ok && t.eqVal(n.val, val)
		}
		return true
	}) && tt._traverse(tt.root.left, "", Ascending, func(k string, n *binaryNode[V]) bool { // tt ⊂ t
		if n.term {
			val, ok := t.Get(k)
			return ok && t.eqVal(n.val, val)
		}
		return true
	})
}

// Min returns the minimum key and its value in the Binary Trie.
func (t *binary[V]) Min() (string, V, bool) {
	var key string
	var val V
	var ok bool

	t._traverse(t.root.left, "", Ascending, func(k string, n *binaryNode[V]) bool {
		if n.term {
			key, val, ok = k, n.val, true
			return false
		}
		return true
	})

	return key, val, ok
}

// Max returns the maximum key and its value in the Binary Trie.
func (t *binary[V]) Max() (string, V, bool) {
	var key string
	var val V
	var ok bool

	t._traverse(t.root.left, "", Descending, func(k string, n *binaryNode[V]) bool {
		if n.term {
			key, val, ok = k, n.val, true
			return false
		}
		return true
	})

	return key, val, ok
}

// Floor returns the largest key in the Binary Trie less than or equal to key.
func (t *binary[V]) Floor(key string) (string, V, bool) {
	var lastKey string
	var lastVal V
	var ok bool

	t._traverse(t.root.left, "", Ascending, func(k string, n *binaryNode[V]) bool {
		if n.term {
			if key < k {
				return false
			}
			lastKey, lastVal, ok = k, n.val, true
		}

		return true
	})

	return lastKey, lastVal, ok
}

// Ceiling returns the smallest key in the Binary Trie greater than or equal to key.
func (t *binary[V]) Ceiling(key string) (string, V, bool) {
	var lastKey string
	var lastVal V
	var ok bool

	t._traverse(t.root.left, "", Descending, func(k string, n *binaryNode[V]) bool {
		if n.term {
			if k < key {
				return false
			}
			lastKey, lastVal, ok = k, n.val, true
		}

		return true
	})

	return lastKey, lastVal, ok
}

// DeleteMin removes the smallest key and associated value from the Binary Trie.
func (t *binary[V]) DeleteMin() (string, V, bool) {
	key, val, ok := t.Min()
	if !ok {
		return key, val, false
	}

	if _, ok = t.Delete(key); !ok {
		return key, val, false
	}

	return key, val, true
}

// DeleteMax removes the largest key and associated value from the Binary Trie.
func (t *binary[V]) DeleteMax() (string, V, bool) {
	key, val, ok := t.Max()
	if !ok {
		return key, val, false
	}

	if _, ok = t.Delete(key); !ok {
		return key, val, false
	}

	return key, val, true
}

// Select returns the k-th smallest key in the Binary Trie.
func (t *binary[V]) Select(rank int) (string, V, bool) {
	var lastKey string
	var lastVal V
	var ok bool

	if rank < 0 || rank >= t.Size() {
		return lastKey, lastVal, false
	}

	i := 0
	t._traverse(t.root.left, "", Ascending, func(k string, n *binaryNode[V]) bool {
		if n.term {
			if i == rank {
				lastKey, lastVal, ok = k, n.val, true
				return false
			}

			i++
		}

		return true
	})

	return lastKey, lastVal, ok
}

// Rank returns the number of keys in the Binary Trie less than key.
func (t *binary[V]) Rank(key string) int {
	i := 0
	t._traverse(t.root.left, "", Ascending, func(k string, n *binaryNode[V]) bool {
		if n.term {
			if k == key {
				return false
			}

			i++
		}

		return true
	})

	return i
}

// RangeSize returns the number of keys in the Binary Trie between two given keys.
func (t *binary[V]) RangeSize(lo, hi string) int {
	i := 0
	t._traverse(t.root.left, "", Ascending, func(k string, n *binaryNode[V]) bool {
		if n.term {
			if lo <= k && k <= hi {
				i++
			} else if k > hi {
				return false
			}
		}

		return true
	})

	return i
}

// Range returns all keys and associated values in the Binary Trie between two given keys.
func (t *binary[V]) Range(lo, hi string) []KeyValue[V] {
	kvs := []KeyValue[V]{}
	t._traverse(t.root.left, "", Ascending, func(k string, n *binaryNode[V]) bool {
		if n.term {
			if lo <= k && k <= hi {
				kvs = append(kvs, KeyValue[V]{k, n.val})
			} else if k > hi {
				return false
			}
		}

		return true
	})

	return kvs
}

// Traverse is used for visiting all key-value pairs in the Binary Trie.
func (t *binary[V]) Traverse(order TraversalOrder, visit VisitFunc[V]) {
	t._traverse(t.root, "", order, func(_ string, n *binaryNode[V]) bool {
		// Special case of empty string
		if n == t.root {
			return visit("", n.val)
		}
		return visit(string(n.char), n.val)
	})
}

func (t *binary[V]) _traverse(n *binaryNode[V], prefix string, order TraversalOrder, visit func(string, *binaryNode[V]) bool) bool {
	if n == nil {
		return true
	}

	next := prefix + string(n.char)

	switch order {
	case VLR, Ascending:
		return visit(next, n) && t._traverse(n.left, next, order, visit) && t._traverse(n.right, prefix, order, visit)
	case VRL:
		return visit(next, n) && t._traverse(n.right, prefix, order, visit) && t._traverse(n.left, next, order, visit)
	case LVR:
		return t._traverse(n.left, next, order, visit) && visit(next, n) && t._traverse(n.right, prefix, order, visit)
	case RVL:
		return t._traverse(n.right, prefix, order, visit) && visit(next, n) && t._traverse(n.left, next, order, visit)
	case LRV:
		return t._traverse(n.left, next, order, visit) && t._traverse(n.right, prefix, order, visit) && visit(next, n)
	case RLV, Descending:
		return t._traverse(n.right, prefix, order, visit) && t._traverse(n.left, next, order, visit) && visit(next, n)
	default:
		return false
	}
}

// Graphviz returns a visualization of the Binary Trie in Graphviz format.
func (t *binary[V]) Graphviz() string {
	// Create a map of node --> id
	var id int
	nodeID := map[*binaryNode[V]]int{}
	t._traverse(t.root, "", VLR, func(_ string, n *binaryNode[V]) bool {
		id++
		nodeID[n] = id
		return true
	})

	graph := graphviz.NewGraph(true, true, false, "Binary Trie", "", "", "", graphviz.ShapeCircle)

	t._traverse(t.root, "", VLR, func(_ string, n *binaryNode[V]) bool {
		name := fmt.Sprintf("%d", nodeID[n])

		var label string
		var style graphviz.Style
		var color, fontColor graphviz.Color

		switch {
		case n == t.root:
			label = "•"
		case !n.term:
			label = string(n.char)
		default:
			label = fmt.Sprintf("%s,%v", string(n.char), n.val)
			style, color, fontColor = graphviz.StyleFilled, graphviz.ColorBlack, graphviz.ColorWhite
		}

		graph.AddNode(graphviz.NewNode(name, "", label, color, style, "", fontColor, ""))

		if n.left != nil {
			left := fmt.Sprintf("%d", nodeID[n.left])
			graph.AddEdge(graphviz.NewEdge(name, left, graphviz.EdgeTypeDirected, "", "", graphviz.ColorBlue, "", "", ""))
		}

		if n.right != nil {
			right := fmt.Sprintf("%d", nodeID[n.right])
			graph.AddEdge(graphviz.NewEdge(name, right, graphviz.EdgeTypeDirected, "", "", graphviz.ColorRed, "", "", ""))
		}

		return true
	})

	return graph.DotCode()
}

// Match returns all the keys and associated values in Binary trie
// that match the given pattern in which * matches any character.
func (t *binary[V]) Match(pattern string) []KeyValue[V] {
	kvs := []KeyValue[V]{}
	t._match(t.root.left, "", pattern, func(k string, n *binaryNode[V]) {
		kvs = append(kvs, KeyValue[V]{k, n.val})
	})

	return kvs
}

func (t *binary[V]) _match(n *binaryNode[V], prefix, pattern string, visit func(string, *binaryNode[V])) {
	if n == nil || len(pattern) == 0 {
		return
	}

	if c := pattern[0]; c == '*' || c == n.char {
		next := prefix + string(n.char)
		if n.term && len(pattern) == 1 {
			visit(next, n)
		}
		t._match(n.left, next, pattern[1:], visit)
	}

	if c := pattern[0]; c == '*' || c != n.char {
		t._match(n.right, prefix, pattern, visit)
	}
}

// WithPrefix returns all the keys and associated values in Binary trie with the given prefix.
func (t *binary[V]) WithPrefix(key string) []KeyValue[V] {
	kvs := []KeyValue[V]{}
	t._withPrefix(t.root.left, "", key, func(k string, n *binaryNode[V]) {
		kvs = append(kvs, KeyValue[V]{k, n.val})
	})

	return kvs
}

func (t *binary[V]) _withPrefix(n *binaryNode[V], prefix, key string, visit func(string, *binaryNode[V])) {
	if n == nil {
		return
	}

	next := prefix + string(n.char)

	if len(key) == 0 {
		if n.term {
			visit(next, n)
		}
		t._withPrefix(n.left, next, key, visit)
		t._withPrefix(n.right, prefix, key, visit)
	} else if key[0] == n.char {
		if n.term && len(key) == 1 {
			visit(next, n)
		}
		t._withPrefix(n.left, next, key[1:], visit)
	} else {
		t._withPrefix(n.right, prefix, key, visit)
	}
}

// LongestPrefix returns the key and associated value in Binary trie
// that is the longest prefix of the given key.
func (t *binary[V]) LongestPrefixOf(key string) (string, V, bool) {
	var lastKey string
	var lastVal V
	var lastOK bool

	t._allPrefixOf(t.root.left, "", key, func(k string, n *binaryNode[V]) {
		lastKey, lastVal, lastOK = k, n.val, true
	})

	return lastKey, lastVal, lastOK
}

func (t *binary[V]) _allPrefixOf(n *binaryNode[V], prefix, key string, visit func(string, *binaryNode[V])) {
	if n == nil || len(key) == 0 {
		return
	}

	if key[0] == n.char {
		next := prefix + string(n.char)
		if n.term {
			visit(next, n)
		}
		t._allPrefixOf(n.left, next, key[1:], visit)
	} else {
		t._allPrefixOf(n.right, prefix, key, visit)
	}
}
